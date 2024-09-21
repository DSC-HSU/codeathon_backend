package conf

import (
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	// create test config file
	jsonConfigTxt := `
	{
		"supabase": {
			"postgres": "postgres",
			"api": "api",
			"anonKey": "anon_key",
			"serviceKey": "service_key",
			"s3Access": "s3_access",
			"s3Secret": "s3_secret",
			"s3Region": "s3_region"
		}
	}
	`
	// write test config file
	jsonConfigPath := "./test.json"
	err := os.WriteFile(jsonConfigPath, []byte(jsonConfigTxt), 0644)
	if err != nil {
		t.Error(err)
	}
	// read test config file
	err = ReadConfig(jsonConfigPath)
	if err != nil {
		t.Error(err)
	}
	// test config
	if Config.Supabase.Postgres != "postgres" {
		t.Error("config.Supabase.Postgres != postgres")
	}
	if Config.Supabase.Api != "api" {
		t.Error("config.Supabase.Api != api")
	}
	if Config.Supabase.AnonKey != "anon_key" {
		t.Error("config.Supabase.AnonKey != anon_key")
	}
	if Config.Supabase.ServiceKey != "service_key" {
		t.Error("config.Supabase.ServiceKey != service_key")
	}
	// delete test config file
	err = os.Remove(jsonConfigPath)
	if err != nil {
		t.Error(err)
	}
}
