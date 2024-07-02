package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ecom/src/config"
	"ecom/src/types"
	"ecom/src/utils"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	} else {
		return tokenString, err
	}
}

func WithJWTAuth(handlerFunc http.HandlerFunc, userStore types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// GET Token From the User Request
		tokenString := utils.GetTokenFromRequest(r)

		// Validate JWT
		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			utils.PermissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			utils.PermissionDenied(w)
			return
		}

		// Validate JWT ~> TRUE then We Need to Fetch User ID from DB (ID from the Token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert userID to int: %v", err)
			utils.PermissionDenied(w)
			return
		}

		user, err := userStore.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			utils.PermissionDenied(w)
			return
		}

		// Set Context "UserID" to the User ID
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user.ID)
		r = r.WithContext(ctx)

		// Call the function if the token is valid
		handlerFunc(w, r)
	}
}

func validateToken(ts string) (*jwt.Token, error) {
	return jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}

	return userID
}
