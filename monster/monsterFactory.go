package monster

type MonsterProductConfig struct {
	MonsterName     string
	InitHealthPoint int
	InitAttackPoint int
}

type MonsterFactory struct {

}

func (mf *MonsterFactory) Produce(config *MonsterProductConfig) MonsterPrototype {
	return &Monster {
		monsterName: config.MonsterName,
		initHealthPoint: config.InitHealthPoint,
		initAttackPoint: config.InitAttackPoint,
		healthPoint: config.InitHealthPoint,
		attackPoint: config.InitAttackPoint,
	}
}

func NewMonsterFactory() *MonsterFactory {
	return &MonsterFactory{
		
	}
}