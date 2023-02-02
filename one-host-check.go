package main

import (
	"crypto/tls"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func onehost() {
	host := "kejati-aceh.kejaksaan.go.id"
	port := 443

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	//fmt.Println("Subject:", cert.Subject)
	//fmt.Println("Issuer: ", cert.Issuer)
	expiredate := cert.NotAfter.Format("02/01/2006")
	duration := cert.NotAfter.Sub(time.Now())
	hari := duration.Round(24*time.Hour).Hours() / 24
	//fmt.Println(host, "Expires:", expiredate, "in", hari, "days")
	fmt.Printf("%s,%s,%.0f\n", host, expiredate, hari)
}
