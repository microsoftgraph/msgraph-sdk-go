package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OpenShiftItem struct {
    ShiftItem
    openSlotCount *int32;
}
func NewOpenShiftItem()(*OpenShiftItem) {
    m := &OpenShiftItem{
        ShiftItem: *NewShiftItem(),
    }
    return m
}
func (m *OpenShiftItem) GetOpenSlotCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.openSlotCount
    }
}
func (m *OpenShiftItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ShiftItem.GetFieldDeserializers()
    res["openSlotCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOpenSlotCount(val)
        return nil
    }
    return res
}
func (m *OpenShiftItem) IsNil()(bool) {
    return m == nil
}
func (m *OpenShiftItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ShiftItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("openSlotCount", m.GetOpenSlotCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OpenShiftItem) SetOpenSlotCount(value *int32)() {
    m.openSlotCount = value
}
