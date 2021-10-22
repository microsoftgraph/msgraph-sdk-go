package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OpenShift struct {
    ChangeTrackedEntity
    draftOpenShift *OpenShiftItem;
    schedulingGroupId *string;
    sharedOpenShift *OpenShiftItem;
}
func NewOpenShift()(*OpenShift) {
    m := &OpenShift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
func (m *OpenShift) GetDraftOpenShift()(*OpenShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.draftOpenShift
    }
}
func (m *OpenShift) GetSchedulingGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroupId
    }
}
func (m *OpenShift) GetSharedOpenShift()(*OpenShiftItem) {
    if m == nil {
        return nil
    } else {
        return m.sharedOpenShift
    }
}
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
func (m *OpenShift) SetDraftOpenShift(value *OpenShiftItem)() {
    m.draftOpenShift = value
}
func (m *OpenShift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
func (m *OpenShift) SetSharedOpenShift(value *OpenShiftItem)() {
    m.sharedOpenShift = value
}
