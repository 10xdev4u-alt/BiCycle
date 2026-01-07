package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// User represents the user data stored in the database.
type User struct {
	ID           int       `json:"id"`
	RegNo        string    `json:"reg_no"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	IsBanned     bool      `json:"is_banned"`
	NoShowCount  int       `json:"no_show_count"`
	Role         string    `json:"role"`
	PasswordHash *string   `json:"password_hash"` // Nullable for OAuth users
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var googleOauthConfig *oauth2.Config

func (a *API) setupGoogleOauth() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  a.Cfg.GoogleRedirectURL,
		ClientID:     a.Cfg.GoogleClientID,
		ClientSecret: a.Cfg.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

// GoogleLogin initiates the Google OAuth login process.
func (a *API) GoogleLogin(c *gin.Context) {
	if googleOauthConfig == nil {
		a.setupGoogleOauth()
	}
	url := googleOauthConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles the Google OAuth callback.
func (a *API) GoogleCallback(c *gin.Context) {
	if googleOauthConfig == nil {
		a.setupGoogleOauth()
	}

	state := c.Query("state")
	if state != "randomstate" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid state parameter."})
		return
	}

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization code not provided."})
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Code exchange failed: %s", err.Error())})
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get user info: %s", err.Error())})
		return
	}
	defer response.Body.Close()

	userData, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read user data: %s", err.Error())})
		return
	}

	var googleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		Hd            string `json:"hd"` // Hosted domain
	}
	if err := json.Unmarshal(userData, &googleUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to parse user data: %s", err.Error())})
		return
	}

	// Verify domain (hd)
	if googleUser.Hd != "svce.ac.in" { // TODO: Make this configurable
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only SVCE students are allowed to register."})
		return
	}

	// TODO: Check if user exists in DB, create if not
	// TODO: Generate JWT and return to frontend

	c.JSON(http.StatusOK, gin.H{
		"message": "Authentication successful. Implement JWT generation and return.",
		"user":    googleUser,
	})
}
