package logger

import "testing"

func TestErrorNil(t *testing.T) {
	Error(nil, nil)
}

func TestSuccessNil(t *testing.T) {
	Success(nil, 0)
}

func TestRecoverFuncNil(t *testing.T) {
	RecoverFunc(nil)
}

func TestRecoverFuncCatch(t *testing.T) {
	defer RecoverFunc(nil)
	panic("test")
}