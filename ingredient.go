package hrstypes

import (
	"strconv"
	"strings"
)

/** INGREDIENT DEFINITION **/

// Ingredient DTO
type Ingredient struct {
	Code        string `json:"code" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Quantity    int    `json:"quantity" bson:"quantity"`
}

// GetObjectInfo - Interface DTOObject Implementations
func (i *Ingredient) GetObjectInfo() string {
	info := []string{
		i.Name,
		i.Description,
	}
	resp := strings.Join(info, ": ")
	quantity := "\nQuantity: " + strconv.Itoa(i.Quantity)
	resp += quantity
	return resp
}

// GetCode function
func (i *Ingredient) GetCode() string {
	return i.Code
}

// SetCode function
func (i *Ingredient) SetCode(code string) {
	i.Code = code
}

// GetName function
func (i *Ingredient) GetName() string {
	return i.Name
}

// SetName function
func (i *Ingredient) SetName(name string) {
	i.Name = name
}

// GetDescription function
func (i *Ingredient) GetDescription() string {
	return i.Description
}

// SetDescription function
func (i *Ingredient) SetDescription(description string) {
	i.Description = description
}

// GetQuantity function
func (i *Ingredient) GetQuantity() int {
	return i.Quantity
}

// SetQuantity function
func (i *Ingredient) SetQuantity(quantity int) {
	i.Quantity = quantity
}
