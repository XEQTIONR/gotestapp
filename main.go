package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

const userkey = "user"

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("home", "dist/index.html")

	r.AddFromFiles("me", "templates/base.html", "templates/me/main.html")
	return r
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		acceptHeader := c.Request.Header.Get("Accept")

		session.Set("to", c.Request.RequestURI)
		session.Save()

		if strings.Contains(acceptHeader, "application/json") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	}

	c.Next()
}

// me is the handler that will return the user information stored in the
// session.
func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	//c.HTML(200, "me", gin.H{"user": user, "title": "titleParam"})
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// status is the handler that will tell the user whether it is logged in or not.
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

type authCreds struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func login(c *gin.Context) {

	var creds authCreds

	session := sessions.Default(c)

	acceptHeader := c.Request.Header.Get("Accept")

	username := ""
	password := ""

	if strings.Contains(acceptHeader, "application/json") {
		c.BindJSON(&creds)
		username = creds.Username
		password = creds.Password
	} else {
		username = c.PostForm("username")
		password = c.PostForm("password")

	}
	to := session.Get("to")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username == "ishtehar" && password == "password" {
		toStr, ok := to.(string)
		session.Delete("to")
		session.Set(userkey, username) // In real world usage you'd set this to the users ID

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		} else {

			acceptHeader := c.Request.Header.Get("Accept")
			if strings.Contains(acceptHeader, "application/json") {
				fmt.Println("login application/json")
				var toLink string

				if to == nil && ok {
					toLink = "/private/me"
				} else {
					toLink = toStr
				}

				c.JSON(http.StatusOK, gin.H{"to": toLink})
				return
			} else {
				if to == nil {
					c.Redirect(http.StatusFound, "/private/me")
					return
				}
				if ok {
					c.Redirect(http.StatusFound, string(toStr))
					return
				}

			}
		}
	} else {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		errorMsg := gin.H{"password": "Password is incorrect or account does not exist"}

		jsonMsg, marshalError := json.Marshal(errorMsg)

		if marshalError == nil {
			session.Set("errors", string(jsonMsg))
		} else {
			session.Set("errors", "Problem ")
		}

		if e := session.Save(); e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Session save error": e})
			return
		}
		toStr, _ := to.(string)
		respond(c, map[string]interface{}{"error": errorMsg, "to": toStr})
		return
	}
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
	c.Redirect(http.StatusFound, "/")
}

func respond(c *gin.Context, data map[string]any) {
	acceptHeader := c.Request.Header.Get("Accept")

	if strings.Contains(acceptHeader, "application/json") {
		c.JSON(http.StatusOK, gin.H(data))
	} else {
		c.HTML(http.StatusOK, "home", gin.H{"data": data})
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	} else {

		secret := []byte(os.Getenv("SESSION_KEY"))
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
			fmt.Println("in /about")
			respond(c, map[string]interface{}{
				"specie": "alien",
				"age":    45,
				"color":  "red",
			})
		})

		r.GET("/login", func(c *gin.Context) {
			session := sessions.Default(c)
			user := session.Get(userkey)
			to := session.Get("to")
			errors := session.Get("errors")
			session.Delete("errors")
			session.Save()

			if user != nil {
				c.Redirect(http.StatusFound, "/private/me")
			} else {
				respond(c, map[string]any{"user": user, "to": to, "errors": errors})
			}
		})

		r.POST("/login", login)
		r.POST("logout", logout)

		private := r.Group("/private")
		private.Use(AuthRequired)
		{
			private.GET("/me", me)
			private.GET("/status", status)
			private.GET("/new", func(c *gin.Context) {
				respond(c, map[string]interface{}{})
			})

			private.GET("/another", func(c *gin.Context) {
				fmt.Println("/another called")
				respond(c, map[string]interface{}{
					"people":  []Person{{Name: "John Doe", Age: 20}, {Name: "Jane Doe", Age: 18}},
					"message": "Pong",
				})
			})
		}

		r.Run("127.0.0.1:9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	}

}
