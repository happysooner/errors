package errors

import "encoding/json"

type HappyCommonError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

func (h *HappyCommonError) Error() string {
	b, _ := json.Marshal(h)
	return string(b)
}

func NewHappyCommonError(code, message, requestId string) error {
	return &HappyCommonError{
		Code:      code,
		Message:   message,
		RequestId: requestId,
	}
}

func (h *HappyCommonError) GetCode() string {
	return h.Code
}

func (h *HappyCommonError) GetMessage() string {
	return h.Message
}

func (h *HappyCommonError) GetRequestId() string {
	return h.RequestId
}
