package api

import (
	"net/http"
	"strings"

	"github.com/anotherandrey/token-rest-api/internal/app/encoding"
	"github.com/anotherandrey/token-rest-api/internal/app/model"
	"github.com/anotherandrey/token-rest-api/internal/app/token"
	"github.com/gorilla/mux"
)

const LoginUrlPath = "/api/v1/login"

type Api struct {
	Router *mux.Router
}

func NewApi() *Api {
	api := &Api{
		Router: mux.NewRouter(),
	}

	subrouter := api.Router.PathPrefix("/api/v1").Subrouter()

	subrouter.Use(api.tokenHandler)

	subrouter.Handle("/whoami", api.handleWhoAmI()).Methods("GET")
	subrouter.Handle("/login", api.handleLogin()).Methods("POST")

	return api
}

func (api *Api) tokenHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		if urlPath == LoginUrlPath {
			next.ServeHTTP(w, r)
			return
		}

		token_ := api.GetToken(r)
		if len(token_) == 0 {
			http.Error(w, "Get out of here without bearer token!!!", http.StatusUnauthorized)
			return
		}

		err := token.ValidToken(token_)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *Api) handleWhoAmI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token_ := api.GetToken(r)

		claims, err := token.ParseToken(token_)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// DB?
		user := &model.UserModel{
			Id:       claims.Id,
			Username: claims.Username,
			Password: claims.Password,
		}

		encoding.Encode(w, &user)
	}
}

func (api *Api) handleLogin() http.HandlerFunc {
	type LoginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		body := &LoginBody{}
		if err := encoding.Decode(r, &body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// DB?
		user := &model.UserModel{
			Id:       1,
			Username: "root",
			Password: "root",
		}

		if body.Username != user.Username || body.Password != user.Password {
			http.Error(w, "Get out of here with wrong credentials!!!", http.StatusUnauthorized)
			return
		}

		token, err := token.CreateToken(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		encoding.Encode(w, map[string]string{"token": token})
	}
}

func (api *Api) GetToken(r *http.Request) string {
	header := r.Header.Get("Authorization")
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return ""
	}
	token := strings.TrimPrefix(header, "Bearer ")
	return token
}
