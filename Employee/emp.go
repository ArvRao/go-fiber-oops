package Employee

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

const (
	FULLTIME uint8 = iota
	CONTRACT
	PARTTIME
	OUTSOURCE
)

type EmployeeStructDb struct {
	ID                  uint   `gorm:"primaryKey"`
	empId               string `gorm:"uniqueIndex;type:char(8);not null"`
	name                string `json:"Name" gorm:"type:varchar(255);not null"`
	basic               string `json:"Basic" gorm:"type:char(6);not null"`
	hra                 string `json:"Basic" gorm:"type:char(6);not null"`
	specialAllowance    string `json:"Basic" gorm:"type:char(6);not null"`
	leaveAllowance      string `json:"Basic" gorm:"type:char(6);not null"`
	nightshiftAllowance string `json:"Basic" gorm:"type:char(6);not null"`
	gratuity            string `json:"Basic" gorm:"type:char(6);not null"`
	providentFund       string `json:"Basic" gorm:"type:char(6);not null"`
	totalEarnings       string `json:"Basic" gorm:"type:char(6);not null"`
}

type EmployeeStruct struct {
	id           uint16
	Name         string `json:"Name"`
	salaryStruct SalaryStruct
}

type SalaryStruct struct {
	salaryEarnings   SalaryEarningsStruct
	salaryDeductions SalaryDeductionsStruct
}

type SalaryEarningsStruct struct {
	basic               uint16
	hra                 uint16
	specialAllowance    uint16
	leaveAllowance      uint16
	nightshiftAllowance uint16
	retirementBenefits  RetirementBenefitsStruct
	totalEarnings       uint16
}

type RetirementBenefitsStruct struct {
	gratuity      uint16
	providentFund uint16
}

type SalaryDeductionsStruct struct {
	taxes           uint16
	cess            uint16
	totalDeductions uint16
}

// collection of signature methods for class
type EmployeeInf interface {
	New(name string, basic uint16, specialAllowance uint16, leaveAllowance uint16, nightshiftAllowance uint16)
	getProfile() (uint8, string)
	getDeductions() uint16
	getGrossSalary() uint16
	getNetSalary() uint16
}

// function that takes necessary data in order to create object/instance. (receiver function)
func New(name string, basic uint16, specialAllowance uint16, leaveAllowance uint16, nightshiftAllowance uint16) EmployeeStruct {
	fmt.Println(basic)
	var instEmployeeStruct EmployeeStruct
	instEmployeeStruct.Name = name
	instEmployeeStruct.salaryStruct.salaryEarnings.basic = basic
	instEmployeeStruct.salaryStruct.salaryEarnings.specialAllowance = specialAllowance
	instEmployeeStruct.salaryStruct.salaryEarnings.leaveAllowance = leaveAllowance
	instEmployeeStruct.salaryStruct.salaryEarnings.nightshiftAllowance = nightshiftAllowance
	return instEmployeeStruct
}

func (instEmployeeStruct EmployeeStruct) GetProfile() string {
	fmt.Println(instEmployeeStruct.salaryStruct.salaryEarnings.basic)
	return ("Name of employee is " + instEmployeeStruct.Name + " with basic salary " + strconv.Itoa(int((instEmployeeStruct.salaryStruct.salaryEarnings.basic))))
}

func (instEmployeeStruct EmployeeStruct) GetGrossSalary() uint16 {
	return instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings - instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions
}

func (instEmployeeStruct *EmployeeStruct) getHra() {
	basic := instEmployeeStruct.salaryStruct.salaryEarnings.basic
	println("basic salary:" + fmt.Sprint(basic))

	switch {
	case basic <= 15000:
		instEmployeeStruct.salaryStruct.salaryEarnings.hra = uint16(float32(basic) * (0.16))
	case basic > 15000 && basic <= 25000:
		instEmployeeStruct.salaryStruct.salaryEarnings.hra = uint16(float32(basic) * (0.14))
	case basic > 25000 && basic <= 35000:
		instEmployeeStruct.salaryStruct.salaryEarnings.hra = uint16(float32(basic) * (0.12))
	case basic > 35000:
		instEmployeeStruct.salaryStruct.salaryEarnings.hra = uint16(float32(basic) * (0.10))
	}
	println("hra salary: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryEarnings.hra))
}

func (instEmployeeStruct *EmployeeStruct) getRetirementBenefits() {
	instEmployeeStruct.salaryStruct.salaryEarnings.retirementBenefits.gratuity = uint16(float32(instEmployeeStruct.salaryStruct.salaryEarnings.basic) * (0.1))
	instEmployeeStruct.salaryStruct.salaryEarnings.retirementBenefits.providentFund = uint16(float32(instEmployeeStruct.salaryStruct.salaryEarnings.basic) * (0.2))
}

func (instEmployeeStruct *EmployeeStruct) GetNetSalary() uint16 {
	instEmployeeStruct.getRetirementBenefits()
	instEmployeeStruct.getHra()
	instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings = instEmployeeStruct.salaryStruct.salaryEarnings.basic + instEmployeeStruct.salaryStruct.salaryEarnings.hra + instEmployeeStruct.salaryStruct.salaryEarnings.leaveAllowance + instEmployeeStruct.salaryStruct.salaryEarnings.nightshiftAllowance + instEmployeeStruct.salaryStruct.salaryEarnings.specialAllowance + instEmployeeStruct.salaryStruct.salaryEarnings.retirementBenefits.gratuity +
		instEmployeeStruct.salaryStruct.salaryEarnings.retirementBenefits.providentFund
	println("basic: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryEarnings.basic))
	println("hra: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryEarnings.hra))
	println("Night shift allowance: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryEarnings.nightshiftAllowance))
	println("Special allowance: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryEarnings.specialAllowance))
	fmt.Println("PF: ", instEmployeeStruct.salaryStruct.salaryEarnings.retirementBenefits.providentFund)
	fmt.Println("Gratuity:", instEmployeeStruct.salaryStruct.salaryEarnings.retirementBenefits.gratuity)
	fmt.Println("Total earnings", instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings)

	return instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings
}

func (instEmployeeStruct *EmployeeStruct) getTaxes() {
	totalEarnings := int(instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings)
	fmt.Println("total earnings in taxes: " + fmt.Sprint(totalEarnings))
	var taxRate uint16
	var cessRate uint16
	switch {
	case totalEarnings <= 50000:
		taxRate = 0
		cessRate = 0
	case totalEarnings > 50000 && totalEarnings <= 70000:
		taxRate = 10
		cessRate = 20
	case totalEarnings > 70000 && totalEarnings <= 90000:
		taxRate = 20
		cessRate = 30
	case totalEarnings > 90000:
		taxRate = 30
		cessRate = 40
	}

	instEmployeeStruct.salaryStruct.salaryDeductions.cess = uint16(float32(instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings * cessRate / 100))
	instEmployeeStruct.salaryStruct.salaryDeductions.taxes = uint16(float32(instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings * taxRate / 100))

}

func (instEmployeeStruct *EmployeeStruct) GetDeductions() uint16 {
	fmt.Println("total earnings: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings))
	instEmployeeStruct.getTaxes()

	fmt.Println("cess: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryDeductions.cess))
	fmt.Println("Taxes: " + fmt.Sprint(instEmployeeStruct.salaryStruct.salaryDeductions.taxes))

	instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions = (instEmployeeStruct.salaryStruct.salaryDeductions.taxes + instEmployeeStruct.salaryStruct.salaryDeductions.cess)
	// fmt.Println("total reduction: ", instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions)
	return instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions
}

func (instEmployeeStruct EmployeeStruct) GetSalarySlip() {
	instEmployeeStruct.GetNetSalary()
	instEmployeeStruct.GetDeductions()

}

type Tabler interface {
	TableName() string
}

func (instEmployeeStruct EmployeeStruct) Tabler() string {
	return "Employees"
}

func RegisterEmployee(app *fiber.Ctx) error {
	var responseMsg string
	instEmployeeStruct := new(EmployeeStruct)
	err := app.BodyParser(instEmployeeStruct)
	log.Info("name: " + instEmployeeStruct.Name)
	if err != nil {
		responseMsg = "Err" + err.Error()
		return app.SendString(responseMsg)
	}
	responseMsg = "This is from /register post method" + fmt.Sprint(instEmployeeStruct.Name)
	log.Info("return srting" + responseMsg)
	return app.SendString(responseMsg)
}
