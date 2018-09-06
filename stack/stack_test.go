package stack

import "testing"

func TestPush(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)

	if stack.Len() != 1 {
		t.Errorf("Len() is incorret, got: %d, want %d", stack.Len(), 1)
	}

	if stack.top.next != nil {
		t.Errorf("top.next incorrect, got non nil, expect nil")
	}

	if stack.top.value != 1 {
		t.Errorf("top.value incorrect, got %d, expect %d", stack.top.value, 1)
	}

	stack.Push(2)

	if stack.top.next == nil {
		t.Errorf("top.next incorrect, got nil, expect not nil")
	}

	if stack.top.next.value != 1 {
		t.Errorf("top.next.value incorrect, got %d, expect %d", stack.top.next.value, 1)
	}
}

func TestPop(t *testing.T) {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)

	value := stack.Pop()
	if stack.Len() != 1 {
		t.Errorf("Len() is incorret, got: %d, want %d", stack.Len(), 1)
	}

	if value != 2 {
		t.Errorf("pop value incorrect, got: %d, want %d", value, 2)
	}

	if stack.top.next != nil {
		t.Errorf("top.next incorrect, got non nil, expect nil")
	}

	value = stack.Pop()
	if value != 1 {
		t.Errorf("pop.value incorrect, got %d, expect %d", value, 1)
	}
	if stack.Len() != 0 {
		t.Errorf("Len() incorrect, got %d, expect %d", stack.Len(), 0)
	}

	value = stack.Pop()
	if value != nil {
		t.Errorf("pop.value incorrect, got not nil, expect nil")
	}

}
