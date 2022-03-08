package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookWorksheetProtectionOptionsable 
type WorkbookWorksheetProtectionOptionsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAllowAutoFilter()(*bool)
    GetAllowDeleteColumns()(*bool)
    GetAllowDeleteRows()(*bool)
    GetAllowFormatCells()(*bool)
    GetAllowFormatColumns()(*bool)
    GetAllowFormatRows()(*bool)
    GetAllowInsertColumns()(*bool)
    GetAllowInsertHyperlinks()(*bool)
    GetAllowInsertRows()(*bool)
    GetAllowPivotTables()(*bool)
    GetAllowSort()(*bool)
    SetAllowAutoFilter(value *bool)()
    SetAllowDeleteColumns(value *bool)()
    SetAllowDeleteRows(value *bool)()
    SetAllowFormatCells(value *bool)()
    SetAllowFormatColumns(value *bool)()
    SetAllowFormatRows(value *bool)()
    SetAllowInsertColumns(value *bool)()
    SetAllowInsertHyperlinks(value *bool)()
    SetAllowInsertRows(value *bool)()
    SetAllowPivotTables(value *bool)()
    SetAllowSort(value *bool)()
}
