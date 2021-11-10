package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Shift struct {
    ChangeTrackedEntity
    // The draft version of this shift that is viewable by managers. Required.
    draftShift *ShiftItem;
    // ID of the scheduling group the shift is part of. Required.
    schedulingGroupId *string;
    // The shared version of this shift that is viewable by both employees and managers. Required.
    sharedShift *ShiftItem;
    // ID of the user assigned to the shift. Required.
    userId *string;
}
// Instantiates a new shift and sets the default values.
func NewShift()(*Shift) {
    m := &Shift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// Gets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) GetDraftShift()(*ShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.draftShift
    }
}
// Gets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) GetSchedulingGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroupId
    }
}
// Gets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) GetSharedShift()(*ShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedShift
    }
}
// Gets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *Shift) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftShift(val.(*ShiftItem))
        }
        return nil
    }
    res["schedulingGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupId(val)
        }
        return nil
    }
    res["sharedShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedShift(val.(*ShiftItem))
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
func (m *Shift) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Shift) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftShift", m.GetDraftShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedShift", m.GetSharedShift())
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
// Sets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
// Parameters:
//  - value : Value to set for the draftShift property.
func (m *Shift) SetDraftShift(value *ShiftItem)() {
    m.draftShift = value
}
// Sets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
// Parameters:
//  - value : Value to set for the schedulingGroupId property.
func (m *Shift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// Sets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
// Parameters:
//  - value : Value to set for the sharedShift property.
func (m *Shift) SetSharedShift(value *ShiftItem)() {
    m.sharedShift = value
}
// Sets the userId property value. ID of the user assigned to the shift. Required.
// Parameters:
//  - value : Value to set for the userId property.
func (m *Shift) SetUserId(value *string)() {
    m.userId = value
}
