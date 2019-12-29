package action

import (
	"context"
	"log"
)

type Authentication struct {
	Token string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	log.Println("执行了验证")
	return map[string]string{"token": a.Token}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
