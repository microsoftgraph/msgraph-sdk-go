package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookRange struct {
    Entity
    address *string;
    addressLocal *string;
    cellCount *int32;
    columnCount *int32;
    columnHidden *bool;
    columnIndex *int32;
    format *WorkbookRangeFormat;
    formulas *Json;
    formulasLocal *Json;
    formulasR1C1 *Json;
    hidden *bool;
    numberFormat *Json;
    rowCount *int32;
    rowHidden *bool;
    rowIndex *int32;
    sort *WorkbookRangeSort;
    text *Json;
    values *Json;
    valueTypes *Json;
    worksheet *WorkbookWorksheet;
}
func NewWorkbookRange()(*WorkbookRange) {
    m := &WorkbookRange{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookRange) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *WorkbookRange) GetAddressLocal()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressLocal
    }
}
func (m *WorkbookRange) GetCellCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cellCount
    }
}
func (m *WorkbookRange) GetColumnCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.columnCount
    }
}
func (m *WorkbookRange) GetColumnHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.columnHidden
    }
}
func (m *WorkbookRange) GetColumnIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.columnIndex
    }
}
func (m *WorkbookRange) GetFormat()(*WorkbookRangeFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookRange) GetFormulas()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulas
    }
}
func (m *WorkbookRange) GetFormulasLocal()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulasLocal
    }
}
func (m *WorkbookRange) GetFormulasR1C1()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulasR1C1
    }
}
func (m *WorkbookRange) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
func (m *WorkbookRange) GetNumberFormat()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.numberFormat
    }
}
func (m *WorkbookRange) GetRowCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowCount
    }
}
func (m *WorkbookRange) GetRowHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.rowHidden
    }
}
func (m *WorkbookRange) GetRowIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowIndex
    }
}
func (m *WorkbookRange) GetSort()(*WorkbookRangeSort) {
    if m == nil {
        return nil
    } else {
        return m.sort
    }
}
func (m *WorkbookRange) GetText()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *WorkbookRange) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *WorkbookRange) GetValueTypes()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.valueTypes
    }
}
func (m *WorkbookRange) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
func (m *WorkbookRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddress(val)
        return nil
    }
    res["addressLocal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddressLocal(val)
        return nil
    }
    res["cellCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCellCount(val)
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
    res["columnHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetColumnHidden(val)
        return nil
    }
    res["columnIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetColumnIndex(val)
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookRangeFormat))
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
    res["hidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHidden(val)
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
    res["rowHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRowHidden(val)
        return nil
    }
    res["rowIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRowIndex(val)
        return nil
    }
    res["sort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeSort() })
        if err != nil {
            return err
        }
        m.SetSort(val.(*WorkbookRangeSort))
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
func (m *WorkbookRange) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("addressLocal", m.GetAddressLocal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cellCount", m.GetCellCount())
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
        err = writer.WriteBoolValue("columnHidden", m.GetColumnHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("columnIndex", m.GetColumnIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
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
        err = writer.WriteBoolValue("hidden", m.GetHidden())
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
        err = writer.WriteBoolValue("rowHidden", m.GetRowHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rowIndex", m.GetRowIndex())
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
    {
        err = writer.WriteObjectValue("worksheet", m.GetWorksheet())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookRange) SetAddress(value *string)() {
    m.address = value
}
func (m *WorkbookRange) SetAddressLocal(value *string)() {
    m.addressLocal = value
}
func (m *WorkbookRange) SetCellCount(value *int32)() {
    m.cellCount = value
}
func (m *WorkbookRange) SetColumnCount(value *int32)() {
    m.columnCount = value
}
func (m *WorkbookRange) SetColumnHidden(value *bool)() {
    m.columnHidden = value
}
func (m *WorkbookRange) SetColumnIndex(value *int32)() {
    m.columnIndex = value
}
func (m *WorkbookRange) SetFormat(value *WorkbookRangeFormat)() {
    m.format = value
}
func (m *WorkbookRange) SetFormulas(value *Json)() {
    m.formulas = value
}
func (m *WorkbookRange) SetFormulasLocal(value *Json)() {
    m.formulasLocal = value
}
func (m *WorkbookRange) SetFormulasR1C1(value *Json)() {
    m.formulasR1C1 = value
}
func (m *WorkbookRange) SetHidden(value *bool)() {
    m.hidden = value
}
func (m *WorkbookRange) SetNumberFormat(value *Json)() {
    m.numberFormat = value
}
func (m *WorkbookRange) SetRowCount(value *int32)() {
    m.rowCount = value
}
func (m *WorkbookRange) SetRowHidden(value *bool)() {
    m.rowHidden = value
}
func (m *WorkbookRange) SetRowIndex(value *int32)() {
    m.rowIndex = value
}
func (m *WorkbookRange) SetSort(value *WorkbookRangeSort)() {
    m.sort = value
}
func (m *WorkbookRange) SetText(value *Json)() {
    m.text = value
}
func (m *WorkbookRange) SetValues(value *Json)() {
    m.values = value
}
func (m *WorkbookRange) SetValueTypes(value *Json)() {
    m.valueTypes = value
}
func (m *WorkbookRange) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
