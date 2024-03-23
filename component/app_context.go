package component

import (
	"gorm.io/gorm"
	"ielts/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvide() uploadprovider.UpLoadProvider
}
type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UpLoadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UpLoadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
func (ctx *appCtx) UploadProvide() uploadprovider.UpLoadProvider {
	return ctx.upProvider

}
