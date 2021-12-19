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
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	saveData := map[string]interface{}{
		"tipe": "produk-1",
		"nama": "Roti Gembong Gede",
		"harga": 12999,
	}

	query := elastic.New(client, ctx)
	query.Index("produk")
	query.Type("doc")
	err = query.Insert(saveData);
	if err != nil {
		fmt.Println("Error insert: ", err.Error())
		return 
	}

	fmt.Println("Result: berhasil insert elaticsearch")
}

