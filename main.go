package main

import (
    "errors"
	"fmt"
	"net/http"
	"os"
    "time"
    
    // "log"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    
    // "github.com/19chonm/461_1_23/database"
    // "github.com/19chonm/461_1_23/url"
)
const serverPort = 4000

func main() {
    app := fiber.New()
    app.Use(cors.New())

    // app.Get("/myEndpoint", func(c *fiber.Ctx) error {
    //     app.Get("GPToken")
    //     app.Get("InputURL")
    //     // c.BaseURL() // return either npm, github, etc .com
    // })

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

    time.Sleep(100 * time.Millisecond)

    requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
}


// func main() {
// 	app := fiber.New()
//     app.Use(cors.New())
//     // database.ConnectDB() // Does not currently use database shit
// 	// defer database.DB.Close()
	
//     api := app.Group("/api")
// 	// url.Register(api, database.DB)

//     log.Fatal(app.Listen((string) serverPort))
// }