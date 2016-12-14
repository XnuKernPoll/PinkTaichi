package utils
//import "reflect"
import "sync"
//import "encoding/json"
//import "strings"
//import "time"
type Operations interface {
	Find(database string, collection string,  doc map[string]interface{}, NumDocs string) []map[string]interface{}
	Insert(database string, collection string,  doc map[string]interface{}, NumDocs string)
	Delete(database string, collection string,  doc map[string]interface{}, NumDocs string)
	Update(database string, collection string, doc map[string]interface{}, NewDoc map[string]interface{})
	CreateCollection(database string, collection string)
	DropCollection(database string, collection string)
	ShowCollections(database string) []string
	ShowDatabases() []string
	DropDatabase(database string)
	CreateDatabase(database string)
}

type DocStore struct {
	State *sync.RWMutex
	DB	map[string]map[string][]map[string]interface{}

	//Collections map[string][]map[string]interface{}
}
func (d *DocStore) checkDatabaseExistance(database string) bool {
//	var key map[string]map[string][]map[string]interface{}
	//var chk bool
	//d.State.RLock()
	if _, ok := d.DB[database]; ok {
		//d.State.RUnlock()
		return true
	} else {
//		d.State.RUnlock()
		return false
	}
}

func (d *DocStore) checkCollectionExistance(database string, collection string) bool {
	//defer d.State.Unlock()
	//key, ok = d.DB[database][collection]
	//d.State.RLock()
	//d.State.RUnlock()


	if _, ok := d.DB[database][collection]; ok {
		return true
	} else {
		return false
	}
}

func (d *DocStore) ShowDatabases() []string {
	//defer d.State.Unlock()
	//d.State.RLock()
	dbs := []string{}
	for k, _ := range d.DB {
		dbs = append(dbs, k)
	}
	//d.State.RUnlock()
	return dbs
}

func (d *DocStore) Retreive(database string, collection string, doc map[string]interface{}) []map[string]interface{} {
	var matchedDocuments []map[string]interface{}
	li := d.AllFinder(database, collection, doc)
	for _, x := range li {
		matchedDocuments = append(matchedDocuments, d.DB[database][collection][x])
	}
	return matchedDocuments
}

func (d *DocStore) FindOne(database string, collection string, doc map[string]interface{}) map[string]interface{} {
	fil := d.OneFinder(database, collection, doc)
	res := d.DB[database][collection][fil]
	return res
}

func (d *DocStore) Find(database string, collection string,  doc map[string]interface{}) []map[string]interface{} {
	var matchedDocuments []map[string]interface{}
	//matchedDocuments = []map[string]interface{}
	var breaker string
	for x := range d.DB[database][collection] {
		for z, y := range doc{
			if d.DB[database][collection][x][z] == y {
					//fmt.Println("cool")
				continue
			} else {
					breaker = "no match"
					break
			}
		}
		if breaker != "no match" {
			matchedDocuments = append(matchedDocuments, d.DB[database][collection][x])

		}
	}
			return matchedDocuments
}

func (d *DocStore) Insert(database string, collection string,  doc map[string]interface{}) {
	//d.State.Lock()
	d.DB[database][collection] = append(d.DB[database][collection], doc)
	return
	//d.State.Unlock()
}
func (d *DocStore) DeleteOne(database string, collection string, doc map[string]interface{}) {
	//d.State.Lock()
	//defer d.State.Unlock()
	of := d.OneFinder(database, collection, doc)
	d.DB[database][collection] = d.DB[database][collection][:of+copy(d.DB[database][collection][of:], d.DB[database][collection][of+1:])]
}

func (d *DocStore) Delete(database string, collection string, doc map[string]interface{}) {
//	d.State.Lock()
	//defer d.State.Unlock()
	s := d.AllFinder(database, collection, doc)
	for _, x := range s {				//fmt.Println(matching[sd])
		d.DB[database][collection] = append(d.DB[database][collection][:x], d.DB[database][collection][x+1:]...)
	}
//	d.State.Unlock()
}
