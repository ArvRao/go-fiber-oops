package Vehicle

import "fmt"

func GetCarName() {
	fmt.Println("getting name")

}

type VehicleInf interface {
	VehicleNew(name string, company string, color string)
	GetNameMthd() string
	GetCompany() string
	GetColor() string
	GetCarInfo() string
	getCarDetails()
	GetPresentPrice() uint32
	calculateDepreciatedPrice()
}

type VehicleStruct struct {
	name            string
	company         string
	color           string
	carDetails      string
	standardPrice   uint32
	carAge          uint8
	depreciatePrice uint32
}

func VehicleNew(name string, company string, color string, standardPrice uint32, carAge uint8) VehicleStruct {
	fmt.Println(name)
	var instVehicleStruct VehicleStruct
	instVehicleStruct.name = name
	instVehicleStruct.color = color
	instVehicleStruct.company = company
	instVehicleStruct.standardPrice = standardPrice
	instVehicleStruct.carAge = carAge
	return instVehicleStruct
}

func (instVehicleStruct VehicleStruct) GetNameMthd() string {
	return instVehicleStruct.name
}

func (instVehicleStruct VehicleStruct) GetCompany() string {
	return instVehicleStruct.company
}

func (instVehicleStruct VehicleStruct) GetColor() string {
	return instVehicleStruct.color
}

func (instVehicleStruct VehicleStruct) GetCarInfo() string {
	instVehicleStruct.getCarDetails()
	return instVehicleStruct.carDetails
}

func (instVehicleStruct *VehicleStruct) getCarDetails() {
	// logic
	instVehicleStruct.carDetails = instVehicleStruct.name + " of " + instVehicleStruct.company + " " + instVehicleStruct.color + " color "
}

func (instVehicleStruct VehicleStruct) GetPresentPrice() uint32 {
	return instVehicleStruct.standardPrice
}

func (instVehicleStruct VehicleStruct) GetDepreciatedPrice() {
	instVehicleStruct.calculateDepreciatedPrice()
}

func (instVehicleStruct *VehicleStruct) calculateDepreciatedPrice() {
	switch {
	case instVehicleStruct.carAge <= 1:
		instVehicleStruct.depreciatePrice = instVehicleStruct.standardPrice - uint32(float32(instVehicleStruct.standardPrice)*(0.1))
	case instVehicleStruct.carAge <= 2:
		instVehicleStruct.depreciatePrice = instVehicleStruct.standardPrice - uint32(float32(instVehicleStruct.standardPrice)*(0.2))
	case instVehicleStruct.carAge <= 3:
		instVehicleStruct.depreciatePrice = instVehicleStruct.standardPrice - uint32(float32(instVehicleStruct.standardPrice)*(0.4))
	case instVehicleStruct.carAge >= 3:
		instVehicleStruct.depreciatePrice = instVehicleStruct.standardPrice - uint32(float32(instVehicleStruct.standardPrice)*(0.7))
	}
	println("Depreciated price: " + fmt.Sprint(instVehicleStruct.depreciatePrice))
}
