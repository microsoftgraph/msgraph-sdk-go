package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookFormatProtection struct {
    Entity
    formulaHidden *bool;
    locked *bool;
}
func NewWorkbookFormatProtection()(*WorkbookFormatProtection) {
    m := &WorkbookFormatProtection{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookFormatProtection) GetFormulaHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.formulaHidden
    }
}
func (m *WorkbookFormatProtection) GetLocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.locked
    }
}
func (m *WorkbookFormatProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["formulaHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFormulaHidden(val)
        return nil
    }
    res["locked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetLocked(val)
        return nil
    }
    return res
}
func (m *WorkbookFormatProtection) IsNil()(bool) {
    return m == nil
}
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
func (m *WorkbookFormatProtection) SetFormulaHidden(value *bool)() {
    m.formulaHidden = value
}
func (m *WorkbookFormatProtection) SetLocked(value *bool)() {
    m.locked = value
}
