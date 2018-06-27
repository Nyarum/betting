package betting

import (
	"time"
	"fmt"
)


type Decimal float64

func (n Decimal) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("%.2f", n)), nil
}

type PlaceInstruction struct {
	OrderType          EOrderType         `json:"orderType,omitempty"`
	SelectionID        int64              `json:"selectionId,omitempty"`
	Handicap           Decimal            `json:"handicap,omitempty"`
	Side               ESide              `json:"side,omitempty"`
	LimitOrder         *LimitOrder         `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   string             `json:"customerOrderRef,omitempty"`
}

type PlaceExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus    `json:"status,omitempty"`
	ErrorCode          EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []PlaceInstructionReport  `json:"instructionReports,omitempty"`
}

type PlaceInstructionReport struct {
	Status              EInstructionReportStatus    `json:"status,omitempty"`
	ErrorCode           EInstructionReportErrorCode `json:"errorCode,omitempty"`
	OrderStatus         EOrderStatus                `json:"orderStatus,omitempty"`
	Instruction         PlaceInstruction            `json:"instruction,omitempty"`
	BetID               string                      `json:"betId,omitempty"`
	PlacedDate          time.Time                   `json:"placedDate,omitempty"`
	AveragePriceMatched float64                     `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64                     `json:"sizeMatched,omitempty"`
}

type LimitOrder struct {
	Size            Decimal          `json:"size,omitempty"`
	Price           Decimal          `json:"price,omitempty"`
	PersistenceType EPersistenceType `json:"persistenceType,omitempty"`
	TimeInForce     ETimeInForce     `json:"timeInForce,omitempty"`
	MinFillSize     Decimal          `json:"minFillSize,omitempty"`
	BetTargetType   EBetTargetType   `json:"betTargetType,omitempty"`
	BetTargetSize   Decimal          `json:"betTargetSize,omitempty"`
}

type LimitOnCloseOrder struct {
	Liability Decimal `json:"liability,omitempty"`
	Price     Decimal `json:"price,omitempty"`
}

type MarketOnCloseOrder struct {
	Liability Decimal `json:"liability,omitempty"`
}

// PlaceOrders to place new orders into market.
func (b *Betting) PlaceOrders(filter Filter) (placeExecutionReport PlaceExecutionReport, err error) {
	err = b.Request(&placeExecutionReport, BettingURL, "placeOrders", &filter)
	if err != nil {
		return
	}

	return
}


type	CancelInstruction	struct	{
	BetID          string         `json:"betId"`
	SizeReduction  Decimal        `json:"sizeReduction,omitempty"`
}

type CancelExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus    `json:"status,omitempty"`
	ErrorCode          EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []CancelInstructionReport  `json:"instructionReports,omitempty"`
}

type CancelInstructionReport struct {
	Status              EInstructionReportStatus    `json:"status,omitempty"`
	ErrorCode           EInstructionReportErrorCode `json:"errorCode,omitempty"`
	Instruction         CancelInstruction            `json:"instruction,omitempty"`
	CancelledDate       time.Time                   `json:"cancelledDate,omitempty"`
	SizeCancelled       float64                     `json:"sizeCancelled,omitempty"`
}

// CancelOrders to place new orders into market.
func (b *Betting) CancelOrders(filter CancelFilter) (cancelExecutionReport CancelExecutionReport, err error) {
	err = b.Request(&cancelExecutionReport, BettingURL, "cancelOrders", &filter)
	if err != nil {
		return
	}

	return
}
