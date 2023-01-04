package models

import post_crud_pb "github.com/acool-kaz/api-gateway-service/pkg/post_crud/pb"

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func FromPostToProto(p Post) *post_crud_pb.Post {
	return &post_crud_pb.Post{
		Id:     int32(p.Id),
		UserId: int32(p.UserId),
		Title:  p.Title,
		Body:   p.Body,
	}
}

func FromProtoToPost(post *post_crud_pb.Post) Post {
	return Post{
		Id:     int(post.GetId()),
		UserId: int(post.GetUserId()),
		Title:  post.GetTitle(),
		Body:   post.GetBody(),
	}
}

type Update struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
