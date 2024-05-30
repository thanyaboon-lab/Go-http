package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// "fmt"
// "log"
// "net/http"

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	fmt.Fprintf(w, "Hello World!")
// }

// func main() {
// 	http.HandleFunc("/hello", helloHandler)

// 	fmt.Println("Starting server at port 8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{Id: 1, Title: "Boon", Author: "Bambino"})
	books = append(books, Book{Id: 2, Title: "Boon2", Author: "Bambino2"})

	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)

	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.Id == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)

	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.Id == bookId {
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			books[i] = book
			return c.JSON(book)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
