package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user on database
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	dbConn, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	usersRepository := repositories.NewUserRepository(dbConn)
	userID, err := usersRepository.CreateUser(user)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	user.ID = userID
	user.Password = ""
	responses.JSON(rw, http.StatusCreated, user)
}

// SearchUsers searches all users on database
func SearchUsers(rw http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))

	dbConn, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
	}
	defer dbConn.Close()

	usersRepository := repositories.NewUserRepository(dbConn)
	users, err := usersRepository.SearchUsers(nameOrNickname)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, users)
}

// SearchUser search an user on database
func SearchUser(rw http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}
	dbConn, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
	}
	defer dbConn.Close()

	repositories := repositories.NewUserRepository(dbConn)
	user, err := repositories.SearchUserById(userID)
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(rw, http.StatusOK, user)
}

// UpdateUser updates an user on database
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(rw, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	dbConn, err := database.Connect()
	if err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
		return
	}

	repositories := repositories.NewUserRepository(dbConn)
	if err = repositories.UpdateUser(userID, user); err != nil {
		responses.Error(rw, http.StatusInternalServerError, err)
	}

	responses.JSON(rw, http.StatusNoContent, nil)
}

// DeleteUser deletes an user from database
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Deleting User User..."))
}
