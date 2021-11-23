package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// openShift 
type OpenShift struct {
    ChangeTrackedEntity
    // An unpublished open shift.
    draftOpenShift *OpenShiftItem;
    // ID for the scheduling group that the open shift belongs to.
    schedulingGroupId *string;
    // A published open shift.
    sharedOpenShift *OpenShiftItem;
}
// NewOpenShift instantiates a new openShift and sets the default values.
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// GetDraftOpenShift gets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) GetDraftOpenShift()(*OpenShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.draftOpenShift
    }
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroupId
    }
}
// GetSharedOpenShift gets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) GetSharedOpenShift()(*OpenShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedOpenShift
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenShift) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftOpenShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftOpenShift(val.(*OpenShiftItem))
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
    res["sharedOpenShift"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedOpenShift(val.(*OpenShiftItem))
        }
        return nil
    }
    return res
}
func (m *OpenShift) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDraftOpenShift sets the draftOpenShift property value. An unpublished open shift.
func (m *OpenShift) SetDraftOpenShift(value *OpenShiftItem)() {
    m.draftOpenShift = value
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID for the scheduling group that the open shift belongs to.
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// SetSharedOpenShift sets the sharedOpenShift property value. A published open shift.
func (m *OpenShift) SetSharedOpenShift(value *OpenShiftItem)() {
    m.sharedOpenShift = value
}
