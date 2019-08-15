package gncdu

import (
	"fmt"
	"testing"
)

func TestToHumanSize(t *testing.T) {
	fmt.Println(ToHumanSize(60))
	fmt.Println(ToHumanSize(60 * 1024))
	fmt.Println(ToHumanSize(60 * 1024 * 1024))
}
