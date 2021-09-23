package sqlindo

import (
	"fmt"
	"testing"
)

// func TestGetConfigs(t *testing.T) {

// 	var db DB
// 	db.Driver = "postgres"
// 	//db.Host = "cfg.rede.tri.com.br"
// 	db.Host = "127.0.0.1"
// 	db.SSLMode = "disable"
// 	db.Port = "5432"
// 	db.Name = "pet3"
// 	db.User = "postgres"
// 	db.Password = ""

// 	db.Connect()

// 	opa := db.GetConfigs("cadin_configuracoes_gerais_pet")
// 	fmt.Println(opa)
// 	db.Conn.Close()
// 	//fmt.Println(opa["parks_config_seis"].(bool))
// }

func TestSelectRow(t *testing.T) {

	var db DB
	db.Driver = "postgres"
	db.Host = "cfg.rede.tri.com.br"
	//db.Host = "127.0.0.1"
	db.SSLMode = "disable"
	db.Port = "5432"
	db.Name = "pet3"
	db.User = "postgres"
	db.Password = ""

	db.Connect()

	opa := db.SelectRow("select * from cadin_configuracoes_gerais_pet")
	fmt.Println(opa)
	fmt.Println(opa["id"].(int64) + 1)
	db.Conn.Close()

}

func TestSelect(t *testing.T) {

	var db DB
	db.Driver = "postgres"
	//db.Host = "cfg.rede.tri.com.br"
	db.Host = "127.0.0.1"
	db.SSLMode = "disable"
	db.Port = "5432"
	db.Name = "pet3"
	db.User = "postgres"
	db.Password = ""

	db.Connect(1)

	opa := db.Select("select * from operations")
	fmt.Println(opa)
	//fmt.Println(opa[0]["id"].(int64) + 1)
	db.Conn.Close()

}

func TestSelectString(t *testing.T) {

	var db DB
	db.Driver = "postgres"
	//db.Host = "cfg.rede.tri.com.br"
	db.Host = "127.0.0.1"
	db.SSLMode = "disable"
	db.Port = "5432"
	db.Name = "pet3"
	db.User = "postgres"
	db.Password = ""

	db.Connect()

	opa := db.SelectString("select * from operations")
	fmt.Println(opa)
	fmt.Println(opa[0]["param"])
	db.Conn.Close()

}
