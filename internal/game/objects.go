package game

type Spaceship struct {
	properties map[string]interface{}
}

func NewSpaceship(x, y, angle, velocity, angularVelocity float64) *Spaceship {
	return &Spaceship{
		properties: map[string]interface{}{
			"location":        []float64{x, y},
			"angle":           angle,
			"velocity":        velocity,
			"angularVelocity": angularVelocity,
		},
	}
}

func (s *Spaceship) GetProperty(name string) interface{} {
	if s == nil {
		return nil
	}
	return s.properties[name]
}

func (s *Spaceship) SetProperty(name string, value interface{}) {
	if s != nil {
		s.properties[name] = value
	}
}
