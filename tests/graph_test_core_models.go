package tests

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompilationOfAddedModels(t *testing.T) {
	assert.NotNil(t, models.NewChangeNotification())
	assert.NotNil(t, models.NewChangeNotificationCollectionResponse())
	assert.NotNil(t, models.NewChangeNotificationEncryptedContent())
	assert.NotNil(t, models.CREATED_CHANGETYPE)
	assert.NotNil(t, models.MISSED_LIFECYCLEEVENTTYPE)
	assert.NotNil(t, models.NewResourceData())
	assert.NotNil(t, models.NewResourcePermission())
}
