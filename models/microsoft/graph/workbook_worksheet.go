package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookWorksheet struct {
    Entity
    charts []WorkbookChart;
    name *string;
    names []WorkbookNamedItem;
    pivotTables []WorkbookPivotTable;
    position *int32;
    protection *WorkbookWorksheetProtection;
    tables []WorkbookTable;
    visibility *string;
}
func NewWorkbookWorksheet()(*WorkbookWorksheet) {
    m := &WorkbookWorksheet{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookWorksheet) GetCharts()([]WorkbookChart) {
    if m == nil {
        return nil
    } else {
        return m.charts
    }
}
func (m *WorkbookWorksheet) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *WorkbookWorksheet) GetNames()([]WorkbookNamedItem) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
func (m *WorkbookWorksheet) GetPivotTables()([]WorkbookPivotTable) {
    if m == nil {
        return nil
    } else {
        return m.pivotTables
    }
}
func (m *WorkbookWorksheet) GetPosition()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
func (m *WorkbookWorksheet) GetProtection()(*WorkbookWorksheetProtection) {
    if m == nil {
        return nil
    } else {
        return m.protection
    }
}
func (m *WorkbookWorksheet) GetTables()([]WorkbookTable) {
    if m == nil {
        return nil
    } else {
        return m.tables
    }
}
func (m *WorkbookWorksheet) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
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
func (m *WorkbookWorksheet) SetCharts(value []WorkbookChart)() {
    m.charts = value
}
func (m *WorkbookWorksheet) SetName(value *string)() {
    m.name = value
}
func (m *WorkbookWorksheet) SetNames(value []WorkbookNamedItem)() {
    m.names = value
}
func (m *WorkbookWorksheet) SetPivotTables(value []WorkbookPivotTable)() {
    m.pivotTables = value
}
func (m *WorkbookWorksheet) SetPosition(value *int32)() {
    m.position = value
}
func (m *WorkbookWorksheet) SetProtection(value *WorkbookWorksheetProtection)() {
    m.protection = value
}
func (m *WorkbookWorksheet) SetTables(value []WorkbookTable)() {
    m.tables = value
}
func (m *WorkbookWorksheet) SetVisibility(value *string)() {
    m.visibility = value
}
