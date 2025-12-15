package fundamentaldata

import itypes "github.com/BeardedWonderDev/alpha-vantage-sdk-go/internal/types"

type FundamentalDataService struct {
	client itypes.Client
}

func NewFundamentalDataService(client itypes.Client) *FundamentalDataService {
	return &FundamentalDataService{
		client: client,
	}
}
