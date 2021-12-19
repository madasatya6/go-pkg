package main 

import (
	"fmt"
	"context"
	"encoding/json"

	elastic "github.com/madasatya6/go-pkg/elasticsearch"
)

func main(){
	esclient, err := elastic.GetESClient("http://localhost:9200")
	if err != nil {
		fmt.Println("Error init: ", err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	saveData := map[string]interface{}{
		"tipe": "produk-1",
		"nama": "Roti Gembong Gede",
		"harga": 12999,
	}

	dataJson, err := json.Marshal(saveData)
	if err != nil {
		panic(err)
	}

	bodyJson := string(dataJson)
	ind, err := esclient.Index().
		Index("students").
		BodyJson(bodyJson).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful ind: ", ind)
}

