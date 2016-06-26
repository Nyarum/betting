package betting

import "errors"

var (
	EOrderProjection = eOrderProjection{
		OP_ALL:                "ALL",
		OP_EXECUTABLE:         "EXECUTABLE",
		OP_EXECUTION_COMPLETE: "EXECUTION_COMPLETE",
	}
	ErrUnknownOrderProjection = errors.New("Unknown orderProjection value")
)

type eOrderProjection struct {
	OP_ALL                eOrderProjectionInternal
	OP_EXECUTABLE         eOrderProjectionInternal
	OP_EXECUTION_COMPLETE eOrderProjectionInternal
}

type eOrderProjectionInternal string
