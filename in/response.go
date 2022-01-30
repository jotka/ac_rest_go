package in

type SamsungResponse struct {
	Results []Results `json:"results"`
}
type Results struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
