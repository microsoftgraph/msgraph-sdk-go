package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookWorksheetProtectionOptions provides operations to manage the collection of drive entities.
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
// NewWorkbookWorksheetProtectionOptions instantiates a new workbookWorksheetProtectionOptions and sets the default values.
func NewWorkbookWorksheetProtectionOptions()(*WorkbookWorksheetProtectionOptions) {
    m := &WorkbookWorksheetProtectionOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbookWorksheetProtectionOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookWorksheetProtectionOptionsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookWorksheetProtectionOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookWorksheetProtectionOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowAutoFilter gets the allowAutoFilter property value. Represents the worksheet protection option of allowing using auto filter feature.
func (m *WorkbookWorksheetProtectionOptions) GetAllowAutoFilter()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAutoFilter
    }
}
// GetAllowDeleteColumns gets the allowDeleteColumns property value. Represents the worksheet protection option of allowing deleting columns.
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteColumns
    }
}
// GetAllowDeleteRows gets the allowDeleteRows property value. Represents the worksheet protection option of allowing deleting rows.
func (m *WorkbookWorksheetProtectionOptions) GetAllowDeleteRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteRows
    }
}
// GetAllowFormatCells gets the allowFormatCells property value. Represents the worksheet protection option of allowing formatting cells.
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatCells()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatCells
    }
}
// GetAllowFormatColumns gets the allowFormatColumns property value. Represents the worksheet protection option of allowing formatting columns.
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatColumns
    }
}
// GetAllowFormatRows gets the allowFormatRows property value. Represents the worksheet protection option of allowing formatting rows.
func (m *WorkbookWorksheetProtectionOptions) GetAllowFormatRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowFormatRows
    }
}
// GetAllowInsertColumns gets the allowInsertColumns property value. Represents the worksheet protection option of allowing inserting columns.
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertColumns
    }
}
// GetAllowInsertHyperlinks gets the allowInsertHyperlinks property value. Represents the worksheet protection option of allowing inserting hyperlinks.
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertHyperlinks()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertHyperlinks
    }
}
// GetAllowInsertRows gets the allowInsertRows property value. Represents the worksheet protection option of allowing inserting rows.
func (m *WorkbookWorksheetProtectionOptions) GetAllowInsertRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowInsertRows
    }
}
// GetAllowPivotTables gets the allowPivotTables property value. Represents the worksheet protection option of allowing using pivot table feature.
func (m *WorkbookWorksheetProtectionOptions) GetAllowPivotTables()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowPivotTables
    }
}
// GetAllowSort gets the allowSort property value. Represents the worksheet protection option of allowing using sort feature.
func (m *WorkbookWorksheetProtectionOptions) GetAllowSort()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowSort
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookWorksheetProtectionOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowAutoFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAutoFilter(val)
        }
        return nil
    }
    res["allowDeleteColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteColumns(val)
        }
        return nil
    }
    res["allowDeleteRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeleteRows(val)
        }
        return nil
    }
    res["allowFormatCells"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowFormatCells(val)
        }
        return nil
    }
    res["allowFormatColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowFormatColumns(val)
        }
        return nil
    }
    res["allowFormatRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowFormatRows(val)
        }
        return nil
    }
    res["allowInsertColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInsertColumns(val)
        }
        return nil
    }
    res["allowInsertHyperlinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInsertHyperlinks(val)
        }
        return nil
    }
    res["allowInsertRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowInsertRows(val)
        }
        return nil
    }
    res["allowPivotTables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPivotTables(val)
        }
        return nil
    }
    res["allowSort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSort(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookWorksheetProtectionOptions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookWorksheetProtectionOptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowAutoFilter sets the allowAutoFilter property value. Represents the worksheet protection option of allowing using auto filter feature.
func (m *WorkbookWorksheetProtectionOptions) SetAllowAutoFilter(value *bool)() {
    if m != nil {
        m.allowAutoFilter = value
    }
}
// SetAllowDeleteColumns sets the allowDeleteColumns property value. Represents the worksheet protection option of allowing deleting columns.
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteColumns(value *bool)() {
    if m != nil {
        m.allowDeleteColumns = value
    }
}
// SetAllowDeleteRows sets the allowDeleteRows property value. Represents the worksheet protection option of allowing deleting rows.
func (m *WorkbookWorksheetProtectionOptions) SetAllowDeleteRows(value *bool)() {
    if m != nil {
        m.allowDeleteRows = value
    }
}
// SetAllowFormatCells sets the allowFormatCells property value. Represents the worksheet protection option of allowing formatting cells.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatCells(value *bool)() {
    if m != nil {
        m.allowFormatCells = value
    }
}
// SetAllowFormatColumns sets the allowFormatColumns property value. Represents the worksheet protection option of allowing formatting columns.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatColumns(value *bool)() {
    if m != nil {
        m.allowFormatColumns = value
    }
}
// SetAllowFormatRows sets the allowFormatRows property value. Represents the worksheet protection option of allowing formatting rows.
func (m *WorkbookWorksheetProtectionOptions) SetAllowFormatRows(value *bool)() {
    if m != nil {
        m.allowFormatRows = value
    }
}
// SetAllowInsertColumns sets the allowInsertColumns property value. Represents the worksheet protection option of allowing inserting columns.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertColumns(value *bool)() {
    if m != nil {
        m.allowInsertColumns = value
    }
}
// SetAllowInsertHyperlinks sets the allowInsertHyperlinks property value. Represents the worksheet protection option of allowing inserting hyperlinks.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertHyperlinks(value *bool)() {
    if m != nil {
        m.allowInsertHyperlinks = value
    }
}
// SetAllowInsertRows sets the allowInsertRows property value. Represents the worksheet protection option of allowing inserting rows.
func (m *WorkbookWorksheetProtectionOptions) SetAllowInsertRows(value *bool)() {
    if m != nil {
        m.allowInsertRows = value
    }
}
// SetAllowPivotTables sets the allowPivotTables property value. Represents the worksheet protection option of allowing using pivot table feature.
func (m *WorkbookWorksheetProtectionOptions) SetAllowPivotTables(value *bool)() {
    if m != nil {
        m.allowPivotTables = value
    }
}
// SetAllowSort sets the allowSort property value. Represents the worksheet protection option of allowing using sort feature.
func (m *WorkbookWorksheetProtectionOptions) SetAllowSort(value *bool)() {
    if m != nil {
        m.allowSort = value
    }
}
