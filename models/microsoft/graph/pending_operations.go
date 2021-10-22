package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PendingOperations struct {
    additionalData map[string]interface{};
    pendingContentUpdate *PendingContentUpdate;
}
func NewPendingOperations()(*PendingOperations) {
    m := &PendingOperations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PendingOperations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PendingOperations) GetPendingContentUpdate()(*PendingContentUpdate) {
    if m == nil {
        return nil
    } else {
        return m.pendingContentUpdate
    }
}
func (m *PendingOperations) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["pendingContentUpdate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPendingContentUpdate() })
        if err != nil {
            return err
        }
        m.SetPendingContentUpdate(val.(*PendingContentUpdate))
        return nil
    }
    return res
}
func (m *PendingOperations) IsNil()(bool) {
    return m == nil
}
func (m *PendingOperations) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pendingContentUpdate", m.GetPendingContentUpdate())
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
func (m *PendingOperations) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PendingOperations) SetPendingContentUpdate(value *PendingContentUpdate)() {
    m.pendingContentUpdate = value
}
