package main

import (
	"database/sql"
	"flag"
	"fmt"
	//"github.com/golang-migrate/migrate"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlserver"
	_ "github.com/golang-migrate/migrate/source/file"
)

var (configPath string)

var (
	name_otdel string
	query      string
)

func init() {
	flag.StringVar(&configPath,"config-path",`C:\Users\v.tereshkov\.GoLand2019.2\config\scratches\apiserver.toml`, "path to file")
}

func main() {

	db, err := sql.Open("sqlserver", "sqlserver://sa:tvv@2012@localhost:1434?database=TT&connection+timeout=30")//"UPP_OPTIMA_Rab")
	if err != nil {
		fmt.Println("Error in connect DB")
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	/*query = "select top 1000 Наименование from dbo.[Справочник.Номенклатура]"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}*/

	// Run migrations
	driver, err := sqlserver.WithInstance(db, &sqlserver.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration... %v", err)
	}

	migrationDir := "e:/TMP/go" //"file:///home/mattes/migrations"

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir), // file://path/to/directory
		"TT", driver)

	/*driver, err := sqlserver.WithInstance(db, &sqlserver.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"TT", driver)
	m.Steps(1)*/

	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	log.Println("Database migrated")
	// actual logic to start your application
	os.Exit(0)

	/*for rows.Next() {
		if err := rows.Scan(&name_otdel); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name_otdel)
	}
	defer rows.Close()*/

	//SendMail()
	
	/*m, err := migrate.New(
		"github://mattes:personal-access-token@mattes/migrate_test",
		"postgres://localhost:5432/database?sslmode=enable")
	m.Steps(2)*/
	
	/*u := &url.URL{
		Scheme:   "sqlserver",
		//Scheme:   "TT",
		User:     url.UserPassword("sa", "tvv@2012"),
		Host:     fmt.Sprintf("%s:%d", "localhost", "ms-sql-m"),
		// Path:  instance, // if connecting to an instance instead of a port
		//RawQuery: query.Encode(),
	}
	
	db, err := sql.Open("sqlserver", u.String())
	/*driver, err := sqlserver.WithInstance(db, &sqlserver.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"sqlserver", driver)
	m.Steps(1)
	errPing  := db.Ping()
	if errPing != nil {
		fmt.Println(errPing)
	}
	
	if err != nil {
		log.Fatal(err)
	}*/

	/*flag.Parse()
	
	
	config := apiserver.NewConfig()
	
	_, err := toml.DecodeFile(configPath,config)
	if err != nil {
		log.Fatal(err)
	}
	
	s := apiserver.New(config)
	if err :=s.Start(); err != nil{
		log.Fatal(err)
	}
	//r := toml.DecodeFile("./config.toml", &conf)*/
	
	
	
}
