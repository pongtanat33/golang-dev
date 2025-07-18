package main

import (
	"log"
	"os"
	"strings"

	_ "time/tzdata" // Importing time zone data to ensure correct time zone handling)

	"github.com/spf13/viper"
)

func init() {

	env := os.Args[1]
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func main() {

	cfg := new(configs.Configs)
	// Fiber configs
	cfg.App.Host = viper.GetString("FIBER.HOST")
	cfg.App.Port = viper.GetString("FIBER.PORT")

	// Database Configs
	cfg.MSSQL.Host = viper.GetString("MSSQL.HOST")
	cfg.MSSQL.Port = viper.GetString("MSSQL.PORT")
	cfg.MSSQL.Username = viper.GetString("MSSQL.USER")
	cfg.MSSQL.Password = viper.GetString("MSSQL.PASSWORD")
	cfg.MSSQL.Database = viper.GetString("MSSQL.DATABASE")

	// New Database
	db, err := databases.NewMSSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	if err := redis.InitRedis(); err != nil {
		log.Fatalf("Redis init failed: %v", err)
	} else {
		log.Println("Redis connected successfully")
	}

	s := servers.NewServer(cfg, db)
	middlewares.NewMiddlewares(s.App, db)
	utils.NewUtils(db)
	// Regiter websocket
	ws.RegisterWsRoutes(s.App)
	s.Start()
}
