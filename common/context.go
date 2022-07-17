package common

import (
	"context"
)

type ContextData struct {
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	IsAdmin   bool   `json:"is_admin"`
	TokenID   string `json:"token_id"`
	TokenExp  int64  `json:"token_exp"`
	TokenIss  string `json:"token_iss"`
	UserAgent string `json:"user_agent"`
	Group     string `json:"group"`
}

func SetContextMetadata(ctx context.Context, data map[string]interface{}) context.Context {
	for k, v := range data {
		ctx = context.WithValue(ctx, k, v)
	}

	return ctx
}

func GetDataFromContext(ctx context.Context) *ContextData {
	ctxData := &ContextData{}
	if ctx.Value("token") != nil {
		ctxData.Token = ctx.Value("token").(string)
	}
	if ctx.Value("x-user-id") != nil {
		ctxData.UserID = ctx.Value("x-user-id").(string)
	}
	if ctx.Value("is_admin") != nil {
		ctxData.IsAdmin = ctx.Value("is_admin").(bool)
	}
	if ctx.Value("email") != nil {
		ctxData.UserEmail = ctx.Value("email").(string)
	}
	if ctx.Value("id") != nil {
		ctxData.TokenID = ctx.Value("id").(string)
	}
	if ctx.Value("exp") != nil {
		ctxData.TokenExp = int64(ctx.Value("exp").(float64))
	}
	if ctx.Value("iss") != nil {
		ctxData.TokenIss = ctx.Value("iss").(string)
	}
	if ctx.Value("user-agent") != nil {
		ctxData.UserAgent = ctx.Value("user-agent").(string)
	}
	if ctx.Value("group") != nil {
		ctxData.Group = ctx.Value("group").(string)
	}

	return ctxData
}
