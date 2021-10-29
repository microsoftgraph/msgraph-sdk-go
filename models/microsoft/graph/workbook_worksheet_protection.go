package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookWorksheetProtection struct {
    Entity
    // Sheet protection options. Read-only.
    options *WorkbookWorksheetProtectionOptions;
    // Indicates if the worksheet is protected.  Read-only.
    protected *bool;
}
// Instantiates a new workbookWorksheetProtection and sets the default values.
func NewWorkbookWorksheetProtection()(*WorkbookWorksheetProtection) {
    m := &WorkbookWorksheetProtection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the options property value. Sheet protection options. Read-only.
func (m *WorkbookWorksheetProtection) GetOptions()(*WorkbookWorksheetProtectionOptions) {
    if m == nil {
        return nil
    } else {
        return m.options
    }
}
// Gets the protected property value. Indicates if the worksheet is protected.  Read-only.
func (m *WorkbookWorksheetProtection) GetProtected()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protected
    }
}
// The deserialization information for the current model
func (m *WorkbookWorksheetProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["options"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheetProtectionOptions() })
        if err != nil {
            return err
        }
        m.SetOptions(val.(*WorkbookWorksheetProtectionOptions))
        return nil
    }
    res["protected"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProtected(val)
        return nil
    }
    return res
}
func (m *WorkbookWorksheetProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookWorksheetProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("options", m.GetOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("protected", m.GetProtected())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the options property value. Sheet protection options. Read-only.
// Parameters:
//  - value : Value to set for the options property.
func (m *WorkbookWorksheetProtection) SetOptions(value *WorkbookWorksheetProtectionOptions)() {
    m.options = value
}
// Sets the protected property value. Indicates if the worksheet is protected.  Read-only.
// Parameters:
//  - value : Value to set for the protected property.
func (m *WorkbookWorksheetProtection) SetProtected(value *bool)() {
    m.protected = value
}
