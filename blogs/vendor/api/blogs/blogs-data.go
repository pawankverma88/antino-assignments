package blogs

import (
	"data"
	"strconv"
	"time"
)

// insertBlog :
func insertBlog(objInputBlog InputPayload) error {

	lastModifiedDate := time.Now().In(data.IST).Format("2006-01-02 15:04:05")

	sqlStr := "INSERT INTO tbl_blogs (author, title, content, last_modified_date) VALUES"
	valueStr := "('" + objInputBlog.Author + "','" + objInputBlog.Title + "', '" + objInputBlog.Content + "', '" + lastModifiedDate + "')"

	sqlStr = sqlStr + valueStr

	stmt, err := data.BlogDB.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// getBlogsDetails :
func getBlogsDetails(blogID string) ([]DBBlogs, error) {
	objBlogs := []DBBlogs{}
	var seasonStr string
	if blogID != "" && blogID != "all" {
		intBlogID, err := strconv.Atoi(blogID)
		if err != nil {
			return nil, err
		}

		seasonStr = " AND blog_id = '" + strconv.Itoa(intBlogID) + "' "
	}

	sqlstr := "SELECT blog_id, author, title, content, creation_date, last_modified_date, status" +
		" FROM tbl_blogs" +
		" WHERE status=?" + seasonStr +
		" ORDER BY creation_date DESC, last_modified_date DESC"

	rows, err := data.BlogDB.Query(sqlstr, "Y")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		singleBlog := DBBlogs{}
		err := rows.Scan(
			&singleBlog.BlogID,
			&singleBlog.Author,
			&singleBlog.Title,
			&singleBlog.Content,
			&singleBlog.CreationDate,
			&singleBlog.LastModifiedDate,
			&singleBlog.Status,
		)
		if err != nil {
			return nil, err
		}
		objBlogs = append(objBlogs, singleBlog)
	}
	return objBlogs, nil
}

// updateBlog :
func updateBlog(objInputBlog InputPayload) error {

	lastModifiedDate := time.Now().In(data.IST).Format("2006-01-02 15:04:05")

	sqlStr := "UPDATE tbl_blogs SET author = ?, title = ?, content = ?, last_modified_date = ? WHERE blog_id = ?"

	stmt, err := data.BlogDB.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(objInputBlog.Author, objInputBlog.Title, objInputBlog.Content, lastModifiedDate, objInputBlog.BlogID)
	if err != nil {
		return err
	}
	return nil
}

// deleteBlog :
func deleteBlog(blogID int) error {

	sqlStr := "DELETE FROM tbl_blogs WHERE blog_id = ?"

	stmt, err := data.BlogDB.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(blogID)
	if err != nil {
		return err
	}
	return nil
}
