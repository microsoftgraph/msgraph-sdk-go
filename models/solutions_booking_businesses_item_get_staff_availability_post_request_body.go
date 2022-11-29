package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody provides operations to call the getStaffAvailability method.
type SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The endDateTime property
    endDateTime DateTimeTimeZoneable
    // The staffIds property
    staffIds []string
    // The startDateTime property
    startDateTime DateTimeTimeZoneable
}
// NewSolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody instantiates a new SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody and sets the default values.
func NewSolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody()(*SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) {
    m := &SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetEndDateTime()(DateTimeTimeZoneable) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetEndDateTime)
    res["staffIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetStaffIds)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue , m.SetStartDateTime)
    return res
}
// GetStaffIds gets the staffIds property value. The staffIds property
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetStaffIds()([]string) {
    return m.staffIds
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) GetStartDateTime()(DateTimeTimeZoneable) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStaffIds() != nil {
        err := writer.WriteCollectionOfStringValues("staffIds", m.GetStaffIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
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
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetEndDateTime(value DateTimeTimeZoneable)() {
    m.endDateTime = value
}
// SetStaffIds sets the staffIds property value. The staffIds property
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetStaffIds(value []string)() {
    m.staffIds = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *SolutionsBookingBusinessesItemGetStaffAvailabilityPostRequestBody) SetStartDateTime(value DateTimeTimeZoneable)() {
    m.startDateTime = value
}
