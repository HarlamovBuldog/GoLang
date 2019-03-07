package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"

	_ "github.com/lib/pq"
)

var mu sync.Mutex

type CatColor string

type Cats struct {
	name           string
	color          string `sql:"type:cat_color"`
	tailLength     int
	whiskersLength int
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/cats", cats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	mu.Unlock()
}

func ping(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Cats Service. Version 0.1")
	mu.Unlock()
}

func cats(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	//do not forget to change host and pass
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

	var restrictions = map[string][]string{
		"attribute": []string{"name", "color", "tail_length", "whiskers_length"},
	}

	selectWhat := r.URL.Query().Get("attribute")
	if selectWhat == "" {
		selectWhat = "*"
	} else {
		//permittedAttrs := []string{"name", "color", "tail_length", "whiskers_length"}
		selectWhat, ok := restrictions["attribute"][selectWhat]
	}

	queryString := "SELECT " + selectWhat + " FROM cats"

	b, err := queryToJSON(db, queryString)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(b)

	/*

		u, err := r.URL.Parse(r.URL.String())
		if err != nil {
			log.Fatal(err)
		}

		//q := u.Query().Get("")


			keyValuePair := r.Form
			for key, value1 := range keyValuePair {
				if len(value1) < 1 {
					continue
				}
				fmt.Fprintf(w, key+" = "+value1[0]+"\n")
			}

			for key, value1 := range keyValuePair {
				if len(value1) < 1 {
					continue
				}
				switch key {
				case "cycles":
					f64, err := strconv.ParseFloat(value1[0], 64)
					if err == nil {
						cycles = f64
						fmt.Fprintf(w, "Success! Value of "+key+
							" changed to "+value1[0]+"\n")
					}
				case "size":
					i, err := strconv.ParseInt(value1[0], 10, 64)
					if err == nil {
						size = int(i)
						fmt.Fprintf(w, "Success! Value of "+key+
							" changed to "+value1[0]+"\n")
					}
				case "nframes":
					i, err := strconv.ParseInt(value1[0], 10, 64)
					if err == nil {
						nframes = int(i)
						fmt.Fprintf(w, "Success! Value of "+key+
							" changed to "+value1[0]+"\n")
					}
				default:
					fmt.Fprintf(w, "Wrong key! "+key+"\n")
				}
			}
	*/
	mu.Unlock()
}

func queryToJSON(db *sql.DB, query string, args ...interface{}) ([]byte, error) {
	// an array of JSON objects
	// the map key is the field name
	var objects []map[string]interface{}
	//Note:guess instead of strColor we can use custom type with Valuer interface
	var strColor *string
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		// figure out what columns were returned
		// the column names will be the JSON object field keys
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}

		// Scan needs an array of pointers to the values it is setting
		// This creates the object and sets the values correctly
		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			if column.Name() != "color" {
				object[column.Name()] = reflect.New(column.ScanType()).Interface()
			} else {
				object[column.Name()] = reflect.New(reflect.TypeOf(strColor).Elem()).Interface()
			}

			values[i] = object[column.Name()]
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		objects = append(objects, object)
	}

	// indent because I want to read the output
	//return json.MarshalIndent(objects, "", "\t")
	return JSONMarshal(objects)
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}