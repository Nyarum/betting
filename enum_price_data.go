package betting

import "errors"

var ErrUnknownPriceData = errors.New("Unknown priceData value")

type PriceData int
type PriceDatas []PriceData

const (
	PD_SP_AVAILABLE   PriceData = iota // Amount available for the BSP auction.
	PD_SP_TRADED                       // Amount traded in the BSP auction.
	PD_EX_BEST_OFFERS                  // Only the best prices available for each runner, to requested price depth.
	PD_EX_ALL_OFFERS                   // trumps EX_BEST_OFFERS if both settings are present
	PD_EX_TRADED                       // Amount traded on the exchange.
)

var priceDataItems = [...]string{
	"SP_AVAILABLE",
	"SP_TRADED",
	"EX_BEST_OFFERS",
	"EX_ALL_OFFERS",
	"EX_TRADED",
}

var priceDataMap = map[string]PriceData{
	priceDataItems[PD_SP_AVAILABLE]:   PD_SP_AVAILABLE,
	priceDataItems[PD_SP_TRADED]:      PD_SP_TRADED,
	priceDataItems[PD_EX_BEST_OFFERS]: PD_EX_BEST_OFFERS,
	priceDataItems[PD_EX_ALL_OFFERS]:  PD_EX_ALL_OFFERS,
	priceDataItems[PD_EX_TRADED]:      PD_EX_TRADED,
}

func (pd PriceData) String() string {
	return priceDataItems[pd]
}

func (code *PriceData) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := priceDataMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownPriceData
	}
	return err

}
