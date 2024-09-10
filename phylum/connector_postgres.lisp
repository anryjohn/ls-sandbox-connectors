(in-package 'sandbox)

(defun postgres-clazz-null ()
  (sorted-map
   "clazz" "POSTGRES_CLAZZ_NULL"
   "representation" ""))

(defun postgres-clazz-boolean (v) ; v is boolean
  (sorted-map
   "clazz" "POSTGRES_CLAZZ_BOOLEAN"
   "representation" (to-string v)))

(defun postgres-clazz-integral (v) ; v is integral
  (sorted-map
   "clazz" "POSTGRES_CLAZZ_INTEGRAL"
   "representation" (to-string v)))

(defun postgres-clazz-floating-point (v) ; v is floating-point
  (sorted-map
   "clazz" "POSTGRES_CLAZZ_FLOATING_POINT"
   "representation" (to-string v)))

(defun postgres-clazz-text (v) ; v is string
  (sorted-map
   "clazz" "POSTGRES_CLAZZ_TEXT"
   "representation" v))

(defun postgres-clazz-blob (v) ; v is bytes
  (sorted-map
   "clazz" "POSTGRES_CLAZZ_BLOB"
   "representation" (to-string (hex:encode v))))

(defun postgres-create-event (request-id query &rest args)
  (sorted-map
   "oid" request-id
   "msp" "Org1MSP"
   "key" request-id
   "pdc" "private"
   "req" (sorted-map
          "request_id" request-id
          "connector_postgres" (sorted-map
                                "query" query
                                "arguments" args))))

(defun postgres-interpret-value (val)
  (let* ([clazz (get val "clazz")]
         [representation (get val "representation")])
    (cond ((string= clazz "POSTGRES_CLAZZ_NULL") ())
          ((string= clazz "POSTGRES_CLAZZ_BOOLEAN")
           (if representation
               (string= representation "true")
               "false"))
          ((string= clazz "POSTGRES_CLAZZ_INTEGRAL")
           (if representation
               (to-int representation)
               0))
          ((string= clazz "POSTGRES_CLAZZ_FLOATING_POINT")
           (if representation
               (to-float representation)
               0))
          ((string= clazz "POSTGRES_CLAZZ_TEXT")
           (if representation
               representation
               ""))
          ((string= clazz "POSTGRES_CLAZZ_BLOB")
           (if representation
               (hex:decode representation)
               ""))
          (:else (error 'postgres-unknown-clazz "unknown representation class")))))

(defun postgres-interpret-row (row)
  (cc:infof (sorted-map "row" row) "CONNECTOR_POSTGRES_INTERPRET_ROW")
  (map 'list #^(postgres-interpret-value %) (get row "values")))
