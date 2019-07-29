package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt"

	"gitlab.com/standard-go/project/internal/app/config/env"
	"gitlab.com/standard-go/project/internal/app/project"
)

/*
|--------------------------------------------------------------------------
| Context
|--------------------------------------------------------------------------
|
*/
type ctxKey struct {
	name string
}

func (k ctxKey) String() string {
	return fmt.Sprintf("tugure-location's context value: %s", k.name)
}

var (
	statusCtxKey      = &ctxKey{"Status"}
	authUserCtxKey    = &ctxKey{"AuthUser"}
)

/*
|--------------------------------------------------------------------------
| Middlewares
|--------------------------------------------------------------------------
|
*/
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		if len(token) <= 0 {
			SendErrorJSON(printErrorMessage(errors.New("401")), w)
			return
		}
		token = splitToken[1]
		hs256 := jwt.NewHS256(env.Get("JWT_SECRET"))

		payload, sig, err := jwt.Parse(token)
		if err != nil {
			SendErrorJSON(printErrorMessage(errors.New("400-Token-Parse")), w)
			return
		}
		if err = hs256.Verify(payload, sig); err != nil {
			SendErrorJSON(printErrorMessage(errors.New("400-Signature")), w)
			return
		}
		var jot project.Token
		if err = jwt.Unmarshal(payload, &jot); err != nil {
			SendErrorJSON(printErrorMessage(errors.New("400-Token-Parse")), w)
			return
		}

		auth := &project.Auth{
			User: 	jot.ToUser(),
			Token: 	token,
		}

		// Validate fields.
		iatValidator := jwt.IssuedAtValidator(now)
		expValidator := jwt.ExpirationTimeValidator(now)
		audValidator := jwt.AudienceValidator("admin")
		if err = jot.Validate(iatValidator, expValidator, audValidator); err != nil {
			switch err {
			case jwt.ErrIatValidation:
				SendErrorJSON(printErrorMessage(errors.New("400-Signature")), w)
				return
			case jwt.ErrExpValidation:
				SendErrorJSON(printErrorMessage(errors.New("401-Expired")), w)
				return
			}
		}

		ctx := context.WithValue(r.Context(), authUserCtxKey, auth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
