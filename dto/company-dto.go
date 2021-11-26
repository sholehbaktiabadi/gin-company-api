package dto

type CompanyUpdateDTO struct {
	ID             uint64 `json:"id" form:"binding"`
	CompanyName    string `json:"companyName" form:"companyName" binding:"required"`
	CompanyProfile string `json:"companyProfile" form:"companyProfile" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type CompanyCreateDTO struct {
	CompanyName    string `json:"companyName" form:"companyName" binding:"required"`
	CompanyProfile string `json:"companyProfile" form:"companyProfile" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
