package api

import (
	"fmt"
	"net/http"
	"strconv"

	"book/config"
	"book/models"
	"book/service"

	"github.com/gin-gonic/gin"
)

var (
	bookService = service.New()
)

func StartServer() {
	r := gin.Default()
	books := r.Group("/")
	{
		books.GET("/book/:id", Book)
		books.GET("/books", Books)
		books.POST("/create/", Create)
		books.PUT("/update/:id", Update)
		books.DELETE("/delete/:id", Delete)
	}
	r.Run(":9090")
}

func Book(ctx *gin.Context) {
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
}

func Books(ctx *gin.Context) {
	books, err := bookService.GetAll()
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't get all books,err: %v", err))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func Create(ctx *gin.Context) {
	var book models.Book
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
}

func Update(ctx *gin.Context) {
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
	newData := models.Book{}
	newData.ID = intId

	err = ctx.ShouldBindJSON(&newData)
	if err != nil {
		config.Logger.Error(fmt.Sprintf("Couldn't parse updated book data from request,err: %v", err))
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	bookService.Update(newData)
	ctx.JSON(200, nil)
}

func Delete(ctx *gin.Context) {
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
}
