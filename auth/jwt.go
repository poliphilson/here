package auth

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

func VerifyAccessToken(c *gin.Context) {
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Fail to get access-token cookie.")
		log.Println(err.Error())
		c.Abort()
		return
	}

	aToken := cookie.Value
	if aToken == "" {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Access token is empty.")
		c.Abort()
		return
	}

	secret := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.ParseWithClaims(aToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		c.Set("uid", claims.Uid)
		c.Set("email", claims.Email)
	} else if v, ok := err.(*jwt.ValidationError); ok {
		if v.Errors == jwt.ValidationErrorExpired {
			response.FailedSignIn(c, status.ExpiredAccessToken)
			log.Println("Access token is expired.")
			log.Println(v.Error())
			c.Abort()
			return
		} else {
			response.FailedSignIn(c, status.FailedSignIn)
			log.Println("Fail to verify access token.")
			log.Println(err.Error())
			c.Abort()
			return
		}
	} else {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Can't handle jwt token error.")
		log.Println(err.Error())
		c.Abort()
		return
	}

	c.Next()
}

func RefreshAccessToken(c *gin.Context) {
	cookie, err := c.Request.Cookie("refresh-token")
	if err != nil {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Fail to get refresh-token cookie.")
		log.Println(err.Error())
		return
	}

	rToken := cookie.Value
	result := VerifyRefreshToken(rToken)
	if !result {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Fail to verify refresh token.")
		return
	}
}

func VerifyRefreshToken(rToken string) bool {
	if rToken == "" {
		log.Println("Refresh token is empty.")
		return false
	}

	secret := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.ParseWithClaims(rToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if _, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		return true
	} else {
		log.Println("Fail to verify refresh token.")
		log.Println(err.Error())
		return false
	}
}

func CreateAccessToken(uid int, email string) (string, error) {
	secret := []byte(os.Getenv("SECRET_KEY"))

	aClaims := AccessTokenClaims{
		Uid:   uid,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	aToken, err := temp.SignedString(secret)
	if err != nil {
		return "", err
	}

	return aToken, nil
}

func CreateRefreshToken() (string, error) {
	secret := []byte(os.Getenv("SECRET_KEY"))

	rClaims := RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	rToken, err := temp.SignedString(secret)
	if err != nil {
		return "", err
	}
	return rToken, nil
}
