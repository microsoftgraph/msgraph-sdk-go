package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookFunctionResult struct {
    Entity
    // 
    error *string;
    // 
    value *Json;
}
// Instantiates a new workbookFunctionResult and sets the default values.
func NewWorkbookFunctionResult()(*WorkbookFunctionResult) {
    m := &WorkbookFunctionResult{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the error property value. 
func (m *WorkbookFunctionResult) GetError()(*string) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the value property value. 
func (m *WorkbookFunctionResult) GetValue()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the error property value. 
// Parameters:
//  - value : Value to set for the error property.
func (m *WorkbookFunctionResult) SetError(value *string)() {
    m.error = value
}
// Sets the value property value. 
// Parameters:
//  - value : Value to set for the value property.
func (m *WorkbookFunctionResult) SetValue(value *Json)() {
    m.value = value
}
