package models

type User struct { //create a model for database
	ID      uint   `json:"id" gorm:"primary_key"` // with tags we can specify json name
	Name    string `json:"name"`
	Contact uint   `json:"contact"`
	Address string `json:"address"`
}

type CreateUser struct {
	Name    string `json:"name"`
	Contact uint   `json:"contact"`
	Address string `json:"address" binding:"required"` // binging required - this is mendatory
}

type UpdateUser struct { // we make this not compulsary to give
	Name    string `json:"name"`
	Contact uint   `json:"contact"`
	Address string `json:"address"`
}
