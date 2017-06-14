package configuration

import (
	"strings"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"

	"github.com/edgedagency/mulungu/core"
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/query"
)

var e *Entry

func init() {
	e = &Entry{}
}

//Entry this is a representation of a configuration entry
//
//e.g. email.from.address = mince@example.com
//e.g. application.name = Mince Inc Super App
type Entry struct {
	core.Model

	Key   string `json:"key" datastore:"key"`
	Value string `json:"value" datastore:"value"`
}

//NewEntryModel instantiates a new user model
func NewEntryModel(context context.Context, namespace string) *Entry {
	m := &Entry{}
	m.Init(context, "ConfigurationEntry", namespace)
	return m
}

//GetAll retireves all configuration
func (e *Entry) GetAll(filter string) ([]*Entry, error) {
	//get configuration entry by key
	return e.FindAll(filter)
}

//GetAll is a convience function so one can say configuration.GetAll(...)
//
func GetAll(context context.Context, namespace, filter string) ([]*Entry, error) {
	e.Init(context, "ConfigurationEntry", namespace)
	return e.FindAll(filter)
}

//Get retireves a configuration by key
func (e *Entry) Get(key, defaultValue string) string {
	entry := NewEntryModel(e.Context, e.Namespace)
	query := datastore.NewQuery(entry.Kind).Filter("key=", strings.TrimSpace(key)).Namespace(e.Namespace).Limit(1)

	logger.Debugf(e.Context, "configuration entry model", "query = %#v", query)
	result := e.Client().Run(e.Context, query)

	for {

		confKey, err := result.Next(entry)
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.Errorf(e.Context, "configuration entry model", "failed to retirve record, error %s return default value %s", err.Error(), defaultValue)
			return defaultValue
		}
		entry.Identify(confKey)
		logger.Debugf(e.Context, "configuration entry model", "Key=%v\n Record=%#v\n", confKey, entry)
	}

	if entry.IsNil() {
		logger.Warningf(e.Context, "configuration entry model", "configuration for key %s not found in namespace %s returning defaultValue %s", key, e.Namespace, defaultValue)
		return defaultValue
	}

	return entry.Value
}

//Get is a convience function so one can say configuration.Get(...)
//
func Get(context context.Context, namespace, key, defaultValue string) string {
	e.Init(context, "ConfigurationEntry", namespace)
	return e.Get(key, defaultValue)
}

// //Set sets or updates a configuration
// func (e *Entry) Set(key string, value string) *Entry {
//
// 	configurations := NewConfigurationEntryModel(e.Context, e.Namespace)
// 	if configurations.Exists(fmt.Sprintf("key=:%s", e.Key)) == true {
// 		return configurations.Update(0, e.Key, nil)
// 	}
// 	if configurations.Exists(fmt.Sprintf("key=:%s", e.Key)) == false {
// 		return configurations.Save(e.Key, nil)
// 	}
// 	//set  configuration entry with key
// 	// overrides existing entry, therefore check ig an entry with key exists update if true create new if false
// 	configurations.Value = value
// }

//FindAll returns all entries from datastore
func (e *Entry) FindAll(filter string) ([]*Entry, error) {
	queryBuilder := query.NewQueryBuilder(e.Context)
	queryBuilder.Build(e.Kind, filter, query.NoSorting)

	entries := make([]*Entry, 0)
	results := e.Run(queryBuilder.Query.Namespace(e.Namespace))

	for {
		resultModel := NewEntryModel(e.Context, e.Namespace)
		key, err := results.Next(resultModel)
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.Errorf(e.Context, "entry model", "failed to obtain results for entry iterator, error %s", err.Error())
		}
		logger.Debugf(e.Context, "entry model", "Key=%v\n Record=#v\n", key, resultModel)
		resultModel.Identify(key)
		entries = append(entries, resultModel)
	}
	return entries, nil
}

// //Exists returns all user from datastore
// func (e *Entry) Exists(filter string) bool {
// 	logger.Debugf(u.Context, "configuration model", "checking if record exists by filter: %s", filter)
//
// 	queryBuilder := query.NewQueryBuilder(e.Context)
// 	queryBuilder.Build(e.Kind, filter, "")
// 	queryBuilder.Query = queryBuilder.Query.Namespace(e.Namespace).KeysOnly().Limit(1)
// 	//only return keys to check if somthing exists based on filter
// 	results, err := e.Client().GetAll(e.Context, queryBuilder.Query, nil)
//
// 	if err != nil {
// 		logger.Errorf(e.Context, "configuration model", "results: %#v error %s", results, err.Error())
// 		return false
// 	}
//
// 	logger.Debugf(u.Context, "configuration model", "results: %#v count : %v", results, len(results))
//
// 	if len(results) > 0 {
// 		return true
// 	}
//
// 	return false
// }
