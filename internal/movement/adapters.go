package movement

import "math"

type IMovingObjectAdapter struct {
	item IGameItem
}

func NewIMovingObjectAdapter(item IGameItem) *IMovingObjectAdapter {
	return &IMovingObjectAdapter{item: item}
}

func (a *IMovingObjectAdapter) GetLocation() (x, y float64) {
	if a.item == nil {
		return 0, 0
	}

	location := a.item.GetProperty("location")
	if location == nil {
		return 0, 0
	}

	if loc, ok := location.([]float64); ok && len(loc) == 2 {
		return loc[0], loc[1]
	}

	return 0, 0
}

func (a *IMovingObjectAdapter) GetVelocity() (vx, vy float64) {
	if a.item == nil {
		return 0, 0
	}

	angleInterface := a.item.GetProperty("angle")
	velocityInterface := a.item.GetProperty("velocity")

	var angle, velocity float64

	if angleInterface != nil {
		if ang, ok := angleInterface.(float64); ok {
			angle = ang
		} else {
			angle = 0
		}
	}

	if velocityInterface != nil {
		if vel, ok := velocityInterface.(float64); ok {
			velocity = vel
		} else if vel, ok := velocityInterface.(int); ok {
			velocity = float64(vel)
		} else {
			velocity = 0
		}
	}

	radians := angle * math.Pi / 180
	vx = velocity * math.Cos(radians)
	vy = velocity * math.Sin(radians)

	return vx, vy
}

func (a *IMovingObjectAdapter) SetLocation(x, y float64) {
	if a.item == nil {
		return
	}
	a.item.SetProperty("location", []float64{x, y})
}

type IRotatingObjectAdapter struct {
	item IGameItem
}

func NewIRotatingObjectAdapter(item IGameItem) *IRotatingObjectAdapter {
	return &IRotatingObjectAdapter{item: item}
}

func (a *IRotatingObjectAdapter) GetAngle() float64 {
	if a.item == nil {
		return 0
	}

	angleInterface := a.item.GetProperty("angle")
	if angleInterface == nil {
		return 0
	}

	if angle, ok := angleInterface.(float64); ok {
		return angle
	}

	return 0
}

func (a *IRotatingObjectAdapter) GetAngularVelocity() float64 {
	if a.item == nil {
		return 0
	}

	angularVelocityInterface := a.item.GetProperty("angularVelocity")
	if angularVelocityInterface == nil {
		return 0
	}

	if angularVelocity, ok := angularVelocityInterface.(float64); ok {
		return angularVelocity
	}

	return 0
}

func (a *IRotatingObjectAdapter) SetAngle(angle float64) {
	if a.item == nil {
		return
	}
	a.item.SetProperty("angle", angle)
}

// Move Движение по прямой
func (a *IMovingObjectAdapter) Move() {
	x, y := a.GetLocation()
	vx, vy := a.GetVelocity()
	a.SetLocation(x+vx, y+vy)
}

// Rotate Поворот объекта на один шаг (угловая скорость - это градусы на шаг)
func (a *IRotatingObjectAdapter) Rotate() {
	currentAngle := a.GetAngle()
	angularVelocity := a.GetAngularVelocity()
	newAngle := currentAngle + angularVelocity

	// Нормализуем угол в диапазон [0, 360)
	newAngle = math.Mod(newAngle, 360)
	if newAngle < 0 {
		newAngle += 360
	}

	a.SetAngle(newAngle)
}
