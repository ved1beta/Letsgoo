package imports

import "fmt"

type Ticket struct {
	ID    int
	Event string
}

func (t Ticket) printEvent() {
	fmt.Println(t.Event)
}
