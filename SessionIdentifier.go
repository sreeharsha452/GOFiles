package session

import(
"net/http"
"time"
)

type SessionIdentifier string

type Session struct{
Id int
Username string
Expired time.Time
}

func InitSession(id sessionIdentifier) Session{
//...
}