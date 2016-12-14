package utils
import (
	"encoding/json"
	"time"
	"fmt"
)
type Command struct {
	Cmd string
	Database string
	Collection string
	Doc map[string]interface{}
	NewDoc map[string]interface{}
}

func HandleCommand(s chan Work) {

	defer func() {
				serv := <- s
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
						serv.Connection.Close()
        }

    }()

	for {
		serv := <- s
		//deadline :=  <- time.After(time.Now().Add(time.Second * 1))
		serv.Connection.SetDeadline(time.Now().Add(time.Second * 1))
		//defer //serv.Connection.Close()
		defer serv.Connection.Close()
		serv.Connection.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		d := json.NewDecoder(serv.Connection)
		var cmd Command
		d.Decode(&cmd)
		switch cmd.Cmd {
		case "Insert" :
			serv.Documents.Verify(cmd.Database)
			serv.Documents.Insert(cmd.Database, cmd.Collection, cmd.Doc)
			serv.Connection.Write([]byte("Added to store"))
			//serv.Connection.Close()
		case "Update One":
			if serv.Documents.checkDatabaseExistance(cmd.Database) && serv.Documents.checkCollectionExistance(cmd.Database, cmd.Collection) {
				serv.Documents.UpdateOne(cmd.Database, cmd.Collection, cmd.Doc, cmd.NewDoc)
				serv.Connection.Write([]byte("Updated"))
				serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}
		case "Update":
			if serv.Documents.checkDatabaseExistance(cmd.Database) && serv.Documents.checkCollectionExistance(cmd.Database, cmd.Collection) {
					serv.Documents.Update(cmd.Database, cmd.Collection, cmd.Doc, cmd.NewDoc)
					serv.Connection.Write([]byte("Updated"))
					serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}
		case "Find":

			v := serv.Documents.Retreive(cmd.Database, cmd.Collection, cmd.Doc)

			res, _ := json.Marshal(v)
			serv.Connection.Write(res)
			serv.Connection.Close()
		case "Find One":
			v := serv.Documents.Find(cmd.Database, cmd.Collection, cmd.Doc)
	 		res, _ := json.Marshal(v)
	 		serv.Connection.Write(res)
			serv.Connection.Close()
		case "Delete":
			if serv.Documents.checkDatabaseExistance(cmd.Database) && serv.Documents.checkCollectionExistance(cmd.Database, cmd.Collection){
				serv.Documents.Delete(cmd.Database, cmd.Collection, cmd.Doc)
				serv.Connection.Write([]byte("Deleted"))
			} else {
				serv.Connection.Close()
			}
		case "Delete One":
			if serv.Documents.checkDatabaseExistance(cmd.Database) && serv.Documents.checkCollectionExistance(cmd.Database, cmd.Collection) {
				serv.Documents.DeleteOne(cmd.Database, cmd.Collection, cmd.Doc)
				serv.Connection.Write([]byte("Deleted"))
				serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}
		case "Create Collection":
			if serv.Documents.checkDatabaseExistance(cmd.Database) {
				serv.Documents.CreateCollection(cmd.Database, cmd.Collection)
				resp := "Collection " + cmd.Collection + " created."
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()
			} else {
					serv.Connection.Close()
				}
		case "Drop Collection":
			if serv.Documents.checkDatabaseExistance(cmd.Database) && serv.Documents.checkCollectionExistance(cmd.Database, cmd.Collection){
				serv.Documents.DropCollection(cmd.Database, cmd.Collection)
				resp := "Collection " + cmd.Collection + " deleted."
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}
		case "Show Collections":
			if serv.Documents.checkDatabaseExistance(cmd.Database) {
				collectionList := serv.Documents.ShowCollections(cmd.Database)
				retList, _ := json.Marshal(collectionList)
				serv.Connection.Write(retList)
				serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}
		case "Create Database":
			if serv.Documents.checkDatabaseExistance(cmd.Database) != true {
				resp := "Database " + cmd.Database + " created."
				serv.Documents.CreateDatabase(cmd.Database)
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}

		case "Drop Database":
			resp := "Database " + cmd.Database + " deleted."
			if serv.Documents.checkDatabaseExistance(cmd.Database) {
				serv.Documents.DropDatabase(cmd.Database)
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()
			} else {
				serv.Connection.Close()
			}
		case "Show Databases":
			dblist := serv.Documents.ShowDatabases()
			marshList, _ := json.Marshal(dblist)
			serv.Connection.Write(marshList)
			serv.Connection.Close()
		case "List All":
			coll := serv.Documents.ListAll(cmd.Database, cmd.Collection)
			ml, _ := json.Marshal(coll)
			serv.Connection.Write(ml)
			serv.Connection.Close()

		default:
			serv.Connection.Close()
		}
	}
	select{}
}
