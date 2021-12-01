package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookFunctionResult 
type WorkbookFunctionResult struct {
    Entity
    // 
    error *string;
    // 
    value *Json;
}
// NewWorkbookFunctionResult instantiates a new workbookFunctionResult and sets the default values.
func NewWorkbookFunctionResult()(*WorkbookFunctionResult) {
    m := &WorkbookFunctionResult{
        Entity: *NewEntity(),
    }
    return m
}
// GetError gets the error property value. 
func (m *WorkbookFunctionResult) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetValue gets the value property value. 
func (m *WorkbookFunctionResult) GetValue()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFunctionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*Json))
        }
        return nil
    }
    return res
}
func (m *WorkbookFunctionResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetError sets the error property value. 
func (m *WorkbookFunctionResult) SetError(value *string)() {
    if m != nil {
        m.error = value
    }
}
// SetValue sets the value property value. 
func (m *WorkbookFunctionResult) SetValue(value *Json)() {
    if m != nil {
        m.value = value
    }
}
