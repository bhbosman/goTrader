package publish

import "github.com/bhbosman/goCommsDefinitions"

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
	StrategyName string
	State        string
	MarketData   []*MarketData
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
