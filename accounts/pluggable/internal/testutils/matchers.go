package testutils

import (
	"fmt"
)

type PointerMatcher struct {
	C chan<- interface{}
}

func (m PointerMatcher) Matches(x interface{}) bool {
	xAddr := fmt.Sprintf("%p", x)
	CAddr := fmt.Sprintf("%p", m.C)
	return xAddr == CAddr
}

func (m PointerMatcher) String() string {
	return fmt.Sprintf("is %v", m.C)
}
