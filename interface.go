package masa

type Masa interface {
	Convert(format string) Masa
	String() string
}
