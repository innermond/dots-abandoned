package rand

import (
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

// letters printable ascii string; lower letters
func Letters(n int) string {
	x := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, n)
	b := 65
	for i := 0; i < n; i++ {
		// only from 65-122 inclusiv; letters lower and upper case
		// except range 90-97 exclusive; printable chars but not letters
		b = x.Intn(123-65) + 65
		for b > 90 && b < 97 {
			b = x.Intn(123-65) + 65
		}
		out = append(out, byte(b))
	}
	return string(out)
}

// Uid return a uuid string
func Uid() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}
