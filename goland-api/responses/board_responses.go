package responses

type BoardResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
