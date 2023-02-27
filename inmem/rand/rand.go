package rand

import (
	"math/rand"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func NewID() int {
	return r.Intn(2 << 15)
}
