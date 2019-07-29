package project

/*
============================
			NOTE
============================
This is responses template
*/

type Message map[string][]string

func NewResponse(data interface{}) *SingleResponse {
	response := &SingleResponse{
		Data: data,
		Meta: SimpleMeta {
			StatusCode: 200,
			Message:    Message{
				"success": []string{"Success"},
			},
		},
	}
	return response
}

type SingleResponse struct {
	Data interface{} `json:"data"`
	Meta SimpleMeta  `json:"meta"`
}

type SimpleMeta struct {
	StatusCode int  	`json:"status_code"`
	Message    Message 	`json:"message"`
}

type Pagination struct {
	CurrentPage int64       `json:"current_page"`
	From        int64       `json:"from"`
	LastPage    int64       `json:"last_page"`
	NextPageUrl string      `json:"next_page_url,omitempty"`
	Path        string      `json:"path"`
	PerPage     int64       `json:"per_page"`
	PrevPageUrl string      `json:"prev_page_url,omitempty"`
	To          int64       `json:"to"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
	Meta        SimpleMeta  `json:"meta"`
}
