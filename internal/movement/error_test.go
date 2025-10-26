package movement

import (
	"testing"

	"space-battle/internal/game"

	"github.com/stretchr/testify/assert"
)

// TestInvalidObjects проверяет работу с невалидными объектами
func TestInvalidObjects(t *testing.T) {
	// Адаптер с nil объектом
	movingAdapter := NewIMovingObjectAdapter(nil)
	rotatingAdapter := NewIRotatingObjectAdapter(nil)

	// Get методы не должны паниковать
	x, y := movingAdapter.GetLocation()
	assert.Equal(t, 0.0, x)
	assert.Equal(t, 0.0, y)

	vx, vy := movingAdapter.GetVelocity()
	assert.Equal(t, 0.0, vx)
	assert.Equal(t, 0.0, vy)

	angle := rotatingAdapter.GetAngle()
	assert.Equal(t, 0.0, angle)

	angularVel := rotatingAdapter.GetAngularVelocity()
	assert.Equal(t, 0.0, angularVel)

	// Set методы не должны паниковать
	movingAdapter.SetLocation(1, 2)
	rotatingAdapter.SetAngle(45)

	// Move и Rotate не должны паниковать
	movingAdapter.Move()
	rotatingAdapter.Rotate()

	t.Log("✅ Работа с nil объектами безопасна")
}

// TestInvalidProperties проверяет работу с невалидными свойствами
func TestInvalidProperties(t *testing.T) {
	// Создаем нормальный корабль, но потом испортим свойства
	ship := game.NewSpaceship(0, 0, 0, 0, 0)

	// Устанавливаем невалидные свойства
	ship.SetProperty("location", "invalid")
	ship.SetProperty("angle", "45")
	ship.SetProperty("velocity", "fast")
	ship.SetProperty("angularVelocity", "spin")

	movingAdapter := NewIMovingObjectAdapter(ship)
	rotatingAdapter := NewIRotatingObjectAdapter(ship)

	// Должно работать без паники
	x, y := movingAdapter.GetLocation()
	assert.Equal(t, 0.0, x)
	assert.Equal(t, 0.0, y)

	vx, vy := movingAdapter.GetVelocity()
	assert.Equal(t, 0.0, vx)
	assert.Equal(t, 0.0, vy)

	angle := rotatingAdapter.GetAngle()
	assert.Equal(t, 0.0, angle)

	angularVel := rotatingAdapter.GetAngularVelocity()
	assert.Equal(t, 0.0, angularVel)

	// Move и Rotate не должны паниковать
	movingAdapter.Move()
	rotatingAdapter.Rotate()

	t.Log("Работа с невалидными свойствами безопасна")
}
