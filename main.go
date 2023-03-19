package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/csv"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	hostdb := os.Getenv("HOST_DB")
	dbuser := os.Getenv("DB_USERNAME")
	dbpass := os.Getenv("DB_PASSWORD")
	portdb := os.Getenv("PORT_DB")
	dbname := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, hostdb, portdb, dbname))
	if err != nil {
		fmt.Println("Tidak berhasil koneksi ke Database")
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	rows, err := db.Query("SELECT domain FROM domain")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	file, err := os.Create("cekssl-report.csv")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Domain", "Expired", "Sisa (Hari)"}
	err = writer.Write(header)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	for rows.Next() {
		var host string
		err := rows.Scan(&host)
		if err != nil {
			panic(err)
		}

		wg.Add(1)
		go func(host string) {
			defer wg.Done()

			conn, err := tls.Dial("tcp", net.JoinHostPort(host, "443"), nil)
			if err != nil {
				return
			}
			defer func(conn *tls.Conn) {
				err := conn.Close()
				if err != nil {
					return
				}
			}(conn)

			cert := conn.ConnectionState().PeerCertificates[0]
			expiredate := cert.NotAfter.Format("02/01/2006")
			durasi := time.Until(cert.NotAfter)
			sisahari := int(durasi.Round(24*time.Hour).Hours() / 24)

			record := []string{host, expiredate, fmt.Sprint(sisahari)}
			err = writer.Write(record)
			if err != nil {
				return
			}

			fmt.Println(host, "Expired:", expiredate, "sisa", sisahari, "hari")
		}(host)
	}
	wg.Wait()
	fmt.Println("Selesai")
}

// TODO Create make file and service to linux
