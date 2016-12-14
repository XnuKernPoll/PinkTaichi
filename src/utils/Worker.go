package utils
import (
	"net"
	//"sync"
)
type Work struct {
	Connection net.Conn
	Documents *DocStore
	CMD *Command
}
