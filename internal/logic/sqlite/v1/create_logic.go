package v1

import (
	"context"
	"fmt"
	"strings"

	"github.com/Ikaros727/sqlite-lab/internal/svc"
	"github.com/Ikaros727/sqlite-lab/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	insertFormat = "INSERT INTO %s (%s) VALUES (%s);"
)

// CreateLogic // 执行 CREATE 方法
type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx).WithFields(logx.Field("Logic", "CreateLogic")),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateReq) (resp *types.CreateResp, err error) {
	resp = &types.CreateResp{}
	resp.Status = types.StatusSuccess()

	db, err := l.svcCtx.UserDatabaseCache.MustLoad(0, "default")
	if err != nil {
		l.Logger.Errorf("UserDatabaseCache.MustLoad(%v, %v) failed, err: %v", 0, "default", err)
		return
	}

	sql, values := buildInsertSQL(req.Database, req.ValPairs)
	tx := db.Exec(sql, values...)
	if err = tx.Error; err != nil {
		l.Logger.Errorf("Exec %v failed, err: %v", sql, err)
		return
	}

	return
}

func buildInsertSQL(database string, valPairs []types.CreateValPair) (sql string, values []interface{}) {
	cols := make([]string, len(valPairs))
	placeholders := make([]string, len(valPairs))
	values = make([]interface{}, len(valPairs))
	for i, pair := range valPairs {
		cols[i] = pair.Col
		placeholders[i] = "?"
		values[i] = pair.Val
	}

	sql = fmt.Sprintf(insertFormat, database, strings.Join(cols, ", "), strings.Join(placeholders, ", "))
	return
}
