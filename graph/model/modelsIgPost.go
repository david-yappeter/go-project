package model

//IgPost IgPost Struct
type IgPost struct {
	ID         int           `json:"id"`
	Caption    string        `json:"caption"`
	Files      []*IgPostFile `json:"files"`
	CreatedAt  *string       `json:"created_at"`
	UpdatedAt  *string       `json:"updated_at"`
	DeletedAt  *string       `json:"deleted_at"`
	IsArchived int           `json:"is_archived"`
	UserID     int           `json:"user_id"`
}
