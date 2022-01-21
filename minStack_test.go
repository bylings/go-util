package util

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	ms := NewMinStack()
	ms.Push(1)
	ms.Push(2)
	ms.Push(3)
	fmt.Println("data1", ms.data)
	fmt.Println("pop", ms.Pop())
	fmt.Println("data2", ms.data)
}
