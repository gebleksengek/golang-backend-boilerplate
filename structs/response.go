package structs

//ResponseStruct ResponseStruct
type ResponseStruct struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message,omitempty"`
	Result  map[string]interface{} `json:"data,omitempty"`
}

//ResponseMultiDataStruct ResponseStruct
type ResponseMultiDataStruct struct {
	Status  bool                     `json:"status"`
	Message string                   `json:"message,omitempty"`
	Result  []map[string]interface{} `json:"data,omitempty"`
}
