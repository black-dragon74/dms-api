package api

type Session struct {
	sid string
}

func (s Session) Validate() bool {
	return false
}

func (s Session) GetID() string {
	return s.sid
}

func NewSession(sessionID string) Session {
	return Session{
		sid: sessionID,
	}
}
