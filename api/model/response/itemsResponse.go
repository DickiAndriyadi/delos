package response

type ItemsResponse struct {
	Data struct {
		Items interface{} `json:"items,omitempty"`
	} `json:"data"`
}
