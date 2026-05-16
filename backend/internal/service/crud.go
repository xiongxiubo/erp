package service

import (
	"erp/internal/model"
	"erp/internal/response"
	"erp/internal/utils"

	"gorm.io/gorm"
)

type CRUDService[T any] struct {
	db          *gorm.DB
	searchCols  []string
	defaultCode string
}

func NewCRUDService[T any](db *gorm.DB, searchCols []string, defaultCode string) *CRUDService[T] {
	return &CRUDService[T]{db: db, searchCols: searchCols, defaultCode: defaultCode}
}

func (s *CRUDService[T]) List(p utils.Pagination) ([]T, int64, error) {
	var items []T
	var total int64
	query := s.db.Model(new(T))
	query = applyListFilters(query, p, s.searchCols)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Limit(p.PageSize).Offset(p.Offset()).Find(&items).Error
	return items, total, err
}

func (s *CRUDService[T]) Get(id uint) (T, error) {
	var item T
	if err := s.db.First(&item, id).Error; err != nil {
		return item, response.NotFound("数据不存在")
	}
	return item, nil
}

func (s *CRUDService[T]) Create(item *T) error {
	return s.db.Create(item).Error
}

func (s *CRUDService[T]) Update(id uint, item *T) error {
	result := s.db.Model(new(T)).Where("id = ?", id).Updates(item)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return response.NotFound("数据不存在")
	}
	return nil
}

func (s *CRUDService[T]) Delete(id uint) error {
	result := s.db.Delete(new(T), id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return response.NotFound("数据不存在")
	}
	return nil
}

func applyListFilters(query *gorm.DB, p utils.Pagination, cols []string) *gorm.DB {
	if p.Keyword != "" && len(cols) > 0 {
		scoped := query.Session(&gorm.Session{})
		for i, col := range cols {
			condition := col + " LIKE ?"
			if i == 0 {
				scoped = scoped.Where(condition, "%"+p.Keyword+"%")
			} else {
				scoped = scoped.Or(condition, "%"+p.Keyword+"%")
			}
		}
		query = query.Where(scoped)
	}
	if p.Status != "" {
		query = query.Where("status = ?", p.Status)
	}
	return query
}

func EnsureStatus(status string) string {
	if status == "" {
		return model.StatusEnabled
	}
	return status
}
