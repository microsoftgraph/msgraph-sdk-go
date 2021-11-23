package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeOffItem 
type TimeOffItem struct {
    ScheduleEntity
    // ID of the timeOffReason for this timeOffItem. Required.
    timeOffReasonId *string;
}
// NewTimeOffItem instantiates a new timeOffItem and sets the default values.
func NewTimeOffItem()(*TimeOffItem) {
    m := &TimeOffItem{
        ScheduleEntity: *NewScheduleEntity(),
    }
    return m
}
// GetTimeOffReasonId gets the timeOffReasonId property value. ID of the timeOffReason for this timeOffItem. Required.
func (m *TimeOffItem) GetTimeOffReasonId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeOffReasonId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOffItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleEntity.GetFieldDeserializers()
    res["timeOffReasonId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeOffReasonId(val)
        }
        return nil
    }
    return res
}
func (m *TimeOffItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeOffItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ScheduleEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("timeOffReasonId", m.GetTimeOffReasonId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTimeOffReasonId sets the timeOffReasonId property value. ID of the timeOffReason for this timeOffItem. Required.
func (m *TimeOffItem) SetTimeOffReasonId(value *string)() {
    m.timeOffReasonId = value
}
