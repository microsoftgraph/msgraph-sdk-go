package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TimeOff struct {
    ChangeTrackedEntity
    // The draft version of this timeOff that is viewable by managers. Required.
    draftTimeOff *TimeOffItem;
    // The shared version of this timeOff that is viewable by both employees and managers. Required.
    sharedTimeOff *TimeOffItem;
    // ID of the user assigned to the timeOff. Required.
    userId *string;
}
// Instantiates a new timeOff and sets the default values.
func NewTimeOff()(*TimeOff) {
    m := &TimeOff{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// Gets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) GetDraftTimeOff()(*TimeOffItem) {
    if m == nil {
        return nil
    } else {
        return m.draftTimeOff
    }
}
// Gets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) GetSharedTimeOff()(*TimeOffItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedTimeOff
    }
}
// Gets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *TimeOff) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftTimeOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffItem() })
        if err != nil {
            return err
        }
        m.SetDraftTimeOff(val.(*TimeOffItem))
        return nil
    }
    res["sharedTimeOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffItem() })
        if err != nil {
            return err
        }
        m.SetSharedTimeOff(val.(*TimeOffItem))
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *TimeOff) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TimeOff) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftTimeOff", m.GetDraftTimeOff())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedTimeOff", m.GetSharedTimeOff())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
// Parameters:
//  - value : Value to set for the draftTimeOff property.
func (m *TimeOff) SetDraftTimeOff(value *TimeOffItem)() {
    m.draftTimeOff = value
}
// Sets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
// Parameters:
//  - value : Value to set for the sharedTimeOff property.
func (m *TimeOff) SetSharedTimeOff(value *TimeOffItem)() {
    m.sharedTimeOff = value
}
// Sets the userId property value. ID of the user assigned to the timeOff. Required.
// Parameters:
//  - value : Value to set for the userId property.
func (m *TimeOff) SetUserId(value *string)() {
    m.userId = value
}
