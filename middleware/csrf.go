package middleware

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const key = "XSRF-TOKEN"

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

func CheckCSRFToken(c *gin.Context) {
	for _, v := range []string{"POST", "PUT", "PATCH", "DELETE"} {
		if c.Request.Method == v {
			header := c.Request.Header.Get("X-XSRF-TOKEN")
			cookie, err := c.Cookie("XSRF-TOKEN")
			acceptHeader := c.Request.Header.Get("Accept")
			referer := c.Request.Header.Get("Referer")

			if err != nil {
				if strings.Contains(acceptHeader, "application/json") {
					c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": gin.H{"csrf": "Form timeout, try again"}})
				} else {
					c.Redirect(http.StatusTemporaryRedirect, referer)
				}
				return
			} else if cookie != header {
				if strings.Contains(acceptHeader, "application/json") {
					c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": gin.H{"csrf": "Form timeout, try again"}})
				} else {
					c.Redirect(http.StatusTemporaryRedirect, referer)
				}
				return
			}
		}
	}
	c.Next()
}
