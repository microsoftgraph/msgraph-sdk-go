package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemReminderViewWithStartDateTimeWithEndDateTimeResponse 
// Deprecated: This class is obsolete. Use reminderViewWithStartDateTimeWithEndDateTimeGetResponse instead.
type ItemReminderViewWithStartDateTimeWithEndDateTimeResponse struct {
    ItemReminderViewWithStartDateTimeWithEndDateTimeGetResponse
}
// NewItemReminderViewWithStartDateTimeWithEndDateTimeResponse instantiates a new ItemReminderViewWithStartDateTimeWithEndDateTimeResponse and sets the default values.
func NewItemReminderViewWithStartDateTimeWithEndDateTimeResponse()(*ItemReminderViewWithStartDateTimeWithEndDateTimeResponse) {
    m := &ItemReminderViewWithStartDateTimeWithEndDateTimeResponse{
        ItemReminderViewWithStartDateTimeWithEndDateTimeGetResponse: *NewItemReminderViewWithStartDateTimeWithEndDateTimeGetResponse(),
    }
    return m
}
// CreateItemReminderViewWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemReminderViewWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemReminderViewWithStartDateTimeWithEndDateTimeResponse(), nil
}
// ItemReminderViewWithStartDateTimeWithEndDateTimeResponseable 
// Deprecated: This class is obsolete. Use reminderViewWithStartDateTimeWithEndDateTimeGetResponse instead.
type ItemReminderViewWithStartDateTimeWithEndDateTimeResponseable interface {
    ItemReminderViewWithStartDateTimeWithEndDateTimeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
