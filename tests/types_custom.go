package tests

import "github.com/negasus/goval"

//go:generate ./../dist/goval -t RequestCustomEmbedItem -t RequestCustomEmbed -t RequestCustom -o goval_request_custom.go

type RequestCustomEmbedItem struct {
	ID int `json:"id" goval:"min=10"`
}

type RequestCustomEmbed struct {
	RequestCustomEmbedItem `goval:"@"`
}

func (r *RequestCustom) customValidationFunc() goval.Errors {
	err := goval.Errors{}

	err.Add("id", goval.NewCustomError(1).AddValue("rule_value", 1))

	return err
}

type RequestCustom struct {
	ID int `json:"id" goval:"@customValidationFunc"`
}
