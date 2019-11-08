package core

import (
	"time"

	"github.com/actdid/mulungu/logger"
	"github.com/actdid/mulungu/provider"
	"github.com/actdid/mulungu/util"
	"golang.org/x/net/context"
)

//DatastoreModel struct representing a cloud function
type DatastoreModel struct {
	Context      context.Context        `json:"-"`
	DataProvider provider.DataProvider  `json:"-"`
	Namespace    string                 `json:"-"`
	Kind         string                 `json:"-"`
	Record       map[string]interface{} `json:"record,omitempty"`
}

//NewDatastoreModel returns a new datastore request struct pointer
func NewDatastoreModel(ctx context.Context, namespace, kind string, record map[string]interface{}) *DatastoreModel {
	datastoreRequest := &DatastoreModel{Context: ctx, Namespace: namespace, Kind: kind, Record: record, DataProvider: provider.NewArangodbDataProvider(ctx, namespace)}
	return datastoreRequest
}

//Save timestamps record
func (ds *DatastoreModel) Save() (map[string]interface{}, error) {
	ds.Timestamp()
	return ds.DataProvider.Save(ds.Kind, ds.RecordAsBytes())
}

//Update update model by id
func (ds *DatastoreModel) Update(id string) (map[string]interface{}, error) {
	ds.Timestamp()
	return ds.DataProvider.Update(ds.Kind, id, ds.RecordAsBytes())
}

//Delete delete model by id
func (ds *DatastoreModel) Delete(id string) (map[string]interface{}, error) {
	return ds.DataProvider.Delete(ds.Kind, id)
}

//Find find record based on provided identifier
func (ds *DatastoreModel) Find(id string) (map[string]interface{}, error) {
	return ds.DataProvider.Find(ds.Kind, id)
}

//FindAll search for and obtain records
func (ds *DatastoreModel) FindAll(searchParams map[string]string) ([]interface{}, error) {
	return ds.DataProvider.FindAll(ds.Kind, searchParams)
}

//Count count records based on filter
func (ds *DatastoreModel) Count(filter string) (int, error) {
	return ds.DataProvider.Count(ds.Kind, filter)
}

//Timestamp timestamps record
func (ds *DatastoreModel) Timestamp() {
	timestamp := time.Now()
	//if no createdDate create it
	if _, ok := ds.Record["createdDate"]; !ok {
		ds.Record["createdDate"] = timestamp
	}
	ds.Record["modifiedDate"] = timestamp

	logger.Debugf(ds.Context, "datastore model", "record timestamped %#v", ds.Record)
}

//JSONBytes returns this as JSONBytes
func (ds *DatastoreModel) JSONBytes() []byte {
	bytes, err := util.InterfaceToByte(ds)
	if err != nil {
		return nil
	}
	return bytes
}

//RecordAsBytes returns this as JSONBytes
func (ds *DatastoreModel) RecordAsBytes() []byte {
	bytes, err := util.InterfaceToByte(ds.Record)
	if err != nil {
		return nil
	}
	return bytes
}
