package types

import (
	"bytes"
	"strconv"
)

type Number []int

func (n Number) Trim() Number {
	for i, elem := range n {
		if elem != 0 {
			return n[i:]
		}
	}
	return []int{0}
}

func (n Number) String() string {
	var b bytes.Buffer
	for _, elem := range n {
		b.WriteString(strconv.Itoa(elem))
	}
	return b.String()
}

func (n *Number) Equal(other *Number) bool {
	if len(*n) != len(*other) {
		return false
	}

	for i, elem := range *n {
		if elem != (*other)[i] {
			return false
		}
	}
	return true
}
