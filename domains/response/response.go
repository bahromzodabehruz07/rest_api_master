package response

type Response struct {
	Status  int64       `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}
