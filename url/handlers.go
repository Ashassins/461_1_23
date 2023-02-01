// // url/handlers.go
package url

// import (
//     "strconv"

//     "github.com/gofiber/fiber/v2"
//     "github.com/jinzhu/gorm"
// )

// type URLHandler struct {
//     repository *URLRepository
// }
// // func (handler *URLHandler) GetAll(c *fiber.Ctx) error {
// //     var urls []URL = handler.repository.FindAll()
// //     return c.JSON(urls)
// // }

// // func (handler *URLHandler) Get(c *fiber.Ctx) error {
// //     id, err := strconv.Atoi(c.Params("id"))
// //     url, err := handler.repository.Find(id)

// //     if err != nil {
// //         return c.Status(404).JSON(fiber.Map{
// //             "status": 404,
// //             "error":  err,
// //         })
// //     }

// //     return c.JSON(url)
// // }

// // func (handler *URLHandler) Create(c *fiber.Ctx) error {
// //     data := new(URL)

// //     if err := c.BodyParser(data); err != nil {
// //         return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
// //     }

// //     item, err := handler.repository.Create(*data)

// //     if err != nil {
// //         return c.Status(400).JSON(fiber.Map{
// //             "status":  400,
// //             "message": "Failed creating item",
// //             "error":   err,
// //         })
// //     }

// //     return c.JSON(item)
// // }

// // func (handler *URLHandler) Update(c *fiber.Ctx) error {
// //     id, err := strconv.Atoi(c.Params("id"))

// //     if err != nil {
// //         return c.Status(400).JSON(fiber.Map{
// //             "status":  400,
// //             "message": "Item not found",
// //             "error":   err,
// //         })
// //     }

// //     url, err := handler.repository.Find(id)

// //     if err != nil {
// //         return c.Status(404).JSON(fiber.Map{
// //             "message": "Item not found",
// //         })
// //     }

// //     urlData := new(URL)

// //     if err := c.BodyParser(urlData); err != nil {
// //         return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
// //     }

// //     url.StoredURL = urlData.StoredURL
// //     url.License = urlData.License
// //     url.Status = urlData.Status

// //     item, err := handler.repository.Save(url)

// //     if err != nil {
// //         return c.Status(400).JSON(fiber.Map{
// //             "message": "Error updating url",
// //             "error":   err,
// //         })
// //     }

// //     return c.JSON(item)
// // }

// // func (handler *URLHandler) Delete(c *fiber.Ctx) error {
// //     id, err := strconv.Atoi(c.Params("id"))
// //     if err != nil {
// //         return c.Status(400).JSON(fiber.Map{
// //             "status":  400,
// //             "message": "Failed deleting url",
// //             "err":     err,
// //         })
// //     }
// //     RowsAffected := handler.repository.Delete(id)
// //     statusCode := 204
// //     if RowsAffected == 0 {
// //         statusCode = 400
// //     }
// //     return c.Status(statusCode).JSON(nil)
// // }

// // func NewURLHandler(repository *URLRepository) *URLHandler {
// //     return &URLHandler{
// //         repository: repository,
// //     }
// // }

// // func Register(router fiber.Router, database *gorm.DB) {
// //     database.AutoMigrate(&URL{})
// //     urlRepository := NewURLRepository(database)
// //     urlHandler := NewURLHandler(urlRepository)

// //     movieRouter := router.Group("/url")
// //     movieRouter.Get("/", urlHandler.GetAll)
// //     movieRouter.Get("/:id", urlHandler.Get)
// //     movieRouter.Put("/:id", urlHandler.Update)
// //     movieRouter.Post("/", urlHandler.Create)
// //     movieRouter.Delete("/:id", urlHandler.Delete)
// // }