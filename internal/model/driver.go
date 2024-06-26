package model

import (
	"github.com/google/uuid"
)

type Category string

const (
	Junior  Category = "junior"
	Amateur Category = "amateur"
	Pro     Category = "pro"
)

type Driver struct {
	Id          uuid.UUID
	Name        string
	PhoneNumber string
	Age         *int
	Weight      *int
	Category    Category
}

func NewDriver(
	id uuid.UUID,
	name string,
	phoneNumber string,
	age *int,
	weight *int,
) *Driver {
	d := &Driver{
		Id:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
		Age:         age,
		Weight:      weight,
	}

	d.defineCategory()

	return d
}

func (d *Driver) defineCategory() {
	if d.Age != nil && *d.Age < 18 {
		d.Category = Junior

		return
	}

	d.Category = Amateur
}

func (d *Driver) Update(
	name string,
	phoneNumber string,
	age *int,
	weight *int,
) {
	d.Name = name
	d.PhoneNumber = phoneNumber
	d.Age = age
	d.Weight = weight

	d.defineCategory()
}
