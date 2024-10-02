package users

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"sarath/perf_testing/internal/data"
	"sarath/perf_testing/internal/json"
	"sarath/perf_testing/internal/json/validator"
	"sarath/perf_testing/internal/logger"
	"sarath/perf_testing/internal/response"

	"github.com/gorilla/mux"
)

type Handler struct {
	logger logger.ApplicationLogger
	db     *data.Models
}

func New(log logger.ApplicationLogger, db *data.Models) *Handler {
	return &Handler{
		logger: log,
		db:     db,
	}
}

func (handler *Handler) RegisterUsers(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	writer := response.New(handler.logger)

	// read the json
	err := json.ReadJsonFromReq(&input, w, r)
	if err != nil {
		handler.logger.Log(fmt.Sprint(err))
		writer.ErrResponse(err, w)
		return
	}

	// validate the json
	v := validator.New()
	v.Check(len(input.Name) > 3, "name", "The name must be greater than 3 length")
	v.Check(len(input.Name) < 32, "new_url", "The name must be less than 32")
	v.Check(v.Matches(input.Email, validator.EmailRX), "new_url", "Only characters and digits are allowed")

	if !v.Valid() {
		writer.ValidationErrorResponse(v, w)
		return
	}

	// Try inserting into the db
	user := data.User{
		UserName: input.Name,
		Email:    input.Email,
	}
	handler.logger.Log(fmt.Sprint("Username: ", input.Name, " Email: ", input.Email))
	err = handler.db.Users.Insert(&user)
	if err != nil {
		writer.ErrResponse(err, w)
		handler.logger.Log(err.Error())
		return
	}

	// return the created response
	writer.CreatedResponse(json.Envelope{
		"data": user,
	}, w)
}

func (app *Handler) FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	writer := response.New(app.logger)

	if id == "" {
		writer.ErrResponse(errors.New("path can't be empty"), w)
		return
	}

	userId, err := strconv.Atoi(id)
	v := validator.New()
	v.Check(err == nil, id, "invalid user id")

	if !v.Valid() {
		writer.ValidationErrorResponse(v, w)
		return
	}

	user := &data.User{
		Id: int64(userId),
	}

	// find in db
	err = app.db.Users.Find(user)
	if err != nil {
		writer.NotFoundResponse(w)
		return
	}

	writer.WriteJSONResponse(json.Envelope{
		"data": user,
	}, w)
}
