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
	Quote(symbol string) (Quote, error)
	SymbolSearch(keywords string) (*SymbolSearchResponse, error)
}

type OptionsData interface {
}

type AlphaInteligence interface {
	AnalyticsFixedWindow(params AnalyticsFixedWindowParams) (*AnalyticsFixedWindowResponse, error)
	AnalyticsSlidingWindow(params AnalyticsSlidingWindowParams) (*AnalyticsSlidingWindowResponse, error)
}

type FundamentalData interface {
	CompanyOverview(symbol string) (*CompanyOverviewResponse, error)
	IncomeStatement(symbol string) (*IncomeStatementResponse, error)
	BalanceSheet(symbol string) (*BalanceSheetResponse, error)
	CashFlow(symbol string) (*CashFlowResponse, error)
	ETFProfile(symbol string) (*ETFProfile, error)
	Dividends(symbol string) (*DividendsResponse, error)
	Splits(symbol string) (*SplitsResponse, error)
}

type Forex interface {
	ExchangeRate(params ForexExchangeRateParams) (*CurrencyExchangeRateResponse, error)
}

type Crypto interface {
	ExchangeRate(params CryptoExchangeRateParams) (*CurrencyExchangeRateResponse, error)
	Intraday(params CryptoIntradayParams) (*CryptoSeriesResponse, error)
	Daily(params CryptoDailyParams) (*CryptoSeriesResponse, error)
	Weekly(params CryptoWeeklyParams) (*CryptoSeriesResponse, error)
	Monthly(params CryptoMonthlyParams) (*CryptoSeriesResponse, error)
}

type Commodities interface {
}

type EconomicIndicators interface {
}

type TechnicalIndicators interface {
	SMA(params IndicatorParams) (*IndicatorResponse, error)
	WMA(params IndicatorParams) (*IndicatorResponse, error)
	DEMA(params IndicatorParams) (*IndicatorResponse, error)
	TEMA(params IndicatorParams) (*IndicatorResponse, error)
	TRIMA(params IndicatorParams) (*IndicatorResponse, error)
	KAMA(params IndicatorParams) (*IndicatorResponse, error)
	MAMA(params IndicatorParams) (*IndicatorResponse, error)
	VWAP(params IndicatorParams) (*IndicatorResponse, error)
	T3(params IndicatorParams) (*IndicatorResponse, error)
	MACD(params IndicatorParams) (*IndicatorResponse, error)
	MACDEXT(params IndicatorParams) (*IndicatorResponse, error)
	STOCH(params IndicatorParams) (*IndicatorResponse, error)
	STOCHF(params IndicatorParams) (*IndicatorResponse, error)
	RSI(params IndicatorParams) (*IndicatorResponse, error)
	STOCHRSI(params IndicatorParams) (*IndicatorResponse, error)
	WILLR(params IndicatorParams) (*IndicatorResponse, error)
	ADX(params IndicatorParams) (*IndicatorResponse, error)
	ADXR(params IndicatorParams) (*IndicatorResponse, error)
	APO(params IndicatorParams) (*IndicatorResponse, error)
	AROON(params IndicatorParams) (*IndicatorResponse, error)
	AROONOSC(params IndicatorParams) (*IndicatorResponse, error)
	BOP(params IndicatorParams) (*IndicatorResponse, error)
	CCI(params IndicatorParams) (*IndicatorResponse, error)
	CMO(params IndicatorParams) (*IndicatorResponse, error)
	ROC(params IndicatorParams) (*IndicatorResponse, error)
	ROCR(params IndicatorParams) (*IndicatorResponse, error)
	MFI(params IndicatorParams) (*IndicatorResponse, error)
	TRIX(params IndicatorParams) (*IndicatorResponse, error)
	ULTOSC(params IndicatorParams) (*IndicatorResponse, error)
	DX(params IndicatorParams) (*IndicatorResponse, error)
	MINUSDI(params IndicatorParams) (*IndicatorResponse, error)
	PLUSDI(params IndicatorParams) (*IndicatorResponse, error)
	MINUSDM(params IndicatorParams) (*IndicatorResponse, error)
	PLUSDM(params IndicatorParams) (*IndicatorResponse, error)
	PPO(params IndicatorParams) (*IndicatorResponse, error)
	MOM(params IndicatorParams) (*IndicatorResponse, error)
	BBANDS(params IndicatorParams) (*IndicatorResponse, error)
	MIDPOINT(params IndicatorParams) (*IndicatorResponse, error)
	MIDPRICE(params IndicatorParams) (*IndicatorResponse, error)
	SAR(params IndicatorParams) (*IndicatorResponse, error)
	TRANGE(params IndicatorParams) (*IndicatorResponse, error)
	ATR(params IndicatorParams) (*IndicatorResponse, error)
	NATR(params IndicatorParams) (*IndicatorResponse, error)
	AD(params IndicatorParams) (*IndicatorResponse, error)
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
