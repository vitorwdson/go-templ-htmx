package handler

import (
	"bytes"
	"encoding/gob"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vitorwdson/go-templ-htmx/db"
)

type Session struct {
	ID        uuid.UUID
	User      db.User
	NextQuery time.Time
}

func (s server) authenticateUser(w http.ResponseWriter, r *http.Request, user db.User) error {
	sessionId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	session := Session{
		ID:        sessionId,
		User:      user,
		NextQuery: time.Now().Add(time.Minute * 5),
	}

	return s.saveSession(w, r, session)
}

func (s server) saveSession(w http.ResponseWriter, r *http.Request, session Session) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(session)
	if err != nil {
		return err
	}

	err = s.Redis.Set(
		r.Context(),
		"user-session:"+session.ID.String(),
		buf.String(),
		time.Duration(time.Duration.Hours(24)*7),
	).Err()
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "SESSION-KEY",
		Value: session.ID.String(),

		Path:    "/",
		Expires: time.Now().AddDate(0, 0, 7),

		Secure:   !s.DevMode,
		HttpOnly: true,
	})

	return nil
}

func (s server) GetSession(w http.ResponseWriter, r *http.Request) (*Session, error) {
	sessionCookie, err := r.Cookie("SESSION-KEY")
	if err != nil {
		return nil, err
	}

	sessionId := sessionCookie.Value
	sessionGob, err := s.Redis.Get(r.Context(), "user-session:"+sessionId).Result()
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

		user, err := s.DB.GetUserByID(r.Context(), session.User.ID)
		if err != nil {
			return nil, errors.New("Invalid user")
		}

		session.User = user
		session.NextQuery = time.Now().Add(time.Minute * 5)

		err = s.saveSession(w, r, session)
		if err != nil {
			return nil, err
		}
	}

	return &session, nil
}

func (s server) KillSession(w http.ResponseWriter, r *http.Request, session Session) error {
	err := s.Redis.Del(
		r.Context(),
		"user-session:"+session.ID.String(),
	).Err()
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "SESSION-KEY",
		Value: "",

		Path:    "/",
		Expires: time.Unix(0, 0),

		Secure:   !s.DevMode,
		HttpOnly: true,
	})

	return nil
}
