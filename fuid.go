package fuid

import (
    "time"
    "strconv"
    "strings"
    "encoding/base64"
    "crypto/rand"
    "github.com/martinlindhe/base36"
)

type Generator struct {
    counter int
}

func pad(number int) string {
    padding := "0000"

    if number >= 9999 {
	return "9999"
    }

    num := strconv.Itoa(number)

    num_len := len(num)
    pad_len := 4

    padding = padding[:pad_len - num_len]

    padding += num

    return strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(padding)), "=", "")
}

func timeBlock() string {
    return strings.ToLower(base36.Encode(uint64(time.Now().UnixNano())))
}

func (g *Generator) counterBlock() string {
    if g.counter > 9999 {
	g.counter = 0
    }

    g.counter += 1

    return pad(g.counter)
}

func randomBlock() string {
    random_bytes := make([]byte, 6)
    rand.Read(random_bytes)

    return strings.ToLower(base36.EncodeBytes(random_bytes))[:6]
}

func (g *Generator) Fuid() string {
    return timeBlock() + g.counterBlock() + randomBlock()
}

