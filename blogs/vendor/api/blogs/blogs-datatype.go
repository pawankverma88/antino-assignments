package blogs

import "database/sql"

// InputPayload :
type InputPayload struct {
	BlogID  string `json:"blog_id,omitempty"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DBBlogs :
type DBBlogs struct {
	BlogID           sql.NullInt32
	Author           sql.NullString
	Title            sql.NullString
	Content          sql.NullString
	CreationDate     sql.NullString
	LastModifiedDate sql.NullString
	Status           sql.NullString
}

// OutputPayload :
type OutputPayload struct {
	BlogID           int32  `json:"blog_id"`
	Author           string `json:"author"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	CreationDate     string `json:"creation_date,omitempty"`
	LastModifiedDate string `json:"last_modified_date,omitempty"`
	Status           string `json:"status,omitempty"`
}
