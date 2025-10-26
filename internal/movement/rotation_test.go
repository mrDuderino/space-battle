package movement

import (
	"testing"

	"space-battle/internal/game"

	"github.com/stretchr/testify/assert"
)

func TestRotation(t *testing.T) {
	tests := []struct {
		name       string
		startAngle float64
		angularVel float64
		expected   float64
	}{
		{
			name:       "positive rotation",
			startAngle: 45,
			angularVel: 30,
			expected:   75,
		},
		{
			name:       "negative rotation",
			startAngle: 90,
			angularVel: -45,
			expected:   45,
		},
		{
			name:       "rotation with overflow",
			startAngle: 350,
			angularVel: 20,
			expected:   10,
		},
		{
			name:       "rotation with underflow",
			startAngle: 10,
			angularVel: -20,
			expected:   350,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ship := game.NewSpaceship(0, 0, tt.startAngle, 0, tt.angularVel)
			adapter := NewIRotatingObjectAdapter(ship)

			assert.Equal(t, tt.startAngle, adapter.GetAngle())

			adapter.Rotate()

			assert.InDelta(t, tt.expected, adapter.GetAngle(), 0.0000001)
		})
	}
}

func TestRotatingObjectAdapter(t *testing.T) {
	ship := game.NewSpaceship(0, 0, 30, 0, 15)
	adapter := NewIRotatingObjectAdapter(ship)

	assert.Equal(t, 30.0, adapter.GetAngle())
	assert.Equal(t, 15.0, adapter.GetAngularVelocity())

	adapter.SetAngle(60)
	assert.Equal(t, 60.0, adapter.GetAngle())
}
