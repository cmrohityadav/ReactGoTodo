package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error loading .env file")
	}
	PORT:=os.Getenv("PORT")
	todos:=[]Todo{}

	fmt.Println("Welcome to React+ Go")
	app := fiber.New()
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello world",
	"data":todos})
	})



	


	// to create todo
	app.Post("/api/todos",func(c *fiber.Ctx) error{
		todo:=&Todo{}
		if err:=c.BodyParser(todo); err!=nil{
			return err
		}

		if todo.Body==""{
			return c.Status(400).JSON(fiber.Map{"error":"Todo body is required"})

		}

		todo.ID=len(todos)+1;
		todos=append(todos, *todo)
		return c.Status(201).JSON(*todo)
	})

	// update todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")

        
        todoID, err := strconv.Atoi(id)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
        }

        
        for i ,_:= range todos {
            if todos[i].ID == todoID {
                todos[i].Completed = true 
                return c.Status(200).JSON(todos[i])
            }
        }

        return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
    })

	//delete todo
	app.Delete("/api/todos/:id",func(c *fiber.Ctx) error {
		id:=c.Params("id");
		iTodoId,err:=strconv.Atoi(id);
		if err!=nil{
			log.Fatal("Invalid ID")
		}

		for i:=range todos{
			if iTodoId==todos[i].ID{
				todos=append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success":"true"})
			}

		}

		return c.Status(400).JSON(fiber.Map{"message":"id does not found to dlete"})

	})

	log.Fatal(app.Listen(":"+PORT))

}
