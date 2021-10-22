package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookSessionInfo struct {
    additionalData map[string]interface{};
    id *string;
    persistChanges *bool;
}
func NewWorkbookSessionInfo()(*WorkbookSessionInfo) {
    m := &WorkbookSessionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WorkbookSessionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WorkbookSessionInfo) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *WorkbookSessionInfo) GetPersistChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistChanges
    }
}
func (m *WorkbookSessionInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["persistChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPersistChanges(val)
        return nil
    }
    return res
}
func (m *WorkbookSessionInfo) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookSessionInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("persistChanges", m.GetPersistChanges())
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
func (m *WorkbookSessionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WorkbookSessionInfo) SetId(value *string)() {
    m.id = value
}
func (m *WorkbookSessionInfo) SetPersistChanges(value *bool)() {
    m.persistChanges = value
}
