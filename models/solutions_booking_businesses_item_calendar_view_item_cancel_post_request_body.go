package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody provides operations to call the cancel method.
type SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The cancellationMessage property
    cancellationMessage *string
}
// NewSolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody instantiates a new SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody and sets the default values.
func NewSolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody()(*SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) {
    m := &SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCancellationMessage gets the cancellationMessage property value. The cancellationMessage property
func (m *SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) GetCancellationMessage()(*string) {
    return m.cancellationMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cancellationMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancellationMessage(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cancellationMessage", m.GetCancellationMessage())
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
func (m *SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCancellationMessage sets the cancellationMessage property value. The cancellationMessage property
func (m *SolutionsBookingBusinessesItemCalendarViewItemCancelPostRequestBody) SetCancellationMessage(value *string)() {
    m.cancellationMessage = value
}
