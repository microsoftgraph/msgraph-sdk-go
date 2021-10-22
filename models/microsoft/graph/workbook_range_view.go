package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookRangeView struct {
    Entity
    cellAddresses *Json;
    columnCount *int32;
    formulas *Json;
    formulasLocal *Json;
    formulasR1C1 *Json;
    index *int32;
    numberFormat *Json;
    rowCount *int32;
    rows []WorkbookRangeView;
    text *Json;
    values *Json;
    valueTypes *Json;
}
func NewWorkbookRangeView()(*WorkbookRangeView) {
    m := &WorkbookRangeView{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookRangeView) GetCellAddresses()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.cellAddresses
    }
}
func (m *WorkbookRangeView) GetColumnCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.columnCount
    }
}
func (m *WorkbookRangeView) GetFormulas()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulas
    }
}
func (m *WorkbookRangeView) GetFormulasLocal()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulasLocal
    }
}
func (m *WorkbookRangeView) GetFormulasR1C1()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulasR1C1
    }
}
func (m *WorkbookRangeView) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
func (m *WorkbookRangeView) GetNumberFormat()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.numberFormat
    }
}
func (m *WorkbookRangeView) GetRowCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowCount
    }
}
func (m *WorkbookRangeView) GetRows()([]WorkbookRangeView) {
    if m == nil {
        return nil
    } else {
        return m.rows
    }
}
func (m *WorkbookRangeView) GetText()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *WorkbookRangeView) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *WorkbookRangeView) GetValueTypes()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.valueTypes
    }
}
func (m *WorkbookRangeView) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cellAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetCellAddresses(val.(*Json))
        return nil
    }
    res["columnCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetColumnCount(val)
        return nil
    }
    res["formulas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetFormulas(val.(*Json))
        return nil
    }
    res["formulasLocal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetFormulasLocal(val.(*Json))
        return nil
    }
    res["formulasR1C1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetFormulasR1C1(val.(*Json))
        return nil
    }
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIndex(val)
        return nil
    }
    res["numberFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetNumberFormat(val.(*Json))
        return nil
    }
    res["rowCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRowCount(val)
        return nil
    }
    res["rows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeView() })
        if err != nil {
            return err
        }
        res := make([]WorkbookRangeView, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookRangeView))
        }
        m.SetRows(res)
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetText(val.(*Json))
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValues(val.(*Json))
        return nil
    }
    res["valueTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValueTypes(val.(*Json))
        return nil
    }
    return res
}
func (m *WorkbookRangeView) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookRangeView) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cellAddresses", m.GetCellAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("columnCount", m.GetColumnCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulas", m.GetFormulas())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasLocal", m.GetFormulasLocal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("formulasR1C1", m.GetFormulasR1C1())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("numberFormat", m.GetNumberFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rowCount", m.GetRowCount())
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
        err = writer.WriteObjectValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("values", m.GetValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("valueTypes", m.GetValueTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookRangeView) SetCellAddresses(value *Json)() {
    m.cellAddresses = value
}
func (m *WorkbookRangeView) SetColumnCount(value *int32)() {
    m.columnCount = value
}
func (m *WorkbookRangeView) SetFormulas(value *Json)() {
    m.formulas = value
}
func (m *WorkbookRangeView) SetFormulasLocal(value *Json)() {
    m.formulasLocal = value
}
func (m *WorkbookRangeView) SetFormulasR1C1(value *Json)() {
    m.formulasR1C1 = value
}
func (m *WorkbookRangeView) SetIndex(value *int32)() {
    m.index = value
}
func (m *WorkbookRangeView) SetNumberFormat(value *Json)() {
    m.numberFormat = value
}
func (m *WorkbookRangeView) SetRowCount(value *int32)() {
    m.rowCount = value
}
func (m *WorkbookRangeView) SetRows(value []WorkbookRangeView)() {
    m.rows = value
}
func (m *WorkbookRangeView) SetText(value *Json)() {
    m.text = value
}
func (m *WorkbookRangeView) SetValues(value *Json)() {
    m.values = value
}
func (m *WorkbookRangeView) SetValueTypes(value *Json)() {
    m.valueTypes = value
}
