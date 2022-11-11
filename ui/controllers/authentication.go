package controllers

import (
	"fmt"
	"ui/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(bearer string, c *fiber.Ctx) (string, string, error) {

	var jwtUsername string
	var jwtUid string

	// Parse the token to check if it's valid
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.Secret, nil
	})
	if err != nil {
		c.ClearCookie("authToken")
		return jwtUid, jwtUsername, err
	}

	// Decode the second portion of the JWT token for the username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.ClearCookie("authToken")
		return jwtUid, jwtUsername, fmt.Errorf("could not parse claims")
	}

	// usering the UID of the user find if there is any UID where the username and password matches
	// if not throw wrong password error

	jwtUid, ok = claims["Uid"].(string)
	if !ok {
		c.ClearCookie("authToken")
		return jwtUid, jwtUsername, fmt.Errorf("no Uid field in JWT")
	}

	// Return the username from claims
	jwtUsername, ok = claims["Username"].(string)
	if !ok {
		return jwtUid, jwtUsername, fmt.Errorf("no Username in JWT")
	}

	return jwtUid, jwtUsername, nil
}
