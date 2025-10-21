package movement

// Интерфейс игрового объекта
type IGameItem interface {
	GetProperty(name string) interface{}
	SetProperty(name string, value interface{})
}

// Интерфейс для движущегося объекта
type IMovingObject interface {
	GetLocation() (x, y float64)
	GetVelocity() (vx, vy float64)
	SetLocation(x, y float64)
}

// Интерфейс для вращающегося объекта
type IRotatingObject interface {
	GetAngle() float64
	GetAngularVelocity() float64
	SetAngle(angle float64)
}
