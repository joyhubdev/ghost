package ghost

import (
	"context"
	"reflect"
)

type ContextObject struct{
	ctx context.Context
}

func (c ContextObject) SetCtx(ctx context.Context){
	c.ctx = ctx
}

func (c ContextObject) GetCtx() context.Context{
	return c.ctx
}

// DomainObject 领域对象(可以表示聚合根、聚合、实体、值对象和领域服务)
type DomainObject struct{
	ContextObject
}

// DomainModel 领域模型（表示实体）
type DomainModel struct{
	DomainObject
}

// NewFromDbModel
// 使用反射机制将dbModel中的field值复制到domainObject中
func (this *DomainObject) NewFromDbModel(dbModel interface{}){
	siType := reflect.TypeOf(dbModel)
	siValue := reflect.ValueOf(dbModel)
	if siType.Kind() == reflect.Ptr{
		siType = siType.Elem()
		siValue = siValue.Elem()
	}
	diValue := reflect.ValueOf(this).Elem()
	for i:=0; i<siType.NumField(); i++{
		fieldName := siType.Field(i).Name
		diFieldValue := diValue.FieldByName(fieldName)
		siFieldValue := siValue.Field(i)
		diFieldValue.Set(siFieldValue)
	}
}