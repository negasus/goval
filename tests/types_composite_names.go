package tests

//go:generate ./../dist/goval -t Composite2 -t Composite -o goval_composite_names.go

type Composite2 struct {
	ID int `json:"id" goval:"min=1"`
}

type Composite struct {
	Item  Composite2   `json:"item" goval:"@"`
	Items []Composite2 `json:"items" goval:"min=1;@"`
}
