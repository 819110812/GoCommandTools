package auto_selector

import (
	"math/rand"
	"time"
)

func AutoSelect(inputs []string) (selected string, err error) {
	if len(inputs) == 0 {
		return "", nil
	}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomIndex := r.Intn(len(inputs))
	selected = inputs[randomIndex]
	return selected, nil
}
