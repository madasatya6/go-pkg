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
	IndexName string 
	TypeName string 
}

func New(conn *elastic.Client, ctx context.Context) Elastic  {
	return Elastic{conn, ctx}
}

func (s *Elastic) Index(name string) {
	s.IndexName = name 
}

func (s *Elastic) Type(name string) {
	s.TypeName = name 
}

func (s *Elastic) Search(params map[string]interface{}, data *interface{})  {

	var result []data
	searchSource := elastic.NewSearchSource()

	if params == nil {
		searchSource.Query(elastic.NewMatchAllQuery())
	}
	if len(params) > 0 {
		bq := elastic.NewBoolQuery()
		for key, value := range params {
			bq = bq.Must(elastic.NewTermQuery(key, value))
		}
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
	
	if s.TypeName == "" {
		searchService := s.Conn.Search().Index(s.IndexName).SearchSource(searchSource)
	} else {
		searchService := s.Conn.Search().Index(s.IndexName).Type(s.TypeName).SearchSource(searchSource)
	}

	searchResult, err := searchService.Do(s.Ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits { 
		if err := json.Unmarshal(hit.Source, &data); err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		result = append(result, data)
	}

	if err != nil {
		fmt.Println("Fetching data fail: ", err)
		return 
	} else {
		for _, s := range  {
			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}
}