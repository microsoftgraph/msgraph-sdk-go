package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookFunctionResult struct {
    Entity
    error *string;
    value *Json;
}
func NewWorkbookFunctionResult()(*WorkbookFunctionResult) {
    m := &WorkbookFunctionResult{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookFunctionResult) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *WorkbookFunctionResult) GetValue()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *WorkbookFunctionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetError(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValue(val.(*Json))
        return nil
    }
    return res
}
func (m *WorkbookFunctionResult) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookFunctionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookFunctionResult) SetError(value *string)() {
    m.error = value
}
func (m *WorkbookFunctionResult) SetValue(value *Json)() {
    m.value = value
}
