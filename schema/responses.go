package schema

type Response struct {
	Status string `json:"status,string,omitempty"`
	Error  string `json:"error,string,omitempty"`
}

type CountryResponse struct {
	Status   string `json:"status"`
	Response []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"response"`
}

type CategoryResponse struct {
	Status   string `json:"status"`
	Response []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"response"`
}

type LanguageResponse struct {
	Status   string `json:"status"`
	Response []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"response"`
}