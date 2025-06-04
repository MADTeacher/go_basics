package main

import (
	"fmt"
	"golang/car_factory/vehicle"
	"golang/car_factory/vehicle/motor"
)

func printCarInfo(car *vehicle.Car) {
	fmt.Printf("----------- %v -----------\n", car.GetBrand())
	fmt.Printf("Car: %+v\n", car)
	fmt.Printf("motor: %+v\n", car.GetMotorData())
	fmt.Println()
}

func main() {
	car := vehicle.NewDefaultCar()
	printCarInfo(car)
	car.StartMove()
	car.StartMotor()
	car.StartMove()
	car.StoptMotor()
	fmt.Println("Is car move? ", car.IsMove())

	// меняем двигатель
	newMotor := motor.NewHyundaiMotor("1.8")
	car.ChangeMotor(newMotor)

	printCarInfo(car)
	printCarInfo(car)
	car.StartMove()
	car.StartMotor()
	car.StartMove()
	fmt.Println("Is car move? ", car.IsMove())
}
