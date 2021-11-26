package entity

// book struct represent book table
type Company struct {
	ID             uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CompanyName    string `gorm:"type:varchar(255)" json:"companyName"`
	CompanyProfile string `gorm:"type:text" json:"companyProfile"`
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignkey:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
