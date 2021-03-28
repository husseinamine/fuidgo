package fuid

import "testing"

func TestId(t *testing.T) {
    gen := Generator{}

    t.Log(gen.Fuid(), len(gen.Fuid()))
}
