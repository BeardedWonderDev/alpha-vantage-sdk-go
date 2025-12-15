package technicalindicators

import (
	"fmt"
	"net/url"

	itypes "github.com/masonJamesWheeler/alpha-vantage-go-wrapper/internal/types"
	"github.com/masonJamesWheeler/alpha-vantage-go-wrapper/types"
)

type TechnicalIndicatorsService struct {
	client itypes.Client
}

func NewTechnicalIndicatorsService(client itypes.Client) *TechnicalIndicatorsService {
	return &TechnicalIndicatorsService{client: client}
}

// SMA retrieves SMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) SMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "SMA"
	return c.getIndicator(params)
}

// EMA retrieves EMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) EMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "EMA"
	return c.getIndicator(params)
}

// WMA retrieves WMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) WMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "WMA"
	return c.getIndicator(params)
}

// DEMA retrieves DEMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) DEMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "DEMA"
	return c.getIndicator(params)
}

// TEMA retrieves TEMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) TEMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "TEMA"
	return c.getIndicator(params)
}

// TRIMA retrieves TRIMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) TRIMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "TRIMA"
	return c.getIndicator(params)
}

// KAMA retrieves KAMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) KAMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "KAMA"
	return c.getIndicator(params)
}

// MAMA retrieves MAMA data based on the provided parameters.
func (c *TechnicalIndicatorsService) MAMA(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MAMA"
	return c.getIndicator(params)
}

// VWAP retrieves VWAP data based on the provided parameters.
func (c *TechnicalIndicatorsService) VWAP(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "VWAP"
	return c.getIndicator(params)
}

// T3 retrieves T3 data based on the provided parameters.
func (c *TechnicalIndicatorsService) T3(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "T3"
	return c.getIndicator(params)
}

// MACD retrieves MACD data based on the provided parameters.
func (c *TechnicalIndicatorsService) MACD(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MACD"
	return c.getIndicator(params)
}

// MACDEXT retrieves MACDEXT data based on the provided parameters.
func (c *TechnicalIndicatorsService) MACDEXT(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MACDEXT"
	return c.getIndicator(params)
}

// STOCH retrieves STOCH data based on the provided parameters.
func (c *TechnicalIndicatorsService) STOCH(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "STOCH"
	return c.getIndicator(params)
}

// STOCHF retrieves STOCHF data based on the provided parameters.
func (c *TechnicalIndicatorsService) STOCHF(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "STOCHF"
	return c.getIndicator(params)
}

// RSI retrieves RSI data based on the provided parameters.
func (c *TechnicalIndicatorsService) RSI(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "RSI"
	return c.getIndicator(params)
}

// STOCHRSI retrieves STOCHRSI data based on the provided parameters.
func (c *TechnicalIndicatorsService) STOCHRSI(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "STOCHRSI"
	return c.getIndicator(params)
}

// WILLR retrieves WILLR data based on the provided parameters.
func (c *TechnicalIndicatorsService) WILLR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "WILLR"
	return c.getIndicator(params)
}

// ADX retrieves ADX data based on the provided parameters.
func (c *TechnicalIndicatorsService) ADX(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ADX"
	return c.getIndicator(params)
}

// ADXR retrieves ADXR data based on the provided parameters.
func (c *TechnicalIndicatorsService) ADXR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ADXR"
	return c.getIndicator(params)
}

// APO retrieves APO data based on the provided parameters.
func (c *TechnicalIndicatorsService) APO(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "APO"
	return c.getIndicator(params)
}

// PPO retrieves PPO data based on the provided parameters.
func (c *TechnicalIndicatorsService) PPO(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "PPO"
	return c.getIndicator(params)
}

// MOM retrieves MOM data based on the provided parameters.
func (c *TechnicalIndicatorsService) MOM(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MOM"
	return c.getIndicator(params)
}

// BOP retrieves BOP data based on the provided parameters.
func (c *TechnicalIndicatorsService) BOP(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "BOP"
	return c.getIndicator(params)
}

// CCI retrieves CCI data based on the provided parameters.
func (c *TechnicalIndicatorsService) CCI(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "CCI"
	return c.getIndicator(params)
}

// CMO retrieves CMO data based on the provided parameters.
func (c *TechnicalIndicatorsService) CMO(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "CMO"
	return c.getIndicator(params)
}

// ROC retrieves ROC data based on the provided parameters.
func (c *TechnicalIndicatorsService) ROC(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ROC"
	return c.getIndicator(params)
}

// ROCR retrieves ROCR data based on the provided parameters.
func (c *TechnicalIndicatorsService) ROCR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ROCR"
	return c.getIndicator(params)
}

// AROON retrieves AROON data based on the provided parameters.
func (c *TechnicalIndicatorsService) AROON(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "AROON"
	return c.getIndicator(params)
}

// AROONOSC retrieves AROONOSC data based on the provided parameters.
func (c *TechnicalIndicatorsService) AROONOSC(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "AROONOSC"
	return c.getIndicator(params)
}

// MFI retrieves MFI data based on the provided parameters.
func (c *TechnicalIndicatorsService) MFI(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MFI"
	return c.getIndicator(params)
}

// TRIX retrieves TRIX data based on the provided parameters.
func (c *TechnicalIndicatorsService) TRIX(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "TRIX"
	return c.getIndicator(params)
}

// ULTOSC retrieves ULTOSC data based on the provided parameters.
func (c *TechnicalIndicatorsService) ULTOSC(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ULTOSC"
	return c.getIndicator(params)
}

// DX retrieves DX data based on the provided parameters.
func (c *TechnicalIndicatorsService) DX(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "DX"
	return c.getIndicator(params)
}

// MINUSDI retrieves MINUSDI data based on the provided parameters.
func (c *TechnicalIndicatorsService) MINUSDI(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MINUS_DI"
	return c.getIndicator(params)
}

// PLUSDI retrieves PLUSDI data based on the provided parameters.
func (c *TechnicalIndicatorsService) PLUSDI(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "PLUS_DI"
	return c.getIndicator(params)
}

// MINUSDM retrieves MINUSDM data based on the provided parameters.
func (c *TechnicalIndicatorsService) MINUSDM(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MINUS_DM"
	return c.getIndicator(params)
}

// PLUSDM retrieves PLUSDM data based on the provided parameters.
func (c *TechnicalIndicatorsService) PLUSDM(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "PLUS_DM"
	return c.getIndicator(params)
}

// BBANDS retrieves BBANDS data based on the provided parameters.
func (c *TechnicalIndicatorsService) BBANDS(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "BBANDS"
	return c.getIndicator(params)
}

// MIDPOINT retrieves MIDPOINT data based on the provided parameters.
func (c *TechnicalIndicatorsService) MIDPOINT(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MIDPOINT"
	return c.getIndicator(params)
}

// MIDPRICE retrieves MIDPRICE data based on the provided parameters.
func (c *TechnicalIndicatorsService) MIDPRICE(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "MIDPRICE"
	return c.getIndicator(params)
}

// SAR retrieves SAR data based on the provided parameters.
func (c *TechnicalIndicatorsService) SAR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "SAR"
	return c.getIndicator(params)
}

// TRANGE retrieves TRANGE data based on the provided parameters.
func (c *TechnicalIndicatorsService) TRANGE(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "TRANGE"
	return c.getIndicator(params)
}

// ATR retrieves ATR data based on the provided parameters.
func (c *TechnicalIndicatorsService) ATR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ATR"
	return c.getIndicator(params)
}

// NATR retrieves NATR data based on the provided parameters.
func (c *TechnicalIndicatorsService) NATR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "NATR"
	return c.getIndicator(params)
}

// AD retrieves AD data based on the provided parameters.
func (c *TechnicalIndicatorsService) AD(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "AD"
	return c.getIndicator(params)
}

// ADOSC retrieves ADOSC data based on the provided parameters.
func (c *TechnicalIndicatorsService) ADOSC(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "ADOSC"
	return c.getIndicator(params)
}

// OBV retrieves OBV data based on the provided parameters.
func (c *TechnicalIndicatorsService) OBV(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "OBV"
	return c.getIndicator(params)
}

// HTTRENDLINE retrieves HT_TRENDLINE data based on the provided parameters.
func (c *TechnicalIndicatorsService) HTTRENDLINE(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "HT_TRENDLINE"
	return c.getIndicator(params)
}

// HTSINE retrieves HT_SINE data based on the provided parameters.
func (c *TechnicalIndicatorsService) HTSINE(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "HT_SINE"
	return c.getIndicator(params)
}

// HTTRENDMODE retrieves HT_TRENDMODE data based on the provided parameters.
func (c *TechnicalIndicatorsService) HTTRENDMODE(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "HT_TRENDMODE"
	return c.getIndicator(params)
}

// HTDCPERIOD retrieves HT_DCPERIOD data based on the provided parameters.
func (c *TechnicalIndicatorsService) HTDCPERIOD(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "HT_DCPERIOD"
	return c.getIndicator(params)
}

// HTDCPHASE retrieves HT_DCPHASE data based on the provided parameters.
func (c *TechnicalIndicatorsService) HTDCPHASE(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "HT_DCPHASE"
	return c.getIndicator(params)
}

// HTPHASOR retrieves HT_PHASOR data based on the provided parameters.
func (c *TechnicalIndicatorsService) HTPHASOR(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	params.Function = "HT_PHASOR"
	return c.getIndicator(params)
}

func (c *TechnicalIndicatorsService) getIndicator(params types.IndicatorParams) (*types.IndicatorResponse, error) {
	queryParams := url.Values{}
	queryParams.Add("symbol", params.Symbol)
	queryParams.Add("interval", params.Interval)
	queryParams.Add("time_period", fmt.Sprintf("%d", params.TimePeriod))
	queryParams.Add("series_type", params.SeriesType)

	if params.Month != "" {
		queryParams.Add("month", params.Month)
	}

	if params.OutputSize != "" {
		queryParams.Add("outputsize", params.OutputSize)
	}

	data, err := c.client.Do(params.Function, queryParams)
	if err != nil {
		return nil, err
	}

	var indicatorResponse types.IndicatorResponse
	if err := types.UnmarshalIndicatorJSON(&indicatorResponse, data, params.Function); err != nil {
		return nil, err
	}

	return &indicatorResponse, nil
}
