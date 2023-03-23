package services

import (
	"boluswizard/models"
	"boluswizard/restapi/config"
	"boluswizard/restapi/operations"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func VerifyCredentialsWithToken(bearer string) (string, error) {

	// Parse the token to check if it's valid
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// ErrorLogger <- fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.Secret, nil
	})
	if err != nil {
		// ErrorLogger <- err
		// c.ClearCookie("authToken")
		return "nil", err
	}

	// Decode the second portion of the JWT token for the username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		//c.ClearCookie("authToken")
		return "nil", fmt.Errorf("could not parse claims")
	}

	// usering the UID of the user find if there is any UID where the username and password matches
	// if not throw wrong password error

	jwtUid, ok := claims["Uid"].(string)
	if !ok {
		// c.ClearCookie("authToken")
		return "nil", fmt.Errorf("no Uid field in JWT")
	}

	user, err := User(jwtUid)
	if err != nil {
		return "nil", err
	}

	// Return the username from claims
	jwtUsername, ok := claims["Username"].(string)
	if !ok {
		return "nil", fmt.Errorf("no Username in JWT")
	}

	if user.Username != jwtUsername {
		// c.ClearCookie("authToken")
		return "nil", err
	}

	return jwtUid, nil
}
func SignUp(body *models.Users) (string, error) {
	var authToken string

	err := UsersExists(body.Username)
	if err != nil {
		return authToken, err
	}

	uid := (uuid.New()).String()

	// Create the Claims
	claims := jwt.MapClaims{
		"Username": body.Username,
		"Uid":      uid,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	authToken, err = token.SignedString(config.Secret)
	if err != nil {
		ErrorLogger <- err
		return authToken, err
	}

	body.UID = uid

	// save user account in redis
	err = SaveUser(body)
	if err != nil {
		return authToken, err
	}

	return authToken, err
}
func SignIn(body models.Users) (string, error) {

	// Create the Claims
	claims := jwt.MapClaims{
		"Username": body.Username,
		"Uid":      body.UID,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	authToken, err := token.SignedString(config.Secret)
	if err != nil {
		ErrorLogger <- err
		return authToken, err
	}

	return authToken, err
}

func Authenticated(username string, token string) middleware.ResponderFunc {

	authenticateResonse := func(rw http.ResponseWriter, pr runtime.Producer) {

		homeCookie := &http.Cookie{
			Name:   "authToken",
			Value:  token,
			Path:   "/u/" + username + "/home",
			MaxAge: 86400,
		}
		wizardToken := &http.Cookie{
			Name:   "wizardToken",
			Value:  token,
			Path:   "/wizard",
			MaxAge: 86400,
		}

		response := operations.SignUpOKBody{
			Status: http.StatusOK,
			Error:  nil,
			Data: &models.Token{
				Username: username,
				Token:    token,
			},
		}
		http.SetCookie(rw, homeCookie)
		http.SetCookie(rw, wizardToken)

		rw.WriteHeader(http.StatusOK)

		pr.Produce(rw, response)
	}

	return authenticateResonse

}
