package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:                "/posts",
		Method:             http.MethodPost,
		Function:           controllers.CreatePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts",
		Method:             http.MethodGet,
		Function:           controllers.GetPosts,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postID}",
		Method:             http.MethodGet,
		Function:           controllers.GetPost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postID}",
		Method:             http.MethodPut,
		Function:           controllers.UpdatePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postID}",
		Method:             http.MethodDelete,
		Function:           controllers.DeletePost,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}/posts",
		Method:             http.MethodGet,
		Function:           controllers.GetPostsByUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postID}/like",
		Method:             http.MethodPost,
		Function:           controllers.Like,
		NeedAuthentication: true,
	},
	{
		URI:                "/posts/{postID}/dislike",
		Method:             http.MethodPost,
		Function:           controllers.Dislike,
		NeedAuthentication: true,
	},
}
