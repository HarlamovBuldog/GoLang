package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
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

	var data struct {
		Attribute string `http:"attribute"`
		Order     string `http:"order"`
		Limit     int    `http:"limit"`
		Offset    int    `http:"offset"`
	}
	data.Order = "ASC"
	/*
		var restrictions = map[string][]string{
			"attribute": []string{"name", "color", "tail_length", "whiskers_length"},
		}
	*/
	if err := Unpack(r, &data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Search: %+v\n", data)
	/*
			selectWhat := r.URL.Query().Get("attribute")
			if selectWhat == "" {
				selectWhat = "*"
			} else {
				//permittedAttrs := []string{"name", "color", "tail_length", "whiskers_length"}
				selectWhat, ok := restrictions["attribute"][selectWhat]
			}
				i := 1
			for counter, attr := range data.Attributes {
			queryString = queryString + attr
			if counter-i > 1 {
				queryString = queryString + ", "
			}
		}
		if len(data.Attributes) == 0 {
			queryString = queryString + "*"
		}
	*/
	queryString := "SELECT * FROM cats"

	if data.Attribute != "" {
		queryString = queryString + " ORDER BY " + data.Attribute
		if data.Order != "" {
			queryString = queryString + " " + data.Order
		}
	}

	if data.Offset != 0 {
		queryString = queryString + " OFFSET " +
			strconv.Itoa(data.Offset)
	}
	if data.Limit != 0 {
		queryString = queryString + " LIMIT " +
			strconv.Itoa(data.Limit)
	}

	b, err := queryToJSON(db, queryString)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(b)
	mu.Unlock()
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	default:
		return fmt.Errorf("bad view %s", v.Type())
	}
	return nil
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
