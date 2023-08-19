package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"example.com/test/rss_argg/interval/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(httpWiter http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(httpWiter, 400, fmt.Sprintf("Failed to parse json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(httpWiter, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	responseWithJSON(httpWiter, 200, databaseUserToUserModel(user))
}
