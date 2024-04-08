package random

import (
	"fmt"
	"math/rand"
)

func RandomNumberInRange() string {
	return fmt.Sprintf("%v", rand.Intn(10000000))
}
