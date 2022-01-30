package in

type State struct {
	DeviceID               string       `json:"deviceId"`
	Name                   string       `json:"name"`
	Label                  string       `json:"label"`
	ManufacturerName       string       `json:"manufacturerName"`
	PresentationID         string       `json:"presentationId"`
	DeviceManufacturerCode string       `json:"deviceManufacturerCode"`
	LocationID             string       `json:"locationId"`
	OwnerID                string       `json:"ownerId"`
	RoomID                 string       `json:"roomId"`
	DeviceTypeName         string       `json:"deviceTypeName"`
	Components             []Components `json:"components"`
	CreateTime             string       `json:"createTime"`
	Type                   string       `json:"type"`
	RestrictionTier        int64        `json:"restrictionTier"`
}
type Capabilities struct {
	ID      string `json:"id"`
	Version int64  `json:"version"`
}
type Categories struct {
	Name         string `json:"name"`
	CategoryType string `json:"categoryType"`
}
type Components struct {
	ID           string         `json:"id"`
	Label        string         `json:"label"`
	Capabilities []Capabilities `json:"capabilities"`
	Categories   []Categories   `json:"categories"`
}
