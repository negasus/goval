package string_in

//go:generate ./../../dist/goval -d -t Request -t Inline

type Inline struct {
	ID int `json:"id" goval:"min=0"`
}

type Request struct {
	Inline `goval:"@"`
}
