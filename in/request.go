package in

type Request struct {
	Value string `form:"value" json:"value" xml:"value"  binding:"required"`
}
