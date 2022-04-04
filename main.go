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

	"github.com/joho/godotenv"
)

type Items struct {
	ID              int    `json:"id"`
	NamaProduk      string `json:"nama_produk"`
	DeskripsiProduk string `json:"deskripsi_produk"`
	Harga           int    `json:"harga"`
}

type AddOrder struct {
	IDProduk     int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}

type GetOrder struct {
	IDOrder      int `json:"id_order"`
	IDProduk     int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}

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
		"/example",
	)

	bivrostSvc.Get("/", bivrostSvc.WithMiddleware(welcomeHandler, exampleMiddleware))
	// bivrostSvc.Get("/ping-error", pingHandlerWithError)
	bivrostSvc.Get("/list", bivrostSvc.WithMiddleware(itemsHandler, exampleMiddleware))
	bivrostSvc.Post("/createorder", bivrostSvc.WithMiddleware(addOrderHandler, exampleMiddleware))
	bivrostSvc.Get("/orders", bivrostSvc.WithMiddleware(ordersHandler, exampleMiddleware))

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

func welcomeHandler(ctx *bv.Context) bv.Result {

	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome to example API",
			"id": "Selamat datang di example API",
		},
	})

}

// func pingHandlerWithError(ctx *bv.Context) bv.Result {
// 	err := raiseError(1)
// 	if err != nil {
// 		ctx.CaptureSErrors(serror.NewFromErrorc(err, "[asgard-service-example][bivrost] error raised on handler"))
// 		return ctx.JSONResponse(http.StatusServiceUnavailable, bv.ResponseBody{
// 			Message: map[string]string{
// 				"en": "Ping API raised an error",
// 				"id": "Ping API mengalami kegagalan",
// 			},
// 		})
// 	}

// 	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
// 		Message: map[string]string{
// 			"en": "Ping API successfully invoked",
// 			"id": "Ping API berhasil dipanggil",
// 		},
// 	})
// }

// func raiseError(errorCode int) error {
// 	return fmt.Errorf("error number: %d", errorCode)
// }

func itemsHandler(ctx *bv.Context) bv.Result {
	var itemList []Items
	db := dbConn()
	data := db.Find(&itemList)
	if data.Error != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"en": "API cant find item list on database",
				"id": "API tidak dapat menemukan list item di database",
			},
			Data: data.Error,
		})
	}
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "API succes find item list on database",
			"id": "API berhasil menemukan list item di database",
		},
		Data: itemList,
	})
}

func addOrderHandler(ctx *bv.Context) bv.Result {
	var order AddOrder
	err := ctx.BodyJSONBind(&order)
	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"en": "error when bind data",
				"id": "error ketika membungkus data",
			},
			Data: err,
		})
	}

	db := dbConn()
	data := db.Table("orders").Create(&order)
	if data.Error != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"en": "API cant find item list on database",
				"id": "API tidak dapat menemukan list item di database",
			},
			Data: data.Error,
		})
	}
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "API succes find item list on database",
			"id": "API berhasil menemukan list item di database",
		},
		Data: order,
	})
}

func ordersHandler(ctx *bv.Context) bv.Result {
	var orders []GetOrder
	db := dbConn()
	data := db.Table("orders").Find(&orders)
	if data.Error != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"en": "API cant find item list on database",
				"id": "API tidak dapat menemukan list item di database",
			},
			Data: data.Error,
		})
	}
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "API succes find item list on database",
			"id": "API berhasil menemukan list item di database",
		},
		Data: orders,
	})
}
