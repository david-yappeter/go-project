package tools

import gorm "gorm.io/gorm"

//QueryMaker Pagination Query Tools
func QueryMaker(query *gorm.DB, limit *int, page *int, ascending *bool, sortBy *string, scopes *bool) {
	sortBy = OrderBy(sortBy, ascending)

	if limit != nil && page != nil && *limit > 0 && *page > 0 {
		offset := (*page - 1) * *limit
		*query = *query.Offset(offset).Limit(*limit)
	}

	DeletedAt(query, scopes)

	if sortBy != nil {
		*query = *query.Order(*sortBy)
	}
}

//OrderBy Order By Generator (ASC,DESC)
func OrderBy(sortBy *string, ascending *bool) *string {
	if sortBy == nil {
		return nil
	}
	if ascending != nil {
		if *ascending {
			sortSyntax := *sortBy + " ASC"
			return &sortSyntax
		}
		sortSyntax := *sortBy + " DESC"
		return &sortSyntax
	}
	return nil
}

//DeletedAt Filtering By Deleted At
func DeletedAt(query *gorm.DB, scopes *bool) {
	if scopes == nil || *scopes != true {
		query.Where("deleted_at IS NULL")
	}
}

//FilterByUserID Filter By User ID
func FilterByUserID(query *gorm.DB, userID *int) {
	if userID != nil {
		query.Where("user_id = ?", *userID)
	}
}
