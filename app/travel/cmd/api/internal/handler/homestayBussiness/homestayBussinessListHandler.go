package homestayBussiness

import (
	"net/http"

	"looklook/app/travel/cmd/api/internal/logic/homestayBussiness"
	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
	"looklook/common/result"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HomestayBussinessListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessListLogic(r.Context(), ctx)
		resp, err := l.HomestayBussinessList(req)
		result.HttpResult(r, w, resp, err)
	}
}
