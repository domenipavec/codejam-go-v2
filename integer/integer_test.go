package integer

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 3))
	assert.Equal(t, 2, Min(15, 2))
	assert.Equal(t, 0, Min(0, 1))
	assert.Equal(t, -1, Min(1, -1))
	assert.Equal(t, 0, Min(5, 4, 3, 2, 1, 0))
	assert.Equal(t, -1, Min(5, 4, 3, -1, 1, 0))
	assert.Equal(t, 1, Min(1))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max(1, 3))
	assert.Equal(t, 15, Max(15, 2))
	assert.Equal(t, 1, Max(0, 1))
	assert.Equal(t, -1, Max(-20, -1))
	assert.Equal(t, 5, Max(5, 4, 3, 2, 1, 0))
	assert.Equal(t, 10, Max(5, 4, 3, 10, 1, 0))
	assert.Equal(t, 1, Max(1))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 1, Abs(1))
	assert.Equal(t, 1, Abs(-1))
	assert.Equal(t, 1234, Abs(1234))
	assert.Equal(t, 1234, Abs(-1234))
	assert.Equal(t, 0, Abs(0))
}

func TestCeil(t *testing.T) {
	assert.Equal(t, 1, Ceil(5, 5))
	assert.Equal(t, 2, Ceil(6, 5))
	assert.Equal(t, 2, Ceil(8, 5))
	assert.Equal(t, 2, Ceil(10, 5))
	assert.Equal(t, 3, Ceil(11, 5))
}

func TestRange(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, Range(3))
	assert.Equal(t, []int{2, 3, 4, 5, 6}, Range(2, 7))
	assert.Equal(t, []int{3, 5, 7}, Range(3, 8, 2))
	assert.Equal(t, []int{5, 10, 15, 20, 25}, Range(5, 30, 5))
}

func TestGcd(t *testing.T) {
	assert.Equal(t, 2, Gcd(40, 6, 1000))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 30, Lcm(2, 5, 10, 3))
}

func BenchmarkGcd(b *testing.B) {
	as := make([]int, b.N)
	bs := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		as[i] = rand.Int()
		bs[i] = rand.Int()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcd2(as[i], bs[i])
	}
}
