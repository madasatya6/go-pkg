package main 

import (
	"fmt"
	"context"

	elastic "github.com/madasatya6/go-pkg/elasticsearch"
)

func main(){
	client, err := elastic.GetESClient("http://localhost:9200")
	if err != nil {
		fmt.Println("Error init: ", err.Error())
	}

	ctx := context.Background()
	
	search := elastic.New(client, ctx)
	search.Index("students")
	search.Type("doc")
	resJson := search.Search(nil)

	fmt.Println("Result: ", resJson)
}

