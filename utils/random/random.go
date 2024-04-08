package random

import (
	"fmt"
	"math/rand"
)

// Below function can take range parameters and generate a randomNumber alphanumberic string based on a range specified.
func RandomNumberInRange() string {
	return fmt.Sprintf("%v", rand.Intn(10000000))
}
