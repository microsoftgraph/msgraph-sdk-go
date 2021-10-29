package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookWorksheet struct {
    Entity
    // Returns collection of charts that are part of the worksheet. Read-only.
    charts []WorkbookChart;
    // The display name of the worksheet.
    name *string;
    // Returns collection of names that are associated with the worksheet. Read-only.
    names []WorkbookNamedItem;
    // Collection of PivotTables that are part of the worksheet.
    pivotTables []WorkbookPivotTable;
    // The zero-based position of the worksheet within the workbook.
    position *int32;
    // Returns sheet protection object for a worksheet. Read-only.
    protection *WorkbookWorksheetProtection;
    // Collection of tables that are part of the worksheet. Read-only.
    tables []WorkbookTable;
    // The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
    visibility *string;
}
// Instantiates a new workbookWorksheet and sets the default values.
func NewWorkbookWorksheet()(*WorkbookWorksheet) {
    m := &WorkbookWorksheet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the charts property value. Returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) GetCharts()([]WorkbookChart) {
    if m == nil {
        return nil
    } else {
        return m.charts
    }
}
// Gets the name property value. The display name of the worksheet.
func (m *WorkbookWorksheet) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the names property value. Returns collection of names that are associated with the worksheet. Read-only.
func (m *WorkbookWorksheet) GetNames()([]WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// Gets the pivotTables property value. Collection of PivotTables that are part of the worksheet.
func (m *WorkbookWorksheet) GetPivotTables()([]WorkbookPivotTable) {
    if m == nil {
        return nil
    } else {
        return m.pivotTables
    }
}
// Gets the position property value. The zero-based position of the worksheet within the workbook.
func (m *WorkbookWorksheet) GetPosition()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// Gets the protection property value. Returns sheet protection object for a worksheet. Read-only.
func (m *WorkbookWorksheet) GetProtection()(*WorkbookWorksheetProtection) {
    if m == nil {
        return nil
    } else {
        return m.protection
    }
}
// Gets the tables property value. Collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) GetTables()([]WorkbookTable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
// Gets the visibility property value. The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
func (m *WorkbookWorksheet) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// The deserialization information for the current model
func (m *WorkbookWorksheet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["charts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChart() })
        if err != nil {
            return err
        }
        res := make([]WorkbookChart, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookChart))
        }
        m.SetCharts(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookNamedItem() })
        if err != nil {
            return err
        }
        res := make([]WorkbookNamedItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookNamedItem))
        }
        m.SetNames(res)
        return nil
    }
    res["pivotTables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookPivotTable() })
        if err != nil {
            return err
        }
        res := make([]WorkbookPivotTable, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookPivotTable))
        }
        m.SetPivotTables(res)
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPosition(val)
        return nil
    }
    res["protection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheetProtection() })
        if err != nil {
            return err
        }
        m.SetProtection(val.(*WorkbookWorksheetProtection))
        return nil
    }
    res["tables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTable() })
        if err != nil {
            return err
        }
        res := make([]WorkbookTable, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookTable))
        }
        m.SetTables(res)
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVisibility(val)
        return nil
    }
    return res
}
func (m *WorkbookWorksheet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookWorksheet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCharts()))
        for i, v := range m.GetCharts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("charts", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPivotTables()))
        for i, v := range m.GetPivotTables() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("pivotTables", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("position", m.GetPosition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("protection", m.GetProtection())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTables()))
        for i, v := range m.GetTables() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tables", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the charts property value. Returns collection of charts that are part of the worksheet. Read-only.
// Parameters:
//  - value : Value to set for the charts property.
func (m *WorkbookWorksheet) SetCharts(value []WorkbookChart)() {
    m.charts = value
}
// Sets the name property value. The display name of the worksheet.
// Parameters:
//  - value : Value to set for the name property.
func (m *WorkbookWorksheet) SetName(value *string)() {
    m.name = value
}
// Sets the names property value. Returns collection of names that are associated with the worksheet. Read-only.
// Parameters:
//  - value : Value to set for the names property.
func (m *WorkbookWorksheet) SetNames(value []WorkbookNamedItem)() {
    m.names = value
}
// Sets the pivotTables property value. Collection of PivotTables that are part of the worksheet.
// Parameters:
//  - value : Value to set for the pivotTables property.
func (m *WorkbookWorksheet) SetPivotTables(value []WorkbookPivotTable)() {
    m.pivotTables = value
}
// Sets the position property value. The zero-based position of the worksheet within the workbook.
// Parameters:
//  - value : Value to set for the position property.
func (m *WorkbookWorksheet) SetPosition(value *int32)() {
    m.position = value
}
// Sets the protection property value. Returns sheet protection object for a worksheet. Read-only.
// Parameters:
//  - value : Value to set for the protection property.
func (m *WorkbookWorksheet) SetProtection(value *WorkbookWorksheetProtection)() {
    m.protection = value
}
// Sets the tables property value. Collection of tables that are part of the worksheet. Read-only.
// Parameters:
//  - value : Value to set for the tables property.
func (m *WorkbookWorksheet) SetTables(value []WorkbookTable)() {
    m.tables = value
}
// Sets the visibility property value. The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
// Parameters:
//  - value : Value to set for the visibility property.
func (m *WorkbookWorksheet) SetVisibility(value *string)() {
    m.visibility = value
}
