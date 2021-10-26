package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TimeConstraint struct {
    // The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
    activityDomain *ActivityDomain;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    timeSlots []TimeSlot;
}
// Instantiates a new timeConstraint and sets the default values.
func NewTimeConstraint()(*TimeConstraint) {
    m := &TimeConstraint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the activityDomain property value. The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
func (m *TimeConstraint) GetActivityDomain()(*ActivityDomain) {
    if m == nil {
        return nil
    } else {
        return m.activityDomain
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the timeSlots property value. 
func (m *TimeConstraint) GetTimeSlots()([]TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
// The deserialization information for the current model
func (m *TimeConstraint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activityDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivityDomain)
        if err != nil {
            return err
        }
        cast := val.(ActivityDomain)
        m.SetActivityDomain(&cast)
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeSlot() })
        if err != nil {
            return err
        }
        res := make([]TimeSlot, len(val))
        for i, v := range val {
            res[i] = *(v.(*TimeSlot))
        }
        m.SetTimeSlots(res)
        return nil
    }
    return res
}
func (m *TimeConstraint) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TimeConstraint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetActivityDomain() != nil {
        cast := m.GetActivityDomain().String()
        err := writer.WriteStringValue("activityDomain", &cast)
        if err != nil {
            return err
        }
    }
    {
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
// Sets the activityDomain property value. The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown.
// Parameters:
//  - value : Value to set for the activityDomain property.
func (m *TimeConstraint) SetActivityDomain(value *ActivityDomain)() {
    m.activityDomain = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TimeConstraint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the timeSlots property value. 
// Parameters:
//  - value : Value to set for the timeSlots property.
func (m *TimeConstraint) SetTimeSlots(value []TimeSlot)() {
    m.timeSlots = value
}
