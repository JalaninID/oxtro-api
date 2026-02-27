package sample_crm

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, contact CRMContact) (CRMContact, error) {
	err := r.db.WithContext(ctx).Create(&contact).Error
	return contact, err
}

func (r *repository) Detail(ctx context.Context, uuid string) (CRMContact, error) {
	var contact CRMContact
	err := r.db.WithContext(ctx).Where("uuid = ? AND deleted_at IS NULL", uuid).First(&contact).Error
	return contact, err
}

func (r *repository) List(ctx context.Context, search string, page, perPage int) ([]CRMContact, int64, error) {
	var contacts []CRMContact
	var total int64

	q := r.db.WithContext(ctx).Model(&CRMContact{}).Where("deleted_at IS NULL")
	if search != "" {
		q = q.Where("name ILIKE ? OR email ILIKE ? OR company ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	q.Count(&total)

	if perPage <= 0 {
		perPage = 20
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * perPage

	err := q.Order("created_at DESC").Offset(offset).Limit(perPage).Find(&contacts).Error
	return contacts, total, err
}

func (r *repository) Update(ctx context.Context, uuid string, contact CRMContact) (CRMContact, error) {
	err := r.db.WithContext(ctx).Where("uuid = ? AND deleted_at IS NULL", uuid).Updates(&contact).Error
	if err != nil {
		return CRMContact{}, err
	}
	return r.Detail(ctx, uuid)
}

func (r *repository) Delete(ctx context.Context, uuid string) error {
	return r.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(&CRMContact{}).Error
}
