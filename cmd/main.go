package main

import (
	"fmt"
	"math"

	"space-battle/internal/game"
	"space-battle/internal/movement"
)

func main() {

	// Демонстрация основного тестового случая
	fmt.Println("1. Основной тест движения (12,5) -> (5,8):")

	angle := math.Atan2(3, -7) * 180 / math.Pi
	velocity := math.Sqrt(7*7 + 3*3)

	ship1 := game.NewSpaceship(12, 5, angle, velocity, 0)
	movingAdapter1 := movement.NewIMovingObjectAdapter(ship1)

	x, y := movingAdapter1.GetLocation()
	vx, vy := movingAdapter1.GetVelocity()
	fmt.Printf("Начальная позиция: (%.0f, %.0f)\n", x, y)
	fmt.Printf("Вычисленная скорость: (%.3f, %.3f)\n", vx, vy)

	movingAdapter1.Move()

	x, y = movingAdapter1.GetLocation()
	fmt.Printf("Конечная позиция: (%.0f, %.0f)\n", x, y)

	success := math.Abs(x-5) < 0.001 && math.Abs(y-8) < 0.001
	fmt.Printf("Ожидается: (5, 8) - %v\n", success)

	// Демонстрация вращения
	fmt.Println("\n2. Демонстрация вращения:")
	ship2 := game.NewSpaceship(0, 0, 45, 0, 30)
	rotatingAdapter := movement.NewIRotatingObjectAdapter(ship2)

	fmt.Printf("Начальный угол: %.0f°\n", rotatingAdapter.GetAngle())

	rotatingAdapter.Rotate()

	fmt.Printf("Конечный угол: %.0f°\n", rotatingAdapter.GetAngle())

	// Демонстрация различных направлений движения
	fmt.Println("\n3. Демонстрация движения в разных направлениях:")

	directions := []struct {
		name  string
		angle float64
	}{
		{"Вправо", 0},
		{"Вверх", 90},
		{"Влево", 180},
		{"Вниз", 270},
	}

	for _, dir := range directions {
		ship := game.NewSpaceship(0, 0, dir.angle, 5, 0)
		adapter := movement.NewIMovingObjectAdapter(ship)

		vx, vy := adapter.GetVelocity()
		fmt.Printf("%s: скорость (%.1f, %.1f)\n", dir.name, vx, vy)
	}
}
