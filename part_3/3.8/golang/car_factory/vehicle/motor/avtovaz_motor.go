package motor

type AvtoVazMotor struct {
	Motor
}

func NewAutoVazMotor(power string) *AvtoVazMotor {
	motor := &AvtoVazMotor{
		Motor: Motor{
			manufacturer: autoVAZ,
			isRun:        false,
		},
	}
	switch power {
	case "1.6":
		motor.amountСylinders = 4
		motor.power = power
	default:
		motor.amountСylinders = 4
		motor.power = defaultPower
	}
	return motor
}
