package api

type PagedCollectionResponse struct {
	Elements interface{} `json:"elements"`
	Page     PageInfo    `json:"page"`
}

type PageInfo struct {
	PageNumber    int  `json:"page_number"`
	IsLastPage    bool `json:"is_last_page"`
	TotalElements int  `json:"total_elements"`
	TotalPages    int  `json:"total_pages"`
	PageSize      int  `json:"page_size"`
}
