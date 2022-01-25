package Api

import "fmt"

type Item struct {
	Label string
}

func (item *Item) PrintLabel() {
	fmt.Println(item.Label)
}
