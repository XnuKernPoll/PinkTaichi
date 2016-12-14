package utils
type IOMaster struct {
  Readers chan Work
  Writers chan Work
  Supervisors chan Work
}
type NetJobs interface {
  CreateReadWritePools()
  CreateSupervisorPools()
}

func (yeti IOMaster) CreateReadWritePools() {
  for x := 0 ; x < 4; x++  {
    go Writer(yeti.Writers)
  }
  for x := 0 ; x < 40; x ++ {
    go Reader(yeti.Readers)
  }
}
func (yeti IOMaster) CreateSupervisorPools() {
  for x := 0 ; x < 100; x ++ {
    go yeti.Supervisor(yeti.Supervisors)
  }
}

func CreateIOMaster() IOMaster {
  read := make(chan Work, 40)
  write := make(chan Work)
  Supervisors := make(chan Work, 100)
  massah := IOMaster{Readers: read, Writers: write, Supervisors: Supervisors}
  return massah
}
