package utils

import "github.com/magiconair/properties"

func getPropertiesFile() *properties.Properties {
	p := properties.MustLoadFile("application.properties", properties.UTF8)
	return p
}

func GetStringProperty(key string) string {
	p := getPropertiesFile()
	stringProperty := p.GetString(key, "")
	return stringProperty
}

func GetIntProperty(key string) int {
	p := getPropertiesFile()
	intProperty := p.GetInt(key, 0)
	return intProperty
}
