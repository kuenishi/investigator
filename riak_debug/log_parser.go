package riak_debug

import (
	//"database/sql"
	"bufio"
	"log"
	"os"
	"strings"
)

func Parse(file string) ([]string, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fp)
	var lines []string
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		log.Printf("length: %d", len(tokens))
		log.Printf("-------")
		if len(tokens) > 4 {
			date := tokens[0] + " " + tokens[1]
			log.Printf(date)
			log.Printf("-------")
			level := tokens[2]
			log.Printf(level)
			log.Printf("-------")
			pid_code := tokens[3]
			log.Printf(pid_code)
			log.Printf("-------")
			msg := strings.Join(tokens[4:], " ")
			log.Printf(msg)
			log.Printf("-------")
		}
	}
	if e := scanner.Err(); e != nil {
		log.Fatal(e)
	}
	return lines, nil
}
