package tgstat

import (
	"context"
	"encoding/json"
	"net/http"
	"tgstat/endpoints"
	//"github.com/helios-ag/tgstat-go/endpoints"
	//"github.com/helios-ag/tgstat-go/schema"
	"tgstat/schema"
)

func (c *Client) CountriesGet(ctx context.Context, lang string) (*schema.CountryResponse, *http.Response, error) {
	path := endpoints.DatabaseCountries

	body := make(map[string]string)
	body["lang"] = lang
	req, err := c.NewRestRequest(ctx, "GET", path, body)

	if err != nil {
		return nil, nil, err
	}

	var response schema.CountryResponse
	result, err := c.Do(req, &response)
	if err != nil {
		return nil, result, err
	}
	_ = json.NewDecoder(result.Body).Decode(&response)

	return &response, result, err
}

func (c *Client) CategoriesGet(ctx context.Context, lang string) (*schema.CategoryResponse, *http.Response, error) {
	path := endpoints.DatabaseCategories

	body := make(map[string]string)
	body["lang"] = lang

	req, err := c.NewRestRequest(ctx, "GET", path, body)

	if err != nil {
		return nil, nil, err
	}

	var response schema.CategoryResponse
	result, err := c.Do(req, &response)
	if err != nil {
		return nil, result, err
	}
	_ = json.NewDecoder(result.Body).Decode(&response)

	return &response, result, err
}

func (c *Client) LanguagesGet(ctx context.Context, lang string) (*schema.LanguageResponse, *http.Response, error) {
	path := endpoints.DatabaseLanguages

	body := make(map[string]string)
	body["lang"] = lang
	req, err := c.NewRestRequest(ctx, "GET", path, body)

	if err != nil {
		return nil, nil, err
	}

	var response schema.LanguageResponse
	result, err := c.Do(req, &response)
	if err != nil {
		return nil, result, err
	}
	_ = json.NewDecoder(result.Body).Decode(&response)

	return &response, result, err
}