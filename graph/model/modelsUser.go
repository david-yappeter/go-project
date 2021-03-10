package model

//User User Struct
type User struct {
	ID                      int     `gorm:"type:int;AUTO_INCREMENT;not null"`
	Name                    string  `gorm:"type:text;not null"`
	Password                string  `gorm:"type:text;not null"`
	Email                   string  `gorm:"type:text;not null"`
	Address                 *string `gorm:"type:text;null;default:null"`
	TelephoneNumber         *string `gorm:"type:text;null;default:null"`
	CreatedAt               string  `gorm:"type:timestamp;not null"`
	UpdatedAt               *string `gorm:"type:timestamp;null;default:null"`
	DeletedAt               *string `gorm:"type:timestamp;null;default:null"`
	Avatar                  *string `gorm:"type:text;null;default:null"`
	IsAdmin                 int     `gorm:"type:int;not null;default: 0"`
	EmailVerificationHash   *string `gorm:"type:text;null;default:null"`
	EmailVerificationStatus int     `gorm:"type:int;not null;default:0"`
	GoogleID                *string `gorm:"type:text;null;default:null"`
	LocationCode            *string `gorm:"type:text;null;default:null"`
}
