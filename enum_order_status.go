package betting

import (
	"errors"
)

var ErrOrderStatus = errors.New("Unknown orderStatus value")

// OrderStatus
type OrderStatus int

const (
	OS_EXECUTION_COMPLETE OrderStatus = iota
	OS_EXECUTABLE
)

var orderStatusItems = [...]string{
	"EXECUTION_COMPLETE",
	"EXECUTABLE",
}

var orderStatusMap = map[string]OrderStatus{
	orderProjectionItems[OS_EXECUTION_COMPLETE]: OS_EXECUTION_COMPLETE,
	orderProjectionItems[OS_EXECUTABLE]:         OS_EXECUTABLE,
}

func (ostatus OrderStatus) String() string {
	return orderStatusItems[ostatus]
}

func (code *OrderStatus) Check(enum string) error {
	val, ok := orderStatusMap[enum]
	if !ok {
		return ErrOrderStatus
	}

	*code = val

	return nil
}
