package twocrystalballs

import (
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBalls(t *testing.T) {
	rand.Seed(time.Now().UnixNano() % 1000)
	idx := rand.Intn(1000)
	data := make([]bool, 1000)
	empty := make([]bool, 1000)

	for i := idx; i < len(data); i++ {
		data[i] = true
	}
	for i := 0; i < len(empty); i++ {
		empty[i] = false
	}

	result := TwoCrystalBalls(data)
	if idx != result {
		t.Errorf("TwoCrystalBalls(%d) = %d", idx, result)
	}

	result = TwoCrystalBalls(empty)
	if result != -1 {
		t.Errorf("TwoCrystalBalls(%d) = %d", idx, result)
	}
}
