package model

//IgPost IgPost Struct
type IgPost struct {
	ID         int           `gorm:"type:int;AUTO_INCREMENT;not null"`
	Caption    string        `gorm:"type:text;not null"`
	CreatedAt  string        `gorm:"type:timestamp;not null"`
	UpdatedAt  *string       `gorm:"type:timestamp;null;default:null"`
	DeletedAt  *string       `gorm:"type:timestamp;null;default:null"`
	IsArchived int           `gorm:"type:int;not null;default:0"`
	UserID     int           `gorm:"type:int;not null"`
	Files      []*IgPostFile `gorm:"-"`
}
