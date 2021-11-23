package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeOff 
type TimeOff struct {
    ChangeTrackedEntity
    // The draft version of this timeOff that is viewable by managers. Required.
    draftTimeOff *TimeOffItem;
    // The shared version of this timeOff that is viewable by both employees and managers. Required.
    sharedTimeOff *TimeOffItem;
    // ID of the user assigned to the timeOff. Required.
    userId *string;
}
// NewTimeOff instantiates a new timeOff and sets the default values.
func NewTimeOff()(*TimeOff) {
    m := &TimeOff{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// GetDraftTimeOff gets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) GetDraftTimeOff()(*TimeOffItem) {
    if m == nil {
        return nil
    } else {
        return m.draftTimeOff
    }
}
// GetSharedTimeOff gets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) GetSharedTimeOff()(*TimeOffItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedTimeOff
    }
}
// GetUserId gets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOff) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftTimeOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftTimeOff(val.(*TimeOffItem))
        }
        return nil
    }
    res["sharedTimeOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedTimeOff(val.(*TimeOffItem))
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *TimeOff) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDraftTimeOff sets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) SetDraftTimeOff(value *TimeOffItem)() {
    m.draftTimeOff = value
}
// SetSharedTimeOff sets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) SetSharedTimeOff(value *TimeOffItem)() {
    m.sharedTimeOff = value
}
// SetUserId sets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) SetUserId(value *string)() {
    m.userId = value
}
