package delivery

import (
	"net/http"

	"github.com/evrintobing/bivrost_example_task2/modules/orders"
	bv "github.com/koinworks/asgard-bivrost/service"
)

type handler struct {
	UC orders.OrderUsecase
}

func NewOrderHandler(UC orders.OrderUsecase) orders.OrderDelivery {
	return &handler{
		UC: UC,
	}

}

func (handle *handler) GetList(ctx *bv.Context) bv.Result {
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

func (handle *handler) AddOrder(ctx *bv.Context) bv.Result {
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
	dataOrder, err := handle.UC.AddOrder(order.IDProduk, order.JumlahProduk)
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
		Data: dataOrder,
	})
}

func (handle *handler) GetOrder(ctx *bv.Context) bv.Result {
	itemList, err := handle.UC.GetOrder()
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
