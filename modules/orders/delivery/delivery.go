package delivery

import (
	"net/http"

	"github.com/evrintobing/bivrost_example_task2/modules/orders"
	bv "github.com/koinworks/asgard-bivrost/service"
)

type handler struct {
	UC orders.OrderUsecase
}

func NewOrderHandler(bv bv.Server, UC orders.OrderUsecase) {
	handlers := handler{
		UC: UC,
	}
	bivrostSvc := bv.AsGatewayService(
		"/orders",
	)

	bivrostSvc.Get("/list", handlers.getList)
	// bivrostSvc.Post("/createorder", bivrostSvc.WithMiddleware(addOrderHandler, exampleMiddleware))
	// bivrostSvc.Get("/orders", bivrostSvc.WithMiddleware(ordersHandler, exampleMiddleware))

}

func (handle *handler) getList(ctx *bv.Context) bv.Result {
	itemList, err := handle.UC.GetItems()
	if err != nil {
		return ctx.JSONResponse(http.StatusBadRequest, bv.ResponseBody{
			Message: map[string]string{
				"en": "API cant find item list on database",
				"id": "API tidak dapat menemukan list item di database",
			},
			Data: err,
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
