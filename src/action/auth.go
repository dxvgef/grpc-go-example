package action

import (
	"context"
	"errors"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 元数据结构
type MetaData struct {
	TokenStr string       // token字符串
	Token    *AccessToken // token对象
}

// token结构
type AccessToken struct {
	Data struct {
		Id       int64  `json:"id,omitempty"`
		Username string `json:"username,omitempty"`
	} `json:"data,omitempty"`
	jwt.Payload
}

// // 获得元数据中的参数
// func (self *MetaData) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
// 	log.Println("执行了GetRequestMetadata")
// 	return map[string]string{"token": self.TokenStr}, nil
// }
//
// // 是否要求启用安全传输
// func (self *MetaData) RequireTransportSecurity() bool {
// 	return false
// }

// 获得token以及token校验结果
func (self *MetaData) AuthToken(ctx context.Context) error {
	// 获得元数据中的token参数值(token字符串)
	tokenStr, err := getMateData(ctx, "token")
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	// 校验token字符串并获得token对象
	token, err := verifyJWT(tokenStr)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	// token有效时赋值
	self.TokenStr = tokenStr
	self.Token = token

	return nil
}

// 获得元数据里的值
func getMateData(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("context不正确")
	}
	value := md.Get(key)
	if len(value) == 0 || value[0] == "" {
		return "", errors.New("无法获得token")
	}
	return value[0], nil
}

// 生成JWT
func makeJWT(t *AccessToken) (string, error) {
	alg := jwt.NewHS256([]byte("secret"))
	token, err := jwt.Sign(t, alg)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// 校验JWT
func verifyJWT(tokenStr string) (*AccessToken, error) {
	var accessToken AccessToken
	alg := jwt.NewHS256([]byte("secret"))
	_, err := jwt.Verify([]byte(tokenStr), alg, &accessToken,
		jwt.ValidatePayload(&accessToken.Payload, jwt.ExpirationTimeValidator(time.Now())),
	)
	if err != nil {
		return nil, err
	}
	return &accessToken, nil
}
