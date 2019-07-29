package api

import (
	"encoding/json"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"gitlab.com/standard-go/project/internal/app/project"
)

func IndexUser(w http.ResponseWriter, r *http.Request) {
	pageRequest := r.Context().Value(pageRequestCtxKey).(*project.PageRequest)

	fetch, count, err := srv.FetchIndexUser(pageRequest)
	if err != nil {
		SendErrorJSON(err, w)
		return
	}

	res := setPaginate(r, pageRequest, fetch, count, project.UserUrl)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtxKey).(*project.User)

	res := project.NewResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func StoreUser(w http.ResponseWriter, r *http.Request) {
	// auth := r.Context().Value(authUserCtxKey).(*claim.Auth)
	var userReq *project.User
	error := json.NewDecoder(r.Body).Decode(&userReq)
	if error != nil {
		SendErrorJSON(printErrorMessage(error), w)
		return
	}

	fetch, err := srv.FetchStoreUser(userReq)
	if err != nil {
		SendErrorJSON(err, w)
		return
	}

	res := project.NewResponse(fetch)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtxKey).(*project.User)
	userId := user.ID.(primitive.ObjectID).Hex()

	var userReq *project.User
	error := json.NewDecoder(r.Body).Decode(&userReq)
	if error != nil {
		SendErrorJSON(printErrorMessage(error), w)
		return
	}

	user.FullName = userReq.FullName

	user, err := srv.FetchUpdateUser(userId, user)
	if err != nil {
		SendErrorJSON(err, w)
		return
	}

	res := project.NewResponse(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DestroyUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userCtxKey).(*project.User)
	userId := user.ID.(primitive.ObjectID).Hex()

	err := srv.FetchDestroyUser(userId)
	if err != nil {
		SendErrorJSON(err, w)
		return
	}

	res := project.NewResponse(nil)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

