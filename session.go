package butterfly

import ( 
	"github.com/googollee/go-socket.io"
	_ "github.com/kr/pty"
	"os"
)

type IoSession struct  {
	ns *socketio.NameSpace
	pty *os.File
}