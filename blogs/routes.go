package main

import (
	"api/blogs"

	"github.com/julienschmidt/httprouter"
)

// addRouteHandlers adds routes for various APIs.
func addRouteHandlers(router *httprouter.Router) {

	router.POST("/api/blogs/create", blogs.CreateBlog)
	router.GET("/api/blogs", blogs.GetAllBlogs)
	router.GET("/api/blog/:blogID", blogs.GetBlog)
	router.POST("/api/blog/update", blogs.UpdateBlog)
	router.POST("/api/blog/delete/:blogID", blogs.DeleteBlog)

}
