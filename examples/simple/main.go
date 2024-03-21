package main

import (
	"fmt"

	"github.com/negasus/goval"
)

//go:generate ../../dist/goval -d -t UsersRequest -t Custom -t Pagination

type Pagination struct {
	Page    int `json:"page" goval:"min=0"`
	PerPage int `json:"perpage" goval:"in=10,50,100"`
}

type Custom struct {
	Index int `json:"index" goval:"min=50"`
}

var customMessageID = 1

func (c *Custom) customValidate() goval.Errors {
	errors := goval.Errors{}

	e := goval.NewCustomError(customMessageID).
		AddValue("value", c.Index)

	errors.Add("Custom2.Index", e)

	return errors
}

func (u *UsersRequest) customValidate() goval.Errors {
	errors := goval.Errors{}

	e := goval.NewError(goval.ErrorTypeInvalid).
		AddValue("field", "Custom")

	errors.Add("Custom", e)

	return errors
}

var (
	fields = []string{"id", "custom", "search", "s"}
)

// UsersRequest is a request for users
type UsersRequest struct {
	Pagination `goval:"@"` // @ alias for `goval:"@Validate"`

	Field   string `json:"field" goval:"in={fields}"`
	Type    string `json:"type" goval:"in=foo,bar"`
	ID      int    `json:"id" goval:"min=1"`
	Custom  Custom `json:"custom" goval:"@customValidate"`
	Custom2 Custom `json:"custom2" goval:"@Custom2.customValidate"`

	Price float64 `json:"price" goval:"min=0.1"`
}

func main() {
	r := &UsersRequest{}
	r.Page = -1
	r.PerPage = 17
	r.ID = 42
	r.Field = "id"
	r.Type = "foo"
	r.Custom.Index = 42
	r.Custom2.Index = 17

	goval.AddCustomMessage(customMessageID, "en", "Bad value for Custom2.Index: {value}")
	goval.AddCustomMessage(customMessageID, "ru", "Неверное значение для поля Custom2.Index: {value}")

	goval.SetDefaultLang("ru")

	errs := r.Validate()

	for k, v := range errs {
		for _, e := range v {
			fmt.Printf("%s: %+v\n", k, e)
		}
	}
}
