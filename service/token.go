package service

import (
	"context"
	"fmt"

	"github.com/davidyap2002/user-go/graph/model"

	"time"

	"github.com/davidyap2002/user-go/tools"

	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

//UserClaims Claims Object (Token)
type UserClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("SECRET"))

//CreateToken Create Token
func CreateToken(ctx context.Context, email string, password string) (*model.TokenData, error) {
	checkUser, err := UserGetByEmail(ctx, email, nil)

	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil, &gqlerror.Error{
			Message: "Email Not Found",
			Extensions: map[string]interface{}{
				"code": "DENIED",
			},
		}
	}

	if !tools.CompareHashedPassword(password, checkUser.Password) {
		fmt.Println("Wrong Password")
		return nil, &gqlerror.Error{
			Message: "Wrong Password",
			Extensions: map[string]interface{}{
				"code": "DENIED",
			},
		}
	}

	signingMethod := jwt.SigningMethodHS256
	expiredTime := time.Now().AddDate(0, 0, 1).UnixNano() / int64(time.Millisecond)

	customClaim := UserClaims{
		ID: checkUser.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime,
		},
	}

	token := jwt.NewWithClaims(signingMethod, customClaim)

	signedToken, err := token.SignedString(jwtSecret)

	if err != nil {
		fmt.Println(err)
		return nil, gqlerror.Errorf(fmt.Sprintf("%s", err))
	}

	return &model.TokenData{
		Type:  "Bearer",
		Token: signedToken,
	}, nil
}

//ValidateToken Validate Token
func ValidateToken(t string) (*jwt.Token, error) {
	token, _ := jwt.ParseWithClaims(t, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("There was an error")
		}
		return jwtSecret, nil
	})

	return token, nil
}
