package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/koinworks/asgard-bivrost/libs"
	bv "github.com/koinworks/asgard-bivrost/service"
	hmodels "github.com/koinworks/asgard-heimdal/models"

	OrderDelivery "github.com/evrintobing/bivrost_example_task2/modules/orders/delivery"
	OrderRepository "github.com/evrintobing/bivrost_example_task2/modules/orders/repository"
	OrderUsecase "github.com/evrintobing/bivrost_example_task2/modules/orders/usecase"

	"github.com/joho/godotenv"
)

func main() {

	hostname, _ := os.Hostname()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	portNumber, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	serviceConfig := &hmodels.Service{
		Class:     "product-service",
		Key:       os.Getenv("APP_KEY"),
		Name:      os.Getenv("APP_NAME"),
		Version:   os.Getenv("APP_VERSION"),
		Host:      hostname,
		Port:      portNumber,
		Namespace: os.Getenv("K8S_NAMESPACE"),
		Metas:     make(hmodels.ServiceMetas),
	}

	registry, err := libs.InitRegistry(libs.RegistryConfig{
		Address:  os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Service:  serviceConfig,
	})

	if err != nil {
		log.Fatal(err)
	}

	server, err := libs.NewServer(registry)
	if err != nil {
		log.Fatal(err)
	}
	bivrostSvc := server.AsGatewayService(
		"/v1",
	)

	DB := dbConn()

	orderRepo := OrderRepository.NewOrderRepository(DB)
	orderUC := OrderUsecase.NewOrderUsecase(orderRepo)
	orderHttp := OrderDelivery.NewOrderHandler(orderUC)

	bivrostSvc.Get("/list", orderHttp.GetList)
	bivrostSvc.Get("/orders", orderHttp.GetOrder)
	bivrostSvc.Post("/createorder", orderHttp.AddOrder)

	err = server.Start()
	if err != nil {
		panic(err)
	}

}

func Initialize(Dbdriver, DbUser, DbPassword, DbHost, DbName string, DbPort int) (DB *gorm.DB) {

	var err error

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
		fmt.Println(DBURL)
		DB, err = gorm.Open(Dbdriver, DBURL)
		// DB, err = gorm.Open(Dbdriver, "postgres://postrgres:root@redis/example?sslmode=disable")
		if err != nil {
			fmt.Println("Cannot connect to database")
			log.Fatal("Error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", Dbdriver)
		}
	}

	return DB

}

func dbConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db_driver := os.Getenv("DB_DRIVER")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	db_port, _ := strconv.Atoi(port)
	DB := Initialize(db_driver, db_user, db_password, db_host, db_name, db_port)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaulTableName string) string {
		return "" + defaulTableName
	}
	return DB
}

func exampleMiddleware(next bv.HandlerFunc) bv.HandlerFunc {
	return func(ctx *bv.Context) bv.Result {
		log.Println("This is some middleware")
		ctx.SetHeader("X-Middleware", "Message From Middleware")
		return next(ctx)
	}
}

func pingHandler(ctx *bv.Context) bv.Result {

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome to Ping API",
			"id": "Selamat datang di Ping API",
		},
	})

}

// func addOrderHandler(ctx *bv.Context) bv.Result {
// 	var order AddOrder
// 	err := ctx.BodyJSONBind(&order)
// 	if err != nil {
// 		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
// 			Message: map[string]string{
// 				"en": "error when bind data",
// 				"id": "error ketika membungkus data",
// 			},
// 			Data: err,
// 		})
// 	}

// 	db := dbConn()
// 	data := db.Table("orders").Create(&order)
// 	if data.Error != nil {
// 		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
// 			Message: map[string]string{
// 				"en": "API cant find item list on database",
// 				"id": "API tidak dapat menemukan list item di database",
// 			},
// 			Data: data.Error,
// 		})
// 	}
// 	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
// 		Message: map[string]string{
// 			"en": "API succes find item list on database",
// 			"id": "API berhasil menemukan list item di database",
// 		},
// 		Data: order,
// 	})
// }

// func ordersHandler(ctx *bv.Context) bv.Result {
// 	var orders []GetOrder
// 	db := dbConn()
// 	data := db.Table("orders").Find(&orders)
// 	if data.Error != nil {
// 		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
// 			Message: map[string]string{
// 				"en": "API cant find item list on database",
// 				"id": "API tidak dapat menemukan list item di database",
// 			},
// 			Data: data.Error,
// 		})
// 	}
// 	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
// 		Message: map[string]string{
// 			"en": "API succes find item list on database",
// 			"id": "API berhasil menemukan list item di database",
// 		},
// 		Data: orders,
// 	})
// }
