package services

import (
	"boluswizard/restapi/config"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
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
