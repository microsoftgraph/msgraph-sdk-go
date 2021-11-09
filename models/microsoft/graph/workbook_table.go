package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookTable struct {
    Entity
    // Represents a collection of all the columns in the table. Read-only.
    columns []WorkbookTableColumn;
    // Indicates whether the first column contains special formatting.
    highlightFirstColumn *bool;
    // Indicates whether the last column contains special formatting.
    highlightLastColumn *bool;
    // Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
    legacyId *string;
    // Name of the table.
    name *string;
    // Represents a collection of all the rows in the table. Read-only.
    rows []WorkbookTableRow;
    // Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
    showBandedColumns *bool;
    // Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
    showBandedRows *bool;
    // Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
    showFilterButton *bool;
    // Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
    showHeaders *bool;
    // Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
    showTotals *bool;
    // Represents the sorting for the table. Read-only.
    sort *WorkbookTableSort;
    // Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
    style *string;
    // The worksheet containing the current table. Read-only.
    worksheet *WorkbookWorksheet;
}
// Instantiates a new workbookTable and sets the default values.
func NewWorkbookTable()(*WorkbookTable) {
    m := &WorkbookTable{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the columns property value. Represents a collection of all the columns in the table. Read-only.
func (m *WorkbookTable) GetColumns()([]WorkbookTableColumn) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// Gets the highlightFirstColumn property value. Indicates whether the first column contains special formatting.
func (m *WorkbookTable) GetHighlightFirstColumn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.highlightFirstColumn
    }
}
// Gets the highlightLastColumn property value. Indicates whether the last column contains special formatting.
func (m *WorkbookTable) GetHighlightLastColumn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.highlightLastColumn
    }
}
// Gets the legacyId property value. Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
func (m *WorkbookTable) GetLegacyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legacyId
    }
}
// Gets the name property value. Name of the table.
func (m *WorkbookTable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the rows property value. Represents a collection of all the rows in the table. Read-only.
func (m *WorkbookTable) GetRows()([]WorkbookTableRow) {
    if m == nil {
        return nil
    } else {
        return m.rows
    }
}
// Gets the showBandedColumns property value. Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
func (m *WorkbookTable) GetShowBandedColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBandedColumns
    }
}
// Gets the showBandedRows property value. Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
func (m *WorkbookTable) GetShowBandedRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBandedRows
    }
}
// Gets the showFilterButton property value. Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
func (m *WorkbookTable) GetShowFilterButton()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showFilterButton
    }
}
// Gets the showHeaders property value. Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
func (m *WorkbookTable) GetShowHeaders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showHeaders
    }
}
// Gets the showTotals property value. Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
func (m *WorkbookTable) GetShowTotals()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showTotals
    }
}
// Gets the sort property value. Represents the sorting for the table. Read-only.
func (m *WorkbookTable) GetSort()(*WorkbookTableSort) {
    if m == nil {
        return nil
    } else {
        return m.sort
    }
}
// Gets the style property value. Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
func (m *WorkbookTable) GetStyle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.style
    }
}
// Gets the worksheet property value. The worksheet containing the current table. Read-only.
func (m *WorkbookTable) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// The deserialization information for the current model
func (m *WorkbookTable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTableColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTableColumn, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookTableColumn))
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["highlightFirstColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighlightFirstColumn(val)
        }
        return nil
    }
    res["highlightLastColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighlightLastColumn(val)
        }
        return nil
    }
    res["legacyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegacyId(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["rows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTableRow() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTableRow, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookTableRow))
            }
            m.SetRows(res)
        }
        return nil
    }
    res["showBandedColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowBandedColumns(val)
        }
        return nil
    }
    res["showBandedRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowBandedRows(val)
        }
        return nil
    }
    res["showFilterButton"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowFilterButton(val)
        }
        return nil
    }
    res["showHeaders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowHeaders(val)
        }
        return nil
    }
    res["showTotals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowTotals(val)
        }
        return nil
    }
    res["sort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTableSort() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSort(val.(*WorkbookTableSort))
        }
        return nil
    }
    res["style"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStyle(val)
        }
        return nil
    }
    res["worksheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorksheet(val.(*WorkbookWorksheet))
        }
        return nil
    }
    return res
}
func (m *WorkbookTable) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookTable) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("highlightFirstColumn", m.GetHighlightFirstColumn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("highlightLastColumn", m.GetHighlightLastColumn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("legacyId", m.GetLegacyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRows()))
        for i, v := range m.GetRows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showBandedColumns", m.GetShowBandedColumns())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showBandedRows", m.GetShowBandedRows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showFilterButton", m.GetShowFilterButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showHeaders", m.GetShowHeaders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showTotals", m.GetShowTotals())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sort", m.GetSort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("style", m.GetStyle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("worksheet", m.GetWorksheet())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the columns property value. Represents a collection of all the columns in the table. Read-only.
// Parameters:
//  - value : Value to set for the columns property.
func (m *WorkbookTable) SetColumns(value []WorkbookTableColumn)() {
    m.columns = value
}
// Sets the highlightFirstColumn property value. Indicates whether the first column contains special formatting.
// Parameters:
//  - value : Value to set for the highlightFirstColumn property.
func (m *WorkbookTable) SetHighlightFirstColumn(value *bool)() {
    m.highlightFirstColumn = value
}
// Sets the highlightLastColumn property value. Indicates whether the last column contains special formatting.
// Parameters:
//  - value : Value to set for the highlightLastColumn property.
func (m *WorkbookTable) SetHighlightLastColumn(value *bool)() {
    m.highlightLastColumn = value
}
// Sets the legacyId property value. Legacy Id used in older Excle clients. The value of the identifier remains the same even when the table is renamed. This property should be interpreted as an opaque string value and should not be parsed to any other type. Read-only.
// Parameters:
//  - value : Value to set for the legacyId property.
func (m *WorkbookTable) SetLegacyId(value *string)() {
    m.legacyId = value
}
// Sets the name property value. Name of the table.
// Parameters:
//  - value : Value to set for the name property.
func (m *WorkbookTable) SetName(value *string)() {
    m.name = value
}
// Sets the rows property value. Represents a collection of all the rows in the table. Read-only.
// Parameters:
//  - value : Value to set for the rows property.
func (m *WorkbookTable) SetRows(value []WorkbookTableRow)() {
    m.rows = value
}
// Sets the showBandedColumns property value. Indicates whether the columns show banded formatting in which odd columns are highlighted differently from even ones to make reading the table easier.
// Parameters:
//  - value : Value to set for the showBandedColumns property.
func (m *WorkbookTable) SetShowBandedColumns(value *bool)() {
    m.showBandedColumns = value
}
// Sets the showBandedRows property value. Indicates whether the rows show banded formatting in which odd rows are highlighted differently from even ones to make reading the table easier.
// Parameters:
//  - value : Value to set for the showBandedRows property.
func (m *WorkbookTable) SetShowBandedRows(value *bool)() {
    m.showBandedRows = value
}
// Sets the showFilterButton property value. Indicates whether the filter buttons are visible at the top of each column header. Setting this is only allowed if the table contains a header row.
// Parameters:
//  - value : Value to set for the showFilterButton property.
func (m *WorkbookTable) SetShowFilterButton(value *bool)() {
    m.showFilterButton = value
}
// Sets the showHeaders property value. Indicates whether the header row is visible or not. This value can be set to show or remove the header row.
// Parameters:
//  - value : Value to set for the showHeaders property.
func (m *WorkbookTable) SetShowHeaders(value *bool)() {
    m.showHeaders = value
}
// Sets the showTotals property value. Indicates whether the total row is visible or not. This value can be set to show or remove the total row.
// Parameters:
//  - value : Value to set for the showTotals property.
func (m *WorkbookTable) SetShowTotals(value *bool)() {
    m.showTotals = value
}
// Sets the sort property value. Represents the sorting for the table. Read-only.
// Parameters:
//  - value : Value to set for the sort property.
func (m *WorkbookTable) SetSort(value *WorkbookTableSort)() {
    m.sort = value
}
// Sets the style property value. Constant value that represents the Table style. The possible values are: TableStyleLight1 thru TableStyleLight21, TableStyleMedium1 thru TableStyleMedium28, TableStyleStyleDark1 thru TableStyleStyleDark11. A custom user-defined style present in the workbook can also be specified.
// Parameters:
//  - value : Value to set for the style property.
func (m *WorkbookTable) SetStyle(value *string)() {
    m.style = value
}
// Sets the worksheet property value. The worksheet containing the current table. Read-only.
// Parameters:
//  - value : Value to set for the worksheet property.
func (m *WorkbookTable) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
