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

    /*go func() {
        f.Write([]byte("ls \n"))
        f.Write([]byte("date\n"))
        f.Write([]byte("time\n"))
        f.Write([]byte{4}) // EOT
    }()*/
    //go io.Copy(os.Stdout, f)

    c = exec.Command("bash")
    f2, err := pty.Start(c)
    f2.Write( []byte("ls -l \n"))
    
    const READSIZE = 1024
    /*go func(){
        for {
            bytes     := make([]byte, READSIZE)
            read, err := f2.Read(bytes)
            if (err != nil) {
                os.Stdout.Write( []byte("finish ----- ") )
                return
            }
            os.Stdout.Write(bytes[:read])
        }
    }()
    */
    go func() {
        log.Println( io.Copy( os.Stdout, f2 ) ) 

    }()

    io.Copy(f2, os.Stdin)
}