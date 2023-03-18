package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/csv"
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
	defer db.Close()

	rows, err := db.Query("SELECT domain FROM domain")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	file, err := os.Create("cekssl-report.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Domain", "Expired", "Sisa (Hari)"}
	writer.Write(header)

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
		durasi := time.Until(cert.NotAfter)
		sisahari := int(durasi.Round(24*time.Hour).Hours() / 24)

		record := []string{host, expiredate, fmt.Sprint(sisahari)}
		writer.Write(record)

		fmt.Println(host, "Expired:", expiredate, "sisa", sisahari, "hari")
	}
}
