package main

import (
    "github.com/julianzz/butterfly/deps/pty"
    "io"
    "os"
    "os/exec"
    "log"
)

func main() {
    c := exec.Command("bash")
    f, err := pty.Start(c)
    cols, rows , _ := pty.Getsize( f )
    log.Println( "fetch size:", cols, rows )
    pty.SetSize( f, 200,300 )
    cols, rows , _ = pty.Getsize( f )

    log.Println( "fetch size:", cols, rows )
    if err != nil {
        panic(err)
    }   

    go func() {
        f.Write([]byte("ls \n"))
        f.Write([]byte("date\n"))
        f.Write([]byte("time\n"))
        f.Write([]byte{4}) // EOT 
    }() 
    io.Copy(os.Stdout, f)

    c = exec.Command("date ")
    f2, err := pty.Start(c)
    f2.Write( []byte("ls -l \n"))
    io.Copy( os.Stdout, f2 );
}