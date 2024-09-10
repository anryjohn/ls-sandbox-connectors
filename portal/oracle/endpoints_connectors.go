package oracle

import (
	"context"

	pb "github.com/luthersystems/sandbox/api/pb/v1"
	"github.com/luthersystems/svc/oracle"
)

func (p *portal) TestTrigger(ctx context.Context, req *pb.TestTriggerRequest) (*pb.TestTriggerResponse, error) {
	return oracle.Call(p.orc, ctx, "test-trigger", req, &pb.TestTriggerResponse{})
}
