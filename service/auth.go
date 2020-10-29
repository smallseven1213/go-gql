package service

import (
	"context"
	mysqldbmodel "gql/mysqldb/model"
	"gql/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type CustomerInfo struct {
	Name string
}

type Claims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
	CustomerInfo
}

var (
	jwtSecret  = []byte("secret")
	userCtxKey = "user"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth != "" {
				token := strings.Split(auth, "Bearer ")[1]

				// --> 補上token是否在redis
				tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
					return jwtSecret, nil
				})

				if tokenClaims != nil && err == nil {
					claims := tokenClaims.Claims.(*Claims)

					user, err := GetUserByUid(&claims.Id)

					if user != nil && err == nil {
						ctx := context.WithValue(r.Context(), userCtxKey, user)

						// and call the next with our new context
						r = r.WithContext(ctx)
					}
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *mysqldbmodel.User {
	raw, _ := ctx.Value(userCtxKey).(*mysqldbmodel.User)
	return raw
}

func GetUserByToken(uid *string) *mysqldbmodel.User {
	db := utils.GetSQLDB()

	user := mysqldbmodel.User{}

	db.Where(&mysqldbmodel.User{Uid: *uid}).First(&user)

	return &user
}

func Login() bool {
	return false
}

func Logout() bool {
	return false
}
