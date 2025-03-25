package motor

type HyundaiMotor struct {
	Motor
	turboMod bool
}

func (m *HyundaiMotor) IsTurboModOn() bool {
	return m.turboMod
}

func (m *HyundaiMotor) TurboModOn() {
	m.turboMod = true
}

func (m *HyundaiMotor) TurboModoff() {
	m.turboMod = false
}

func NewHyundaiMotor(power string) *HyundaiMotor {
	motor := &HyundaiMotor{
		Motor: Motor{
			manufacturer: hyundai,
			isRun:        false,
		},
		turboMod: false,
	}
	switch power {
	case "1.8":
		motor.amountСylinders = 6
		motor.power = power
	case "1.6":
		motor.amountСylinders = 4
		motor.power = power
	default:
		motor.amountСylinders = 4
		motor.power = defaultPower
	}
	return motor
}
