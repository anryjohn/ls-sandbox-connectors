(in-package 'sandbox)

(defun email-create-event (request-id recipient title body)
  (sorted-map
   "oid" request-id
   "msp" "Org1MSP"
   "key" request-id
   "pdc" "private"
   "req" (sorted-map
          "request_id" request-id
          "connector_email" (sorted-map
			     "recipient" recipient
                             "title" title
                             "body" body))))
