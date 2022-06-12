package Application

import (
	"fmt"
	"playground/MainUrlShortener/Core/APIServer"
	"playground/MainUrlShortener/Core/Database"
	"playground/MainUrlShortener/Core/Router"
)

func Run() {
	fmt.Println("Application running")
	Database.ConnectDatabase()
	Router.Routes(Database.GetConnection())
	APIServer.Run()
}
