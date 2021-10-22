package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookTable struct {
    Entity
    columns []WorkbookTableColumn;
    highlightFirstColumn *bool;
    highlightLastColumn *bool;
    legacyId *string;
    name *string;
    rows []WorkbookTableRow;
    showBandedColumns *bool;
    showBandedRows *bool;
    showFilterButton *bool;
    showHeaders *bool;
    showTotals *bool;
    sort *WorkbookTableSort;
    style *string;
    worksheet *WorkbookWorksheet;
}
func NewWorkbookTable()(*WorkbookTable) {
    m := &WorkbookTable{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookTable) GetColumns()([]WorkbookTableColumn) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
func (m *WorkbookTable) GetHighlightFirstColumn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.highlightFirstColumn
    }
}
func (m *WorkbookTable) GetHighlightLastColumn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.highlightLastColumn
    }
}
func (m *WorkbookTable) GetLegacyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.legacyId
    }
}
func (m *WorkbookTable) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *WorkbookTable) GetRows()([]WorkbookTableRow) {
    if m == nil {
        return nil
    } else {
        return m.rows
    }
}
func (m *WorkbookTable) GetShowBandedColumns()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBandedColumns
    }
}
func (m *WorkbookTable) GetShowBandedRows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showBandedRows
    }
}
func (m *WorkbookTable) GetShowFilterButton()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showFilterButton
    }
}
func (m *WorkbookTable) GetShowHeaders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showHeaders
    }
}
func (m *WorkbookTable) GetShowTotals()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showTotals
    }
}
func (m *WorkbookTable) GetSort()(*WorkbookTableSort) {
    if m == nil {
        return nil
    } else {
        return m.sort
    }
}
func (m *WorkbookTable) GetStyle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.style
    }
}
func (m *WorkbookTable) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
func (m *WorkbookTable) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTableColumn() })
        if err != nil {
            return err
        }
        res := make([]WorkbookTableColumn, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookTableColumn))
        }
        m.SetColumns(res)
        return nil
    }
    res["highlightFirstColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHighlightFirstColumn(val)
        return nil
    }
    res["highlightLastColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHighlightLastColumn(val)
        return nil
    }
    res["legacyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLegacyId(val)
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
    res["rows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTableRow() })
        if err != nil {
            return err
        }
        res := make([]WorkbookTableRow, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookTableRow))
        }
        m.SetRows(res)
        return nil
    }
    res["showBandedColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowBandedColumns(val)
        return nil
    }
    res["showBandedRows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowBandedRows(val)
        return nil
    }
    res["showFilterButton"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowFilterButton(val)
        return nil
    }
    res["showHeaders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowHeaders(val)
        return nil
    }
    res["showTotals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowTotals(val)
        return nil
    }
    res["sort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookTableSort() })
        if err != nil {
            return err
        }
        m.SetSort(val.(*WorkbookTableSort))
        return nil
    }
    res["style"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStyle(val)
        return nil
    }
    res["worksheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        m.SetWorksheet(val.(*WorkbookWorksheet))
        return nil
    }
    return res
}
func (m *WorkbookTable) IsNil()(bool) {
    return m == nil
}
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
func (m *WorkbookTable) SetColumns(value []WorkbookTableColumn)() {
    m.columns = value
}
func (m *WorkbookTable) SetHighlightFirstColumn(value *bool)() {
    m.highlightFirstColumn = value
}
func (m *WorkbookTable) SetHighlightLastColumn(value *bool)() {
    m.highlightLastColumn = value
}
func (m *WorkbookTable) SetLegacyId(value *string)() {
    m.legacyId = value
}
func (m *WorkbookTable) SetName(value *string)() {
    m.name = value
}
func (m *WorkbookTable) SetRows(value []WorkbookTableRow)() {
    m.rows = value
}
func (m *WorkbookTable) SetShowBandedColumns(value *bool)() {
    m.showBandedColumns = value
}
func (m *WorkbookTable) SetShowBandedRows(value *bool)() {
    m.showBandedRows = value
}
func (m *WorkbookTable) SetShowFilterButton(value *bool)() {
    m.showFilterButton = value
}
func (m *WorkbookTable) SetShowHeaders(value *bool)() {
    m.showHeaders = value
}
func (m *WorkbookTable) SetShowTotals(value *bool)() {
    m.showTotals = value
}
func (m *WorkbookTable) SetSort(value *WorkbookTableSort)() {
    m.sort = value
}
func (m *WorkbookTable) SetStyle(value *string)() {
    m.style = value
}
func (m *WorkbookTable) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
