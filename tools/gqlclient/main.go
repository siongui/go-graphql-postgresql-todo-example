package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

var strReq1 = `
query ($count: Int!, $page: Int!) {
  TodoPages(paginationInput: {count: $count, page: $page}) {
    pagination_info {
      total_count
      total_pages
      current_page
    }
    todos {
      id
      content_code
      created_date
      updated_date
      content_name
      description
      start_date
      end_date
      status
      created_by
      updated_by
    }
  }
}
`

func main() {
	client := graphql.NewClient("http://localhost:3005/query")

	req := graphql.NewRequest(strReq1)
	req.Var("count", 5)
	req.Var("page", 1)

	tokenStr := "1234"
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenStr))

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	fmt.Println(respData)
	//fmt.Println(respData["TodoPages"].(string))
}
