package errors

import (
	"testing"
)

var e = HappyCommonError{
	Code:      "500",
	Message:   "test",
	RequestID: "testReq",
}

func TestNewHappyCommonError(t *testing.T) {
	r := NewHappyCommonError(e.Code, e.Message, e.RequestID)
	str := r.Error()
	if str != `{"code":"500","message":"test","requestId":"testReq"}` {
		t.Error("TestNewHappyCommonError failed")
	}
}

func TestHappyCommonError_GetCode(t *testing.T) {
	if e.GetCode() != "500" {
		t.Error("TestHappyCommonError_GetCode failed")
	}
}

func TestHappyCommonError_GetMessage(t *testing.T) {
	if e.GetMessage() != "test" {
		t.Error("TestHappyCommonError_GetMessage failed")
	}
}

func TestHappyCommonError_GetRequestId(t *testing.T) {
	if e.GetRequestID() != "testReq" {
		t.Error("TestHappyCommonError_GetRequestId failed")
	}
}
