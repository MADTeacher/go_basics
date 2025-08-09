package motor

type IMotor interface {
	GetPower() string
	RunMotor()
	StopMotor()
	IsRun() bool
	AmountСylinders() uint8
	GetManufacturerName() string
}
