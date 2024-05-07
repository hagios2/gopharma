package services

type RegistrationRequest struct {
	FirstName    string `form:"user" json:"firstname"  binding:"required"`
	LastName     string `form:"lastname" json:"lastname" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required,email"`
	CompanyName  string `form:"company_name" json:"company_name" binding:"required"`
	CompanyEmail string `form:"company_email" json:"company_email" binding:"required,email"`
	CompanyPhone string `form:"company_phone" json:"company_phone" binding:"required,e164"`
}
