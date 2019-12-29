package action

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"local/pb"
)

type User struct {
	pb.UnimplementedUserServer
}

func (self *User) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	var resp pb.TokenResponse
	if req.Username == "" || req.Password == "" {
		return &resp, status.Error(codes.InvalidArgument, "参数不完整")
	}
	resp.Token = "abcdefghijklmnopqrstuvwxyz"
	return &resp, nil
}

func (self *User) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
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
