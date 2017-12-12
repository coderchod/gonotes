package main

import(
        "database/sql"
        "fmt"
        _"github.com/mattn/go-sqlite3"
	"os"
	"os/user"
	"strings"
)

func main(){

        db, err := sql.Open("sqlite3", user.HomeDir + "/gonotes/notes.db")
        checkErr(err)

        count := len(os.Args)

        if count > 1{
        firstArg := os.Args[1]

        switch firstArg{
        case "-list":
                //fmt.Println("List all notes")
		if count == 2 {
                	rows, err := db.Query("SELECT *  FROM lists")
                	checkErr(err)
                        fmt.Println("\n Existing lists are: ")

                	for rows.Next(){
				var id int
				var listName string
                        	err := rows.Scan(&id, &listName)
                        	checkErr(err)
				fmt.Println("\n " + listName)

                	}
		fmt.Println("\n ``````````````````````````````")
		fmt.Println("\n For notes in the list, type: ")
		fmt.Println("gonotes -list <list name> \n")
		} else{
			list := os.Args[2]
			processList(list, db)
			}
        default:
                processCrap(os.Args[1:], db)
        }
        }else{
                helpText()
        }
	db.Close()
}
func processCrap(args []string, db *sql.DB){
        c := len(args)
        
	firstId := 0
	arg := args[firstId]	
	
	var name string
	var Lid int
	if string(arg[0]) == "@"{
 		
		if c > 2 {
			listTag := arg
			fmt.Println("\n Added to list " + listTag)
		
         	       rows, err := db.Query("SELECT uid  FROM lists where name =?", string(listTag))
               	       checkErr(err)
		
		        for rows.Next(){
				err := rows.Scan(&Lid)
				//fmt.Println(Lid)
				checkErr(err)
			} 
			if Lid == 0{
				//fmt.Println("not found...creating")
				stmt, err := db.Prepare("INSERT INTO lists('name') values (?)")
				checkErr(err)
			
				_, err = stmt.Exec(string(listTag))
				checkErr(err)
			}else{
				//fmt.Println("List found.. ID: " + string(Lid))
			}
			firstId += 1
			name = args[firstId]
		
		} else {
			helpText()
		}
		}else{
		name = args[firstId]
		}

        if c>1{
        msg := args[firstId +1 :]
	messagewa := strings.Join(msg," ")
        ps := name + ": " + messagewa
        
	fmt.Println("\n Added: \n " + ps)

	//check if exists, if yes, update
        stmt, err := db.Prepare("update notes set message=?,created=CURRENT_TIMESTAMP,listid=? where name=?")
        checkErr(err)

        res,err := stmt.Exec(messagewa, Lid, name)
        checkErr(err)

        affect, err := res.RowsAffected()
        checkErr(err)

	if affect == 0{
		stmt, err := db.Prepare("INSERT INTO notes(name, message, created, listid) values(?,?,CURRENT_TIMESTAMP, ?)")
        	checkErr(err)

        	_, err = stmt.Exec(name, messagewa, Lid)
        	checkErr(err)

	        //id, err := res.LastInsertId()
        	//checkErr(err)

	        //fmt.Println(id)
	}
        }else{
                //fmt.Println("Check db for msg.. if not found open new multiline")
		rows, err := db.Query("SELECT * FROM notes where name =?", name)
        	checkErr(err)

        for rows.Next(){
                var uid int
                var name string
                var message string
                var created string
		var listid int
                err = rows.Scan(&uid, &name, &message, &created, &listid)
                checkErr(err)
                fmt.Println("\n " + name + " , " + created)
		fmt.Println("\n `````````````````````````````````")
		fmt.Println("\n " + message + "\n ")
        }

        }
}

func processList(list string, db *sql.DB){
		listTag:= list
		fmt.Println("\n " + listTag)
		fmt.Println("\n `````````````````````````\n")
		
                var Lid int
		rows, err := db.Query("SELECT uid  FROM lists where name =?", listTag)
                checkErr(err)

                for rows.Next(){
                        err = rows.Scan(&Lid)
                        //fmt.Println(Lid)
                        checkErr(err)
                }
		
		notesRows, err:= db.Query("SELECT * FROM notes where listid =?", Lid)
		checkErr(err)

        	for notesRows.Next(){
                	var uid int
                	var name string
                	var message string
                	var created string
			var listid int
                	err = notesRows.Scan(&uid, &name, &message, &created, &listid)
                	checkErr(err)
			var messagePrompt string
			if len(message) < 10 {
				messagePrompt = message
			}else{
				messagePrompt = message[:10]
			}
                	fmt.Println("\n " + name + ": " + messagePrompt + "... ")
        	}
		fmt.Println("\n")
}

func checkErr(err error){
        if err!= nil{
                panic(err)
        }
}

func helpText(){
                fmt.Println("\n GoNotes - Random notes and stuff \n")
                fmt.Println("``````````````````````````````````` \n")
                fmt.Println(" Usage: \n")
                fmt.Println(" gonotes <list tag> (optional) <name of the note> <note/message> \n")
                fmt.Println(" gonotes <name of the note> to display the note \n")
                fmt.Println(" gonotes -list <tag> to list all notes under tag \n")
                fmt.Println(" ```````````````````````````````````````````````` \n")



}
