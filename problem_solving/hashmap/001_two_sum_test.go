package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	one []int
	two int
}

type expected struct {
	one []int
}

type config struct {
	p input
	a expected
}

func Test_twoSum(t *testing.T) {
	ast := assert.New(t)

	qs := []config{
		config{
			p: input{
				one: []int{3, 2, 4},
				two: 6,
			},
			a: expected{
				one: []int{1, 2},
			},
		},
		config{
			p: input{
				one: []int{3, 2, 4},
				two: 8,
			},
			a: expected{
				one: nil,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum(p.one, p.two), "输入:%v", p)
	}
}
