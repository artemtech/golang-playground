package main

import (
	"fmt"
	"belajar-viper/config"
	"reflect"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	configtype := reflect.TypeOf(configuration)
	for i := 0; i < configtype.NumField(); i++ {
		fieldName := configtype.Field(i).Name
		fieldValue := reflect.ValueOf(configuration).Field(i)

		// Print the field name and value
		fmt.Printf("%s: %v\n", fieldName, fieldValue.Interface())
	}

	fmt.Printf("nama database: %s \n", configuration.DBName)
}
