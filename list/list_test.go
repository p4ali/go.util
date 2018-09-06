package list

import (
	"fmt"
	"testing"
)

func TestList_Append(t *testing.T) {
	list := &List{}
	list.Append(1)
	list.Append("hi")
	list.Append(2.0)

	if list.Len() != 3 {
		t.Errorf("list.len() failed. got %d, expect %d", list.Len(), 3)
	}

	cur := list.head
	for cur != nil {
		fmt.Println(cur)
		cur = cur.next
	}
}

