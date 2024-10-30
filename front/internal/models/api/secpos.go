package api

type ApiData struct {
	SecurityPositionKey ApiSecurityPositionKey
	SearchFilter        ApiSearchFilter
}

type ApiSecurityPositionKey struct {
	Clientbic       string `uri:"clientbic" binding:"required"`
	Isin            string `uri:"isin" binding:"required"`
	Account         string `uri:"account" binding:"required"`
	Restrictiontype string `uri:"restrictiontype" binding:"required"`
}

type ApiSearchFilter struct {
	BusinessPeriodType string `form:"businessperiodtype"`
	BusinessDate       string `form:"businessdate"`
	// ApiSecurityPositionKey ApiSecurityPositionKey
}

type SecurityPosition struct {
	SecurityPositionKey   SecurityPositionKey   `json:"securityPositionKey"`
	SecurityPositionValue SecurityPositionValue `json:"securityPositionValue"`
}

type SecurityPositionKey struct {
	Isin            string `json:"isin"`
	Account         string `json:"account"`
	RestrictionType string `json:"restrictionType"`
	ClientId        string `json:"clientId"`
}

type SecurityPositionValue struct {
	Position_quantity     float64 `json:"position_quantity"`
	Position_quantity_sod float64 `json:"position_quantity_sod"`
	Period_evt_reference  string  `json:"period_evt_reference"`
	Sett_position_ts      string  `json:"sett_position_ts"`
}
