package action

import (
	"context"
	"log"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"local/pb"
)

type User struct {
	pb.UnimplementedUserServer
	MetaData *MetaData // 组合元数据类进来
}

func (self *User) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	var resp pb.TokenResponse

	// 验证请求参数
	if req.Username == "" || req.Password == "" {
		return &resp, status.Error(codes.InvalidArgument, "参数不完整")
	}

	// 生成accessToken
	var accessToken AccessToken
	accessToken.Data.Id = 123
	accessToken.Data.Username = "dxvgef"
	accessToken.ExpirationTime = jwt.NumericDate(time.Now().Add(time.Minute))
	token, err := makeJWT(&accessToken)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 响应数据
	resp.Token = token
	log.Println("服务端生成的token:", resp.Token)
	return &resp, nil
}

func (self *User) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	log.Println("收到客户端请求 User.List()")

	// 执行并获得认证结果
	if err := self.MetaData.AuthToken(ctx); err != nil {
		return nil, err
	}

	log.Println("客户端请求的token:", self.MetaData.TokenStr)

	if len(req.Columns) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Columns参数值无效")
	}

	var resp pb.ListResponse

	total, err := DB.Model().Table("users").
		Column(req.Columns...).
		SelectAndCount(&resp.Rows)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp.Total = int32(total)
	return &resp, nil
}
