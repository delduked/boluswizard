package services

import (
	"fmt"
	"time"

	"api/config"
	"api/types"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func SignUp(body *types.Users) (string, error) {
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

	// save user account in redis
	err = SaveUser(body)
	if err != nil {
		return authToken, err
	}

	return authToken, err
}
func SignIn(body types.Users) (string, error) {

	// Create the Claims
	claims := jwt.MapClaims{
		"Username": body.Username,
		"Uid":      body.Uid,
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

func VerifyCredentialsWithToken(bearer string, c *fiber.Ctx) (string, error) {

	// Parse the token to check if it's valid
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrorLogger <- fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.Secret, nil
	})
	if err != nil {
		ErrorLogger <- err
		c.ClearCookie("authToken")
		return "nil", err
	}

	// Decode the second portion of the JWT token for the username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.ClearCookie("authToken")
		return "nil", fmt.Errorf("could not parse claims")
	}

	// usering the UID of the user find if there is any UID where the username and password matches
	// if not throw wrong password error

	jwtUid, ok := claims["Uid"].(string)
	if !ok {
		c.ClearCookie("authToken")
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
		c.ClearCookie("authToken")
		return "nil", err
	}

	return jwtUid, nil
}

func VerifyMiddleWare(bearer string) (string, error) {

	// Parse the token to check if it's valid
	token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			ErrorLogger <- fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.Secret, nil
	})
	if err != nil {
		return "", err
	}

	// Decode the second portion of the JWT token for the username
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("could not verify token")
	}

	// might need to add expiration logic

	return claims["Uid"].(string), nil
}

func CheckSecret(body *types.Users) error {
	if body.Password == "n4th4n43l" {
		return nil
	}
	if body.Password != "n4th4n43l" {
		return fmt.Errorf("incorrect secret")
	}
	return fmt.Errorf("System error")
}
