package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookSessionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Id of the workbook session.
    id *string;
    // true for persistent session. false for non-persistent session (view mode)
    persistChanges *bool;
}
// Instantiates a new workbookSessionInfo and sets the default values.
func NewWorkbookSessionInfo()(*WorkbookSessionInfo) {
    m := &WorkbookSessionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookSessionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. Id of the workbook session.
func (m *WorkbookSessionInfo) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the persistChanges property value. true for persistent session. false for non-persistent session (view mode)
func (m *WorkbookSessionInfo) GetPersistChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.persistChanges
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkbookSessionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. Id of the workbook session.
// Parameters:
//  - value : Value to set for the id property.
func (m *WorkbookSessionInfo) SetId(value *string)() {
    m.id = value
}
// Sets the persistChanges property value. true for persistent session. false for non-persistent session (view mode)
// Parameters:
//  - value : Value to set for the persistChanges property.
func (m *WorkbookSessionInfo) SetPersistChanges(value *bool)() {
    m.persistChanges = value
}
