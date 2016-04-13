package betting

import "errors"

var ErrUnknownOrderProjection = errors.New("Unknown orderProjection value")

type OrderProjection int

const (
	OP_ALL                OrderProjection = iota // 	EXECUTABLE and EXECUTION_COMPLETE orders
	OP_EXECUTABLE                                // 	An order that has a remaining unmatched portion
	OP_EXECUTION_COMPLETE                        // 	An order that does not have any remaining unmatched portion
)

var orderProjectionItems = [...]string{
	"ALL",
	"EXECUTABLE",
	"EXECUTION_COMPLETE",
}

var orderProjectionMap = map[string]OrderProjection{
	orderProjectionItems[OP_ALL]:                OP_ALL,
	orderProjectionItems[OP_EXECUTABLE]:         OP_EXECUTABLE,
	orderProjectionItems[OP_EXECUTION_COMPLETE]: OP_EXECUTION_COMPLETE,
}

func (op OrderProjection) String() string {
	return orderProjectionItems[op]
}

func (code *OrderProjection) Check(enum string) error {
	val, ok := orderProjectionMap[enum]
	if !ok {
		return ErrUnknownOrderProjection
	}

	*code = val

	return nil
}
