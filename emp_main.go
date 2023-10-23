package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"learning_go.com/employee/Employee"
	// "learning_go.com/employeeObj/User"
	// "learning_go.com/employeeObj/Vehicle"
)

func main() {
	fmt.Println("Hello world")
	// employeeObj := Employee.New("Arvind", 25000, 10000, 5000, 1000)
	// fmt.Println("Profile method " + employeeObj.GetProfile())

	// fmt.Println("Net Salary is ", employeeObj.GetNetSalary())
	// fmt.Println("total deductions is ", employeeObj.GetDeductions())
	// employeeObj.GetSalarySlip()
	// var instAreaShapeRectInf areaShapesInf
	// var instAreaShapeCircleInf areaShapesInf
	// instAreaShapeRectInf = RectangleStruct{5, 5}
	// instAreaShapeCircleInf = CircleStruct{8}
	// fmt.Println("Area of rectangle : " + fmt.Sprint(instAreaShapeRectInf.getArea()))
	// fmt.Println("Area of circle : " + fmt.Sprint(instAreaShapeCircleInf.getArea()))

	// vehicle
	// vehicle := Vehicle.VehicleNew("Aventador", "Lambo", "Orange", 1000, 6)
	// fmt.Println(vehicle)
	// name := vehicle.GetName()
	// vehicle.GetCompany()
	// vehicle.GetColor()
	// carDetails := vehicle.GetCarInfo()
	// fmt.Println("name of car: " + name)
	// fmt.Println("Car details: " + carDetails)
	// fmt.Println("Standard price: " + fmt.Sprint(vehicle.GetPresentPrice()))
	// vehicle.GetDepreciatedPrice()

	// // user
	// // user := User.GetName()
	// user := User.UserNew("Arvind", "124", "arvind@email.com")
	// fmt.Println("user:" + fmt.Sprint(user))
	// username := user.GetName()
	// fmt.Println("username:" + username)
	empRegister()

}

func empRegister() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		vehello := "hello world from fiber grpc combo"
		return c.Send([]byte(vehello))
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		employeeObj, rspMsg := Employee.New(c)
		empProfile := employeeObj.GetProfile()
		if rspMsg == "Err" {
			return c.Send([]byte(rspMsg))
		}
		return c.Send([]byte(empProfile))
		// err := Employee.RegisterEmployee(c)
		// log.Print("error" + fmt.Sprint(err))
		// return err

	})
	log.Fatal(app.Listen(":8080"))
}

type areaShapesInf interface {
	getArea() uint16
}

type RectangleStruct struct {
	length  uint16
	breadth uint16
}

type CircleStruct struct {
	radius uint16
}

func (instRectangleStruct RectangleStruct) getArea() uint16 {
	return instRectangleStruct.breadth * instRectangleStruct.length
}

func (insCircleStruct CircleStruct) getArea() uint16 {
	return uint16(float32(insCircleStruct.radius) * 2 * 3.14)
}
