package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/exec"
	"strings"
)

// subcommands :=

func usage() {
	str := `
$ investigator import <tag> [riak-debug | riak-cs-debug] 
$ investigator diag <tag>
$ investigator query <tag> <query>`
	fmt.Println(str[1:])
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
		return
	}
	switch args[0] {
	case "import":
		do_import(args[1], args[2])
	case "diag":
		do_diag(args[1])
	case "query":
		do_query(args[1], args[2])
	case "help":
		usage()
	default:
		usage()
	}
}

func ensure_dir(dir string) {
	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func do_import(tag, file string) {
	ensure_dir(tag)
	cmd := exec.Command("tar", "xzf", file, "-C", tag)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	tokens := strings.Split(file, "/")
	filename := tokens[len(tokens)-1]
	import_target := tag + "/" + filename[:len(filename)-7]
	fmt.Println("import target: %s\n", import_target)

	//c, _ := sqlite3.Open("sqlite.db")
	c, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	_, e := c.Exec("CREATE TABLE x(a integer, b integer, c integer)")
	if e != nil {
		log.Fatal(e)
	}

	tx, _ := c.Begin()
	stmt, _ := tx.Prepare("INSERT INTO x VALUES(?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(1, 1, 0)
	tx.Commit()

	rows, _ := c.Query("SELECT rowid, * FROM x")
	defer rows.Close()
	for rows.Next() {
		var a, b, c int
		rows.Scan(&a, &b, &c)
		fmt.Println("%v %v %v", a, b, c)
	}
	c.Close()
}
func do_diag(tag string) {
}
func do_query(tag, query string) {}
