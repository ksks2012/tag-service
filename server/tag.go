package server

import (
	"context"
	"encoding/json"
	// "net/http"

	bapi "github.com/tag-service/pkg/bapi"
	"github.com/tag-service/pkg/errcode"
	pb "github.com/tag-service/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:18080/api/v1/tags?name=" + r.GetName())
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	// resp, err := http.Get("http://127.0.0.1:18080/api/v1/tags?name=" + r.GetName())
	// if err != nil {
	// 	return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	// }

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}
