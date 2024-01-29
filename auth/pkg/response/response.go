package response

import "strconv"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *Response) Error() string {
	resultResponse := `{"status":"` + r.Status + `","message":"` + r.Message + `","code":` + strconv.Itoa(r.Code) + `}`
	return resultResponse
}
