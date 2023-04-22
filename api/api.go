package api

import (
	"log"
	"net/http"
	"strconv"

	"book/dto"
	"book/service"

	"github.com/gin-gonic/gin"
)

var bookService = service.New()

func StartServer() {
	r := gin.Default()

	r.GET("/book/:id", func(ctx *gin.Context) {
		bookId, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		book, err := bookService.Get(bookId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, "Not Found")
			return
		}
		ctx.JSON(200, book)
	})
	r.GET("/books", func(ctx *gin.Context) {
		books, err := bookService.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, books)
	})
	r.POST("/create/", func(ctx *gin.Context) {
		var book dto.BookDTO
		err := ctx.ShouldBindJSON(&book)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		id, err := bookService.Create(book)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(200, id)
	})
	r.PUT("update/:id", func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		newData := dto.BookDTO{}
		newData.ID = intId

		err = ctx.ShouldBindJSON(&newData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		bookService.Update(newData)
	})
	r.DELETE("delete/:id", func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		err := bookService.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(200, nil)
	})

	r.Run(":9090")
}
