package config

import "github.com/gorilla/sessions"

const SESSION_ID = "go-auth-sess"

var Store = sessions.NewCookieStore([]byte("asdfghjk1234"))
