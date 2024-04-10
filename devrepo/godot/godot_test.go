package godot

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	idx, er := fetchIndex()
	fmt.Println(idx)
	fmt.Println(er)
}
