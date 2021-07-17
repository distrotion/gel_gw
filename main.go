package main

import (
	"context"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"gogatewaydemo/jwt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/machinebox/graphql"
)

type Tokenschema struct {
	Token string `json:"token"`
}

type Gqlquery struct {
	Query string `json:"query"`
}

func main() {
	r := gin.Default()

	r.POST("/test1", func(c *gin.Context) {

		var input Gqlquery
		c.ShouldBind(&input)
		//-----------------------------------------

		str := fmt.Sprintf("%v", input.Query)
		token, err := jwt.GenerateToken(str)
		if err != nil {

		}

		//=========================================

		c.JSON(200, token)
	})

	r.POST("/test2", func(c *gin.Context) {

		var input Tokenschema
		c.ShouldBind(&input)

		//-----------------------------------------
		fmt.Println(input.Token)
		// str := fmt.Sprintf("%v", input)

		tokenStr := input.Token
		output, err := jwt.ParseToken(tokenStr)
		if err != nil {
			return
		}

		//=========================================

		c.JSON(200, output)
	})

	r.POST("/test3", func(c *gin.Context) {

		var input Gqlquery
		//c.ShouldBind(&input) BindJSON
		c.BindJSON(&input)

		//-----------------------------------------
		//fmt.Println(reflect.TypeOf(input.Query))
		fmt.Println(input.Query)
		// str := fmt.Sprintf("%v", input)

		graphqlClient := graphql.NewClient("http://34.101.230.112:27017/query")
		graphqlRequest := graphql.NewRequest(input.Query)
		var graphqlResponse interface{}
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			panic(err)
		}
		fmt.Println(graphqlResponse)

		//=========================================

		c.JSON(200, graphqlResponse)
	})

	r.POST("/test4", func(c *gin.Context) {

		var input Gqlquery
		c.ShouldBind(&input)

		//-----------------------------------------  decode

		tokenStr := input.Query
		output, err := jwt.ParseToken(tokenStr)
		if err != nil {
			return
		}

		//-----------------------------------------

		graphqlClient := graphql.NewClient("http://34.101.230.112:27017/query")
		graphqlRequest := graphql.NewRequest(output)
		var graphqlResponse interface{}
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			panic(err)
		}
		fmt.Println(graphqlResponse)

		//=========================================

		c.JSON(200, graphqlResponse)
	})

	r.POST("/test5", func(c *gin.Context) {

		var input Gqlquery
		c.ShouldBind(&input)

		//-----------------------------------------
		fmt.Println(reflect.TypeOf(input.Query))
		// str := fmt.Sprintf("%v", input)

		graphqlClient := graphql.NewClient("http://34.101.230.112:27017/query")
		graphqlRequest := graphql.NewRequest(input.Query)
		var graphqlResponse interface{}
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			panic(err)
		}
		fmt.Println(graphqlResponse)

		//=========================================

		c.JSON(200, graphqlResponse)
	})

	r.POST("/testinput1", func(c *gin.Context) {
		fmt.Println(c)
		var input interface{}
		c.BindJSON(&input)

		// var input Gqlquery
		//-----------------------------------------

		// input.Query = `test na ja`
		fmt.Println(input)

		//=========================================
		// b, _ := json.Marshal(input)
		// s := string(b)

		c.JSON(200, input)

	})

	r.OPTIONS("/testinput2", func(c *gin.Context) {

		var input interface{}
		c.BindJSON(&input)

		// var input Gqlquery
		//-----------------------------------------

		// input.Query = `test na ja`
		fmt.Println(input)

		//=========================================

		c.JSON(200, input)
	})

	r.Run(":9201")
}
