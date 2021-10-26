package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookWorksheetProtectionOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents the worksheet protection option of allowing using auto filter feature.
    allowAutoFilter *bool;
    // Represents the worksheet protection option of allowing deleting columns.
    allowDeleteColumns *bool;
    // Represents the worksheet protection option of allowing deleting rows.
    allowDeleteRows *bool;
    // Represents the worksheet protection option of allowing formatting cells.
    allowFormatCells *bool;
    // Represents the worksheet protection option of allowing formatting columns.
    allowFormatColumns *bool;
    // Represents the worksheet protection option of allowing formatting rows.
    allowFormatRows *bool;
    // Represents the worksheet protection option of allowing inserting columns.
    allowInsertColumns *bool;
    // Represents the worksheet protection option of allowing inserting hyperlinks.
    allowInsertHyperlinks *bool;
    // Represents the worksheet protection option of allowing inserting rows.
    allowInsertRows *bool;
    // Represents the worksheet protection option of allowing using pivot table feature.
    allowPivotTables *bool;
    // Represents the worksheet protection option of allowing using sort feature.
    allowSort *bool;
}
// Instantiates a new workbookWorksheetProtectionOptions and sets the default values.
func NewWorkbookWorksheetProtectionOptions()(*WorkbookWorksheetProtectionOptions) {
    m := &WorkbookWorksheetProtectionOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookWorksheetProtectionOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowAutoFilter property value. Represents the worksheet protection option of allowing using auto filter feature.
func (m *WorkbookWorksheetProtectionOptions) GetAllowAutoFilter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAutoFilter
    }
}
// Gets the allowDeleteColumns property value. Represents the worksheet protection option of allowing deleting columns.
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteColumns
    }
}
// Gets the allowDeleteRows property value. Represents the worksheet protection option of allowing deleting rows.
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteRows
    }
}
// Gets the allowFormatCells property value. Represents the worksheet protection option of allowing formatting cells.
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatCells()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatCells
    }
}
// Gets the allowFormatColumns property value. Represents the worksheet protection option of allowing formatting columns.
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatColumns
    }
}
// Gets the allowFormatRows property value. Represents the worksheet protection option of allowing formatting rows.
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatRows
    }
}
// Gets the allowInsertColumns property value. Represents the worksheet protection option of allowing inserting columns.
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertColumns
    }
}
// Gets the allowInsertHyperlinks property value. Represents the worksheet protection option of allowing inserting hyperlinks.
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertHyperlinks()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertHyperlinks
    }
}
// Gets the allowInsertRows property value. Represents the worksheet protection option of allowing inserting rows.
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertRows
    }
}
// Gets the allowPivotTables property value. Represents the worksheet protection option of allowing using pivot table feature.
func (m *WorkbookWorksheetProtectionOptions) GetAllowPivotTables()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPivotTables
    }
}
// Gets the allowSort property value. Represents the worksheet protection option of allowing using sort feature.
func (m *WorkbookWorksheetProtectionOptions) GetAllowSort()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowSort
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkbookWorksheetProtectionOptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowAutoFilter property value. Represents the worksheet protection option of allowing using auto filter feature.
// Parameters:
//  - value : Value to set for the allowAutoFilter property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowAutoFilter(value *bool)() {
    m.allowAutoFilter = value
}
// Sets the allowDeleteColumns property value. Represents the worksheet protection option of allowing deleting columns.
// Parameters:
//  - value : Value to set for the allowDeleteColumns property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteColumns(value *bool)() {
    m.allowDeleteColumns = value
}
// Sets the allowDeleteRows property value. Represents the worksheet protection option of allowing deleting rows.
// Parameters:
//  - value : Value to set for the allowDeleteRows property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteRows(value *bool)() {
    m.allowDeleteRows = value
}
// Sets the allowFormatCells property value. Represents the worksheet protection option of allowing formatting cells.
// Parameters:
//  - value : Value to set for the allowFormatCells property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatCells(value *bool)() {
    m.allowFormatCells = value
}
// Sets the allowFormatColumns property value. Represents the worksheet protection option of allowing formatting columns.
// Parameters:
//  - value : Value to set for the allowFormatColumns property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatColumns(value *bool)() {
    m.allowFormatColumns = value
}
// Sets the allowFormatRows property value. Represents the worksheet protection option of allowing formatting rows.
// Parameters:
//  - value : Value to set for the allowFormatRows property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatRows(value *bool)() {
    m.allowFormatRows = value
}
// Sets the allowInsertColumns property value. Represents the worksheet protection option of allowing inserting columns.
// Parameters:
//  - value : Value to set for the allowInsertColumns property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertColumns(value *bool)() {
    m.allowInsertColumns = value
}
// Sets the allowInsertHyperlinks property value. Represents the worksheet protection option of allowing inserting hyperlinks.
// Parameters:
//  - value : Value to set for the allowInsertHyperlinks property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertHyperlinks(value *bool)() {
    m.allowInsertHyperlinks = value
}
// Sets the allowInsertRows property value. Represents the worksheet protection option of allowing inserting rows.
// Parameters:
//  - value : Value to set for the allowInsertRows property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertRows(value *bool)() {
    m.allowInsertRows = value
}
// Sets the allowPivotTables property value. Represents the worksheet protection option of allowing using pivot table feature.
// Parameters:
//  - value : Value to set for the allowPivotTables property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowPivotTables(value *bool)() {
    m.allowPivotTables = value
}
// Sets the allowSort property value. Represents the worksheet protection option of allowing using sort feature.
// Parameters:
//  - value : Value to set for the allowSort property.
func (m *WorkbookWorksheetProtectionOptions) SetAllowSort(value *bool)() {
    m.allowSort = value
}
