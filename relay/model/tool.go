package model

type Tool struct {
	Id       string   `json:"id,omitempty"`
	Type     string   `json:"type,omitempty"` // when splicing claude tools stream messages, it is empty
	Function Function `json:"function"`
}

type Function struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`       // when splicing claude tools stream messages, it is empty
	Parameters  any    `json:"parameters,omitempty"` // request
	Arguments   any    `json:"arguments,omitempty"`  // response
}

type ZhiPuTool struct {
	Id        string     `json:"id,omitempty"`
	Type      string     `json:"type,omitempty"` // when splicing claude tools stream messages, it is empty
	Function  *Function  `json:"function,omitempty"`
	WebSearch *WebSearch `json:"web_search,omitempty"` // zhipu request param
}

type WebSearch struct {
	Enable       bool   `json:"enable,omitempty"`
	SearchQuery  string `json:"search_query,omitempty"`
	SearchResult bool   `json:"search_result,omitempty"`
}
