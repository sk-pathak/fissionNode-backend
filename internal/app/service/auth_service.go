package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type AuthService struct {
	Store *sessions.CookieStore
}

func NewAuthService(secret []byte) *AuthService {
	store := sessions.NewCookieStore(secret)
	return &AuthService{
		Store: store,
	}
}

func (s *AuthService) BeginAuth(provider string, c *gin.Context) (string, error) {
	gothic.Store = s.Store

	url, err := gothic.GetAuthURL(c.Writer, c.Request)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *AuthService) CompleteAuth(c *gin.Context) (goth.User, error) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		return goth.User{}, err
	}
	return user, nil
}

func (s *AuthService) SaveUserSession(c *gin.Context, user goth.User) error {
	session, err := s.Store.Get(c.Request, "user-session")
	if err != nil {
		return err
	}

	session.Values["user"] = user
	if err := session.Save(c.Request, c.Writer); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) GetUserSession(c *gin.Context) (goth.User, bool) {
	session, err := s.Store.Get(c.Request, "user-session")
	if err != nil {
		return goth.User{}, false
	}

	user, ok := session.Values["user"].(goth.User)
	return user, ok
}

func (s *AuthService) Logout(c *gin.Context) error {
	session, err := s.Store.Get(c.Request, "user-session")
	if err != nil {
		return err
	}

	delete(session.Values, "user")
	if err := session.Save(c.Request, c.Writer); err != nil {
		return err
	}
	return nil
}
