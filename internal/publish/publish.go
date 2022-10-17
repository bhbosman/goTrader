package publish

import (
	"github.com/bhbosman/goCommsDefinitions"
	"time"
)

type IStrategy interface {
	GetStrategyName() string
}

type PricePoint struct {
	Price  float64
	Volume float64
}
type PriceLine struct {
	Bid PricePoint
	Ask PricePoint
}

type MarketData struct {
	Lines [5]PriceLine
}

type PublishData struct {
	Date         time.Time
	StrategyName string
	State        string
	MarketData   []*MarketData
	Actions      []string
}

func (self *PublishData) GetStrategyName() string {
	return self.StrategyName
}

type DeleteStrategy struct {
	StrategyName string
}

type StrategyList struct {
	Strategies []string
}

type RequestStrategyList struct {
	Callback goCommsDefinitions.IAdder
}

type RequestStrategyItem struct {
	Item     string
	Callback goCommsDefinitions.IAdder
}
