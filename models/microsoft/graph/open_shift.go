package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OpenShift struct {
    ChangeTrackedEntity
    // An unpublished open shift.
    draftOpenShift *OpenShiftItem;
    // ID for the scheduling group that the open shift belongs to.
    schedulingGroupId *string;
    // A published open shift.
    sharedOpenShift *OpenShiftItem;
}
// Instantiates a new openShift and sets the default values.
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// Gets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) GetDraftOpenShift()(*OpenShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.draftOpenShift
    }
}
// Gets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroupId
    }
}
// Gets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) GetSharedOpenShift()(*OpenShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedOpenShift
    }
}
// The deserialization information for the current model
func (m *OpenShift) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftOpenShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftItem() })
        if err != nil {
            return err
        }
        m.SetDraftOpenShift(val.(*OpenShiftItem))
        return nil
    }
    res["schedulingGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSchedulingGroupId(val)
        return nil
    }
    res["sharedOpenShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftItem() })
        if err != nil {
            return err
        }
        m.SetSharedOpenShift(val.(*OpenShiftItem))
        return nil
    }
    return res
}
func (m *OpenShift) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OpenShift) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftOpenShift", m.GetDraftOpenShift())
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
        err = writer.WriteObjectValue("sharedOpenShift", m.GetSharedOpenShift())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the draftOpenShift property value. An unpublished open shift.
// Parameters:
//  - value : Value to set for the draftOpenShift property.
func (m *OpenShift) SetDraftOpenShift(value *OpenShiftItem)() {
    m.draftOpenShift = value
}
// Sets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
// Parameters:
//  - value : Value to set for the schedulingGroupId property.
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// Sets the sharedOpenShift property value. A published open shift.
// Parameters:
//  - value : Value to set for the sharedOpenShift property.
func (m *OpenShift) SetSharedOpenShift(value *OpenShiftItem)() {
    m.sharedOpenShift = value
}
