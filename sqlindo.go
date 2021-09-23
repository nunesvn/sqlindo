package sqlindo

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DebugLevel int

type DB struct {
	Conn     *sql.DB
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (d *DB) uri() string {
	uri := "host=" + d.Host + " "
	uri = uri + "dbname=" + d.Name + " "
	uri = uri + "port=" + d.Port + " "
	uri = uri + "user=" + d.User + " "
	uri = uri + "password=" + d.User + " "
	uri = uri + "sslmode=" + d.SSLMode + " "
	return uri
}

func (d *DB) Connect(debuglevel int) {

	var err error

	DebugLevel = debuglevel

	d.Conn, err = sql.Open(d.Driver, d.uri())

	if err != nil {
		panic(err)
	}

	err = d.Conn.Ping()
	if err != nil {
		panic(err)
	}

}

func (d *DB) SelectRow(query string) map[string]interface{} {
	if DebugLevel > 5 {
		log.Printf("%s", query)
	}

	result := d.Select(query)
	//if len(result) == 0 {
	//	return make(map[string]interface{})
	//}
	if len(result) == 0 {
		return make(map[string]interface{})
	}
	return result[0]

}

func (d *DB) SelectString(query string) []map[string]string {

	var result []map[string]string

	r := d.Select(query)

	for _, m := range r {
		var rowResult = map[string]string{}

		for name, value := range m {

			switch v := value.(type) {
			case string:
				rowResult[name] = v
			case int64:
				rowResult[name] = strconv.FormatInt(v, 10)
			case int:
				rowResult[name] = strconv.Itoa(v)
			case bool:
				rowResult[name] = strconv.FormatBool(v)
			case time.Time:
				rowResult[name] = v.String()
			default:
				rowResult[name] = "NULL"
			}
		}
		result = append(result, rowResult)
	}
	return result

}

func (d *DB) SelectStringValues(query string) []map[string]string {

	var result []map[string]string

	r := d.Select(query)

	for _, m := range r {
		var rowResult = map[string]string{}

		for name, value := range m {

			switch v := value.(type) {
			case string:
				rowResult[name] = v
			case int64:
				rowResult[name] = strconv.FormatInt(v, 10)
			case int:
				rowResult[name] = strconv.Itoa(v)
			case bool:
				rowResult[name] = strconv.FormatBool(v)
			case time.Time:
				rowResult[name] = v.String()
			default:
				rowResult[name] = "NULL"
			}
		}
		result = append(result, rowResult)
	}
	return result

}

// Select recebe uma query e retorna um slice de maps com o resultado
func (d *DB) Select(query string) []map[string]interface{} {

	var result []map[string]interface{}

	if DebugLevel > 5 {
		log.Printf("%s", query)
	}

	rows, err := d.Conn.Query(query)
	if err != nil {
		log.Printf("%s", query)
		log.Fatal(err)
	}

	defer rows.Close()

	columns, err := rows.Columns()

	for rows.Next() {

		values := make([]interface{}, len(columns))
		for i := range columns {
			values[i] = &values[i]
		}

		if err := rows.Scan(values...); err != nil {
			log.Fatal(err)
		}
		//fmt.Println(err)

		var rowResult = make(map[string]interface{})
		for i := range values {

			name := columns[i]
			switch v := values[i].(type) {
			case string:
				rowResult[name] = v
			case int64:
				rowResult[name] = v
			case int:
				rowResult[name] = v
			case bool:
				rowResult[name] = v
			case time.Time:
				rowResult[name] = v
			default:
				rowResult[name] = "NULL"
			}
		}
		result = append(result, rowResult)
	}

	return result

}
