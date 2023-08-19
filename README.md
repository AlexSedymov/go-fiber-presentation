# Go fiber presentation

### GO Setup & First Project

* https://go.dev/dl/ - download and install GO language

* Make a new folder for your project

* In the new folder initialize your go module: 
    
    `go mod init <module-name>`

* Install the Fiber package by running the following command:

    `go get -u github.com/gofiber/fiber/v2`

* Create a new Go file in your project directory: `main.go`
* 

```
package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")
}
```