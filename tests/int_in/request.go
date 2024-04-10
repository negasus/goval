package string_in

//go:generate ./../../dist/goval -d -t Request

var enum = []int{1, 2, 3}

type Request struct {
	TagValues int `json:"tagValues" goval:"in=1,2,3"`
	TagVar    int `json:"tagVar" goval:"in={enum}"`
}
