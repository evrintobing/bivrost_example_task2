package orders

import bv "github.com/koinworks/asgard-bivrost/service"

type OrderDelivery interface {
	AddOrder(ctx *bv.Context) bv.Result
	GetList(ctx *bv.Context) bv.Result
	GetOrder(ctx *bv.Context) bv.Result
}
