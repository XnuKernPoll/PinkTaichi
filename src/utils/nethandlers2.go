package utils
//import "time"
func Writer(ch chan Work) {

  for {
    serv := <- ch
    defer func() {
    if r := recover(); r != nil {
        serv.Documents.State.Unlock()
        serv.Connection.Close()
    }
  }()
    switch serv.CMD.Cmd {
    case "Delete":
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) && serv.Documents.checkCollectionExistance(serv.CMD.Database, serv.CMD.Collection){
        serv.Documents.State.Lock()
        serv.Documents.Delete(serv.CMD.Database, serv.CMD.Collection, serv.CMD.Doc)
        serv.Documents.State.Unlock()
        serv.Connection.Write([]byte("Deleted"))
        serv.Connection.Close()
      } else {
        serv.Connection.Close()
      }
    case "Delete One":
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) && serv.Documents.checkCollectionExistance(serv.CMD.Database, serv.CMD.Collection) {
        serv.Documents.State.Lock()
        serv.Documents.DeleteOne(serv.CMD.Database, serv.CMD.Collection, serv.CMD.Doc)
        serv.Documents.State.Unlock()
        serv.Connection.Write([]byte("Deleted"))
        serv.Connection.Close()
      } else {
        serv.Connection.Close()
      }
    case "Create Collection":
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) {
        serv.Documents.State.Lock()
				serv.Documents.CreateCollection(serv.CMD.Database, serv.CMD.Collection)
        serv.Documents.State.Unlock()
				resp := "Collection " + serv.CMD.Collection + " created."
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()

			} else {
				    serv.Connection.Close()
			}
    case "Update One":
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) && serv.Documents.checkCollectionExistance(serv.CMD.Database, serv.CMD.Collection) {
        serv.Documents.State.Lock()
        serv.Documents.UpdateOne(serv.CMD.Database, serv.CMD.Collection, serv.CMD.Doc, serv.CMD.NewDoc)
        serv.Documents.State.Unlock()
        serv.Connection.Write([]byte("updated"))
        serv.Connection.Close()
      } else {
        serv.Connection.Close()
      }
    case "Update":
      if  serv.Documents.checkDatabaseExistance(serv.CMD.Database) && serv.Documents.checkCollectionExistance(serv.CMD.Database, serv.CMD.Collection) {
        serv.Documents.State.Lock()
        serv.Documents.Update(serv.CMD.Database, serv.CMD.Collection, serv.CMD.Doc, serv.CMD.NewDoc)
        serv.Documents.State.Unlock()
        serv.Connection.Write([]byte("Updated"))
        serv.Connection.Close()
      } else {
        serv.Connection.Close()
      }
    case "Drop Collection":
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) && serv.Documents.checkCollectionExistance(serv.CMD.Database, serv.CMD.Collection){
        serv.Documents.State.Lock()
        serv.Documents.DropCollection(serv.CMD.Database, serv.CMD.Collection)
        serv.Documents.State.Unlock()
				resp := "Collection " + serv.CMD.Collection + " deleted."
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()
      } else {
				serv.Connection.Close()
			}

    case "Create Database":
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) != true {
				resp := "Database " + serv.CMD.Database + " created."
        serv.Documents.State.Lock()
				serv.Documents.CreateDatabase(serv.CMD.Database)
        serv.Documents.State.Unlock()
				serv.Connection.Write([]byte(resp))
				serv.Connection.Close()
      } else {
				serv.Connection.Close()
			}

    case "Drop Database":
      resp := "Database " + serv.CMD.Database + " deleted."
      if serv.Documents.checkDatabaseExistance(serv.CMD.Database) {
        serv.Documents.State.Lock()
        serv.Documents.DropDatabase(serv.CMD.Database)
        serv.Documents.State.Unlock()
        serv.Connection.Write([]byte(resp))
        serv.Connection.Close()
        } else {
        serv.Connection.Close()
      }
    case "Insert":
        serv.Documents.Verify(serv.CMD.Database)
        serv.Documents.State.Lock()
        serv.Documents.Insert(serv.CMD.Database, serv.CMD.Collection, serv.CMD.Doc)
        serv.Documents.State.Unlock()
        serv.Connection.Write([]byte("Added to store"))
        serv.Connection.Close()

    default:
      serv.Connection.Close()
    }
  }
}
