package response

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
