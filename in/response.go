package in

type SamsungResponse struct {
	RequestID string `json:"requestId"`
	Error     *Error `json:"error"`
}
type Details struct {
	Code    string        `json:"code"`
	Target  string        `json:"target"`
	Message string        `json:"message"`
	Details []interface{} `json:"details"`
}
type Error struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Details []Details `json:"details"`
}
