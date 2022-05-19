package main

import (
	"errors"
	"fmt"
)

func main() {

}

type Account struct {
	LastName  string
	FirstName string
}

func (a *Account) changeName(newname string) {
	a.FirstName = newname
}

type Employee struct {
	Account
	Credits float64
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}
func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("Invaliad credit amount")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has.")
	}
	return 0.0, errors.New("You can't remove negative numbers.")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}
