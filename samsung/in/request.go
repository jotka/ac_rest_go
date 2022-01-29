package in

type req struct {
	value string `form:"value" json:"value" xml:"value"  binding:"required"`
}
