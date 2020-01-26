package strategy

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	a "github.com/alextonkonogov/tonkonogov_progress/pkg/strategy/addition"
	m "github.com/alextonkonogov/tonkonogov_progress/pkg/strategy/multiplication"
)

type data struct {
	first, second, m_expect, a_expect int
}

var tests = []data{
	{2, 3, 6, 5},
	{1, 7, 7, 8},
	{6, 2, 12, 8},
	{0, 5, 0, 5},
	{-6, 5, -30, -1},
}

func TestStrategy(t *testing.T) {
	for i, tt := range tests {
		name := strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			f, s, m_expect, a_expect := tt.first, tt.second, tt.m_expect, tt.a_expect

			ctx := NewContext()

			ctx.Action(m.NewMultiplication())
			m_result := ctx.Apply(f, s)
			assert.EqualValues(t, m_expect, m_result, "Result is not equal to the expected one")

			ctx.Action(a.NewAddition())
			a_result := ctx.Apply(f, s)
			assert.EqualValues(t, a_expect, a_result, "Result is not equal to the expected one")
		})
	}
}
