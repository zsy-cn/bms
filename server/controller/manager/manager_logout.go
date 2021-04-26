package manager

import (
	"github.com/henrylee2cn/faygo"

	"github.com/zsy-cn/bms/server/controller"
)

// LogoutParams ...
type LogoutParams struct{}

// Serve ...
func (p *LogoutParams) Serve(ctx *faygo.Context) (err error) {
	logger.Debug("request manager logout controller")
	result := controller.NewResult()
	defer ctx.JSON(200, result, true)

	manager := ctx.GetSession("Manager")
	if manager != nil {
		// result.Data = manager
		ctx.DelSession("Manager")
	}
	return
}
