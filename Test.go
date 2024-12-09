package main

import(
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysgl"
	"gorm.io/gorm"
)

	var dbClient *gorm.DB

	func main(){
		
	//intialize fiber

		app :=fiber.New()

		//Load.env file

		err :=godotenv Load(".env")

		if err !=nil{
				
			
log.Println("warning: .env file not found.Defaulting to environment variadles")

	}

	//connect to detadase
	
	dsn :=os.Getenv("DB-DSN")
	//dsn :="user:pasword@tcp(127.0.0.1:3306)/datadase?charset=utf8mb4&parseTime=True&loc=local"
	dbClient, err =gorm.Open(mysgl.Open(dsn), &gorm.config{})
	if err != nil{
		log.fatal("Error connecting to the database:", err)
	}
	log.Println("Database connected")

	//Get the underlying SQL database connection and defer its closure
	sqlDB, err :=dbClient.DB()
	if err != nil {
		log.Fatal("Error getting database connection: ", err)
	}
	defer sqlDB.Close()

	//set up a simple GET route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&Fiber.Map{
			"massage": "In the neam of god"
		})
	})
	//Start the fiber app
	port :=Getenv("PORT")
	if port =="" {
		port ="3000" //Default to port 3000 if PORT is not set  
		
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}