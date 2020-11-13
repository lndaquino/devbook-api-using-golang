package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Function:           controllers.SearchUsers,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}",
		Method:             http.MethodGet,
		Function:           controllers.SearchUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}/follow",
		Method:             http.MethodPost,
		Function:           controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}/unfollow",
		Method:             http.MethodPost,
		Function:           controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}/followers",
		Method:             http.MethodGet,
		Function:           controllers.SearchUserFollowers,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}/following",
		Method:             http.MethodGet,
		Function:           controllers.FollowingUsers,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userID}/updatePassword",
		Method:             http.MethodPost,
		Function:           controllers.UpdatePassword,
		NeedAuthentication: true,
	},
}
