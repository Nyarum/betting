package betting

import "errors"

var ErrUnknownPersistenceType = errors.New("Unknown persistenceType value")

type PersistenceType int

const (
	PT_LAPSE           PersistenceType = iota // Lapse the order at turn-in-play
	PT_PERSIST                                // Persist the order to in-play. The bet will be place automatically into the in-play market at the start of the event.
	PT_MARKET_ON_CLOSE                        // Put the order into the auction (SP) at turn-in-play
)

var persistenceTypeItems = [...]string{
	"LAPSE",
	"PERSIST",
	"MARKET_ON_CLOSE",
}

var persistenceTypeMap = map[string]PersistenceType{
	persistenceTypeItems[PT_LAPSE]:           PT_LAPSE,
	persistenceTypeItems[PT_PERSIST]:         PT_PERSIST,
	persistenceTypeItems[PT_MARKET_ON_CLOSE]: PT_MARKET_ON_CLOSE,
}

func (code PersistenceType) String() string {
	return persistenceTypeItems[code]
}

func (code *PersistenceType) Check(enum string) error {
	val, ok := persistenceTypeMap[enum]
	if !ok {
		return ErrUnknownPersistenceType
	}

	*code = val

	return nil
}
