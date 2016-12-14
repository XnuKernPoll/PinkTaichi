package utils
//import "fmt"
func (d *DocStore) DropCollection(database string, collection string) {
  //d.State.Lock()
  //d.State.Unlock()
  delete(d.DB[database], collection)
  //d.State.Unlock()
}

func (d *DocStore) CreateCollection(database string, collection string) {
  //d.State.Lock()
  emptyColl := []map[string]interface{}{}
  d.DB[database][collection] = emptyColl
  //d.State.Unlock()
}
func (d *DocStore) CreateDatabase(database string) {
  //d.State.Lock()
  emptyDB  := map[string][]map[string]interface{}{}
  d.DB[database] = emptyDB
  //d.State.Unlock()
}
func (d *DocStore) DropDatabase(database string) {
  //d.State.Lock()
  delete(d.DB, database)
  //d.State.Unlock()
}

func (d *DocStore) ShowCollections(database string) []string {
  //d.State.RLock()
  var retList []string
  for k := range d.DB[database] {
    retList = append(retList, k)
  }
  //d.State.RUnlock()
  return retList
}
func (d *DocStore) Verify(database string){
  if d.checkDatabaseExistance(database) {
    return
  } else {
    d.CreateDatabase(database)
  }
}
