(in-package 'sandbox)

(defun camunda-start-create-event (request-id process-id &optional variables)
  (sorted-map
   "oid" request-id
   "msp" "Org1MSP"
   "key" request-id
   "pdc" "private"
   "req" (sorted-map
          "request_id" request-id
          "connector_camunda_start" (sorted-map
				     "process_id" process-id
				     "variables" (to-string (hex:encode (json:dump-bytes (or variables (sorted-map)))))))))
