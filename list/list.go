package list

type item struct {
	value interface{}
	next *item
}

type List struct {
	head *item
	size int
}

func (list *List) Len() int {
	return list.size
}

func (list *List) Append(value interface{}) {
	newItem := &item{ value: value }
	if list.Len() ==0 {
		list.head = newItem
	} else {
		cur := list.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = newItem
	}
	list.size ++
}
