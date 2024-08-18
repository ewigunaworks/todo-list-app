package helpers

type RequestList struct {
	Page        int32  `json:"Page"`
	Limit       int32  `json:"Limit"`
	Sort        string `json:"Sort"`
	SortBy      string `json:"SortBy"`
	FilterField string `json:"FilterField"`
	FilterValue string `json:"FilterValue"`
	Search      string `json:"Search"`
}

func (f *RequestList) IsFilterFieldEmpty() bool {
	return f.FilterField == ""
}

func (f *RequestList) IsFilterValueEmpty() bool {
	return f.FilterValue == ""
}
