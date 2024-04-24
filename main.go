package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const userkey = "user"

var secret = []byte("Secret123")

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// r.AddFromFiles("home", "templates/base.html", "templates/index/main.html", "templates/index/person.html")
	r.AddFromFiles("home", "dist/index.html")

	r.AddFromFiles("me", "templates/base.html", "templates/me/main.html")
	return r
}

func AuthRequired(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()

}

// me is the handler that will return the user information stored in the
// session.
func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.HTML(200, "me", gin.H{"user": user, "title": "titleParam"})
	// c.JSON(http.StatusOK, gin.H{"user": user})
}

// status is the handler that will tell the user whether it is logged in or not.
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "hello" || password != "itsme" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

// logout is the handler called for the user to log out.
func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func respond(c *gin.Context, data map[string]any) {
	if c.Request.Header.Get("AJAXRequest") == "true" {
		c.JSON(http.StatusOK, gin.H(data))
	} else {
		c.HTML(http.StatusOK, "home", gin.H{
			"data": data,
		})
	}
}

func main() {
	r := gin.New()
	r.Static("/assets", "dist/assets")
	r.Static("/dist", "dist")
	r.Use(sessions.Sessions("mysessions", cookie.NewStore(secret)))

	r.HTMLRender = createMyRender()

	r.GET("/", func(c *gin.Context) {
		people := []Person{
			{Name: "John Doe", Age: 20},
			{Name: "Jane Doe", Age: 18},
			{Name: "Stan Doe", Age: 3},
		}

		two := []Person{
			people[rand.Int()%3],
			people[rand.Int()%3],
		}

		respond(c, map[string]interface{}{
			"people": two})
	})

	r.GET("/about", func(c *gin.Context) {
		respond(c, map[string]interface{}{
			"specie": "alien",
			"age":    45,
			"color":  "red",
		})
	})

	r.GET("/another", func(c *gin.Context) {
		fmt.Println("/another called")
		respond(c, map[string]interface{}{
			"people":  []Person{{Name: "John Doe", Age: 20}, {Name: "Jane Doe", Age: 18}},
			"message": "Pong",
		})
	})

	r.POST("login", login)
	r.POST("logout", logout)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}

	r.Run("127.0.0.1:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
