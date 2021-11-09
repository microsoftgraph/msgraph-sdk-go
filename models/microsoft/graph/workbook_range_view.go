package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookRangeView struct {
    Entity
    // Represents the cell addresses
    cellAddresses *Json;
    // Returns the number of visible columns. Read-only.
    columnCount *int32;
    // Represents the formula in A1-style notation.
    formulas *Json;
    // Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
    formulasLocal *Json;
    // Represents the formula in R1C1-style notation.
    formulasR1C1 *Json;
    // Index of the range.
    index *int32;
    // Represents Excel's number format code for the given cell. Read-only.
    numberFormat *Json;
    // Returns the number of visible rows. Read-only.
    rowCount *int32;
    // Represents a collection of range views associated with the range. Read-only. Read-only.
    rows []WorkbookRangeView;
    // Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
    text *Json;
    // Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values *Json;
    // Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
    valueTypes *Json;
}
// Instantiates a new workbookRangeView and sets the default values.
func NewWorkbookRangeView()(*WorkbookRangeView) {
    m := &WorkbookRangeView{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the cellAddresses property value. Represents the cell addresses
func (m *WorkbookRangeView) GetCellAddresses()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.cellAddresses
    }
}
// Gets the columnCount property value. Returns the number of visible columns. Read-only.
func (m *WorkbookRangeView) GetColumnCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.columnCount
    }
}
// Gets the formulas property value. Represents the formula in A1-style notation.
func (m *WorkbookRangeView) GetFormulas()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulas
    }
}
// Gets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRangeView) GetFormulasLocal()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulasLocal
    }
}
// Gets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRangeView) GetFormulasR1C1()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.formulasR1C1
    }
}
// Gets the index property value. Index of the range.
func (m *WorkbookRangeView) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// Gets the numberFormat property value. Represents Excel's number format code for the given cell. Read-only.
func (m *WorkbookRangeView) GetNumberFormat()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.numberFormat
    }
}
// Gets the rowCount property value. Returns the number of visible rows. Read-only.
func (m *WorkbookRangeView) GetRowCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowCount
    }
}
// Gets the rows property value. Represents a collection of range views associated with the range. Read-only. Read-only.
func (m *WorkbookRangeView) GetRows()([]WorkbookRangeView) {
    if m == nil {
        return nil
    } else {
        return m.rows
    }
}
// Gets the text property value. Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
func (m *WorkbookRangeView) GetText()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Gets the values property value. Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookRangeView) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// Gets the valueTypes property value. Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
func (m *WorkbookRangeView) GetValueTypes()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.valueTypes
    }
}
// The deserialization information for the current model
func (m *WorkbookRangeView) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cellAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellAddresses(val.(*Json))
        }
        return nil
    }
    res["columnCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnCount(val)
        }
        return nil
    }
    res["formulas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulas(val.(*Json))
        }
        return nil
    }
    res["formulasLocal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasLocal(val.(*Json))
        }
        return nil
    }
    res["formulasR1C1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasR1C1(val.(*Json))
        }
        return nil
    }
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["numberFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberFormat(val.(*Json))
        }
        return nil
    }
    res["rowCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowCount(val)
        }
        return nil
    }
    res["rows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookRangeView() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookRangeView, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookRangeView))
            }
            m.SetRows(res)
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(*Json))
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(*Json))
        }
        return nil
    }
    res["valueTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueTypes(val.(*Json))
        }
        return nil
    }
    return res
}
func (m *WorkbookRangeView) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the cellAddresses property value. Represents the cell addresses
// Parameters:
//  - value : Value to set for the cellAddresses property.
func (m *WorkbookRangeView) SetCellAddresses(value *Json)() {
    m.cellAddresses = value
}
// Sets the columnCount property value. Returns the number of visible columns. Read-only.
// Parameters:
//  - value : Value to set for the columnCount property.
func (m *WorkbookRangeView) SetColumnCount(value *int32)() {
    m.columnCount = value
}
// Sets the formulas property value. Represents the formula in A1-style notation.
// Parameters:
//  - value : Value to set for the formulas property.
func (m *WorkbookRangeView) SetFormulas(value *Json)() {
    m.formulas = value
}
// Sets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale. For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
// Parameters:
//  - value : Value to set for the formulasLocal property.
func (m *WorkbookRangeView) SetFormulasLocal(value *Json)() {
    m.formulasLocal = value
}
// Sets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
// Parameters:
//  - value : Value to set for the formulasR1C1 property.
func (m *WorkbookRangeView) SetFormulasR1C1(value *Json)() {
    m.formulasR1C1 = value
}
// Sets the index property value. Index of the range.
// Parameters:
//  - value : Value to set for the index property.
func (m *WorkbookRangeView) SetIndex(value *int32)() {
    m.index = value
}
// Sets the numberFormat property value. Represents Excel's number format code for the given cell. Read-only.
// Parameters:
//  - value : Value to set for the numberFormat property.
func (m *WorkbookRangeView) SetNumberFormat(value *Json)() {
    m.numberFormat = value
}
// Sets the rowCount property value. Returns the number of visible rows. Read-only.
// Parameters:
//  - value : Value to set for the rowCount property.
func (m *WorkbookRangeView) SetRowCount(value *int32)() {
    m.rowCount = value
}
// Sets the rows property value. Represents a collection of range views associated with the range. Read-only. Read-only.
// Parameters:
//  - value : Value to set for the rows property.
func (m *WorkbookRangeView) SetRows(value []WorkbookRangeView)() {
    m.rows = value
}
// Sets the text property value. Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
// Parameters:
//  - value : Value to set for the text property.
func (m *WorkbookRangeView) SetText(value *Json)() {
    m.text = value
}
// Sets the values property value. Represents the raw values of the specified range view. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
// Parameters:
//  - value : Value to set for the values property.
func (m *WorkbookRangeView) SetValues(value *Json)() {
    m.values = value
}
// Sets the valueTypes property value. Represents the type of data of each cell. Read-only. The possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error.
// Parameters:
//  - value : Value to set for the valueTypes property.
func (m *WorkbookRangeView) SetValueTypes(value *Json)() {
    m.valueTypes = value
}
