package provider

//DataProvider interface for data providers
type DataProvider interface {
	//Save save record
	Save(collectionName string, data []byte) (map[string]interface{}, error)

	//Update update record
	Update(collectionName string, id string, data []byte) (map[string]interface{}, error)

	//Delete delete record
	Delete(collectionName string, id string) (map[string]interface{}, error)

	//Find find record based on provided identifier
	Find(collectionName string, id string) (map[string]interface{}, error)

	//FindAll search for and obtain records
	FindAll(collectionName string, searchParams map[string]string) ([]interface{}, error)

	//FindAll search for and obtain records
	Count(collectionName string, filter string) (int, error)
}
