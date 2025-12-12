package public

import "github.com/starchou/go-okx/rest/api"

func NewGetInstruments(param *GetInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/public/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetInstrumentsResponse{}
}

type GetInstrumentsParam struct {
	api.PagingParam
	InstType   string `url:"instType,omitempty"`
	InstFamily string `url:"instFamily,omitempty"`
	InstId     string `url:"instId,omitempty"`
}

type GetInstrumentsResponse struct {
	api.Response
	Data []Instrument `json:"data"`
}

type Instrument struct {
	Alias             string   `json:"alias"`
	AuctionEndTime    string   `json:"auctionEndTime"`
	BaseCcy           string   `json:"baseCcy"`
	Category          string   `json:"category"`
	CtMult            string   `json:"ctMult"`
	CtType            string   `json:"ctType"`
	CtVal             string   `json:"ctVal"`
	CtValCcy          string   `json:"ctValCcy"`
	ContTdSwTime      string   `json:"contTdSwTime"`
	ExpTime           string   `json:"expTime"`
	FutureSettlement  bool     `json:"futureSettlement"`
	GroupID           string   `json:"groupId"`
	InstFamily        string   `json:"instFamily"`
	InstID            string   `json:"instId"`
	InstType          string   `json:"instType"`
	Lever             string   `json:"lever"`
	ListTime          string   `json:"listTime"`
	LotSz             string   `json:"lotSz"`
	MaxIcebergSz      string   `json:"maxIcebergSz"`
	MaxLmtAmt         string   `json:"maxLmtAmt"`
	MaxLmtSz          string   `json:"maxLmtSz"`
	MaxMktAmt         string   `json:"maxMktAmt"`
	MaxMktSz          string   `json:"maxMktSz"`
	MaxStopSz         string   `json:"maxStopSz"`
	MaxTriggerSz      string   `json:"maxTriggerSz"`
	MaxTwapSz         string   `json:"maxTwapSz"`
	MinSz             string   `json:"minSz"`
	OptType           string   `json:"optType"`
	OpenType          string   `json:"openType"`
	PreMktSwTime      string   `json:"preMktSwTime"`
	QuoteCcy          string   `json:"quoteCcy"`
	TradeQuoteCcyList []string `json:"tradeQuoteCcyList"`
	SettleCcy         string   `json:"settleCcy"`
	State             string   `json:"state"`
	RuleType          string   `json:"ruleType"`
	Stk               string   `json:"stk"`
	TickSz            string   `json:"tickSz"`
	Uly               string   `json:"uly"`
	InstIDCode        int      `json:"instIdCode"`
}
