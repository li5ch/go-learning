package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)


var (
	defaultIDFields = []string{"id"}
	fieldNameKey    = "column"
	gormTagName     = "gorm"
	ErrTagFormat    = errors.New("tag format error")
)

// BuildBulkInsertSQL build a sql for bulk insert
// Param:
//    tab:  table name
//    objs: object slice
func BuildBulkInsertSQL(tab string, objs interface{}) (string, []interface{}, error) {
	sql := ""
	objVal := reflect.ValueOf(objs)
	if objVal.Kind() != reflect.Slice {
		errMsg := fmt.Sprintf("Expected slice but got: %v", objVal.Kind())
		return sql, nil, errors.New(errMsg)
	}
	if objVal.Len() == 0 {
		errMsg := fmt.Sprintf("The list must not be empty")
		return sql, nil, errors.New(errMsg)
	}
	// build field name
	firstObj := objVal.Index(0)
	fields, err := BuildTableFields(firstObj.Interface())
	if err != nil {
		errMsg := fmt.Sprintf("build sql error: %v", err)
		return sql, nil, errors.New(errMsg)
	}
	// build field value
	var placeHolders []string
	var values []interface{}
	for i := 0; i < objVal.Len(); i++ {
		placeHolder, val := BuildTableFieldsValue(objVal.Index(i).Interface())
		values = append(values, val...)
		placeHolders = append(placeHolders, placeHolder)
	}
	sql = "INSERT INTO " + tab + " " + fields +
		" VALUES " + strings.Join(placeHolders, ",")
	return sql, values, nil
}

// BuildTableFields build a table fields string for bulk operation, except auto increment id.
// e.g. (username,password)
func BuildTableFields(obj interface{}) (result string, err error) {
	fields, err := ParseTableFields(obj)
	if err != nil {
		return result, err
	}
	fields = removeFields(fields, defaultIDFields)
	result = "(" + strings.Join(fields, ",") + ")"
	return result, nil
}

func removeFields(originFields, skipFields []string) (result []string) {
	for _, field := range originFields {
		if inStringSlice(field, skipFields) {
			continue
		}
		result = append(result, field)
	}
	return
}

func inStringSlice(haystack string, strs []string) bool {
	for _, str := range strs {
		if haystack == str {
			return true
		}
	}
	return false
}

func ParseTableFields(obj interface{}) ([]string, error) {
	var fields []string
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		column, err := getColumnName(field.Tag)
		if err != nil {
			return nil, err
		}
		fields = append(fields, column)
	}
	return fields, nil
}

// BuildTableFieldsValue build a place holder and a value list for bulk operation, except auto increment id.
// e.g. return `(?,?)`,[1,'michael'],nil
func BuildTableFieldsValue(obj interface{}) (string, []interface{}) {
	var placeHolder = "("
	var values []interface{}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 1; i < val.NumField(); i++ {
		placeHolder = placeHolder + "?"
		if i+1 == val.NumField() {
			placeHolder = placeHolder + ")"
		} else {
			placeHolder = placeHolder + ","
		}
		field := val.Field(i)
		values = append(values, field.Interface())
	}
	return placeHolder, values
}

// e.g. ["username","password"]
// 此处默认第一个字段是id, 会跳过
func TableFields(obj interface{}) ([]string, error) {
	var fields []string
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	for i := 1; i < typ.NumField(); i++ {
		field := typ.Field(i)
		column, err := getColumnName(field.Tag)
		if err != nil {
			return nil, err
		}
		fields = append(fields, column)
	}
	return fields, nil
}

func getColumnName(tag reflect.StructTag) (string, error) {
	// column tag format: `gorm:"column:id" json:"id"`
	gormTag := tag.Get(gormTagName)
	if gormTag == "" {
		return "", ErrTagFormat
	}

	gormTags := strings.Split(gormTag, ";")
	if len(gormTags) == 0 {
		return "", ErrTagFormat
	}
	for _, kvPairStr := range gormTags {
		kvPair := strings.Split(kvPairStr, ":")
		if len(kvPair) != 2 {
			continue
		}

		if strings.TrimSpace(kvPair[0]) == fieldNameKey {
			return strings.TrimSpace(kvPair[1]), nil
		}
	}

	return "", ErrTagFormat
}

