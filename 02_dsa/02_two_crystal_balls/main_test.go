package twocrystalballs

import (
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBalls(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	for i := 0; i < idx; i++ {
		data[i] = true
	}

	result := TwoCrystalBalls(data)
	if idx != result {
		t.Errorf("TwoCrystalBalls(%d) = %d", idx, result)
	}

	result = TwoCrystalBalls(make([]bool, 881))
	if result != -1 {
		t.Errorf("TwoCrystalBalls(%d) = %d", idx, result)
	}
}
