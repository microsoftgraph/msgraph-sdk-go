package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TimeConstraint struct {
    activityDomain *ActivityDomain;
    additionalData map[string]interface{};
    timeSlots []TimeSlot;
}
func NewTimeConstraint()(*TimeConstraint) {
    m := &TimeConstraint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TimeConstraint) GetActivityDomain()(*ActivityDomain) {
    if m == nil {
        return nil
    } else {
        return m.activityDomain
    }
}
func (m *TimeConstraint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TimeConstraint) GetTimeSlots()([]TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
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
func (m *TimeConstraint) SetActivityDomain(value *ActivityDomain)() {
    m.activityDomain = value
}
func (m *TimeConstraint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TimeConstraint) SetTimeSlots(value []TimeSlot)() {
    m.timeSlots = value
}
