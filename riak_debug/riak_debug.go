package riak_debug

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	CREATE_COMMANDS_TABLE = `create table node_commands(
node text not null,
cpuinfo text,
date text,
debian_version text,
df text,
df_i text,
disk_by_id text,
diskstats text,
dmesg text,
dpkg text,
free text,
hostname text,
iostat_linux text,
last text,
limits_conf text,
lsb_release text,
meminfo text,
messages text,
mount text,
netstat_an text,
netstat_i text,
netstat_rn text,
ps text,
riak_aae_status text,
riak_diag text,
riak_member_status text,
riak_ping text,
riak_repl_status text,
riak_ring_status text,
riak_status text,
riak_transfers text,
riak_version text,
rx_crc_errors text,
schedulers text,
uname text,
vmstat text,
w text)`
	CREATE_CONSOLE_LOG_TABLE = `create table console_logs (date string, timestamp string,
level string,
pid string,
code string, message string)`
	CREATE_CRASH_LOG_TABLE  = "crash_logs"
	CREATE_ERLANG_LOG_TABLE = "erlang_logs"
	CREATE_ERROR_LOG_TABLE  = "error_logs"
	CREATE_NODE_TABLE       = "create table nodes(appconfig string, vmargs string)"
	// TODO: ring
)

func MaybeCreateAllTables(db *sql.DB) {
	MaybeCreateTable(CREATE_COMMANDS_TABLE, db)
	MaybeCreateTable(CREATE_CONSOLE_LOG_TABLE, db)
	//MaybeCreateTable(CREATE_CRASH_LOG_TABLE, db)
	//MaybeCreateTable(CREATE_ERLANG_LOG_TABLE, db)
	//MaybeCreateTable(CREATE_ERROR_LOG_TABLE, db)
	MaybeCreateTable(CREATE_NODE_TABLE, db)
}

func MaybeCreateTable(q string, db *sql.DB) {
	_, e := db.Exec(q)
	if e != nil {
		log.Print(e)
	}
}

func Content(path string) string {
	bytes, err := ioutil.ReadFile(path)
	log.Print("loading %s", path)
	if err != nil {
		log.Print(err.Error())
		return ""
	}
	log.Print("=> %v bytes", len(bytes))
	return string(bytes)
}

func ImportCommandsResult(base_path string, db *sql.DB) {
	tx, etx := db.Begin()
	if etx != nil {
		log.Fatal(etx)
	}
	log.Println("importing %v", base_path)

	stmt, estmt := tx.Prepare(`
INSERT INTO node_commands
(node, cpuinfo, date, debian_version, df, df_i, disk_by_id, diskstats,
 dmesg, dpkg, free, hostname, iostat_linux, last, limits_conf, lsb_release,
 meminfo, messages, mount, netstat_an, netstat_i, netstat_rn, ps, riak_aae_status,
 riak_diag, riak_member_status, riak_ping, riak_repl_status, riak_status, riak_transfers, riak_version, rx_crc_errors,
 schedulers, uname, vmstat, w
)
 VALUES
(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
 ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if estmt != nil {
		//tx.Abort()
		log.Fatal(etx)
	}
	stmt.Exec(
		base_path,
		Content(base_path+"/cpuinfo"),
		Content(base_path+"/date"),
		Content(base_path+"/debian_version"),
		Content(base_path+"/df"),
		Content(base_path+"/df_i"),
		Content(base_path+"/disk_by_id"),
		Content(base_path+"/diskstats"),
		Content(base_path+"/dmesg"),
		Content(base_path+"/dpkg"),
		Content(base_path+"/free"),
		Content(base_path+"/hostname"),
		Content(base_path+"/iostat_linux"),
		Content(base_path+"/last"),
		Content(base_path+"/limits.conf"),
		Content(base_path+"/lsb_release"),
		Content(base_path+"/meminfo"),
		Content(base_path+"/messages"),
		Content(base_path+"/mount"),
		Content(base_path+"/netstat_an"),
		Content(base_path+"/netstat_i"),
		Content(base_path+"/netstat_rn"),
		Content(base_path+"/ps"),
		Content(base_path+"/riak_aae_status"),
		Content(base_path+"/riak_diag"),
		Content(base_path+"/riak_member_status"),
		Content(base_path+"/riak_ping"),
		Content(base_path+"/riak_repl_status"),
		Content(base_path+"/riak_status"),
		Content(base_path+"/riak_transfers"),
		Content(base_path+"/riak_version"),
		Content(base_path+"/rx_crc_errors"),
		Content(base_path+"/schedulers"),
		Content(base_path+"/uname"),
		Content(base_path+"/vmstat"),
		Content(base_path+"/w"),
	)
	defer stmt.Close()

	e := tx.Commit()
	if e != nil {
		log.Fatal(e)
		tx.Rollback()
	}
	log.Println("importing %v successfully done.", base_path)
}

func ImportLogsResult(base_path string, db *sql.DB) {
	console_logs, err := filepath.Glob(base_path + "/console.log*")
	if err != nil {
		log.Printf("no console.log files: " + err.Error())
	}
	tx, e := db.Begin()
	for _, console_log := range console_logs {

		log.Printf(console_log)
		tokens, e := Parse(console_log)
		if e != nil {
			log.Println("errrrrrrrrrrrrrRR")
		}
		log.Println("%v", tokens)
	}
}
