package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"

	chpb "github.com/luthersystems/sandbox/api/chpb/v1"

	"github.com/lib/pq"
)

type PostgresConnector struct {
	db *sql.DB
}

func NewPostgresConnector() (*PostgresConnector, error) {
	s := &PostgresConnector{}
	db, err := sql.Open("postgres", "postgres://lutheran:supersecret@sandbox_postgres/lutherize?sslmode=disable")
	if err != nil {
		return nil, err
	}
	s.db = db
	return s, nil
}

type Scantron struct {
	out []*chpb.PostgresValue
}

func (s *Scantron) Scan(src any) error {
	var val *chpb.PostgresValue

	if src == nil {
		val = &chpb.PostgresValue{
			Clazz: chpb.PostgresClazz_POSTGRES_CLAZZ_NULL,
		}
	} else if x, ok := src.(bool); ok {
		representation := "false"
		if x {
			representation = "true"
		}
		val = &chpb.PostgresValue{
			Clazz:          chpb.PostgresClazz_POSTGRES_CLAZZ_BOOLEAN,
			Representation: representation,
		}
	} else if x, ok := src.(int64); ok {
		val = &chpb.PostgresValue{
			Clazz:          chpb.PostgresClazz_POSTGRES_CLAZZ_INTEGRAL,
			Representation: strconv.FormatInt(x, 10),
		}
	} else if x, ok := src.(float64); ok {
		val = &chpb.PostgresValue{
			Clazz:          chpb.PostgresClazz_POSTGRES_CLAZZ_FLOATING_POINT,
			Representation: strconv.FormatFloat(x, 'g', -1, 64),
		}
	} else if x, ok := src.(string); ok {
		val = &chpb.PostgresValue{
			Clazz:          chpb.PostgresClazz_POSTGRES_CLAZZ_TEXT,
			Representation: x,
		}
	} else if x, ok := src.([]byte); ok {
		val = &chpb.PostgresValue{
			Clazz:          chpb.PostgresClazz_POSTGRES_CLAZZ_BLOB,
			Representation: hex.EncodeToString(x),
		}
	} else {
		return fmt.Errorf("unexpected runtime type of data item")
	}

	s.out = append(s.out, val)

	return nil
}

func (s *PostgresConnector) postgresHelper(ctx context.Context, in *chpb.PostgresRequest) (*chpb.PostgresResponse, error) {
	toAny := (func(v *chpb.PostgresValue) (any, error) {
		switch clazz := v.GetClazz(); clazz {
		case chpb.PostgresClazz_POSTGRES_CLAZZ_NULL:
			if v.GetRepresentation() != "" {
				return nil, fmt.Errorf("bad representation for null (should be empty string)")
			}
			return nil, nil // nil
		case chpb.PostgresClazz_POSTGRES_CLAZZ_BOOLEAN:
			if v.GetRepresentation() == "true" {
				return true, nil
			} else if v.GetRepresentation() == "false" {
				return false, nil
			} else {
				return nil, fmt.Errorf("bad representation for boolean (should be 'true' or 'false')")
			}
		case chpb.PostgresClazz_POSTGRES_CLAZZ_INTEGRAL:
			w, err := strconv.ParseInt(v.GetRepresentation(), 10, 64)
			if err != nil {
				return nil, err
			}
			return w, nil
		case chpb.PostgresClazz_POSTGRES_CLAZZ_FLOATING_POINT:
			w, err := strconv.ParseFloat(v.GetRepresentation(), 64)
			if err != nil {
				return nil, err
			}
			return w, nil
		case chpb.PostgresClazz_POSTGRES_CLAZZ_TEXT:
			return v.GetRepresentation(), nil
		case chpb.PostgresClazz_POSTGRES_CLAZZ_BLOB:
			w, err := hex.DecodeString(v.GetRepresentation())
			if err != nil {
				return nil, err
			}
			return w, nil
		default:
			return nil, fmt.Errorf("unrecognized representation clazz")
		}
	})

	toAnyAll := (func(w []*chpb.PostgresValue) ([]any, error) {
		out := make([]any, 0, len(w))
		for _, v := range w {
			val, err := toAny(v)
			if err != nil {
				return nil, err
			}
			out = append(out, val)
		}
		return out, nil
	})

	arguments, err := toAnyAll(in.GetArguments())
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(in.GetQuery(), arguments...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	out := make([]*chpb.PostgresRow, 0)

	for rows.Next() {
		scantron := &Scantron{out: make([]*chpb.PostgresValue, 0)}

		scantrons := make([]any, len(columns))
		for i := 0; i < len(columns); i++ {
			scantrons[i] = scantron
		}

		err := rows.Scan(scantrons...)
		if err != nil {
			return nil, err
		}

		out = append(out, &chpb.PostgresRow{
			Values: scantron.out,
		})
	}

	return &chpb.PostgresResponse{
		ColumnNames: columns,
		Rows:        out,
		Metadata:    in.GetMetadata(),
	}, nil
}

func (s *PostgresConnector) Handle(ctx context.Context, req *chpb.PostgresRequest) (*chpb.PostgresResponse, error) {
	res, err := s.postgresHelper(ctx, req)
	if err != nil {
		if postgresError, ok := err.(*pq.Error); ok {
			return &chpb.PostgresResponse{
				Error: &chpb.PostgresError{
					ErrorCode:    string(postgresError.Code),
					ErrorMessage: err.Error(),
				},
			}, nil
		}
		return &chpb.PostgresResponse{
			Error: &chpb.PostgresError{
				ErrorCode:    "CONNR", // connector-sourced error
				ErrorMessage: err.Error(),
			},
		}, nil
	}
	return res, nil
}
