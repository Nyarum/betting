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

func (ostatus OrderStatus) String() string {
	return orderStatusItems[ostatus]
}

func (code *OrderStatus) UnmarshalJSON(buf []byte) error {
	var err error
	switch {
	case string(buf) == orderStatusItems[OS_EXECUTION_COMPLETE]:
		*code = OS_EXECUTION_COMPLETE
	case string(buf) == orderStatusItems[OS_EXECUTABLE]:
		*code = OS_EXECUTABLE
	default:
		err = ErrOrderStatus
	}
	return err
}
