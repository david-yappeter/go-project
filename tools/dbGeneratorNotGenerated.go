package tools

import gorm "gorm.io/gorm"

//DeletedAt Filtering By Deleted At
func DeletedAt(query *gorm.DB, scopes *bool) {
	if scopes != nil && *scopes == true {
		query.Where("deleted_at IS NULL")
	}
}

//FilterByUserID Filter By User ID
func FilterByUserID(query *gorm.DB, userID *int) {
	if userID != nil {
		query.Where("user_id = ?", *userID)
	}
}
