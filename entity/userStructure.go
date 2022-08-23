package entity

type User struct {
	Id       int
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Rating   string `json:"rating"`
}

type InputSignIn struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProfileData struct {
	Name   string
	Phone  string
	Email  string
	Rating float32
}
