package model

// Binding from JSON
type Todo struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" binding:"required"`
	Completed bool   `json:"completed"`
}
