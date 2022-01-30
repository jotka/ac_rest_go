package in

type Request struct {
	Value string `json:"value" binding:"required"`
}
