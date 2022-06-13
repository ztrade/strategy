package strategy

import (
	"fmt"

	"github.com/ztrade/indicator"
	. "github.com/ztrade/trademodel"
)

type MACD struct {
	engine   Engine
	macd     indicator.CommonIndicator
	position float64

	binSize string
	fast    int
	slow    int
	dea     int
	amount  float64
}

func NewMACD() *MACD {
	s := new(MACD)
	return s
}

func (s *MACD) Param() (paramInfo []Param) {
	paramInfo = []Param{
		StringParam("bin", "Kline binsize", "kline binsize", "15m", &s.binSize),
		IntParam("fast", "macd fast", "macs fast", 12, &s.fast),
		IntParam("slow", "macd slow", "macs slow", 26, &s.slow),
		IntParam("dea", "macd dea", "macs dea", 9, &s.dea),
		FloatParam("amount", "amount", "amount", 1, &s.amount),
	}
	return
}

// OnCandleLarge call when binSize candle
func (s *MACD) OnCandleLarge(candle *Candle) {
	// update macd indicator
	s.macd.Update(candle.Close)
	// get macd indicator: crossDown,crossUp,fast,slow,result
	inds := s.macd.Indicator()
	down := inds["crossDown"]
	up := inds["crossUp"]
	if up == 1 {
		// macd cross up
		if s.position < 0 {
			s.engine.CloseShort(candle.Close, s.amount)
		} else if s.position > 0 {
			return
		}
		s.engine.OpenLong(candle.Close, s.amount)
	} else if down == 1 {
		// macd cross down
		if s.position > 0 {
			s.engine.CloseLong(candle.Close, s.amount)
		} else if s.position < 0 {
			return
		}
		s.engine.OpenShort(candle.Close, s.amount)
	}
}
func (s *MACD) Init(engine Engine, params ParamData) {
	s.engine = engine
	s.macd = engine.AddIndicator("macd", s.fast, s.slow, s.dea)
	engine.Merge("1m", s.binSize, s.OnCandleLarge)
	fmt.Printf("fast: %d, slow: %d, dea: %d, binsize: %s, amount: %f\n", s.fast, s.slow, s.dea, s.binSize, s.amount)
}

// OnCandle call every 1m candle
func (s *MACD) OnCandle(candle *Candle) {
	// fmt.Println(candle)
}

func (s *MACD) OnPosition(pos, price float64) {
	s.position = pos
}

// OnTrade call when own order filled
func (s *MACD) OnTrade(trade *Trade) {
	// fmt.Println("trade:", trade)
}

// OnTradeMarket call when market has new trade
// only called in real trading
func (s *MACD) OnTradeMarket(trade *Trade) {
	// fmt.Println("tradeHistory:", trade)
}

// OnDepth call when order book update
// only called in real trading
func (s *MACD) OnDepth(depth *Depth) {
	// fmt.Println("depth:", depth)
}
