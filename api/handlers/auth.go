package handlers

import (
	"fmt"
	"net/http"

	"github.com/cla-github/file_manager/db"
	"github.com/cla-github/file_manager/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hash), err
}

func comparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}

func Signup(c *gin.Context) {
	var user types.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not hash password, try again."})
		return
	}

	formattedUser := types.User{
		Username: user.Username,
		Password: hashedPass,
		Email:    user.Email,
	}

	//TODO: Add formatted user to db
	updatedUsers := append(db.Users, formattedUser)
	c.IndentedJSON(http.StatusOK, updatedUsers)

}

func Signin(c *gin.Context) {
	var data types.Auth

	if err := c.BindJSON(&data); err != nil {
		c.JSON(500, gin.H{"error": "Could not extract username and password from body."})
		return
	}

	var foundUser *types.User

	for _, u := range db.Users {
		if u.Username == data.Username {

			fmt.Println(u)
			foundUser = &u
			break
		}
	}
	fmt.Println(foundUser)

	if foundUser == nil {
		c.JSON(400, gin.H{"error": "Invalid username or password."})
		return
	}

	passwordMatches := comparePassword(data.Password, foundUser.Password)

	if !passwordMatches {
		c.JSON(400, gin.H{"error": "Incorrect username or password."})
		return
	}

	c.JSON(200, gin.H{"error": nil})
}
