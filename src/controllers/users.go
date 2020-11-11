package controllers

import "net/http"

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser"))
}

// SearchAllUsers searchs all users
func SearchAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearchAllUsers"))
}

// SearchUser searchs an user by Id
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SearchUser"))
}

// UpdateUser updates an user infos by Id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

// DeleteUser deletes an user by Id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}
