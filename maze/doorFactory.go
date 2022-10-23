package maze

type DoorProductConfig struct {
	DoorType   string
	DoorId     string
}

type DoorFactory struct {

}

func (df *DoorFactory) Produce(config *DoorProductConfig) DoorPrototype {
	return &Door{
		connection: make(map[string]string),
		doorType: config.DoorType,
		doorId: config.DoorId,
		preloads: make([]DoorMiddleware, 0),
		callbacks: make([]DoorMiddleware, 0),
		visited: false,
	}
}

func NewDoorFactory() *DoorFactory {
	return &DoorFactory{

	}
}