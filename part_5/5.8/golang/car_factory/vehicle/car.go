package vehicle

import (
	"fmt"
	"golang/car_factory/vehicle/motor"
)

type Car struct {
	motor           motor.IMotor
	isLeftHandDrive bool
	brand           string
	isMove          bool
}

func (c *Car) GetBrand() string {
	return c.brand
}

func (c *Car) ChangeMotor(motor motor.IMotor) {
	fmt.Println()
	fmt.Println("-------- Change Motor --------")
	fmt.Printf("Old motor: %+v\n", c.motor)
	c.motor = motor
	fmt.Printf("New motor: %+v\n", c.motor)
	fmt.Println()
}

func (c *Car) IsLeftHandDrive() bool {
	return c.isLeftHandDrive
}

func (c *Car) StartMotor() {
	c.motor.RunMotor()
}

func (c *Car) StoptMotor() {
	c.motor.StopMotor()
	c.isMove = false
}

func (c *Car) StartMove() {
	if c.motor.IsRun() {
		c.isMove = true
		fmt.Printf("Сar '%v' started moving\n", c.brand)
	} else {
		fmt.Printf("Сar '%v' can't start moving\n", c.brand)
	}
}

func (c *Car) StopMove() {
	if c.isMove {
		c.isMove = false
		fmt.Printf("Сar '%v' stopped \n", c.brand)
	} else {
		fmt.Println("To stop first start moving")
	}
}

func (c *Car) GetMotorData() string {
	return fmt.Sprintf("%+v", c.motor)
}

func (c *Car) GetMotorPower() string {
	return c.motor.GetPower()
}

func (c *Car) IsMove() bool {
	return c.isMove
}

func NewCar(brand string, motor motor.IMotor, isLeftHandDrive bool) *Car {
	return &Car{
		motor:           motor,
		brand:           brand,
		isLeftHandDrive: isLeftHandDrive,
		isMove:          false,
	}
}

func NewDefaultCar() *Car {
	return &Car{
		motor:           motor.NewAutoVazMotor("1.6"),
		brand:           "Lada",
		isLeftHandDrive: false,
		isMove:          false,
	}
}
