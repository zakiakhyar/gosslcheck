package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	pass := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("mysql", fmt.Sprintf("cekssl:%s@tcp(127.0.0.1:3308)/cekssl", pass))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT domain FROM domain")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var host string

	for rows.Next() {
		err := rows.Scan(&host)
		if err != nil {
			panic(err)
		}
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", host), nil)
		if err != nil {
			continue
		}
		defer conn.Close()

		cert := conn.ConnectionState().PeerCertificates[0]
		expiredate := cert.NotAfter.Format("02/01/2006")
		durasi := cert.NotAfter.Sub(time.Now())
		sisahari := int(durasi.Round(24*time.Hour).Hours() / 24)
		fmt.Println(host, "Expired:", expiredate, "sisa", sisahari, "hari")
	}
}
