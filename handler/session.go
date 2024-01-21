package handler

import (
	"bytes"
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	userModel "github.com/vitorwdson/go-templ-htmx/model/user"
)

type Session struct {
	ID        uuid.UUID
	User      userModel.User
	NextQuery time.Time
}

func (h Handler) authenticateUser(c echo.Context, user userModel.User) error {
	sessionId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	session := Session{
		ID:        sessionId,
		User:      user,
		NextQuery: time.Now().Add(time.Minute * 5),
	}

	return h.SaveSession(c, session)
}

func (h Handler)  SaveSession(c echo.Context, session Session) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(session)
	if err != nil {
		return err
	}

	err = h.Redis.Set(
		c.Request().Context(),
		"user-session:"+session.ID.String(),
		buf.String(),
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

func (h Handler)  GetSession(c echo.Context) (*Session, error) {
	sessionCookie, err := c.Cookie("SESSION-KEY")
	if err != nil {
		return nil, err
	}

	sessionId := sessionCookie.Value
	sessionGob, err := h.Redis.Get(
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

	if session.NextQuery.Compare(time.Now()) == -1 {

		user, err := userModel.GetByID(h.DB, session.User.ID)
		if err != nil || user == nil {
			return nil, errors.New("Invalid user")
		}

		session.User = *user
		session.NextQuery = time.Now().Add(time.Minute * 5)

		err = h.SaveSession(c, session)
		if err != nil {
			return nil, err
		}
	}

	return &session, nil
}

func (h Handler)  KillSession(c echo.Context, session Session) error {
	err := h.Redis.Del(
		c.Request().Context(),
		"user-session:"+session.ID.String(),
	).Err()
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:  "SESSION-KEY",
		Value: "",

		Path:    "/",
		Expires: time.Unix(0, 0),

		Secure:   false, // TODO: Get this from env variable
		HttpOnly: true,
	})

	return nil
}
