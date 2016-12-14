package utils
import (
  "encoding/json"
  //"runtime"

)
func (Mas *IOMaster)Supervisor(ch chan Work) {


  for {
    serving := <- ch
    defer func() {
    if r := recover(); r != nil {
        serving.Connection.Close()
    }
  }()
    decoder := json.NewDecoder(serving.Connection)
    var cmd Command
    decoder.Decode(&cmd)
    if (cmd.Cmd == "Insert")||(cmd.Cmd == "Update")||(cmd.Cmd == "Update One")||(cmd.Cmd == "Delete One")|| (cmd.Cmd == "Delete") || (cmd.Cmd == "Create Collection")||(cmd.Cmd == "Drop Collection")||(cmd.Cmd == "Create Database")||(cmd.Cmd == "Drop Database") {
      job := Work{CMD: &cmd, Connection: serving.Connection, Documents: serving.Documents}
      Mas.Writers <- job
    } else if (cmd.Cmd == "Find") || (cmd.Cmd == "Find One") || (cmd.Cmd == "Show Collections") || (cmd.Cmd == "Show Databases") || (cmd.Cmd == "List All") {
      job := Work{CMD: &cmd, Connection: serving.Connection, Documents: serving.Documents}
      Mas.Readers <- job
    } else {
      serving.Connection.Close()
    }
  }
}


func Reader(c chan Work) {
  for {

    server := <- c
    defer func() {
    if r := recover(); r != nil {
        server.Documents.State.RUnlock()
        server.Connection.Close()
    }
  }()
    switch server.CMD.Cmd {
    case "Find":
      server.Documents.State.RLock()
      //defer server.Documents.State.RUnlock()
      docs := server.Documents.Retreive(server.CMD.Database, server.CMD.Collection, server.CMD.Doc)
      server.Documents.State.RUnlock()
      results, _ := json.Marshal(docs)
      server.Connection.Write(results)
      server.Connection.Close()
    case "Find One":
      server.Documents.State.RLock()
      //defer server.Documents.State.RUnlock()
      docs := server.Documents.FindOne(server.CMD.Database, server.CMD.Collection, server.CMD.Doc)
      server.Documents.State.RUnlock()
      results, _ := json.Marshal(docs)
      server.Connection.Write(results)
      server.Connection.Close()
    case "Show Collections":
      if server.Documents.checkDatabaseExistance(server.CMD.Database) {
        server.Documents.State.RLock()
        //defer server.Documents.State.RUnlock()
				collectionList := server.Documents.ShowCollections(server.CMD.Database)
        server.Documents.State.RUnlock()
				retList, _ := json.Marshal(collectionList)
				server.Connection.Write(retList)
        server.Connection.Close()
			} else {
				server.Connection.Close()
			}
    case "Show Databases":
      server.Documents.State.RLock()
      //defer server.Documents.State.RUnlock()
      dblist := server.Documents.ShowDatabases()
      server.Documents.State.RUnlock()
      marshList, _ := json.Marshal(dblist)
      server.Connection.Write(marshList)
      server.Connection.Close()

    case "List All":
      server.Documents.State.RLock()
      //defer server.Documents.State.RUnlock()
    //  fmt.Println("ayyy")
      dc := server.Documents.ListAll(server.CMD.Database, server.CMD.Collection)
      server.Documents.State.RUnlock()
      res, _ := json.Marshal(dc)
      server.Connection.Write(res)
      server.Connection.Close()
    default:
      server.Connection.Close()
  }
 }
}
