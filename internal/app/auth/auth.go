package auth

import (
	"log"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	config "github.com/sk-pathak/fissionNode-backend/configs"
)

const (
	MaxAge = 86400 * 30
	isProd = false
)

func NewAuth() *sessions.CookieStore {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	key := cfg.SessionSecret

	googleClientID := cfg.GoogleClientID
	googleClientSecret := cfg.GoogleClientSecret

	githubClientID := cfg.GithubClientID
	githubClientSecret := cfg.GithubClientSecret

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store
	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:8080/auth/google/complete"),
	)

	goth.UseProviders(
		github.New(githubClientID, githubClientSecret, "http://localhost:8080/auth/github/complete"),
	)

	return store
}
