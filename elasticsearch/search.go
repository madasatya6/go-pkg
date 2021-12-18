package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

/******
* @author madasatya6
* you now for search
*/

type Elastic struct{
	Conn *elastic.Client
	Ctx context.Context
}

func New(conn *elastic.Client, ctx context.Context) Elastic  {
	return Elastic{conn, ctx}
}

func (s *Elastic) Search(params map[string]interface{}, data *interface{}) {

	var result []data
	searchSource := elastic.NewSearchSource()

	if params == nil {
		searchSource.Query(elastic.NewMatchAllQuery())
	}
	if len(params) > 0 {
		bq := elastic.NewBoolQuery()
		bq = bq.Must(elastic.NewTermQuery("nama", search))
		bq = bq.Must(elastic.NewTermQuery("divisi", "backend"))
		searchSource.Query(bq)
	}

	/* this block will basically print out the es query */
	queryStr, err := searchSource.Source()
	if err != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal= ", err)
	}

	queryJs, err := json.Marshal(queryStr)
	if err != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal= ", err)
	}

	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* Until this block */

	searchService := esclient.Search().Index("students").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student Student 
		if err := json.Unmarshal(hit.Source, &student); err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, s := range students {
			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}
}