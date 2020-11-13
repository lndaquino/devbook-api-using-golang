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
}
