package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mod/server/pkg/check_func/valid"
	"go.mod/server/pkg/db/pgsql"
)

type AuthInfo struct {
	Username string
	Email    string
	Password string
}

type ClaimsStru struct {
	IDUser   int
	Username string
	Password string
	Email    string
	jwt.StandardClaims
}

// API register
func (api *API) register(w http.ResponseWriter, r *http.Request) {
	var auth AuthInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if valid.CheckEmail(auth.Email) && valid.CheckUsrname(auth.Username) && valid.CheckPassLen(auth.Password) {
		user := pgsql.User{
			Username: string([]rune(auth.Username)),
			Password: auth.Password,
			Email:    auth.Email,
		}
		res := pgsql.AddNewUser(user)
		if res {
			w.WriteHeader(http.StatusCreated)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

// API login
func (api *API) login(w http.ResponseWriter, r *http.Request) {
	var auth AuthInfo

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user := pgsql.User{
		Username: string([]rune(auth.Username)),
		Password: auth.Password,
	}

	res := pgsql.ChUserExist(user)
	if res {
		timeToken := time.Now().Add(2 * time.Minute)
		claims := &ClaimsStru{
			Username: auth.Username,
			Password: auth.Password,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: timeToken.Unix(),
			},
		}
		tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		tokenSign, err := tokenJWT.SignedString([]byte("secretSecret")) //лучше доставать из окружения env
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			w.Header().Set("Authorization", string([]byte(tokenSign)))
			w.WriteHeader(http.StatusOK)
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// API profile
func (api *API) profile(w http.ResponseWriter, r *http.Request) {

	token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte("secretSecret"), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !ok {
			return
		}
		timeTok := claims["exp"].(float64)
		userN := claims["Username"].(string)
		auth := pgsql.User{
			Username: userN,
		}
		insTime := (int64(timeTok) - time.Now().Unix())
		if insTime <= 0 {
			w.WriteHeader(http.StatusForbidden)
			return
		} else {
			ud, res := pgsql.SelDataUser(auth)
			if res {
				timeToken := time.Now().Add(2 * time.Minute)
				claims := &ClaimsStru{
					IDUser:   ud.IDUser,
					Username: ud.Username,
					Email:    ud.Email,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: timeToken.Unix(),
					},
				}
				tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
				tokenSign, err := tokenJWT.SignedString([]byte("secretSecret"))
				if err != nil {
					http.Error(w, err.Error(), http.StatusForbidden)
				} else {
					w.Header().Set("Authorization", string([]byte(tokenSign)))
					w.WriteHeader(http.StatusOK)
				}
			} else {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
