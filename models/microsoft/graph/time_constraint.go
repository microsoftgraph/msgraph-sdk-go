package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeConstraint 
type TimeConstraint struct {
    // The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
    activityDomain *ActivityDomain;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    timeSlots []TimeSlot;
}
// NewTimeConstraint instantiates a new timeConstraint and sets the default values.
func NewTimeConstraint()(*TimeConstraint) {
    m := &TimeConstraint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActivityDomain gets the activityDomain property value. The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
func (m *TimeConstraint) GetActivityDomain()(*ActivityDomain) {
    if m == nil {
        return nil
    } else {
        return m.activityDomain
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTimeSlots gets the timeSlots property value. 
func (m *TimeConstraint) GetTimeSlots()([]TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeConstraint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activityDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivityDomain)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDomain(val.(*ActivityDomain))
        }
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeSlot() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeSlot, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeSlot))
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    return res
}
func (m *TimeConstraint) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeConstraint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetActivityDomain() != nil {
        cast := (*m.GetActivityDomain()).String()
        err := writer.WriteStringValue("activityDomain", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTimeSlots() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("timeSlots", cast)
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
// SetActivityDomain sets the activityDomain property value. The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
func (m *TimeConstraint) SetActivityDomain(value *ActivityDomain)() {
    if m != nil {
        m.activityDomain = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeConstraint) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTimeSlots sets the timeSlots property value. 
func (m *TimeConstraint) SetTimeSlots(value []TimeSlot)() {
    if m != nil {
        m.timeSlots = value
    }
}
