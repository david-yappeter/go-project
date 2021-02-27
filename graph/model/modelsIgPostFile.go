package model

//IgPostFile IgPostFile Struct
type IgPostFile struct {
	ID     int    `json:"id"`
	FileID string `json:"file_id"`
	PostID int    `json:"post_id"`
}
