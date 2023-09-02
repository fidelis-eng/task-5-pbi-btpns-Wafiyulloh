package helpers

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SECRET = "LKAJLKDFJLJalkjdlfajslkdfj123130124"

func Validation(c *gin.Context, data interface{}) error {
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return err
	}
	return err
}

func Encrypt_password(c *gin.Context, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return ""
	}
	return string(hash)
}

func Check_password(pass_1 string, pass_2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pass_1), []byte(pass_2))
	if err != nil {
		return err
	}
	return err
}

func Initialize_token(userid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(SECRET))
	return tokenStr, err
}
