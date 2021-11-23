package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// timeRange 
type TimeRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // End time for the time range.
    endTime *string;
    // Start time for the time range.
    startTime *string;
}
// NewTimeRange instantiates a new timeRange and sets the default values.
func NewTimeRange()(*TimeRange) {
    m := &TimeRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEndTime gets the endTime property value. End time for the time range.
func (m *TimeRange) GetEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// GetStartTime gets the startTime property value. Start time for the time range.
func (m *TimeRange) GetStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["endTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["startTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    return res
}
func (m *TimeRange) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startTime", m.GetStartTime())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndTime sets the endTime property value. End time for the time range.
func (m *TimeRange) SetEndTime(value *string)() {
    m.endTime = value
}
// SetStartTime sets the startTime property value. Start time for the time range.
func (m *TimeRange) SetStartTime(value *string)() {
    m.startTime = value
}
