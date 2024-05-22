package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
