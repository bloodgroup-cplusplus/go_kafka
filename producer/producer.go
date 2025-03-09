package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {

	app := fiber.New()
	api := app.Group("/api/v1")

	api.Post("/comments", createComment())
	app.Listen(":3000")

}

func ConnectProducer (brokers)

func PushCommentToQueue(topic string,message [] byte) error {
	brokersUrl := [] string {"localhost:29092"}
	producer, err := ConnectProducer(brokersUrl)
}

func createComment(c *fiber.Ctx) error {

	cmt := new (Comment)
	if err := c.BodyParser(cmt); err !=nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success":false,
			"message":err,
		})
		return err
	}
	cmtInBytes,err :=json.Marshal(cmt)
	PushCommentToQueue("comments",cmtInBytes)

	c.JSON(&fiber.Map{
		"success":true,
		"message":"Comment pushed successfully",
		"comment":cmt,
	})
	if err !=nil {
		c.Status(500).JSON(&fiber.Map{
			"success":false,
			"message":"Error creating product",
		})
		return err
	}
	return err

}