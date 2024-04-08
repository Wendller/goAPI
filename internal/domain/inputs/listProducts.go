package inputs

import (
	"net/http"
	"strconv"
)

type ListProductsInput struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

func NewListProductsInput(request *http.Request) (*ListProductsInput, error) {
	var input ListProductsInput

	page := request.URL.Query().Get("page")
	limit := request.URL.Query().Get("limit")
	sort := request.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	input.Page = pageInt
	input.Limit = limitInt
	input.Sort = sort

	return &input, nil
}
