package ring

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRingInit(t *testing.T) {
	r := InitRing(10)
	fmt.Println(r.String())
}

func TestRing_RingAddGet(t *testing.T) {
	r := InitRing(5)

	r.RingAdd(1)
	t1 := time.Now().Unix()
	fmt.Println(t1)
	time.Sleep(time.Second)

	r.RingAdd(2)
	time.Sleep(time.Second)

	r.RingAdd(3)
	time.Sleep(time.Second)

	r.RingAdd(4)
	time.Sleep(time.Second)

	r.RingAdd(5)
	t2 := time.Now().Unix()
	fmt.Println(t2)

	fmt.Println(r.String())

	_, ok := r.RingGet(t1)
	assert.Equal(t, true, ok, r.String())

	_, ok = r.RingGet(t2)
	assert.Equal(t, true, ok, r.String())

	_, ok = r.RingGet(t1 - 1)
	assert.Equal(t, false, ok, r.String())

	time.Sleep(time.Second)
	r.RingAdd(6)
	fmt.Println(r.String())
}
