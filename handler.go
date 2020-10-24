package main

// import (
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo"
// )

// type handler struct{}

// func (h *handler) login(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	if username == "leeeeoy" && password == "123" {
// 		token := jwt.New(jwt.SigningMethodHS256)

// 		// This is the information which frontend can use
// 		// The backend can also decode the token and get admin etc.
// 		claims := token.Claims.(jwt.MapClaims)
// 		claims["name"] = "Yoel"
// 		claims["admin"] = true
// 		claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

// 		// Generate encoded token and send it as response.
// 		// The signing string should be secret (a generated UUID          works too)
// 		t, err := token.SignedString([]byte("secret"))
// 		if err != nil {
// 			return err
// 		}
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"token": t,
// 		})
// 	}
// 	return echo.ErrUnauthorized
// }
