package error

type Error struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
