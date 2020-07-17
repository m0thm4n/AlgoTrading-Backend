package Controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/m0thm4n/AlgoTrading-Backend/auth"
	"github.com/m0thm4n/AlgoTrading-Backend/models"
	"github.com/m0thm4n/AlgoTrading-Backend/Responses"
	"github.com/m0thm4n/AlgoTrading-Backend/Utils/formaterror"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
}