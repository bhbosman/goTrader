package trackMarket

import (
	"fmt"
	"github.com/bhbosman/goMessages/marketData/stream"
)

type PegToPrice struct {
	strategyName     string
	instrument       string
	Side             string
	PegPrice         float64
	Volume           *float64
	Consideration    *float64
	TradingInterface string
}

func (self *PegToPrice) StrategyName() string {
	return self.strategyName
}

func (self *PegToPrice) Instruments() []string {
	return []string{self.instrument}
}

func NewPegToPrice(
	strategyName string,
	instrument string,
	tradingInterface string,
	side string,
	pegPrice float64,
	volume *float64,
	consideration *float64,
) IPricingVolumeCalculation {
	return &PegToPrice{
		strategyName:     strategyName,
		instrument:       instrument,
		Side:             side,
		PegPrice:         pegPrice,
		Volume:           volume,
		Consideration:    consideration,
		TradingInterface: tradingInterface,
	}
}

func (self *PegToPrice) Calculate(activeDataMap map[string]*stream.PublishTop5) ([]*PriceVolumeResponse, error) {
	var volume *float64 = nil
	if self.Volume != nil {
		volume = self.Volume
	} else if self.Consideration != nil {
		v := *self.Consideration / self.PegPrice
		volume = &v
	}
	if volume == nil {
		return nil, fmt.Errorf("invalid values for volume calculation")
	}

	return []*PriceVolumeResponse{
			{
				TradingInterface: self.TradingInterface,
				OrderId:          1,
				Side:             self.Side,
				Price:            self.PegPrice,
				Volume:           *volume,
			},
		},
		nil
}
