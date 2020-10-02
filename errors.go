package errors

import "encoding/json"

// HappyCommonError 错误信息
type HappyCommonError struct {
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}

func (h *HappyCommonError) Error() string {
	b, _ := json.Marshal(h)
	return string(b)
}

// NewHappyCommonError 创建一个新的错误
func NewHappyCommonError(code, message, requestID string) error {
	return &HappyCommonError{
		Code:      code,
		Message:   message,
		RequestID: requestID,
	}
}

// GetCode 获取code代码
func (h *HappyCommonError) GetCode() string {
	return h.Code
}

// GetMessage 获取错误信息
func (h *HappyCommonError) GetMessage() string {
	return h.Message
}

// GetRequestID 获取请求特征ID
func (h *HappyCommonError) GetRequestID() string {
	return h.RequestID
}
