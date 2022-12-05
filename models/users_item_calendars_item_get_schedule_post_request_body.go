package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemCalendarsItemGetSchedulePostRequestBody provides operations to call the getSchedule method.
type UsersItemCalendarsItemGetSchedulePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AvailabilityViewInterval property
    availabilityViewInterval *int32
    // The EndTime property
    endTime DateTimeTimeZoneable
    // The Schedules property
    schedules []string
    // The StartTime property
    startTime DateTimeTimeZoneable
}
// NewUsersItemCalendarsItemGetSchedulePostRequestBody instantiates a new UsersItemCalendarsItemGetSchedulePostRequestBody and sets the default values.
func NewUsersItemCalendarsItemGetSchedulePostRequestBody()(*UsersItemCalendarsItemGetSchedulePostRequestBody) {
    m := &UsersItemCalendarsItemGetSchedulePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemCalendarsItemGetSchedulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemCalendarsItemGetSchedulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemCalendarsItemGetSchedulePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAvailabilityViewInterval gets the availabilityViewInterval property value. The AvailabilityViewInterval property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) GetAvailabilityViewInterval()(*int32) {
    return m.availabilityViewInterval
}
// GetEndTime gets the endTime property value. The EndTime property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) GetEndTime()(DateTimeTimeZoneable) {
    return m.endTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityViewInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityViewInterval(val)
        }
        return nil
    }
    res["endTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["schedules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSchedules(res)
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetSchedules gets the schedules property value. The Schedules property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) GetSchedules()([]string) {
    return m.schedules
}
// GetStartTime gets the startTime property value. The StartTime property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) GetStartTime()(DateTimeTimeZoneable) {
    return m.startTime
}
// Serialize serializes information the current object
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("availabilityViewInterval", m.GetAvailabilityViewInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    if m.GetSchedules() != nil {
        err := writer.WriteCollectionOfStringValues("schedules", m.GetSchedules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startTime", m.GetStartTime())
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
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAvailabilityViewInterval sets the availabilityViewInterval property value. The AvailabilityViewInterval property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) SetAvailabilityViewInterval(value *int32)() {
    m.availabilityViewInterval = value
}
// SetEndTime sets the endTime property value. The EndTime property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) SetEndTime(value DateTimeTimeZoneable)() {
    m.endTime = value
}
// SetSchedules sets the schedules property value. The Schedules property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) SetSchedules(value []string)() {
    m.schedules = value
}
// SetStartTime sets the startTime property value. The StartTime property
func (m *UsersItemCalendarsItemGetSchedulePostRequestBody) SetStartTime(value DateTimeTimeZoneable)() {
    m.startTime = value
}
