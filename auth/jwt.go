package auth

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/poliphilson/here/repository"
	"github.com/poliphilson/here/response"
	"github.com/poliphilson/here/status"
)

type AccessTokenClaims struct {
	Uid          int    `json:"uid"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	Uuid string
	jwt.RegisteredClaims
}

var secret []byte = []byte(os.Getenv("SECRET_KEY"))

func VerifyAccessToken(c *gin.Context) {
	cookie, err := c.Request.Cookie("access_token")
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

	token, err := jwt.ParseWithClaims(aToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		c.Set("uid", claims.Uid)
		c.Set("email", claims.Email)
		c.Set("profile_image", claims.ProfileImage)
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
		log.Println("Can't handle jwt.")
		log.Println(err.Error())
		c.Abort()
		return
	}

	c.Next()
}

func RefreshAccessToken(c *gin.Context) {
	rCookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Fail to get refresh-token cookie.")
		log.Println(err.Error())
		return
	}

	aCookie, err := c.Request.Cookie("access_token")
	if err != nil {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Fail to get access-token cookie.")
		log.Println(err.Error())
		return
	}

	aToken := aCookie.Value
	token, err := jwt.ParseWithClaims(aToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		response.BadRequset(c, status.StatusOK)
		log.Println("Access token is not expired.")
		return
	} else if v, ok := err.(*jwt.ValidationError); ok {
		if v.Errors == jwt.ValidationErrorExpired {
			rToken := rCookie.Value
			result := VerifyRefreshToken(rToken)
			if !result {
				response.FailedSignIn(c, status.FailedSignIn)
				log.Println("Fail to verify refresh token.")
				return
			}

			newAToken, err := CreateAccessToken(claims.Uid)
			if err != nil {
				response.InternalServerError(c, status.InternalError)
				log.Println("Fail to create new access token.")
				log.Println(err.Error())
				return
			}

			c.SetCookie("access_token", newAToken, 60*60*72, "/", "localhost", false, true)
			response.Ok(c, status.StatusOK)
			return
		} else {
			response.FailedSignIn(c, status.FailedSignIn)
			log.Println("Fail to verify access token.")
			log.Println(err.Error())
			return
		}
	} else {
		response.FailedSignIn(c, status.FailedSignIn)
		log.Println("Can't handle jwt.")
		log.Println(err.Error())
		return
	}
}

func VerifyRefreshToken(rToken string) bool {
	if rToken == "" {
		log.Println("Refresh token is empty.")
		return false
	}

	token, err := jwt.ParseWithClaims(rToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		rUuid := claims.Uuid
		redisClient := repository.Redis()
		err = redisClient.Get(context.Background(), rUuid).Err()
		if err != nil {
			log.Println("Fail to get refresh token uuid from redis.")
			log.Println(err.Error())
			return false
		}
		return true
	} else {
		log.Println("Can't handle jwt.")
		log.Println(err.Error())
		return false
	}
}

func CreateAccessToken(uid int) (string, error) {
	aClaims := AccessTokenClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
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
	rUuid := uuid.New().String()
	rClaims := RefreshTokenClaims{
		Uuid: rUuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)
	rToken, err := temp.SignedString(secret)
	if err != nil {
		return "", err
	}
	redisClient := repository.Redis()
	err = redisClient.Set(context.Background(), rUuid, rToken, time.Duration(72)*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return rToken, nil
}

func DeleteRefreshTokenFromRedis(rToken string) {
	token, _ := jwt.ParseWithClaims(rToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		rUuid := claims.Uuid
		redisClient := repository.Redis()
		redisClient.Del(context.Background(), rUuid)
	} else {
		log.Println("Can't handle jwt.")
	}
}
