package config

import (
	"os"
	"reflect"
	"strconv"
)

func LoadConfig(config interface{}) error {
	fields := getStructFields(config)
	for _, field := range fields {
		envVarName := getEnvVarName(field)
		envValue := getEnvValue(envVarName)
		if envValue == "" {
			envValue = getEnvVarDefault(field)
		}
		err := setFieldValue(config, field, envValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func getStructFields(config interface{}) []reflect.StructField {
	fields := make([]reflect.StructField, 0)
	t := reflect.TypeOf(config).Elem()
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i))
	}
	return fields
}

func getEnvVarName(field reflect.StructField) string {
	return field.Tag.Get("env")
}

func getEnvVarDefault(field reflect.StructField) string {
	return field.Tag.Get("envDefault")
}

func getEnvValue(envVarName string) string {
	return os.Getenv(envVarName)
}

func setFieldValue(config interface{}, field reflect.StructField, value string) error {
	v := reflect.ValueOf(config).Elem().FieldByName(field.Name)
	if !v.IsValid() {
		return nil
	}
	if !v.CanSet() {
		return nil
	}
	setValue(v, value)
	return nil
}

func setValue(val reflect.Value, value string) {
	switch val.Kind() {
	case reflect.String:
		val.SetString(value)
	case reflect.Int:
		intVal, _ := strconv.Atoi(value)
		val.SetInt(int64(intVal))
	default:
		panic("unhandled default case")
	}
}
