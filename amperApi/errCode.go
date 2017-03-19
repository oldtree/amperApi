package amperApi

import "fmt"

type StatuCode struct {
	Status int
	Desc   string
}

func NewStatuCode(code int, desc string) *StatuCode {
	return &StatuCode{Status: code, Desc: desc}
}

func (e StatuCode) String() string {
	return fmt.Sprintf("status %d : desc %s", e.Status, e.Desc)
}

var (
	StatusUnKnown              = NewStatuCode(-1, "Status Unknown")
	StatusOk                   = NewStatuCode(200, "Request successful. Object found and returned.")
	StatusCreated              = NewStatuCode(201, "Object successfully created.The message body contains its current state. If present, the location header lists the URL for the new object.")
	StatusBadRequest           = NewStatuCode(400, "System unable to process the request because it was malformed.")
	StatusUnauthorized         = NewStatuCode(401, "Request unsuccessful because of authentication failure. API token not valid.")
	StatusNotFound             = NewStatuCode(404, "A path or ID specified in the URL does not exist. Also returned in cases where an object does exist in the system, but it does not belong to the user account that was authenticated.")
	StatusMethodNotAllowed     = NewStatuCode(405, "Requested method cannot be performed on the specified object.")
	StatusConflict             = NewStatuCode(409, "Object could not be created because it conflicts with an existing object. Ensure that any user-defined project name is unique for that user account.")
	StatusUnsupportedMediaType = NewStatuCode(415, "Content-Type is missing or set to an unsupported value.Ensure that Content-Type is application/json.")
	StatusUnprocessableEntity  = NewStatuCode(422, "POST request was parsable, but failed validation check. A required property was either missing or a value was invalid.")
	StatusTooManyRequest       = NewStatuCode(429, "Request rejected because too many similar requests have been submitted in a short amount of time. Wait and try submitting requests later.")
)
