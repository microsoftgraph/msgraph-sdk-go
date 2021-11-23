package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookWorksheet 
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
// NewWorkbookWorksheet instantiates a new workbookWorksheet and sets the default values.
func NewWorkbookWorksheet()(*WorkbookWorksheet) {
    m := &WorkbookWorksheet{
        Entity: *NewEntity(),
    }
    return m
}
// GetCharts gets the charts property value. Returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) GetCharts()([]WorkbookChart) {
    if m == nil {
        return nil
    } else {
        return m.charts
    }
}
// GetName gets the name property value. The display name of the worksheet.
func (m *WorkbookWorksheet) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetNames gets the names property value. Returns collection of names that are associated with the worksheet. Read-only.
func (m *WorkbookWorksheet) GetNames()([]WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// GetPivotTables gets the pivotTables property value. Collection of PivotTables that are part of the worksheet.
func (m *WorkbookWorksheet) GetPivotTables()([]WorkbookPivotTable) {
    if m == nil {
        return nil
    } else {
        return m.pivotTables
    }
}
// GetPosition gets the position property value. The zero-based position of the worksheet within the workbook.
func (m *WorkbookWorksheet) GetPosition()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// GetProtection gets the protection property value. Returns sheet protection object for a worksheet. Read-only.
func (m *WorkbookWorksheet) GetProtection()(*WorkbookWorksheetProtection) {
    if m == nil {
        return nil
    } else {
        return m.protection
    }
}
// GetTables gets the tables property value. Collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) GetTables()([]WorkbookTable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
// GetVisibility gets the visibility property value. The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
func (m *WorkbookWorksheet) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookWorksheet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["charts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChart() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookChart, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookChart))
            }
            m.SetCharts(res)
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
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookNamedItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookNamedItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookNamedItem))
            }
            m.SetNames(res)
        }
        return nil
    }
    res["pivotTables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookPivotTable() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookPivotTable, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookPivotTable))
            }
            m.SetPivotTables(res)
        }
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    res["protection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheetProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtection(val.(*WorkbookWorksheetProtection))
        }
        return nil
    }
    res["tables"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTable() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookTable, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookTable))
            }
            m.SetTables(res)
        }
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
func (m *WorkbookWorksheet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCharts sets the charts property value. Returns collection of charts that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) SetCharts(value []WorkbookChart)() {
    m.charts = value
}
// SetName sets the name property value. The display name of the worksheet.
func (m *WorkbookWorksheet) SetName(value *string)() {
    m.name = value
}
// SetNames sets the names property value. Returns collection of names that are associated with the worksheet. Read-only.
func (m *WorkbookWorksheet) SetNames(value []WorkbookNamedItem)() {
    m.names = value
}
// SetPivotTables sets the pivotTables property value. Collection of PivotTables that are part of the worksheet.
func (m *WorkbookWorksheet) SetPivotTables(value []WorkbookPivotTable)() {
    m.pivotTables = value
}
// SetPosition sets the position property value. The zero-based position of the worksheet within the workbook.
func (m *WorkbookWorksheet) SetPosition(value *int32)() {
    m.position = value
}
// SetProtection sets the protection property value. Returns sheet protection object for a worksheet. Read-only.
func (m *WorkbookWorksheet) SetProtection(value *WorkbookWorksheetProtection)() {
    m.protection = value
}
// SetTables sets the tables property value. Collection of tables that are part of the worksheet. Read-only.
func (m *WorkbookWorksheet) SetTables(value []WorkbookTable)() {
    m.tables = value
}
// SetVisibility sets the visibility property value. The Visibility of the worksheet. The possible values are: Visible, Hidden, VeryHidden.
func (m *WorkbookWorksheet) SetVisibility(value *string)() {
    m.visibility = value
}
