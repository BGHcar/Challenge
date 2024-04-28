package main

const (
	migrationPath = "file://migrations"
	migrationScriptsVersion = 1
)

func main() {

	_ = logs.InitLogger()

	client := database.getZincSearchClient()