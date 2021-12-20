package homestayBussiness

import (
	"context"

	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
	"looklook/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
)

type HomestayBussinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBussinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayBussinessListLogic {
	return HomestayBussinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBussinessListLogic) HomestayBussinessList(req types.HomestayBussinessListReq) (*types.HomestayBussinessListResp, error) {

	list, err := l.svcCtx.HomestayBusinessModel.FindPageList(req.LastId, req.PageSize)
	if err != nil {
		return nil, errors.Wrapf(xerr.ErrDBError, "req : %+v , err:%v", req,err)
	}

	var resp []types.HomestayBusinessListInfo
	if len(list) > 0 {
		for _, item := range list {
			var typeHomestayBusinessListInfo types.HomestayBusinessListInfo
			_ = copier.Copy(&typeHomestayBusinessListInfo, item)

			resp = append(resp, typeHomestayBusinessListInfo)
		}
	}

	return &types.HomestayBussinessListResp{
		List: resp,
	}, nil
}
