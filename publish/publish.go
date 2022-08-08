package publish

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

type PublishData struct {
	StrategyName string
	Lines        [5]PriceLine
}

func (self *PublishData) GetStrategyName() string {
	return self.StrategyName
}
