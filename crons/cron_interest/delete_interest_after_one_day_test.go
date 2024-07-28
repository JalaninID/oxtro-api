package cron_interest

import (
	"app/domain/mocks"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCronInterest_DeleteInterestAfterOneDay(t *testing.T) {
	mockRepoInterest := new(mocks.RepositoryInterest)
	cron := NewCronInterest(&logrus.Logger{}, mockRepoInterest)
	mockRepoInterest.On("DeletePartnerExpired").Return(nil)
	cron.DeleteInterestAfterOneDay()
	assert.Equal(t, "app/domain.RepositoryInterest.DeletePartnerExpired-fm", cron.CronNameInterestExpired)
	time.Sleep(10 * time.Second)
}
