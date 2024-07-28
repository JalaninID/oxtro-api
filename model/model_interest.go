package model

import "time"

type Interest struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	UserID         string
	IsInterest     bool
	InterestUserID string
}

func (i *Interest) TableName() string {
	return "interests"
}

type FilterInterest struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Gender         string
	UserID         string
	IsInterest     bool
	InterestUserID []string
	PerPage        int
	Page           int
	Offset         int
	Sort           string
}

func (f *FilterInterest) GetSort() string {
	if f.Sort == "" {
		f.Sort = "created_at desc"
	} else {
		f.Sort = f.Sort + " desc"
	}
	return f.Sort
}
