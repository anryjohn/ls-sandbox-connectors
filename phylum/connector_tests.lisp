(defun without-last (l)
  (reverse 'list (cdr (reverse 'list l))))

;; mk-postgres-test implements the state machine
(defun mk-postgres-test (state)
  ;(cc:infof (sorted-map "state" state) "mk-postgres-test")
  (labels
    ([handle (resp)
       (let* ([next-uuid (mk-uuid)]
              [next-state (assoc state "test_id" next-uuid)]
	      [resp-body (get resp "response")]
              [resp-err (get resp-body "error")])
	 (statedb:put "connector_test_flipper" next-uuid)
         (if resp-err
             (progn
               (cc:errorf resp-err "response error"))
             (progn
               (cc:infof resp "got connector resp")
               (cc:infof state "got connector state")
               (if (= (get state "step") 1)
                   (sorted-map
                    "del" false
                    "put" (assoc next-state "step" 2)
                    "events" (list (postgres-create-event next-uuid "SELECT * FROM pg_catalog.pg_tables;")))
                   (if (= (get state "step") 2)
                       (progn
                         (cc:infof (sorted-map "foo" (string:join (keys resp-body) ",")) "DEBUG")
                         (sorted-map
                          "del" false
                          "put" (assoc (assoc (assoc next-state "step" 3) "resp_cached" resp) "seen" (* (length (get resp-body "column_names")) (length (get resp-body "rows"))))
                          "events" (list (postgres-create-event next-uuid "CREATE TABLE IF NOT EXISTS summary (stas varchar(80) PRIMARY KEY, prefix varchar(80), seen integer NOT NULL, x_a integer, x_b boolean NOT NULL, x_c integer NOT NULL, x_d double precision NOT NULL, x_e text, x_f bytea);"))))
                       (if (= (get state "step") 3)
                           (sorted-map
                            "del" false
                            "put" (assoc next-state "step" 4)
                            "events" (list (postgres-create-event next-uuid
                                                                  "INSERT INTO summary VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
                                                                  (postgres-clazz-text (cc:timestamp (cc:now)))
                                                                  (postgres-clazz-text (get state "prefix"))
                                                                  (postgres-clazz-integral (get state "seen"))
								  (postgres-clazz-null)
								  (postgres-clazz-boolean true)
								  (postgres-clazz-integral 7)
								  (postgres-clazz-floating-point 7.89)
								  (postgres-clazz-text "HELLO_WORLD_ONE")
								  (postgres-clazz-blob (to-bytes "HELLO_WORLD_TWO")))))
			   (if (= (get state "step") 4)
                               (sorted-map
				"del" false
				"put" (assoc next-state "step" 5)
				"events" (list (postgres-create-event next-uuid "SELECT * FROM summary LIMIT 1")))
			       (let* ([row (cdr (cdr (cdr (postgres-interpret-row (first (get resp-body "rows"))))))])
				 ;; need to peel of the bytes value here and special case its comparison as it is not supported by deep-equal?
				 (if (and (deep-equal? (without-last row) (list '() true 7 7.89 "HELLO_WORLD_ONE"))
					  (string= (to-string (nth row (- (length row) 1))) "HELLO_WORLD_TWO"))
				     (cc:infof (sorted-map) "PASS")
				     (cc:infof (sorted-map) "FAIL"))
				 (sorted-map
				  "del" false
				  "put" ()
				  "events" ())))))))))])
    (lambda (op &rest args)
        (cond ((equal? op 'handle) (apply handle args))
              (:else (error 'unknown-operation op))))))

;; mk-postgres-test-storage implements the "storage" aspect
(defun mk-postgres-test-storage ()
  (labels
    ([name () "postgres-test"]

     [mk-test-storage-key (test-id)
       (join-index-cols "sandbox" "postgres-test" test-id)]

     [storage-put-test (test)
       (cc:infof (sorted-map "test" test) "PUTTING")
       (sidedb:put (mk-test-storage-key (get test "test_id")) test)]

     [storage-get-test (test-id)
       (cc:infof (sorted-map "test-id" test-id) "GETTING")
       (thread-first
         (mk-test-storage-key test-id)
         (sidedb:get)
         (mk-postgres-test))]

     [storage-del-test (test-id)
       (sidedb:del (mk-test-storage-key test-id))]

     [init (prefix)
       (let* ([test-id (mk-uuid)])
         (cc:infof (sorted-map "test-id" test-id) "INIT")
         (sorted-map
	  "del" false
	  "put" (sorted-map "test_id" test-id "step" 2 "prefix" prefix)
	  "events" (list (postgres-create-event test-id "SELECT * FROM pg_catalog.pg_tables LIMIT 1;"))))])

    (lambda (op &rest args)
        (cond ((equal? op 'name) (apply name args))
              ((equal? op 'init) (apply init args))
              ((equal? op 'get) (apply storage-get-test args))
              ((equal? op 'del) (apply storage-del-test args))
              ((equal? op 'put) (apply storage-put-test args))
              (:else (error 'unknown-operation op))))))

(set 'postgres-test (singleton mk-postgres-test-storage))

(register-connector-factory postgres-test)

;; mk-email-test implements the state machine
(defun mk-email-test (state)
  (labels
    ([handle (resp)
       (let* ([next-uuid (mk-uuid)]
              [next-state (assoc state "test_id" next-uuid)]
	      [resp-body (get resp "response")]
              [resp-err (get resp-body "error")])
         (if resp-err
             (progn
               (cc:errorf resp-err "response error"))
             (progn
               (if (= (get state "step") 2)
		 (progn
		   (cc:infof (sorted-map) "EMAIL: GOT TO STEP 2")
                   (sorted-map
                    "del" false
                    "put" ()
                    "events" ()))
		 (error 'unexpected "unexpected")))))])
    (lambda (op &rest args)
        (cond ((equal? op 'handle) (apply handle args))
              (:else (error 'unknown-operation op))))))

;; mk-email-test-storage implements the "storage" aspect
(defun mk-email-test-storage ()
  (labels
    ([name () "email-test"]

     [mk-test-storage-key (test-id)
       (join-index-cols "sandbox" "email-test" test-id)]

     [storage-put-test (test)
       (sidedb:put (mk-test-storage-key (get test "test_id")) test)]

     [storage-get-test (test-id)
       (thread-first
         (mk-test-storage-key test-id)
         (sidedb:get)
         (mk-email-test))]

     [storage-del-test (test-id)
       (sidedb:del (mk-test-storage-key test-id))]

     [init (recipient)
       (let* ([test-id (mk-uuid)])
         (sorted-map
	  "del" false
	  "put" (sorted-map "test_id" test-id "step" 2)
	  "events" (list (email-create-event test-id recipient "ABC Company Newsletter" "Dear friend, we at ABC Company are committed to the future."))))])

    (lambda (op &rest args)
        (cond ((equal? op 'name) (apply name args))
              ((equal? op 'init) (apply init args))
              ((equal? op 'get) (apply storage-get-test args))
              ((equal? op 'del) (apply storage-del-test args))
              ((equal? op 'put) (apply storage-put-test args))
              (:else (error 'unknown-operation op))))))

(set 'email-test (singleton mk-email-test-storage))

(register-connector-factory email-test)

;; mk-camunda-test implements the state machine
(defun mk-camunda-test (state)
  (labels
    ([handle (resp)
       (let* ([next-uuid (mk-uuid)]
              [next-state (assoc state "test_id" next-uuid)]
	      [resp-body (get resp "response")]
              [resp-err (get resp-body "error")])
         (if resp-err
             (progn
               (cc:errorf resp-err "response error"))
             (progn
	       (cond ((= (get state "step") 2)
		      (let* ([ikey (get resp-body "process_instance_key")])
			(cc:infof (sorted-map "ikey" ikey) "CAMUNDA_STEP_2")
			(sorted-map
			 "del" false
			 "put" (assoc (assoc next-state "step" 3) "ikey" ikey)
			 "events" (list (camunda-inspect-create-event next-uuid ikey "COMPLETED")))))
		     ((= (get state "step") 3)
		      (if (and (camunda-inspect-ok? resp-body) (string= (get (camunda-inspect-unpack resp-body) "state") "COMPLETED"))
			  (progn
			    (cc:infof (sorted-map) "CAMUNDA_COMPLETED")
			    (sorted-map
			     "del" false
			     "put" ()
			     "events" ()))
			  (progn
			    (cc:infof (sorted-map) "CAMUNDA_SPIN")
			    (sorted-map
			     "del" false
			     "put" next-state
			     "events" (list (camunda-inspect-create-event next-uuid (get state "ikey") "COMPLETED"))))))
		     (:else (error 'unexpected "unexpected"))))))])
    (lambda (op &rest args)
        (cond ((equal? op 'handle) (apply handle args))
              (:else (error 'unknown-operation op))))))

;; mk-camunda-test-storage implements the "storage" aspect
(defun mk-camunda-test-storage ()
  (labels
    ([name () "camunda-test"]

     [mk-test-storage-key (test-id)
       (join-index-cols "sandbox" "camunda-test" test-id)]

     [storage-put-test (test)
       (sidedb:put (mk-test-storage-key (get test "test_id")) test)]

     [storage-get-test (test-id)
       (thread-first
         (mk-test-storage-key test-id)
         (sidedb:get)
         (mk-camunda-test))]

     [storage-del-test (test-id)
       (sidedb:del (mk-test-storage-key test-id))]

     [init (process_id)
       (let* ([test-id (mk-uuid)])
         (sorted-map
	  "del" false
	  "put" (sorted-map "test_id" test-id "step" 2)
	  "events" (list (camunda-start-create-event test-id process_id))))])

    (lambda (op &rest args)
        (cond ((equal? op 'name) (apply name args))
              ((equal? op 'init) (apply init args))
              ((equal? op 'get) (apply storage-get-test args))
              ((equal? op 'del) (apply storage-del-test args))
              ((equal? op 'put) (apply storage-put-test args))
              (:else (error 'unknown-operation op))))))

(set 'camunda-test (singleton mk-camunda-test-storage))

(register-connector-factory camunda-test)

(defendpoint "test-trigger" (req)
  (let* ([kind (get req "kind")])
    (cond ((string= kind "postgres1") (do-transition postgres-test (postgres-test 'init "apple_sauce")))
	  ((string= kind "email1") (do-transition email-test (email-test 'init "demo@example.com")))
	  ((string= kind "email2") (do-transition email-test (email-test 'init "demo+other@example.com")))
	  ((string= kind "camunda1") (do-transition camunda-test (camunda-test 'init "c8-sdk-demo"))))
    (route-success '())))
