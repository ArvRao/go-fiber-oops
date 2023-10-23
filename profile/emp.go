package Employee

import (
	"fmt"
	"strconv"
)

const (
	FULLTIME uint8 = iota
	CONTRACT
	PARTTIME
	OUTSOURCE
)

type EmployeeStruct struct {
	id           uint16
	name         string
	salaryStruct SalaryStruct
}

type SalaryStruct struct {
	salaryEarnings   SalaryEarningsStruct
	salaryDeductions SalaryDeductionsStruct
}

type SalaryEarningsStruct struct {
	base                uint16
	hra                 uint16
	specialAllowance    uint16
	leaveAllowance      uint16
	nightshiftAllowance uint16
	totalEarnings       uint16
}

type SalaryDeductionsStruct struct {
	gratuity        uint16
	providentFund   uint16
	taxes           uint16
	salaryCess      uint16
	totalDeductions uint16
}

// collection of signature methods for class
type EmployeeInf interface {
	employeeNew(name string, baseSalary uint16, specialAllowance uint16, leaveAllowance uint16, nightshiftAllowance uint16)
	getProfileMthd() (uint8, string)
	getDeductionsMthd() uint16
	getGrossSalaryMthd() uint16
	getNetSalaryMthd() uint16
}

// function that takes necessary data in order to create object/instance. (receiver function)
func EmployeeNew(name string, baseSalary uint16, specialAllowance uint16, leaveAllowance uint16, nightshiftAllowance uint16) EmployeeStruct {
	fmt.Println(baseSalary)
	var instEmployeeStruct EmployeeStruct
	instEmployeeStruct.name = name
	instEmployeeStruct.salaryStruct.salaryEarnings.base = baseSalary
	instEmployeeStruct.salaryStruct.salaryEarnings.specialAllowance = specialAllowance
	instEmployeeStruct.salaryStruct.salaryEarnings.leaveAllowance = leaveAllowance
	instEmployeeStruct.salaryStruct.salaryEarnings.nightshiftAllowance = nightshiftAllowance
	return instEmployeeStruct
}

func (instEmployeeStruct EmployeeStruct) GetProfileMthd() string {
	fmt.Println(instEmployeeStruct.salaryStruct.salaryEarnings.base)
	return ("Name of employee is " + instEmployeeStruct.name + " with base salary " + strconv.Itoa(int((instEmployeeStruct.salaryStruct.salaryEarnings.base))))
}

func (instEmployeeStruct EmployeeStruct) GetGrossSalaryMthd() uint16 {
	return instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings - instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions
}

func (instEmployeeStruct EmployeeStruct) GetNetSalaryMthd() uint16 {
	instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings = instEmployeeStruct.salaryStruct.salaryEarnings.base + instEmployeeStruct.salaryStruct.salaryEarnings.hra + instEmployeeStruct.salaryStruct.salaryEarnings.leaveAllowance + instEmployeeStruct.salaryStruct.salaryEarnings.nightshiftAllowance + instEmployeeStruct.salaryStruct.salaryEarnings.specialAllowance
	return instEmployeeStruct.salaryStruct.salaryEarnings.totalEarnings
}

func (instEmployeeStruct EmployeeStruct) GetDeductionsMthd() uint16 {
	instEmployeeStruct.salaryStruct.salaryDeductions.gratuity = (10 * (instEmployeeStruct.salaryStruct.salaryEarnings.base)) / 100
	instEmployeeStruct.salaryStruct.salaryDeductions.providentFund = (20 * (instEmployeeStruct.salaryStruct.salaryEarnings.base)) / 100
	instEmployeeStruct.salaryStruct.salaryDeductions.salaryCess = 0
	instEmployeeStruct.salaryStruct.salaryDeductions.taxes = 0
	instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions = (instEmployeeStruct.salaryStruct.salaryDeductions.gratuity + instEmployeeStruct.salaryStruct.salaryDeductions.providentFund + instEmployeeStruct.salaryStruct.salaryDeductions.taxes + instEmployeeStruct.salaryStruct.salaryDeductions.salaryCess)
	// fmt.Println("total reduction: ", instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions)
	return instEmployeeStruct.salaryStruct.salaryDeductions.totalDeductions
}
