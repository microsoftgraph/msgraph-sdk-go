package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookWorksheetProtection 
type WorkbookWorksheetProtection struct {
    Entity
    // Sheet protection options. Read-only.
    options WorkbookWorksheetProtectionOptionsable;
    // Indicates if the worksheet is protected.  Read-only.
    protected *bool;
}
// NewWorkbookWorksheetProtection instantiates a new workbookWorksheetProtection and sets the default values.
func NewWorkbookWorksheetProtection()(*WorkbookWorksheetProtection) {
    m := &WorkbookWorksheetProtection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookWorksheetProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookWorksheetProtectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookWorksheetProtection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookWorksheetProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["options"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetProtectionOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptions(val.(WorkbookWorksheetProtectionOptionsable))
        }
        return nil
    }
    res["protected"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtected(val)
        }
        return nil
    }
    return res
}
// GetOptions gets the options property value. Sheet protection options. Read-only.
func (m *WorkbookWorksheetProtection) GetOptions()(WorkbookWorksheetProtectionOptionsable) {
    if m == nil {
        return nil
    } else {
        return m.options
    }
}
// GetProtected gets the protected property value. Indicates if the worksheet is protected.  Read-only.
func (m *WorkbookWorksheetProtection) GetProtected()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protected
    }
}
// Serialize serializes information the current object
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
// SetOptions sets the options property value. Sheet protection options. Read-only.
func (m *WorkbookWorksheetProtection) SetOptions(value WorkbookWorksheetProtectionOptionsable)() {
    if m != nil {
        m.options = value
    }
}
// SetProtected sets the protected property value. Indicates if the worksheet is protected.  Read-only.
func (m *WorkbookWorksheetProtection) SetProtected(value *bool)() {
    if m != nil {
        m.protected = value
    }
}
