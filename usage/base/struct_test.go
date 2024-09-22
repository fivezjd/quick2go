package base

type User interface {
	GetName() string
}

type Teacher struct {
}

type T1 interface {
	~int
	String() string
}
