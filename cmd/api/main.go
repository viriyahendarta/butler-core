package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/viriyahendarta/butler-core/config"
	userdb "github.com/viriyahendarta/butler-core/database/user"
	"github.com/viriyahendarta/butler-core/infra/database"
	businessresource "github.com/viriyahendarta/butler-core/resource/business"
	reporesource "github.com/viriyahendarta/butler-core/resource/repo"
	serviceresource "github.com/viriyahendarta/butler-core/resource/service"
	"github.com/viriyahendarta/butler-core/server"
)

func init() {
	if err := config.Init(); err != nil {
		log.Fatalln("Failed to initialize config", err)
	}
}

func initDatabase(dbConfig config.Database) []*sql.DB {
	urls := []string{dbConfig.MasterURL}
	for _, url := range dbConfig.SlaveURLs {
		urls = append(urls, url)
	}

	dbs, err := database.OpenConnections(dbConfig.Driver, urls...)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to open database connection [%s]:", dbConfig.Name), err)
	}

	for _, db := range dbs {
		db.SetMaxOpenConns(dbConfig.MaxOpenConnections)
		db.SetMaxIdleConns(dbConfig.MaxIdleConnections)
		db.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnectionMaxLifetime))
	}

	log.Printf("Database connection [%s] successfuly established\n", dbConfig.Name)
	return dbs
}

func main() {
	coreDBConfig := config.Get().Databases.CoreDatabase
	coreDB := database.New(coreDBConfig.Driver, initDatabase(coreDBConfig))

	repoResource := initRepoResource(coreDB)
	businessResource := initBusinessResource(repoResource)
	serviceResource := initServiceResource(businessResource)

	port := config.Get().HTTPServer.Port
	portEnv := os.Getenv("PORT")
	if p, err := strconv.Atoi(portEnv); err == nil {
		port = p
	}

	httpServer := server.InitHTTPServer(serviceResource, port)
	err := httpServer.Run(config.GetEnv())
	if err != nil {
		log.Fatalln("Failed to start HTTP server", err)
	}
}

func initRepoResource(coreDB *database.DB) *reporesource.Resource {
	return &reporesource.Resource{
		CoreDB: coreDB,
	}
}

func initBusinessResource(repoResource *reporesource.Resource) *businessresource.Resource {
	return &businessresource.Resource{
		UserDB: userdb.GetDatabase(repoResource.CoreDB),
	}
}

func initServiceResource(businessResource *businessresource.Resource) *serviceresource.Resource {
	return &serviceresource.Resource{
		BusinessResource: businessResource,
		Router:           mux.NewRouter(),
	}
}
