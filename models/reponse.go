package models

// ResponseData 返回JSON模型
type ResponseData struct {
	Success bool        `json:"result"`
	Msg     string      `json:"desc"`
	Data    interface{} `json:"data"`
}
