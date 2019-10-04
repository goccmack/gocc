package ast

var (
	StringGetter func(interface{}) string
)

func getString(v interface{}) string {
	return StringGetter(v)
}
