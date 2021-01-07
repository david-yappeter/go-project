package model

//FileUpload File Upload Struct
type FileUpload struct {
	ID        int     `gorm:"type:int;AUTO_INCREMENT;not null"`
	FileID    string  `gorm:"type:text;not null"`
	CreatedAt string  `gorm:"type:timestamp;not null"`
	UpdatedAt *string `gorm:"type:timestamp;null;default:null"`
	DeletedAt *string `gorm:"type:timestamp;null;default:null"`
	UserID    int     `gorm:"type:int;not null"`
}

//NewFileUpload New File Upload Struct
type NewFileUpload struct {
	FileID string `json:"file_id"`
	UserID int    `json:"user_id"`
}

//UpdateFileUpload Update File Upload Struct
type UpdateFileUpload struct {
	ID     int    `json:"id"`
	FileID string `json:"file_id"`
	UserID int    `json:"user_id"`
}
