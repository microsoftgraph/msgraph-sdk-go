package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeRange 
type TimeRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // End time for the time range.
    endTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly;
    // Start time for the time range.
    startTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly;
}
// NewTimeRange instantiates a new timeRange and sets the default values.
func NewTimeRange()(*TimeRange) {
    m := &TimeRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTimeRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEndTime gets the endTime property value. End time for the time range.
func (m *TimeRange) GetEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeRange) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["startTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
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
// GetStartTime gets the startTime property value. Start time for the time range.
func (m *TimeRange) GetStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// Serialize serializes information the current object
func (m *TimeRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeOnlyValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("startTime", m.GetStartTime())
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
func (m *TimeRange) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEndTime sets the endTime property value. End time for the time range.
func (m *TimeRange) SetEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.endTime = value
    }
}
// SetStartTime sets the startTime property value. Start time for the time range.
func (m *TimeRange) SetStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.startTime = value
    }
}
