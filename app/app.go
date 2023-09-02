package app

type AuthRegister struct {
	Id       string `valid:"required" json:"id"`
	Username string `valid:"required" json:"username"`
	Email    string `valid:"required,email" json:"email"`
	Password string `valid:"required,minstringlength(6)" json:"password"`
}

type AuthLogin struct {
	Email    string `valid:"required,email" json:"email"`
	Password string `valid:"required,minstringlength(6)" json:"password"`
}
