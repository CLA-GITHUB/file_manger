package user

import (
	"net/http"

	model "github.com/cla-github/file_manager/internal/user"
	userService "github.com/cla-github/file_manager/internal/user/service"
	"github.com/cla-github/file_manager/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not hash password, try again.",
			"data":  nil,
		})
		return
	}

	user.Password = hashedPass

	err = userService.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": nil,
		"data":  "User creation successful.",
	})
}

func Signin(c *gin.Context) {
	type Auth struct {
		Email    string
		Password string
	}
	var data Auth

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(500, gin.H{
			"error": "Could not extract username and password from body.",
			"data":  nil,
		})
		return
	}

	foundUser := userService.GetUserByEmail(data.Email)
	if foundUser == nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password.",
			"data":  nil,
		})
		return
	}

	passwordMatches := comparePassword(data.Password, foundUser.Password)

	if !passwordMatches {
		c.JSON(400, gin.H{
			"error": "Invalid username or password.",
			"data":  nil,
		})
		return
	}

	// generate jwt token
	token, err := utils.GenerateJwt(foundUser.ID)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Error creating jwt token",
			"data":  nil,
		})
		return
	}

	c.JSON(200, gin.H{"error": nil, "data": token})
}

func Me(c *gin.Context) {
	userId, err := uuid.Parse(c.GetString("uid"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"data":  nil,
		})
		return
	}

	user := userService.GetMe(userId)
	if user == nil {
		c.JSON(200, gin.H{
			"error": "User not found.",
			"data":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"error": nil,
		"data":  user,
	})
}

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
