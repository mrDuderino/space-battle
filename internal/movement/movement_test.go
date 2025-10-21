package movement

import (
	"github.com/stretchr/testify/assert"
	"math"
	"space-battle/internal/game"
	"testing"
)

func TestMovementFrom12_5To5_8(t *testing.T) {
	angle := math.Atan2(3, -7) * 180 / math.Pi
	velocity := math.Sqrt(7*7 + 3*3)

	ship := game.NewSpaceship(12, 5, angle, velocity, 0)
	adapter := NewIMovingObjectAdapter(ship)

	x, y := adapter.GetLocation()
	assert.Equal(t, 12.0, x)
	assert.Equal(t, 5.0, y)

	vx, vy := adapter.GetVelocity()
	assert.InEpsilon(t, -7.0, vx, 0.001)
	assert.InEpsilon(t, 3.0, vy, 0.001)

	adapter.Move()

	x, y = adapter.GetLocation()
	assert.InEpsilon(t, 5.0, x, 0.001)
	assert.InEpsilon(t, 8.0, y, 0.001)
}
