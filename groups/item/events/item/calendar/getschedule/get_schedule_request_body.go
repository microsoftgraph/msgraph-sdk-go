package getschedule

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GetScheduleRequestBody 
type GetScheduleRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    availabilityViewInterval *int32;
    // 
    endTime *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
    // 
    schedules []string;
    // 
    startTime *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
}
// NewGetScheduleRequestBody instantiates a new getScheduleRequestBody and sets the default values.
func NewGetScheduleRequestBody()(*GetScheduleRequestBody) {
    m := &GetScheduleRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetScheduleRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAvailabilityViewInterval gets the availabilityViewInterval property value. 
func (m *GetScheduleRequestBody) GetAvailabilityViewInterval()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.availabilityViewInterval
    }
}
// GetEndTime gets the endTime property value. 
func (m *GetScheduleRequestBody) GetEndTime()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// GetSchedules gets the schedules property value. 
func (m *GetScheduleRequestBody) GetSchedules()([]string) {
    if m == nil {
        return nil
    } else {
        return m.schedules
    }
}
// GetStartTime gets the startTime property value. 
func (m *GetScheduleRequestBody) GetStartTime()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetScheduleRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["availabilityViewInterval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityViewInterval(val)
        }
        return nil
    }
    res["endTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
        }
        return nil
    }
    res["schedules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["startTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *GetScheduleRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetScheduleRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
func (m *GetScheduleRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAvailabilityViewInterval sets the availabilityViewInterval property value. 
func (m *GetScheduleRequestBody) SetAvailabilityViewInterval(value *int32)() {
    if m != nil {
        m.availabilityViewInterval = value
    }
}
// SetEndTime sets the endTime property value. 
func (m *GetScheduleRequestBody) SetEndTime(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    if m != nil {
        m.endTime = value
    }
}
// SetSchedules sets the schedules property value. 
func (m *GetScheduleRequestBody) SetSchedules(value []string)() {
    if m != nil {
        m.schedules = value
    }
}
// SetStartTime sets the startTime property value. 
func (m *GetScheduleRequestBody) SetStartTime(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    if m != nil {
        m.startTime = value
    }
}
