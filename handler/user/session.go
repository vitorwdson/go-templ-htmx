package user

import (
	"bytes"
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	userModel "github.com/vitorwdson/go-templ-htmx/model/user"
)

type Session struct {
	ID   uuid.UUID
	User userModel.User
}

func (h UserHandler) authenticateUser(c echo.Context, user *userModel.User) error {
	sessionId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	session := Session{
		ID:   sessionId,
		User: *user,
	}

	return SaveSession(c, h.Redis, session)
}

func SaveSession(c echo.Context, r *redis.Client, session Session) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(session)
	if err != nil {
		return err
	}

	err = r.Set(
		c.Request().Context(),
		"user-session:"+session.ID.String(),
		buf,
		time.Duration(time.Duration.Hours(24)*7),
	).Err()
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:  "SESSION-KEY",
		Value: session.ID.String(),

		Path:    "/",
		Expires: time.Now().AddDate(0, 0, 7),

		Secure:   false, // TODO: Get this from env variable
		HttpOnly: true,
	})

	return nil
}

func GetSession(c echo.Context, r *redis.Client) (*Session, error) {
	sessionCookie, err := c.Cookie("SESSION-KEY")
	if err != nil {
		return nil, err
	}

	sessionId := sessionCookie.Value
	sessionGob, err := r.Get(
		c.Request().Context(),
		"user-session:"+sessionId,
	).Result()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte(sessionGob))
	dec := gob.NewDecoder(buf)

	var session Session
	err = dec.Decode(&session)
	if err != nil {
		return nil, err
	}

	if session.ID.String() != sessionId {
		return nil, errors.New("Invalid session ID")
	}

	return &session, nil
}
