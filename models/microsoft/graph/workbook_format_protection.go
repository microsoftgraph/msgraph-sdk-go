package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// workbookFormatProtection 
type WorkbookFormatProtection struct {
    Entity
    // Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn't have uniform formula hidden setting.
    formulaHidden *bool;
    // Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn't have uniform lock setting.
    locked *bool;
}
// NewWorkbookFormatProtection instantiates a new workbookFormatProtection and sets the default values.
func NewWorkbookFormatProtection()(*WorkbookFormatProtection) {
    m := &WorkbookFormatProtection{
        Entity: *NewEntity(),
    }
    return m
}
// GetFormulaHidden gets the formulaHidden property value. Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn't have uniform formula hidden setting.
func (m *WorkbookFormatProtection) GetFormulaHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.formulaHidden
    }
}
// GetLocked gets the locked property value. Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn't have uniform lock setting.
func (m *WorkbookFormatProtection) GetLocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.locked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFormatProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["formulaHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulaHidden(val)
        }
        return nil
    }
    res["locked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocked(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookFormatProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookFormatProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("formulaHidden", m.GetFormulaHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locked", m.GetLocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormulaHidden sets the formulaHidden property value. Indicates if Excel hides the formula for the cells in the range. A null value indicates that the entire range doesn't have uniform formula hidden setting.
func (m *WorkbookFormatProtection) SetFormulaHidden(value *bool)() {
    m.formulaHidden = value
}
// SetLocked sets the locked property value. Indicates if Excel locks the cells in the object. A null value indicates that the entire range doesn't have uniform lock setting.
func (m *WorkbookFormatProtection) SetLocked(value *bool)() {
    m.locked = value
}
