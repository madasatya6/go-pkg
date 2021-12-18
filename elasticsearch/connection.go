package elasticsearch

import (
	"fmt"
	elastic "github.com/olivere/elastic/v7"
)

/******
* @author madasatya6
*/
func GetESClient(url string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false))
	
	fmt.Println("Elastic Search initialized...")
	return client, err 
}

