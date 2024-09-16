package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/article"
	"github.com/geekengineers/Microservice-Project-Demo/protobuf/article/articleconnect"
	"github.com/geekengineers/Microservice-Project-Demo/services/blog/utils"
)

func main() {
	client := articleconnect.NewArticleServiceClient(http.DefaultClient, "http://localhost:8001")

	log.Println("Client established")

	res, err := client.Create(context.TODO(), &connect.Request[article.CreateRequest]{
		Msg: &article.CreateRequest{
			Title:       "Article 2",
			Description: "adads",
			Content:     "adads",
			CoverImage:  "assdasd",
		},
	})
	// res, err := client.Find(context.TODO(), &connect.Request[article.FindRequest]{
	// 	Msg: &article.FindRequest{
	// 		Id: 1,
	// 	},
	// })
	utils.HandleError(err)

	log.Println(res)
}
