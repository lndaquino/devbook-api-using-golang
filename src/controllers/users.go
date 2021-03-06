package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("create"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers searchs users by name or nick
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// SearchUser searchs an user by Id
func SearchUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	user, err := repository.SearchByID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates an user infos by Id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDfromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDfromToken {
		responses.Error(w, http.StatusForbidden, errors.New("Invalid user"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Update(userID, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes an user by Id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDfromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDfromToken {
		responses.Error(w, http.StatusForbidden, errors.New("Invalid user"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Delete(userID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func getUserID(r *http.Request) (uint64, error) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)

	return userID, err
}

// FollowUser allows an user to follow other one
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Unable to follow your own id"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Follow(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser allows to unfollow an user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.Error(w, http.StatusForbidden, errors.New("Unable to unfollow your own id"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Unfollow(userID, followerID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

//SearchUserFollowers search all followers from an user
func SearchUserFollowers(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	followers, err := repository.SearchFollowers(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// FollowingUsers returns the users an user is following
func FollowingUsers(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.Following(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// UpdatePassword updates an user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDfromToken, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if userIDfromToken != userID {
		responses.Error(w, http.StatusForbidden, errors.New("Unable to update password of other user"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	var password models.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	savedPassword, err := repository.SearchPassword(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPassword(password.Current, savedPassword); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Invalid password"))
		return
	}

	hashedPassword, err := security.Hash(password.New)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
