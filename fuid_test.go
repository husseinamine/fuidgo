package fuid

import (
    "fmt"
    "testing"
)

func TestId(t *testing.T) {
    gen := Generator{}

    fmt.Println(gen.Fuid())
}
