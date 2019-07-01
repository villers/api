package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/villers/api/datamodel"
	"github.com/villers/api/datasource"
	"github.com/villers/api/repository"
	"github.com/villers/api/service"
	"strconv"
)

func main() {
	app := newApp()

	app.Run(
		// Start the web server at localhost:8080
		iris.Addr(":8080"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

func newApp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		usersRoutes.Get("/", getUsers)
		usersRoutes.Get("/{id:uint64 min(1)}", getUserByID)
		usersRoutes.Delete("/{id:uint64 min(1)}", deleteUserByID)
		usersRoutes.Post("/create", createUser)
	}

	return app
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

func getUsers(ctx iris.Context) {
	userRepository := repository.NewUserRepository(datasource.Users)
	userService := service.NewUserService(userRepository)

	ctx.JSON(userService.GetAll())
}

func getUserByID(ctx iris.Context) {
	userID, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	userRepository := repository.NewUserRepository(datasource.Users)
	userService := service.NewUserService(userRepository)
	user, found := userService.GetByID(userID)

	if found == false {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	ctx.JSON(user)
}

func deleteUserByID(ctx iris.Context) {
	userID, _ := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	userRepository := repository.NewUserRepository(datasource.Users)
	userService := service.NewUserService(userRepository)
	deleted := userService.DeleteByID(userID)

	if deleted == false {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	ctx.StatusCode(iris.StatusOK)
}

func createUser(ctx iris.Context) {
	var user datamodel.User
	err := ctx.ReadJSON(&user)

	if err != nil {
		ctx.Values().Set("error", "parsing user, read and parse json failed. "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	userRepository := repository.NewUserRepository(datasource.Users)
	userService := service.NewUserService(userRepository)
	updateUser, err := userService.InsertOrUpdate(user)

	if err != nil {
		ctx.Values().Set("error", err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(updateUser)
}
