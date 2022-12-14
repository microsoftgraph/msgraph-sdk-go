package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody provides operations to call the getSchedule method.
type ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AvailabilityViewInterval property
    availabilityViewInterval *int32
    // The EndTime property
    endTime iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable
    // The Schedules property
    schedules []string
    // The StartTime property
    startTime iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable
}
// NewItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody instantiates a new ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody()(*ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) {
    m := &ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAvailabilityViewInterval gets the availabilityViewInterval property value. The AvailabilityViewInterval property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) GetAvailabilityViewInterval()(*int32) {
    return m.availabilityViewInterval
}
// GetEndTime gets the endTime property value. The EndTime property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) GetEndTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable) {
    return m.endTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable))
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
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetSchedules gets the schedules property value. The Schedules property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) GetSchedules()([]string) {
    return m.schedules
}
// GetStartTime gets the startTime property value. The StartTime property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) GetStartTime()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable) {
    return m.startTime
}
// Serialize serializes information the current object
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAvailabilityViewInterval sets the availabilityViewInterval property value. The AvailabilityViewInterval property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) SetAvailabilityViewInterval(value *int32)() {
    m.availabilityViewInterval = value
}
// SetEndTime sets the endTime property value. The EndTime property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) SetEndTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable)() {
    m.endTime = value
}
// SetSchedules sets the schedules property value. The Schedules property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) SetSchedules(value []string)() {
    m.schedules = value
}
// SetStartTime sets the startTime property value. The StartTime property
func (m *ItemCalendarGroupsItemCalendarsItemGetSchedulePostRequestBody) SetStartTime(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DateTimeTimeZoneable)() {
    m.startTime = value
}
