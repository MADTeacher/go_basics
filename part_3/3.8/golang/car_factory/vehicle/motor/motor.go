package motor

import "fmt"

const hyundai = "Hyundai"
const autoVAZ = "Автоваз"
const defaultPower = "1.4"

type Motor struct {
	power           string
	amountСylinders uint8
	isRun           bool
	manufacturer    string
}

func (m *Motor) GetPower() string {
	return m.power
}

func (m *Motor) RunMotor() {
	fmt.Println("Motor is run")
	m.isRun = true
}

func (m *Motor) StopMotor() {
	fmt.Println("Motor is stop")
	m.isRun = false
}

func (m *Motor) IsRun() bool {
	return m.isRun
}

func (m *Motor) AmountСylinders() uint8 {
	return m.amountСylinders
}

func (m *Motor) GetManufacturerName() string {
	return m.manufacturer
}
