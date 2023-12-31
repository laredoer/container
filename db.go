package container

import (
	"context"
	"errors"
	"time"

	db "git.5th.im/lb-public/gear/db/v2"
	"git.5th.im/lb-public/gear/log"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type DB interface {
	GetDBName() string
	GetDialect() string
}

// GetDB get db
func GetDB[T DB](ctx context.Context) *gorm.DB {
	var t T
	v, err := do.InvokeNamed[*db.DB](container.injector, t.GetDBName())
	if err != nil {
		log.Panicf("GetDB %s err:%v", t.GetDBName(), err)
	}

	return v.WithContext(ctx).Debug()
}

// ProcessWithPagination processes the query with pagination.
func ProcessWithPagination[T any](query *gorm.DB, pageSize int, processFunc func(rows *T) error) error {
	var errs []error
	page := 1
	for {
		var records []*T
		err := query.Limit(pageSize).Offset((page - 1) * pageSize).Find(&records).Error
		if err != nil {
			return err
		}

		if len(records) == 0 {
			break
		}

		for i := 0; i < len(records); i++ {
			time.Sleep(time.Millisecond * 500)
			err := processFunc(records[i])
			if err != nil {
				log.Errorf("processFunc error: %v", err)
				errs = append(errs, err)
				continue
			}
		}
		page++
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}
	return nil
}
