(in-package 'sandbox)

(defun camunda-inspect-create-event (request-id process-instance-key &optional wait-for-state)
  (sorted-map
   "oid" request-id
   "msp" "Org1MSP"
   "key" request-id
   "pdc" "private"
   "req" (sorted-map
          "request_id" request-id
          "connector_camunda_inspect" (sorted-map
				       "process_instance_key" process-instance-key
				       "wait_for_state" (or wait-for-state "")))))

(defun camunda-inspect-ok? (resp-body)
  (to-bool (get resp-body "success")))

(defun camunda-inspect-unpack (resp-body)
  (json:load-bytes (hex:decode (get resp-body "content"))))
