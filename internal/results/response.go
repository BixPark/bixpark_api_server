package results

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type ListResponse struct {
	Data       interface{} `json:"data"`
	Message    Message     `json:"message"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	TotalCount int         `json:"totalCount"`
}

type SingleResponse struct {
	Data    interface{} `json:"data"`
	Message Message     `json:"message"`
}
