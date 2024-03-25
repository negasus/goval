package tests

//go:generate ./../dist/goval -t RequestInt -t RequestFloat -t RequestString -t RequestStringArray -t RequestIntSlice -o goval_request_compare.go

type RequestInt struct {
	ID int `json:"id" goval:"min=1;max=10"`
}

type RequestFloat struct {
	Price float64 `json:"price" goval:"min=10.5;max=20.5"`
}

type RequestString struct {
	Name string `json:"name" goval:"min=3;max=10"`
}

type RequestStringArray struct {
	Keys []string `json:"keys" goval:"min=3;max=10"`
}

type RequestIntSlice struct {
	IDs []int `json:"ids" goval:"min=1;max=10"`
}
