package api

import (
	"fmt"
	"net/http"
	"strconv"

	dto "book/DTO"
	"book/config"
	"book/service"

	"github.com/gin-gonic/gin"
)

var (
	bookService = service.New()
)

func StartServer() {
	r := gin.Default()

	r.GET("/book/:id", func(ctx *gin.Context) {
		bookId, ok := ctx.Params.Get("id")
		if !ok {
			config.Logger.Error(fmt.Sprintf("Couldn't get id parameter from context path:%v method:%v", ctx.Request.URL.Path, ctx.Request.Method))
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		book, err := bookService.Get(bookId)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Coudn't get book by it's id,err: %v", err))
			ctx.JSON(http.StatusNotFound, "Not Found")
			return
		}
		ctx.JSON(200, book)
	})
	r.GET("/books", func(ctx *gin.Context) {
		books, err := bookService.GetAll()
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't get all books,err: %v", err))
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, books)
	})
	r.POST("/create/", func(ctx *gin.Context) {
		var book dto.BookDTO
		err := ctx.ShouldBindJSON(&book)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't parse book data from request,err: %v", err))
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		id, err := bookService.Create(book)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't create new book instace, err: %v", err))
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(200, id)
	})
	r.PUT("/update/:id", func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			config.Logger.Error(fmt.Sprintf("Couldn't get id parameter from context path:%v method:%v", ctx.Request.URL.Path, ctx.Request.Method))
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't convert id from string to int,err: %v", err))
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		newData := dto.BookDTO{}
		newData.ID = intId

		err = ctx.ShouldBindJSON(&newData)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't parse updated book data from request,err: %v", err))
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		bookService.Update(newData)
	})
	r.DELETE("/delete/:id", func(ctx *gin.Context) {
		id, ok := ctx.Params.Get("id")
		if !ok {
			config.Logger.Error(fmt.Sprintf("Couldn't get id parameter from context path:%v method:%v", ctx.Request.URL.Path, ctx.Request.Method))
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		err := bookService.Delete(id)
		if err != nil {
			config.Logger.Error(fmt.Sprintf("Couldn't delete book from db,err:%v", err))
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(200, nil)
	})

	r.Run(":9090")
}
