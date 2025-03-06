package v1

import (
	"context"

	"github.com/Ikaros727/sqlite-lab/internal/svc"
	"github.com/Ikaros727/sqlite-lab/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RawLogic
type RawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RawLogic {
	return &RawLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("Logic", "RawLogic")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RawLogic) Raw(req *types.RawReq) (resp *types.RawResp, err error) {
	resp = &types.RawResp{}
	resp.Status = types.StatusSuccess()

	// 获取数据库操作对象
	db, err := l.svcCtx.UserDatabaseCache.MustLoad(1, req.Database)
	if err != nil {
		l.Logger.Errorf("UserDatabaseCache.MustLoad(%v, %v) failed, err: %v", 1, req.Database, err)
		return
	}

	// 参数转换
	args := make([]interface{}, len(req.Args))
	for i, arg := range req.Args {
		args[i] = arg
	}

	// 执行操作
	db = db.Raw(req.SQL, args...).Scan(&resp.Data)
	if err = db.Error; err != nil {
		l.Logger.Errorf("db.Exec failed, err: %v", err)
		return
	}

	return
}
