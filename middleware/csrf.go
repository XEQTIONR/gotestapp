package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckCSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		for _, v := range []string{"POST", "PUT", "PATCH", "DELETE"} {
			if c.Request.Method == v {
				header := c.Request.Header.Get("X-XSRF-TOKEN")
				cookie, _ := c.Cookie("XSRF-TOKEN")
				acceptHeader := c.Request.Header.Get("Accept")
				referer := c.Request.Header.Get("Referer")
				session := sessions.Default(c)

				if strings.Contains(acceptHeader, "application/json") {
					c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": gin.H{"csrf": "Invalid CSRF token (AJAX)"}})
					return
				} else {
					if cookie != header {
						session.Set("errors", map[string]string{"csrf": "Invalid CSRF TOKEN (non AJAX)"})
						session.Save()

						c.Redirect(http.StatusFound, referer)
						c.Abort()
					}
				}
			}
		}
		c.Next()
	}

}
