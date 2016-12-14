package main
import (
  "utils"
  "sync"
  "net"
  //fmt"
  "time"
)

func ServerInit() utils.IOMaster {

  handle := utils.CreateIOMaster()
  handle.CreateReadWritePools()
  handle.CreateSupervisorPools()
  time.Sleep(time.Second * 3)

  return handle

}

func main() {
  handler := ServerInit()
  docs := map[string]map[string][]map[string]interface{}{}
	docStore := utils.DocStore{DB: docs, State: &sync.RWMutex{}}
  c, _ := net.Listen("tcp", ":2000")
  for {
    conn, _ := c.Accept()
    job := utils.Work{Connection: conn, Documents: &docStore}
    handler.Supervisors <- job
  }
}
