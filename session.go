package butterfly

import ( 
	"github.com/googollee/go-socket.io"
	"github.com/kr/pty"
	"os"
)

type IoSession struct  {
	ns *socketio.NameSpace
	pty *os.File
}