package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookWorksheetProtectionOptions struct {
    additionalData map[string]interface{};
    allowAutoFilter *bool;
    allowDeleteColumns *bool;
    allowDeleteRows *bool;
    allowFormatCells *bool;
    allowFormatColumns *bool;
    allowFormatRows *bool;
    allowInsertColumns *bool;
    allowInsertHyperlinks *bool;
    allowInsertRows *bool;
    allowPivotTables *bool;
    allowSort *bool;
}
func NewWorkbookWorksheetProtectionOptions()(*WorkbookWorksheetProtectionOptions) {
    m := &WorkbookWorksheetProtectionOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WorkbookWorksheetProtectionOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowAutoFilter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAutoFilter
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteColumns
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteRows
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatCells()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatCells
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatColumns
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatRows
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertColumns
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertHyperlinks()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertHyperlinks
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertRows
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowPivotTables()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPivotTables
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetAllowSort()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowSort
    }
}
func (m *WorkbookWorksheetProtectionOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowAutoFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAutoFilter(val)
        return nil
    }
    res["allowDeleteColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowDeleteColumns(val)
        return nil
    }
    res["allowDeleteRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowDeleteRows(val)
        return nil
    }
    res["allowFormatCells"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowFormatCells(val)
        return nil
    }
    res["allowFormatColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowFormatColumns(val)
        return nil
    }
    res["allowFormatRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowFormatRows(val)
        return nil
    }
    res["allowInsertColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowInsertColumns(val)
        return nil
    }
    res["allowInsertHyperlinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowInsertHyperlinks(val)
        return nil
    }
    res["allowInsertRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowInsertRows(val)
        return nil
    }
    res["allowPivotTables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowPivotTables(val)
        return nil
    }
    res["allowSort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowSort(val)
        return nil
    }
    return res
}
func (m *WorkbookWorksheetProtectionOptions) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookWorksheetProtectionOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAutoFilter", m.GetAllowAutoFilter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteColumns", m.GetAllowDeleteColumns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteRows", m.GetAllowDeleteRows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowFormatCells", m.GetAllowFormatCells())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowFormatColumns", m.GetAllowFormatColumns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowFormatRows", m.GetAllowFormatRows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowInsertColumns", m.GetAllowInsertColumns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowInsertHyperlinks", m.GetAllowInsertHyperlinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowInsertRows", m.GetAllowInsertRows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowPivotTables", m.GetAllowPivotTables())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowSort", m.GetAllowSort())
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
func (m *WorkbookWorksheetProtectionOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowAutoFilter(value *bool)() {
    m.allowAutoFilter = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteColumns(value *bool)() {
    m.allowDeleteColumns = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteRows(value *bool)() {
    m.allowDeleteRows = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatCells(value *bool)() {
    m.allowFormatCells = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatColumns(value *bool)() {
    m.allowFormatColumns = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatRows(value *bool)() {
    m.allowFormatRows = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertColumns(value *bool)() {
    m.allowInsertColumns = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertHyperlinks(value *bool)() {
    m.allowInsertHyperlinks = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertRows(value *bool)() {
    m.allowInsertRows = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowPivotTables(value *bool)() {
    m.allowPivotTables = value
}
func (m *WorkbookWorksheetProtectionOptions) SetAllowSort(value *bool)() {
    m.allowSort = value
}
