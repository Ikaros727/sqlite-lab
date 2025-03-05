package v1

import (
	"context"

	"github.com/Ikaros727/sqlite-lab/internal/svc"
	"github.com/Ikaros727/sqlite-lab/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DDLLogic //
type DDLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDDLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DDLLogic {
	return &DDLLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("Logic", "DDLLogic")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DDLLogic) DDL(req *types.DDLReq) (resp *types.DDLResp, err error) {
	resp = &types.DDLResp{}
	resp.Status = types.StatusSuccess()

	db, err := l.svcCtx.UserDatabaseCache.MustLoad(0, "default")
	if err != nil {
		l.Logger.Errorf("UserDatabaseCache.MustLoad(%v, %v) failed, err: %v", 0, "default", err)
		return
	}

	tx := db.Exec(req.Sql)
	if err = tx.Error; err != nil {
		l.Logger.Errorf("Exec %v failed, err: %v", req.Sql, err)
		return
	}

	return
}
