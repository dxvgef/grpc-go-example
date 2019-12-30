package action

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"local/pb"
)

type Content struct {
	pb.UnimplementedUserServer
	MetaData *MetaData // 组合元数据类进来
}

func (self *Content) Add(ctx context.Context, req *pb.ContentAddRequest) (*pb.ContentAddResponse, error) {
	var resp pb.ContentAddResponse
	resp.Id = 1211590588382576640
	return &resp, nil
}

func (self *Content) List(ctx context.Context, req *pb.ContentListRequest) (*pb.ContentListResponse, error) {

	// 执行并获得认证结果
	if err := self.MetaData.AuthToken(ctx); err != nil {
		return nil, err
	}

	log.Println("客户端请求的token:", self.MetaData.TokenStr)

	if len(req.Columns) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Columns参数值无效")
	}

	var resp pb.ContentListResponse

	total, err := DB.Model().Table("content").
		Column(req.Columns...).
		SelectAndCount(&resp.Rows)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp.Total = int32(total)
	return &resp, nil
}
