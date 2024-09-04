package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"qropen-backend/internal/core/ports"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	authService       ports.AuthService
	oauthService      ports.OAuthService
	repository        ports.Repositories
	googleOauthConfig *oauth2.Config
}

func NewAuthHandler(authService ports.AuthService, oauthService ports.OAuthService) *AuthHandler {
	googleOauthConfig := &oauth2.Config{
		ClientID:     "YOUR_GOOGLE_CLIENT_ID",
		ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return &AuthHandler{
		authService:       authService,
		oauthService:      oauthService,
		googleOauthConfig: googleOauthConfig,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	token, err := h.authService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Đăng xuất thành công"})
}

func (h *AuthHandler) Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Đây là route được bảo vệ"})
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url, err := h.oauthService.GetGoogleAuthURL()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	userInfo, err := h.oauthService.HandleGoogleCallback(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.repository.UserRepo.CreateUser(userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userInfo})
}

func generateStateOauthCookie(c *gin.Context) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, 3600, "/", "localhost", false, true)
	return state
}
