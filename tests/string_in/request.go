package string_in

//go:generate ./../../dist/goval -d -t Request

var enum = []string{"foo", "bar", "b a z"}

type Request struct {
	TagValues string `json:"tagValues" goval:"in=foo,bar,'b a z'"`
	TagVar    string `json:"tagVar" goval:"in={enum}"`
}
