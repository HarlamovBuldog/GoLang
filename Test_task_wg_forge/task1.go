package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type CatColor int64

const (
	black CatColor = iota
	white
	blackWhite
	red
	redWhite
	redBlackWhite
)

type CatsStat struct {
	color CatColor
	count int64
}

func main() {

	connStr := "host=10.10.0.89 port=5432 user=wg_forge password=42a dbname=wg_forge_db sslmode=disable"
	//connStr := "wg_forge:42a@tcp(10.10.0.89)/wg_forge_db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT color, count(*) FROM cats GROUP BY color LIMIT 1000")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		testCatsStat := new(CatsStat)
		if err := rows.Scan(&testCatsStat.color, &testCatsStat.count); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(testCatsStat)
	}
}

/*
package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/ssh"
	"log"
)

type CatColor int64

const (
	black CatColor = iota
	white
	blackWhite
	red
	redWhite
	redBlackWhite
)

type CatsStat struct {
	color CatColor
	count int64
}

func main() {

	sshHost := "10.10.0.89" // SSH Server Hostname/IP
	sshUser := "vlad"       // SSH Username
	sshPass := "1234vlad"   // Empty string for no password
	//dbUser := "wg_forge"    // DB username
	//dbPass := "42a"         // DB Password
	//dbHost := "localhost"   // DB Hostname/IP
	//dbName := "wg_forge_db" // Database name

	//var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", sshHost+":22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

	connStr := "host=10.10.0.89 port=5432 user=wg_forge password=42a dbname=wg_forge_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT color, count(*) FROM cats GROUP BY color LIMIT 1000")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		testCatsStat := new(CatsStat)
		if err := rows.Scan(&testCatsStat.color, &testCatsStat.count); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(testCatsStat)
	}

}
*/
