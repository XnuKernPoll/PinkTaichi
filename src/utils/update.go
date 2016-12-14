package utils
import "reflect"
//import "encoding/json"
func (d *DocStore) Update(database string, collection string, doc map[string]interface{}, NewDoc map[string]interface{}){
//  d.State.Lock()
  var keys []string
  var vals []interface{}
  var NewKeys []string
  var NewVals []interface{}
  //matchedDocuments = []map[string]interface{}
  var breaker string
  for k, _ := range doc {
    keys = append(keys, k)
    vals = append(vals, doc[k])
  }
  for k, _ := range NewDoc {
    NewKeys = append(NewKeys, k)
    NewVals = append(NewVals, NewDoc[k])
  }
  for x := range d.DB[database][collection] {
    for z := range keys {
      if d.DB[database][collection][x][keys[z]] == vals[z] {
      //	fmt.Println("cool")
        continue
        } else {
        breaker = "no match"
        break
      }
    }
    if breaker != "no match" {
      for nk := range NewKeys {
        d.DB[database][collection][x][NewKeys[nk]] = NewVals[nk]
      }
    }
  }
  //d.State.Unlock()
}



func (d *DocStore) AllFinder (database string, collection string, doc map[string]interface{}) []int {
  var ind []int
  //var docty map[string]interface{}

  for x := range d.DB[database][collection] {
    var breaker string
    for k, v := range doc {
        if reflect.DeepEqual(d.DB[database][collection][x][k], v) {
          continue
        } else {
          breaker = "no match"
          //break

        }
      }
    if breaker != "no match" {
      ind = append(ind, x)
    }
  }
  return ind
}
func (d *DocStore) ListAll(database string, collection string) []map[string]interface{} {
  //fmt.Println(d.DB[database][collection])
  return d.DB[database][collection]

}

func (d *DocStore) OneFinder(database string, collection string, doc map[string]interface{}) int {
  var keys []string
  var vals []interface{}
  var breaker string
  for k, v := range doc {
    keys = append(keys, k)
    vals = append(vals, v)
  }
  var ind int
  for x := range  d.DB[database][collection] {
    for z := range keys {
      if d.DB[database][collection][x][keys[z]] == vals[z] {
      //	fmt.Println("cool")
        continue
        } else {
        breaker = "no match"
        break
      }
    }
    if breaker != "no match"{
      ind = x

    }
  }
  return ind
}


func (d *DocStore) UpdateOne(database string, collection string, doc map[string]interface{}, NewDoc map[string]interface{}) {
  //d.State.Lock()
  of := d.OneFinder(database, collection, doc)
  var NewKeys []string
  var NewVals []interface{}
  for k, _ := range NewDoc {
    //fmt.Println(k)
    //fmt.Println(NewDoc[k]
    NewKeys = append(NewKeys, k)
    NewVals = append(NewVals, NewDoc[k])
  }
    for x := range NewKeys {
      d.DB[database][collection][of][NewKeys[x]] = NewVals[x]
    }
}
