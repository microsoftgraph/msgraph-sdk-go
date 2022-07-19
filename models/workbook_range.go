package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRange 
type WorkbookRange struct {
    Entity
    // Represents the range reference in A1-style. Address value will contain the Sheet reference (e.g. Sheet1!A1:B4). Read-only.
    address *string
    // Represents range reference for the specified range in the language of the user. Read-only.
    addressLocal *string
    // Number of cells in the range. Read-only.
    cellCount *int32
    // Represents the total number of columns in the range. Read-only.
    columnCount *int32
    // Represents if all columns of the current range are hidden.
    columnHidden *bool
    // Represents the column number of the first cell in the range. Zero-indexed. Read-only.
    columnIndex *int32
    // Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
    format WorkbookRangeFormatable
    // Represents the formula in A1-style notation.
    formulas Jsonable
    // Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
    formulasLocal Jsonable
    // Represents the formula in R1C1-style notation.
    formulasR1C1 Jsonable
    // Represents if all cells of the current range are hidden. Read-only.
    hidden *bool
    // Represents Excel's number format code for the given cell.
    numberFormat Jsonable
    // Returns the total number of rows in the range. Read-only.
    rowCount *int32
    // Represents if all rows of the current range are hidden.
    rowHidden *bool
    // Returns the row number of the first cell in the range. Zero-indexed. Read-only.
    rowIndex *int32
    // The worksheet containing the current range. Read-only.
    sort WorkbookRangeSortable
    // Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
    text Jsonable
    // Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
    values Jsonable
    // Represents the type of data of each cell. Possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
    valueTypes Jsonable
    // The worksheet containing the current range. Read-only.
    worksheet WorkbookWorksheetable
}
// NewWorkbookRange instantiates a new WorkbookRange and sets the default values.
func NewWorkbookRange()(*WorkbookRange) {
    m := &WorkbookRange{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.workbookRange";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkbookRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRange(), nil
}
// GetAddress gets the address property value. Represents the range reference in A1-style. Address value will contain the Sheet reference (e.g. Sheet1!A1:B4). Read-only.
func (m *WorkbookRange) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetAddressLocal gets the addressLocal property value. Represents range reference for the specified range in the language of the user. Read-only.
func (m *WorkbookRange) GetAddressLocal()(*string) {
    if m == nil {
        return nil
    } else {
        return m.addressLocal
    }
}
// GetCellCount gets the cellCount property value. Number of cells in the range. Read-only.
func (m *WorkbookRange) GetCellCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.cellCount
    }
}
// GetColumnCount gets the columnCount property value. Represents the total number of columns in the range. Read-only.
func (m *WorkbookRange) GetColumnCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.columnCount
    }
}
// GetColumnHidden gets the columnHidden property value. Represents if all columns of the current range are hidden.
func (m *WorkbookRange) GetColumnHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.columnHidden
    }
}
// GetColumnIndex gets the columnIndex property value. Represents the column number of the first cell in the range. Zero-indexed. Read-only.
func (m *WorkbookRange) GetColumnIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.columnIndex
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["addressLocal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressLocal(val)
        }
        return nil
    }
    res["cellCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellCount(val)
        }
        return nil
    }
    res["columnCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnCount(val)
        }
        return nil
    }
    res["columnHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnHidden(val)
        }
        return nil
    }
    res["columnIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnIndex(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookRangeFormatable))
        }
        return nil
    }
    res["formulas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulas(val.(Jsonable))
        }
        return nil
    }
    res["formulasLocal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasLocal(val.(Jsonable))
        }
        return nil
    }
    res["formulasR1C1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormulasR1C1(val.(Jsonable))
        }
        return nil
    }
    res["hidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidden(val)
        }
        return nil
    }
    res["numberFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberFormat(val.(Jsonable))
        }
        return nil
    }
    res["rowCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowCount(val)
        }
        return nil
    }
    res["rowHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowHidden(val)
        }
        return nil
    }
    res["rowIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRowIndex(val)
        }
        return nil
    }
    res["sort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookRangeSortFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSort(val.(WorkbookRangeSortable))
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(Jsonable))
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    res["valueTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueTypes(val.(Jsonable))
        }
        return nil
    }
    res["worksheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookWorksheetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorksheet(val.(WorkbookWorksheetable))
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
func (m *WorkbookRange) GetFormat()(WorkbookRangeFormatable) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetFormulas gets the formulas property value. Represents the formula in A1-style notation.
func (m *WorkbookRange) GetFormulas()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.formulas
    }
}
// GetFormulasLocal gets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRange) GetFormulasLocal()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.formulasLocal
    }
}
// GetFormulasR1C1 gets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRange) GetFormulasR1C1()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.formulasR1C1
    }
}
// GetHidden gets the hidden property value. Represents if all cells of the current range are hidden. Read-only.
func (m *WorkbookRange) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// GetNumberFormat gets the numberFormat property value. Represents Excel's number format code for the given cell.
func (m *WorkbookRange) GetNumberFormat()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.numberFormat
    }
}
// GetRowCount gets the rowCount property value. Returns the total number of rows in the range. Read-only.
func (m *WorkbookRange) GetRowCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowCount
    }
}
// GetRowHidden gets the rowHidden property value. Represents if all rows of the current range are hidden.
func (m *WorkbookRange) GetRowHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.rowHidden
    }
}
// GetRowIndex gets the rowIndex property value. Returns the row number of the first cell in the range. Zero-indexed. Read-only.
func (m *WorkbookRange) GetRowIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rowIndex
    }
}
// GetSort gets the sort property value. The worksheet containing the current range. Read-only.
func (m *WorkbookRange) GetSort()(WorkbookRangeSortable) {
    if m == nil {
        return nil
    } else {
        return m.sort
    }
}
// GetText gets the text property value. Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
func (m *WorkbookRange) GetText()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetValues gets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookRange) GetValues()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetValueTypes gets the valueTypes property value. Represents the type of data of each cell. Possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
func (m *WorkbookRange) GetValueTypes()(Jsonable) {
    if m == nil {
        return nil
    } else {
        return m.valueTypes
    }
}
// GetWorksheet gets the worksheet property value. The worksheet containing the current range. Read-only.
func (m *WorkbookRange) GetWorksheet()(WorkbookWorksheetable) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// Serialize serializes information the current object
func (m *WorkbookRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAddress sets the address property value. Represents the range reference in A1-style. Address value will contain the Sheet reference (e.g. Sheet1!A1:B4). Read-only.
func (m *WorkbookRange) SetAddress(value *string)() {
    if m != nil {
        m.address = value
    }
}
// SetAddressLocal sets the addressLocal property value. Represents range reference for the specified range in the language of the user. Read-only.
func (m *WorkbookRange) SetAddressLocal(value *string)() {
    if m != nil {
        m.addressLocal = value
    }
}
// SetCellCount sets the cellCount property value. Number of cells in the range. Read-only.
func (m *WorkbookRange) SetCellCount(value *int32)() {
    if m != nil {
        m.cellCount = value
    }
}
// SetColumnCount sets the columnCount property value. Represents the total number of columns in the range. Read-only.
func (m *WorkbookRange) SetColumnCount(value *int32)() {
    if m != nil {
        m.columnCount = value
    }
}
// SetColumnHidden sets the columnHidden property value. Represents if all columns of the current range are hidden.
func (m *WorkbookRange) SetColumnHidden(value *bool)() {
    if m != nil {
        m.columnHidden = value
    }
}
// SetColumnIndex sets the columnIndex property value. Represents the column number of the first cell in the range. Zero-indexed. Read-only.
func (m *WorkbookRange) SetColumnIndex(value *int32)() {
    if m != nil {
        m.columnIndex = value
    }
}
// SetFormat sets the format property value. Returns a format object, encapsulating the range's font, fill, borders, alignment, and other properties. Read-only.
func (m *WorkbookRange) SetFormat(value WorkbookRangeFormatable)() {
    if m != nil {
        m.format = value
    }
}
// SetFormulas sets the formulas property value. Represents the formula in A1-style notation.
func (m *WorkbookRange) SetFormulas(value Jsonable)() {
    if m != nil {
        m.formulas = value
    }
}
// SetFormulasLocal sets the formulasLocal property value. Represents the formula in A1-style notation, in the user's language and number-formatting locale.  For example, the English '=SUM(A1, 1.5)' formula would become '=SUMME(A1; 1,5)' in German.
func (m *WorkbookRange) SetFormulasLocal(value Jsonable)() {
    if m != nil {
        m.formulasLocal = value
    }
}
// SetFormulasR1C1 sets the formulasR1C1 property value. Represents the formula in R1C1-style notation.
func (m *WorkbookRange) SetFormulasR1C1(value Jsonable)() {
    if m != nil {
        m.formulasR1C1 = value
    }
}
// SetHidden sets the hidden property value. Represents if all cells of the current range are hidden. Read-only.
func (m *WorkbookRange) SetHidden(value *bool)() {
    if m != nil {
        m.hidden = value
    }
}
// SetNumberFormat sets the numberFormat property value. Represents Excel's number format code for the given cell.
func (m *WorkbookRange) SetNumberFormat(value Jsonable)() {
    if m != nil {
        m.numberFormat = value
    }
}
// SetRowCount sets the rowCount property value. Returns the total number of rows in the range. Read-only.
func (m *WorkbookRange) SetRowCount(value *int32)() {
    if m != nil {
        m.rowCount = value
    }
}
// SetRowHidden sets the rowHidden property value. Represents if all rows of the current range are hidden.
func (m *WorkbookRange) SetRowHidden(value *bool)() {
    if m != nil {
        m.rowHidden = value
    }
}
// SetRowIndex sets the rowIndex property value. Returns the row number of the first cell in the range. Zero-indexed. Read-only.
func (m *WorkbookRange) SetRowIndex(value *int32)() {
    if m != nil {
        m.rowIndex = value
    }
}
// SetSort sets the sort property value. The worksheet containing the current range. Read-only.
func (m *WorkbookRange) SetSort(value WorkbookRangeSortable)() {
    if m != nil {
        m.sort = value
    }
}
// SetText sets the text property value. Text values of the specified range. The Text value will not depend on the cell width. The # sign substitution that happens in Excel UI will not affect the text value returned by the API. Read-only.
func (m *WorkbookRange) SetText(value Jsonable)() {
    if m != nil {
        m.text = value
    }
}
// SetValues sets the values property value. Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
func (m *WorkbookRange) SetValues(value Jsonable)() {
    if m != nil {
        m.values = value
    }
}
// SetValueTypes sets the valueTypes property value. Represents the type of data of each cell. Possible values are: Unknown, Empty, String, Integer, Double, Boolean, Error. Read-only.
func (m *WorkbookRange) SetValueTypes(value Jsonable)() {
    if m != nil {
        m.valueTypes = value
    }
}
// SetWorksheet sets the worksheet property value. The worksheet containing the current range. Read-only.
func (m *WorkbookRange) SetWorksheet(value WorkbookWorksheetable)() {
    if m != nil {
        m.worksheet = value
    }
}
