package middleware

type contextKey string

const (
	ContextKeyUserID    contextKey = "userID"
	ContextKeySessionID contextKey = "sessionID"
)
