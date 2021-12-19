package elasticsearch

import (
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic/v7"
)

/******
* @author madasatya6
* insert data
*/
func (s *Elastic) Insert(data map[string]interface{}) error {
	var service *elastic.IndexService
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err 
	}

	bodyJson := string(dataJson)
	
	if s.TypeName != "" {
		service = s.Conn.Index().Index(s.IndexName).Type(s.TypeName).BodyJson(bodyJson)
	} else {
		service = s.Conn.Index().Index(s.IndexName).BodyJson(bodyJson)
	}
	
	response, err := service.Do(s.Ctx)
	if err != nil {
		return err 
	}
	
	fmt.Println("[Elastic][InsertProduct]Insertion Successful response: ", response)
	return nil
}