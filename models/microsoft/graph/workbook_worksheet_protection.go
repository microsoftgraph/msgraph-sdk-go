package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookWorksheetProtection struct {
    Entity
    options *WorkbookWorksheetProtectionOptions;
    protected *bool;
}
func NewWorkbookWorksheetProtection()(*WorkbookWorksheetProtection) {
    m := &WorkbookWorksheetProtection{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookWorksheetProtection) GetOptions()(*WorkbookWorksheetProtectionOptions) {
    if m == nil {
        return nil
    } else {
        return m.options
    }
}
func (m *WorkbookWorksheetProtection) GetProtected()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protected
    }
}
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
func (m *WorkbookWorksheetProtection) SetOptions(value *WorkbookWorksheetProtectionOptions)() {
    m.options = value
}
func (m *WorkbookWorksheetProtection) SetProtected(value *bool)() {
    m.protected = value
}
