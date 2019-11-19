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
	Code        string   `json:"code" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name,omitempty"`
	Description string   `json:"description" bson:"description,omitempty"`
	Steps       []string `json:"steps" bson:"steps,omitempty"`
	Ingredients []string `json:"ingredients" bson:"ingredients,omitempty"`
}

// GetCode function
func (r *Recipe) GetCode() string {
	return r.Code
}

// SetCode function
func (r *Recipe) SetCode(code string) {
	r.Code = code
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
func (r *Recipe) SetIngredients(ingredients []string) {
	r.Ingredients = ingredients
}

// GetIngredients function
func (r *Recipe) GetIngredients() []string {
	return r.Ingredients
}

// GetObjectInfo - Interface DTOObject Implementation
func (r *Recipe) GetObjectInfo() string {
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
