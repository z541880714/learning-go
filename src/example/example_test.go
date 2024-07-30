package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ret := add(1, 2)
	t.Logf("add(3,4) => %d", ret)
}
