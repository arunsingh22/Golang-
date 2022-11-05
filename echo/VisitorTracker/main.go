package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func count(c echo.Context) error {
	cookie, err := c.Cookie("my-token")
	if err != nil {
		//creating cookie firstTime
		cookie := new(http.Cookie)
		cookie.Name = "my-token"
		cookie.Value = "1"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.String(http.StatusOK, cookie.Value)
	} else {
		a, _ := strconv.Atoi(cookie.Value)
		cookie.Value = strconv.Itoa(a + 1)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, fmt.Sprintf("Visitors: %s", cookie.Value))
	}
}

// Creating visitor counter using cookies
func main() {
	e := echo.New()

	//routes
	e.GET("/", count)
	e.Logger.Fatal(e.Start("localhost:1900"))
}
