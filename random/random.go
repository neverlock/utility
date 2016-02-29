package random

import (
	"github.com/lazybeaver/xorshift"
	"math/rand"
	"time"
)

func Str(l int) string {
	var char string
	buf := make([]byte, l)
	char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < l; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

func Int(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Fint(max uint64) int {
	seed := uint64(time.Now().Nanosecond())
	xor128p := xorshift.NewXorShift128Plus(seed)
	random := xor128p.Next()
	return int(random % max)
}
