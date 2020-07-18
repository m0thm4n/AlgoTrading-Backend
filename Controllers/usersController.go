package Controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/m0thm4n/AlgoTrading-Backend/auth"
	"github.com/m0thm4n/AlgoTrading-Backend/Models"
	"github.com/m0thm4n/AlgoTrading-Backend/Responses"
	"github.com/m0thm4n/AlgoTrading-Backend/Utils/formaterror"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := Models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		Responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	Responses.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := Models.User{}

	users, err := user.FindAllUser(server.DB)
	if err != nil {
		Responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	Responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strcov.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := Models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	Responses.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := Models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != uint32(uid) {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	user.Prepare()

	err = user.Validate("update")
	if err != nil {
		Responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		Responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	Responses.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := Models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		Responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	if tokenID != 0 && tokenID != uint32(uid) {
		Responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		Responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	Responses.JSON(w, http.StatusNoContent, "")
}