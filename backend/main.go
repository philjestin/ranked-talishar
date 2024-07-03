package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/philjestin/ranked-talishar/controllers"
	dbCon "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/routes"
	"github.com/philjestin/ranked-talishar/util"
)

var (
    server *gin.Engine
    db     *dbCon.Queries
    ctx    context.Context

    ContactController controllers.ContactController
    ContactRoutes     routes.ContactRoutes
    UserController    controllers.UserController
    UserRoutes	      routes.UserRoutes
    GameController    controllers.GameController
    GameRoutes        routes.GameRoutes
    FormatController  controllers.FormatController
    FormatRoutes      routes.FormatRoutes
)

func init() {
    ctx = context.TODO()
    config, err := util.LoadConfig(".")

    if err != nil {
        log.Fatalf("could not loadconfig: %v", err)
    }

    conn, err := sql.Open(config.DbDriver, config.DbSource)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }

    db = dbCon.New(conn)

    fmt.Println("PostgreSql connected successfully...")

    // Contact Controller and Routes
    ContactController = *controllers.NewContactController(db, ctx)
    ContactRoutes = routes.NewRouteContact(ContactController)

    // Users Controller and Routes
    UserController = *controllers.NewUserController(db, ctx)
    UserRoutes = routes.NewRouteUser(UserController)

    // Games Controller and Routes
    GameController = *controllers.NewGameController(db, ctx)
    GameRoutes = routes.NewRouteGame(GameController)

    // Formats Controller and Routes
    FormatController = *controllers.NewFormatController(db, ctx)
    FormatRoutes = routes.NewRouteFormat(FormatController)

    server = gin.Default()
}

func main() {
    config, err := util.LoadConfig(".")

    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    router := server.Group("/api")

    router.GET("/healthcheck", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{"message": "The contact APi is working fine"})
    })

    ContactRoutes.ContactRoute(router)
    UserRoutes.UserRoute(router)
    GameRoutes.GameRoute(router)
    FormatRoutes.FormatRoute(router)

    server.NoRoute(func(ctx *gin.Context) {
        ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
    })

    log.Fatal(server.Run(":" + config.ServerAddress))
}