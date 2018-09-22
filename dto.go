package hrstypes

import (
	"strconv"
	"strings"
)

/**
** RECIPE DEFINITION
 */

// Recipe DTO
type Recipe struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Steps       []string     `json:"steps"`
	Ingredients []Ingredient `json:"ingredients"`
}

// GetID function
func (r *Recipe) GetID() string {
	return r.ID
}

// SetID function
func (r *Recipe) SetID(id string) {
	r.ID = id
}

// GetName function
func (r *Recipe) GetName() string {
	return r.Name
}

// SetName function
func (r *Recipe) SetName(name string) {
	r.Name = name
}

// GetDescription function
func (r *Recipe) GetDescription() string {
	return r.Description
}

// SetDescription function
func (r *Recipe) SetDescription(description string) {
	r.Description = description
}

// GetSteps function
func (r *Recipe) GetSteps() []string {
	return r.Steps
}

// SetSteps function
func (r *Recipe) SetSteps(steps []string) {
	r.Steps = steps
}

// SetIngredients function
func (r *Recipe) SetIngredients(ingredients []Ingredient) {
	r.Ingredients = ingredients
}

// GetIngredients function
func (r *Recipe) GetIngredients() []Ingredient {
	return r.Ingredients
}

// Interface ResponseObject Implementation
func (r *Recipe) getObjectInfo() string {
	info := []string{
		r.GetName(),
		r.GetDescription(),
	}
	resp := strings.Join(info, ": ")

	for i, step := range r.GetSteps() {
		resp += "\nStep " + strconv.Itoa(i) + ":" + step
	}

	return resp
}

/** INGREDIENT DEFINITION **/

// Ingredient DTO
type Ingredient struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

// Interface ResponseObject Implementations
func (i *Ingredient) getObjectInfo() string {
	info := []string{
		i.Name,
		i.Description,
	}
	resp := strings.Join(info, ": ")
	quantity := "\nQuantity: " + strconv.Itoa(i.Quantity)
	resp += quantity
	return resp
}

// GetID function
func (i *Ingredient) GetID() string {
	return i.ID
}

// SetID function
func (i *Ingredient) SetID(id string) {
	i.ID = id
}

func (i *Ingredient) getName() string {
	return i.Name
}

func (i *Ingredient) setName(name string) {
	i.Name = name
}

func (i *Ingredient) getDescription() string {
	return i.Description
}

func (i *Ingredient) setDescription(description string) {
	i.Description = description
}

func (i *Ingredient) getQuantity() int {
	return i.Quantity
}

func (i *Ingredient) setQuantity(quantity int) {
	i.Quantity = quantity
}
