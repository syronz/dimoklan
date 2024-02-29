package consts

const (
	MaxRetryForPickColor = 100
	LanguageEn           = "en"
	MaxUserID            = 16777215
	HiddenPassword       = "******"
	IP                   = "ip"
	// Map size
	MaxX = 1000
	MaxY = 1000

	// Used for make hash in registration
	HashSalt = "-Ug*mY.m6-FAuFX/K>cxlFri"

	// Base domain
	TableData = "data"

	// Entities
	RegisterEntity = "register"
	UserEntity     = "user"
	AuthEntity     = "auth"
	MarshalEntity  = "marshal"
	CellEntity     = "cell"

	// Partitions
	ParCell     = "c#"
	ParFraction = "f#"
	ParRegister = "r#"
	ParUser     = "u#"
	ParMarshal  = "m#"
	ParAuth     = "e#"

	// Map Domain
	TableMap = "map"

	// Gameplay
	GoldForNewUser   = 100
	FarrForNewUser   = 10
	ArmyForNewUser   = 100
	StarForNewUser   = 1
	SpeedForNewUser  = 1.
	AttackForNewUser = 1.
)
