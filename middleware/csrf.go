package middleware

import (
	"fmt"
	"math/rand"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const key = "csrf_token"

func GenerateCSRFToken(c *gin.Context) {
	session := sessions.Default(c)
	f := func() (string, error) {
		seed := fmt.Sprintf("%d", rand.Int())
		b, e := bcrypt.GenerateFromPassword([]byte(seed), 14)

		return string(b), e
	}

	if token, ok := session.Get(key).(string); ok {
		if len(token) == 0 {
			if newToken, err := f(); err == nil {
				session.Set(key, newToken)
				if err := session.Save(); err != nil {
					fmt.Println("Error saving session: ", err)
				}
			} else {
				fmt.Println("Session generation error: ", err)
			}
		}
	} else {
		if newToken, err := f(); err == nil {
			session.Set(key, newToken)
			if err := session.Save(); err != nil {
				fmt.Println("Error saving session: ", err)
			}
		} else {
			fmt.Println("Session generation error: ", err)
		}
	}

	c.Next()
}
