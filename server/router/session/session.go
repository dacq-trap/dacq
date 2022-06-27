package session

type ContextKey string

const (
	UserNameKey    ContextKey = "userName"
	AccessTokenKey ContextKey = "accessToken"
)
