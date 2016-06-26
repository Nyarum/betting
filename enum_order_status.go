package betting

import (
	"errors"
)

var (
	EOrderStatus = eOrderStatus{
		OS_EXECUTION_COMPLETE: "EXECUTION_COMPLETE",
		OS_EXECUTABLE:         "EXECUTABLE",
	}
	ErrOrderStatus = errors.New("Unknown orderStatus value")
)

type eOrderStatus struct {
	OS_EXECUTION_COMPLETE eOrderStatusInternal
	OS_EXECUTABLE         eOrderStatusInternal
}

type eOrderStatusInternal string
