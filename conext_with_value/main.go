package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func main() {
	processRequest("jane", "abc123")
}

// UserID - return user id
func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

// AuthToken - returns token
func AuthToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func processRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	handleResponse(ctx)
}

func handleResponse(ctx context.Context) {
	fmt.Printf("handling response for %v %v", UserID(ctx), AuthToken(ctx))
}
