package main

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

// func (a *Integer) Less(b Integer) bool {
// 	return (*a).Less(b)
// }

func debug() {
	intt := Integer(1)
	var inttp LessAdder = &intt

	// intt.Add()
	// intt.Less()

	inttp.Add(1)
	inttp.Less(1)

}
