package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	p "part2/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	name = flag.String("name", "err", "Name to greet")
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := p.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/getUser/:searchquery", func(ctx *gin.Context) {
		searchquery := ctx.Param("searchquery")
		req := &p.SearchQuery{Query: searchquery}
		if response, err := client.GetUserByName(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Name),
			})
		} else {
			fmt.Println(searchquery)
			fmt.Println(err.Error())
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}

	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
