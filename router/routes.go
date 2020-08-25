package router

import (
	"apiservice/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

//Auth0 integration
var jwtMiddleware = AuthRequired()

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			//router.POST(route.Pattern, checkJWT(), route.HandlerFunc)
			router.POST(route.Pattern, route.HandlerFunc) //To protect your route replace this line with the previous line
		case http.MethodPut:
			//router.PUT(route.Pattern, checkJWT(), route.HandlerFunc)
			router.PUT(route.Pattern, route.HandlerFunc) //To protect your route replace this line with the previous line
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It works!!!"})
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},
	{
		"ServiceGet",
		http.MethodGet,
		"/service",
		controllers.ServiceGet,
	},
	{
		"ServicePost",
		http.MethodPost,
		"/service",
		controllers.ServicePost,
	},
	{
		"ServicePut",
		http.MethodPut,
		"/service/:id",
		controllers.ServicePut,
	},
}
