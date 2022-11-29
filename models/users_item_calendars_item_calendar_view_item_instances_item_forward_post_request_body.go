package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody provides operations to call the forward method.
type UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Comment property
    comment *string
    // The ToRecipients property
    toRecipients []Recipientable
}
// NewUsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody instantiates a new UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody and sets the default values.
func NewUsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody()(*UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) {
    m := &UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetComment)
    res["toRecipients"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue , m.SetToRecipients)
    return res
}
// GetToRecipients gets the toRecipients property value. The ToRecipients property
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) GetToRecipients()([]Recipientable) {
    return m.toRecipients
}
// Serialize serializes information the current object
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    if m.GetToRecipients() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetToRecipients())
        err := writer.WriteCollectionOfObjectValues("toRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetToRecipients sets the toRecipients property value. The ToRecipients property
func (m *UsersItemCalendarsItemCalendarViewItemInstancesItemForwardPostRequestBody) SetToRecipients(value []Recipientable)() {
    m.toRecipients = value
}
