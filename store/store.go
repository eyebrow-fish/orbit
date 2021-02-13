package store

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"reflect"
)

type Db struct {
	*sql.DB
}

func DbCtx() (context.Context, error) {
	root := context.Background()
	db, err := sql.Open("postgres", "user=user dbname=orbit sslmode=verify-full")
	if err != nil {
		return nil, err
	}
	return context.WithValue(root, "db", &Db{db}), nil
}

func (s *Db) ExecUnique(sql string, args ...interface{}) error {
	res, err := s.DB.Exec(sql, args...)
	if err != nil {
		return err
	} else if rows, err := res.RowsAffected(); err != nil {
		return err
	} else if rows < 1 {
		return fmt.Errorf("could not create new row")
	}
	return nil
}

func (s *Db) QueryUnique(schema interface{}, sql string, args ...interface{}) (interface{}, error) {
	rows, err := s.DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer cleanup(rows)
	if !rows.Next() {
		return nil, fmt.Errorf("could not find any rows")
	}
	value := reflect.ValueOf(schema)
	var fields []interface{}
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()
		fields = appendConverted(fields, field)
	}
	err = rows.Scan(fields...)
	if err != nil {
		return nil, err
	}
	resp := reflect.New(reflect.TypeOf(schema)).Elem()
	for i := range fields {
		resp.Field(i).Set(reflect.ValueOf(fields[i]).Elem())
	}
	return resp.Interface(), nil
}

func (s *Db) QueryMany(schema interface{}, sql string, args ...interface{}) ([]interface{}, error) {
	rows, err := s.DB.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer cleanup(rows)
	value := reflect.ValueOf(schema)
	var fields []interface{}
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()
		fields = appendConverted(fields, field)
	}
	var items []interface{}
	for rows.Next() {
		err = rows.Scan(fields...)
		if err != nil {
			return nil, err
		}
		resp := reflect.New(reflect.TypeOf(schema)).Elem()
		for i := range fields {
			resp.Field(i).Set(reflect.ValueOf(fields[i]).Elem())
		}
		items = append(items, resp.Interface())
	}
	return items, nil
}

func cleanup(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		panic(err)
	}
}

// This is certainly not my favorite piece of code.
// Let's hide it here. :)
func appendConverted(fields []interface{}, field interface{}) []interface{} {
	switch field.(type) {
	case int:
		x := field.(int)
		fields = append(fields, &x)
	case int8:
		x := field.(int8)
		fields = append(fields, &x)
	case int16:
		x := field.(int16)
		fields = append(fields, &x)
	case int32:
		x := field.(int32)
		fields = append(fields, &x)
	case int64:
		x := field.(int64)
		fields = append(fields, &x)
	case uint:
		x := field.(uint)
		fields = append(fields, &x)
	case uint8:
		x := field.(uint8)
		fields = append(fields, &x)
	case uint16:
		x := field.(uint16)
		fields = append(fields, &x)
	case uint32:
		x := field.(uint32)
		fields = append(fields, &x)
	case uint64:
		x := field.(uint64)
		fields = append(fields, &x)
	case float32:
		x := field.(float32)
		fields = append(fields, &x)
	case float64:
		x := field.(float64)
		fields = append(fields, &x)
	case complex64:
		x := field.(complex64)
		fields = append(fields, &x)
	case complex128:
		x := field.(complex128)
		fields = append(fields, &x)
	case bool:
		x := field.(bool)
		fields = append(fields, &x)
	case uintptr:
		x := field.(uintptr)
		fields = append(fields, &x)
	case string:
		x := field.(string)
		fields = append(fields, &x)
	}
	return fields
}
