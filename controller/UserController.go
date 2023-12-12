package controller

import (
	user2 "demo_project/Repositories/user"
	"demo_project/database"
	"demo_project/dto"
	"demo_project/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

var ur = user2.UserRepositories{}
var bodyUser dto.CreateUser

func SignUp(c *gin.Context) {
	if err := c.ShouldBindJSON(&bodyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(bodyUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser := dto.CreateUser{Name: bodyUser.Name, Email: bodyUser.Email, Password: string(hash)}
	result := ur.Create(&newUser)

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": true,
	})
}

// Hàm xử lý trả về lỗi
func handleError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

// Hàm tạo và ký token
func generateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

// Hàm thiết lập cookie
func setCookie(c *gin.Context, tokenString string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("authorization", tokenString, 3600*24*30, "", "", false, true)
}

// Login Hàm xử lý đăng nhập
func Login(c *gin.Context) {
	// Kiểm tra lỗi khi parse JSON
	if err := c.ShouldBindJSON(&bodyUser); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	// Kiểm tra tài khoản
	var user model.User
	if err := database.GetInstance().GetDB().First(&user, "email = ?", bodyUser.Email).Error; err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	// So sánh mật khẩu
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyUser.Password)); err != nil {
		handleError(c, http.StatusBadRequest, err)
		return
	}

	// Tạo và ký token
	tokenString, err := generateToken(int(user.ID))
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}

	// Thiết lập cookie
	setCookie(c, tokenString)

	// Trả về phản hồi thành công
	c.JSON(http.StatusOK, gin.H{"status": true, "token": tokenString})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
