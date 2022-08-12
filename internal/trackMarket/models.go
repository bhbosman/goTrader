package trackMarket

type Strategies struct {
	modelSettings []modelSettings
	PegToPrice    []*PegToPrice
}

type modelSettings struct {
	Name       string
	instrument string
}

type callbackMessage struct {
	cb func()
}
