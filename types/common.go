package types

type Client interface {
	CoreStocks() CoreStocks
	OptionsData() OptionsData
	AlphaInteligence() AlphaInteligence
	FundamentalData() FundamentalData
	Forex() Forex
	Crypto() Crypto
	Commodities() Commodities
	EconomicIndicators() EconomicIndicators
	TechnicalIndicators() TechnicalIndicators
}

type CoreStocks interface {
	Intraday(params TimeSeriesParams) (TimeSeriesIntraday, error)
	Daily(params TimeSeriesParams) (TimeSeriesDaily, error)
	DailyAdjusted(params TimeSeriesParams) (TimeSeriesDailyAdjusted, error)
	Weekly(params TimeSeriesParams) (TimeSeriesWeekly, error)
	WeeklyAdjusted(params TimeSeriesParams) (TimeSeriesWeekly, error)
	Monthly(params TimeSeriesParams) (TimeSeriesMonthly, error)
	MonthlyAdjusted(params TimeSeriesParams) (TimeSeriesMonthlyAdjusted, error)
	SymbolSearch(keywords string) (*SymbolSearchResponse, error)
}

type OptionsData interface {
}

type AlphaInteligence interface {
}

type FundamentalData interface {
}

type Forex interface {
}

type Crypto interface {
}

type Commodities interface {
}

type EconomicIndicators interface {
}

type TechnicalIndicators interface {
	ADX(params IndicatorParams) (*IndicatorResponse, error)
	ADXR(params IndicatorParams) (*IndicatorResponse, error)
	APO(params IndicatorParams) (*IndicatorResponse, error)
	AROON(params IndicatorParams) (*IndicatorResponse, error)
	AROONOSC(params IndicatorParams) (*IndicatorResponse, error)
	BOP(params IndicatorParams) (*IndicatorResponse, error)
	CCI(params IndicatorParams) (*IndicatorResponse, error)
	CMO(params IndicatorParams) (*IndicatorResponse, error)
	DX(params IndicatorParams) (*IndicatorResponse, error)
	EMA(params IndicatorParams) (*IndicatorResponse, error)
	HTDCPHASE(params IndicatorParams) (*IndicatorResponse, error)
	HTDCPERIOD(params IndicatorParams) (*IndicatorResponse, error)
	HTPHASOR(params IndicatorParams) (*IndicatorResponse, error)
	HTSINE(params IndicatorParams) (*IndicatorResponse, error)
	HTTRENDLINE(params IndicatorParams) (*IndicatorResponse, error)
	HTTRENDMODE(params IndicatorParams) (*IndicatorResponse, error)
	OBV(params IndicatorParams) (*IndicatorResponse, error)
	ADOSC(params IndicatorParams) (*IndicatorResponse, error)
}
