package eigenda

type BaseResponse struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func BaseResource(status bool, code int, data interface{}, msg string) (baseRep *BaseResponse) {
	baseRep = &BaseResponse{Status: status, Code: code, Data: data, Msg: msg}
	return
}
