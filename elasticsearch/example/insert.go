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

	var saveData = map[string]interface{}{
		"id": "produk-1",
		"nama": "Roti Gembong Gede",
		"harga": 12999.00,
	}

	query := elastic.New(client, ctx)
	query.Index("students")
	query.Type("doc")
	if err := query.Insert(saveData); err != nil {
		fmt.Println("Error insert: ", err.Error())
		return 
	}

	fmt.Println("Result: ", resJson)
}

