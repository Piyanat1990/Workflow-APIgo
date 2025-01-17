package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Piyanat1990/workflow/internal/auth"
	"github.com/Piyanat1990/workflow/internal/item"
	"github.com/Piyanat1990/workflow/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
  		if err != nil {
    	log.Fatal("Error loading .env file")
  }
  


	// Connect database
	db, err := gorm.Open(
		postgres.Open(
			os.Getenv("URL_DB"),
		),
	)

	if err != nil {
		log.Panic(err)
	}

	// Controller
	controller := item.NewController(db)

	// Router
	r := gin.Default()

	config := cors.DefaultConfig()
	// frontend URL
	config.AllowOrigins = []string{
		"http://localhost:3001",
		"http://127.0.0.1:3001",
		"http://localhost:3000",
		"http://127.0.0.1:3000",
	}


	
	r.Use(cors.New(config))

	r.GET("/version", func(c *gin.Context) {
		version, err := GetLatestDBVersion(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"version": version})
	})



   userController:=user.NewController(db,os.Getenv("JWT_SECRET"))
	r.POST("/login",userController.Login)


    // Grouping route & Register router
	items := r.Group("/items")

	
//  verify token
items.Use(auth.Guard(os.Getenv("JWT_SECRET")))

	{
		items.GET("", controller.FindItems)
	    items.POST("", controller.CreateItem)
		items.PATCH("/:id", controller.UpdateItemStatus)
		items.GET("/:id", controller.FindItem)
		items.PUT("/:id", controller.UpdateItems)
		items.DELETE("/:id", controller.DeleteItem)
		
	}
 

// POST 	/items
// GET 		/items?status=xxxxx
// PATCH	/items/:id

// GET 		/items/:id
// PUT		/items/:id
// DELETE 	/items/:id

///graceful shutdown
srv := &http.Server{
	Addr:    (":" + os.Getenv("DBPORT")),
	Handler: r.Handler(),
}

go func() {
	// service connections
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}()

// Wait for interrupt signal to gracefully shutdown the server with
// a timeout of 5 seconds.
quit := make(chan os.Signal, 1)
// kill (no param) default send syscall.SIGTERM
// kill -2 is syscall.SIGINT
// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
log.Println("Shutdown Server ...")

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := srv.Shutdown(ctx); err != nil {
	log.Fatal("Server Shutdown:", err)
}
// catching ctx.Done(). timeout of 5 seconds.
select {

case <-ctx.Done():
	log.Println("timeout of 5 seconds.")
}
log.Println("Server exiting")




// ดึง DSN จาก .env
dsn := os.Getenv("POSTGRES_DSN")
if dsn == "" {
 log.Fatal("POSTGRES_DSN not found in environment")
 }

// // Start server on port 2024
// if err := r.Run(":" + os.Getenv("DBPORT")); err != nil {
//  log.Panic(err)
//  }


	
	// Start server
	// if err := r.Run(); err != nil {
	// 	log.Panic(err)
	// }

   

	
}

type GooseDBVersion struct {
	ID        int
	VersionID int
	IsApplied bool
	Tstamp    string
}

// TableName overrides the table name used by User to `profiles`
func (GooseDBVersion) TableName() string {
	return "goose_db_version"
}

// GetLatestDBVersion returns the latest applied version from the goose_db_version table.
func GetLatestDBVersion(db *gorm.DB) (int, error) {
	var version GooseDBVersion

	// Query to get the latest version applied
	err := db.Order("version_id desc").Where("is_applied = ?", true).First(&version).Error
	if err != nil {
		return 0, err
	}

	return version.VersionID, nil
}
