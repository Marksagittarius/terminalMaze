package item

type ItemFactory struct {

}

type ItemProductConfig struct {
	ItemName string
	RationalPoint int
}

func NewItemFactory() *ItemFactory {
	return &ItemFactory {
		
	}
}

func (ipc *ItemFactory) Produce(config *ItemProductConfig) ItemPrototype {
	return &Item {
		itemName: config.ItemName,
		rationalPoint: config.RationalPoint,
	}
}

