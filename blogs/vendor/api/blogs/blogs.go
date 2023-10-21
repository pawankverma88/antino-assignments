package blogs

import (
	"data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var StatusBadRequest = 400          // StatusBadRequest is an HTTP status code indicating a client error (400 Bad Request).
var StatusNotFound = 404            // StatusNotFound is an HTTP status code indicating that the requested resource was not found (404 Not Found).
var StatusInternalServerError = 500 // StatusInternalServerError is an HTTP status code indicating a server error (500 Internal Server Error).
var StatusOK = 200                  // StatusOK is an HTTP status code indicating a successful request (200 OK).

// CreateBlog :
func CreateBlog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//------------------ Check Input Body Part -------------------//

	if r.Body == nil {
		data.RespondJSON(w, StatusBadRequest, "Invalid request. Please review your input data.")
	}

	//------------------ Decode Input Json And move Into Struct -------------------//

	decoder := json.NewDecoder(r.Body)
	var objInputBlog InputPayload
	var err error

	err = decoder.Decode(&objInputBlog)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "Invalid request - please check your input data."+err.Error())
		return
	}

	//------------------ clean/Sanitize Group Input Fields -------------------//

	objInputBlog.Author = data.CleanText(objInputBlog.Author, false, true)
	objInputBlog.Title = data.CleanText(objInputBlog.Title, false, true)
	objInputBlog.Content = data.CleanText(objInputBlog.Content, false, true)

	//------------------ Validate All Input Fields Start -------------------//

	//------------------ Check Blog Author Empty or Not -------------------//

	if objInputBlog.Author == "" {
		data.RespondJSON(w, StatusBadRequest, "Author is required. Please provide the name of the author for the blog post.")
		return
	}

	//------------------ Check Blog Title Empty or Not -------------------//

	if objInputBlog.Title == "" {
		data.RespondJSON(w, StatusBadRequest, "Title is required. Please provide the title for the blog post")
		return
	}

	//------------------ Check Blog Content Empty or Not -------------------//

	if objInputBlog.Content == "" {
		data.RespondJSON(w, StatusBadRequest, "Content is required. Please provide the content for the blog post")
		return
	}

	//------------------ Validate All Input Fields End -------------------//

	//------------------ Inserting Blog Details -------------------//

	err = insertBlog(objInputBlog)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "An error occurred while fetching blogs :"+err.Error()+" Please check your request and try again")
		return
	}

	data.RespondJSON(w, StatusOK, "Blog Successfully Added")
	return

}

// GetAllBlogs : retrieves a list of all blogs that have been inserted with a 'status' of 'Y'.
func GetAllBlogs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//------------------ Getting Blogs Details -------------------//

	objAllBlogs, err := getBlogsDetails("all")
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "Error while fetching Blogs :"+err.Error())
		return
	}

	//------------------ Validate Fetched Blogs Data Exist or Not -------------------//
	if len(objAllBlogs) == 0 {
		data.RespondJSON(w, StatusNotFound, "Blogs Not Found")
		return
	}

	//------------------ Binding Blogs Details -------------------//

	finalBlogs := []OutputPayload{}
	for _, singleBlog := range objAllBlogs {
		blog := OutputPayload{}
		blog.BlogID = singleBlog.BlogID.Int32
		blog.Author = singleBlog.Author.String
		blog.Title = singleBlog.Title.String
		blog.Content = singleBlog.Content.String
		blog.CreationDate = singleBlog.CreationDate.String
		blog.LastModifiedDate = singleBlog.LastModifiedDate.String
		finalBlogs = append(finalBlogs, blog)
	}

	//------------------ Validate Final binded Blogs Data Exist or Not -------------------//
	if len(finalBlogs) == 0 {
		data.RespondJSON(w, StatusNotFound, "Blogs Not Found")
		return
	}

	data.RespondJSONObject(w, StatusOK, finalBlogs)
	return
}

// GetBlog : retrieves a specific blog with the requested blog id, but only if its 'status' is set to 'Y'.
func GetBlog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//------------------ Sanitize Param -------------------//

	blogID := data.CleanText(p.ByName("blogID"), false, true)

	//------------------ Validate Blog ID Empty or Not -------------------//

	if blogID == "" {
		data.RespondJSON(w, StatusBadRequest, "Blog ID is required. Please provide a valid blog ID to retrieve the blog.")
		return
	}

	//------------------ Getting Blog Details As Requested -------------------//

	objAllBlogs, err := getBlogsDetails(blogID)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "An error occurred while fetching blogs :"+err.Error()+" Please check your request and try again")
		return
	}

	//------------------ Validate Fetched Blog Data Exist or Not -------------------//

	if len(objAllBlogs) == 0 {
		data.RespondJSON(w, StatusNotFound, "No blog found with the specified identifier. Please verify the blog ID and try again.")
		return
	}

	//------------------ Binding Blog Details -------------------//

	finalBlogs := OutputPayload{}
	for _, singleBlog := range objAllBlogs {

		finalBlogs.BlogID = singleBlog.BlogID.Int32
		finalBlogs.Author = singleBlog.Author.String
		finalBlogs.Title = singleBlog.Title.String
		finalBlogs.Content = singleBlog.Content.String
		finalBlogs.CreationDate = singleBlog.CreationDate.String
		finalBlogs.LastModifiedDate = singleBlog.LastModifiedDate.String

	}

	//------------------ Validate Final binded Blog Data Exist or Not -------------------//

	if (finalBlogs == OutputPayload{}) {
		data.RespondJSON(w, StatusNotFound, "No blogs were found")
		return
	}

	data.RespondJSONObject(w, StatusOK, finalBlogs)
	return
}

// UpdateBlog : updates the information of an existing blog with new data
func UpdateBlog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//------------------ Check Input Body Part -------------------//

	if r.Body == nil {
		data.RespondJSON(w, StatusBadRequest, "Invalid request. Please review your input data.")
	}

	//------------------ Decode Input Json And move Into Struct -------------------//

	decoder := json.NewDecoder(r.Body)
	var objInputBlog InputPayload
	var err error

	err = decoder.Decode(&objInputBlog)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "Invalid request - please check your input data."+err.Error())
		return
	}

	//------------------ clean/Sanitize Group Input Fields -------------------//

	objInputBlog.BlogID = data.CleanText(objInputBlog.BlogID, false, true)
	objInputBlog.Author = data.CleanText(objInputBlog.Author, false, true)
	objInputBlog.Title = data.CleanText(objInputBlog.Title, false, true)
	objInputBlog.Content = data.CleanText(objInputBlog.Content, false, true)

	//------------------ Validate All Input Fields Start -------------------//

	//------------------ Check Blog Id Empty or Not -------------------//

	if objInputBlog.BlogID == "" {
		data.RespondJSON(w, StatusBadRequest, "blogID is required. Please provide the blogID for the blog post.")
		return
	}

	//------------------ Check Blog Author Empty or Not -------------------//
	if objInputBlog.Author == "" {
		data.RespondJSON(w, StatusBadRequest, "Author is required. Please provide the name of the author for the blog post.")
		return
	}

	//------------------ Check Blog Title Empty or Not -------------------//

	if objInputBlog.Title == "" {
		data.RespondJSON(w, StatusBadRequest, "Title is required. Please provide the title for the blog post")
		return
	}

	//------------------ Check Blog Content Empty or Not -------------------//

	if objInputBlog.Content == "" {
		data.RespondJSON(w, StatusBadRequest, "Content is required. Please provide the content for the blog post")
		return
	}

	//------------------ Check Blog Exist or Not as requested -------------------//

	objBlog, err := getBlogsDetails(objInputBlog.BlogID)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "An error occurred while fetching blogs :"+err.Error()+" Please check your request and try again")
		return
	}

	if len(objBlog) == 0 {
		data.RespondJSON(w, StatusNotFound, "No blog found with the specified identifier. Please verify the blog ID and try again")
		return
	}

	//------------------ Validate All Input Fields End -------------------//

	//------------------ Update Blog Details -------------------//

	err = updateBlog(objInputBlog)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "An error occurred while fetching blogs :"+err.Error()+" Please check your request and try again")
		return
	}

	data.RespondJSON(w, StatusOK, "Blog Update Successfully")
	return

}

// DeleteBlog : deletes a blog with the specified identifier from the database
func DeleteBlog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//------------------ Sanitize Param -------------------//

	blogID := data.CleanText(p.ByName("blogID"), false, true)

	//------------------ Validate Blog ID Empty or Not -------------------//

	if blogID == "" {
		data.RespondJSON(w, StatusBadRequest, "Blog ID is required. Please provide a valid blog ID to retrieve the blog.")
		return
	}

	//------------------ Validate Blog ID Should Be Numaric value -------------------//

	intBlogID, err := strconv.Atoi(blogID)
	if err != nil {
		data.RespondJSON(w, StatusBadRequest, "Blog ID is required. Please provide a valid blog ID to retrieve the blog.")
		return
	}

	//------------------ Getting Blog Details As Requested BlogID -------------------//

	objBlog, err := getBlogsDetails(blogID)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "An error occurred while fetching blogs :"+err.Error()+" Please check your request and try again")
		return
	}

	//------------------ Verfiy BlogID exist or Not -------------------//
	if len(objBlog) == 0 {
		data.RespondJSON(w, StatusBadRequest, "No blog found with the specified identifier. Please verify the blog ID and try again.")
		return
	}

	//------------------ Validate All Input Fields End -------------------//

	//------------------ Inserting Blog Details -------------------//

	err = deleteBlog(intBlogID)
	if err != nil {
		data.RespondJSON(w, StatusInternalServerError, "An error occurred while delete blog :"+err.Error()+" Please check your request and try again")
		return
	}

	data.RespondJSON(w, StatusOK, "Blog Removed Successfully")
	return

}
