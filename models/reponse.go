package models

// ResponseData 返回JSON模型
type ResponseData struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
