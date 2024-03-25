package tests

//go:generate ./../dist/goval -t CompareInIntVar -t CompareInIntVarSlice -t CompareInInt -t CompareInIntSlice -t CompareInStringVar -t CompareInString -t CompareInStringVarSlice -t CompareInStringSlice -o goval_request_in.go

var compareIntVarData = []int{1, 2, 3}

type CompareInIntVar struct {
	ID int `json:"id" goval:"in={compareIntVarData}"`
}

type CompareInIntVarSlice struct {
	ID []int `json:"id" goval:"in={compareIntVarData}"`
}

type CompareInInt struct {
	ID int `json:"id" goval:"in=10,20,30"`
}

type CompareInIntSlice struct {
	ID []int `json:"id" goval:"in=10,20,30"`
}

var compareStringVarData = []string{"a", "b", "c"}

type CompareInStringVar struct {
	Name string `json:"name" goval:"in={compareStringVarData}"`
}

type CompareInStringVarSlice struct {
	Name []string `json:"name" goval:"in={compareStringVarData}"`
}

type CompareInString struct {
	Name string `json:"name" goval:"in=aa,bb,cc"`
}

type CompareInStringSlice struct {
	Name []string `json:"name" goval:"in=aa,bb,cc"`
}
