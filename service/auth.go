package service

import (
	mysqldbmodel "gql/mysqldb/model"
	"gql/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

var jwtSecret = []byte("secret")

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			token := strings.Split(auth, "Bearer ")[1]

			tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
				return jwtSecret, nil
			})

			if err != nil || tokenClaims == nil {
				next.ServeHTTP(w, r)
				return
			}
			/*

				// Allow unauthenticated users in
				if err != nil || c == nil {
					next.ServeHTTP(w, r)
					return
				}

				userId, err := validateAndGetUserID(c)
				if err != nil {
					http.Error(w, "Invalid cookie", http.StatusForbidden)
					return
				}

				// get the user from the database
				user := getUserByID(db, userId)

				// put it in context
				ctx := context.WithValue(r.Context(), userCtxKey, user)

				// and call the next with our new context
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			*/

		})
	}
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
