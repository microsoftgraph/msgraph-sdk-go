package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookTableable 
type WorkbookTableable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetColumns()([]WorkbookTableColumnable)
    GetHighlightFirstColumn()(*bool)
    GetHighlightLastColumn()(*bool)
    GetLegacyId()(*string)
    GetName()(*string)
    GetRows()([]WorkbookTableRowable)
    GetShowBandedColumns()(*bool)
    GetShowBandedRows()(*bool)
    GetShowFilterButton()(*bool)
    GetShowHeaders()(*bool)
    GetShowTotals()(*bool)
    GetSort()(WorkbookTableSortable)
    GetStyle()(*string)
    GetWorksheet()(WorkbookWorksheetable)
    SetColumns(value []WorkbookTableColumnable)()
    SetHighlightFirstColumn(value *bool)()
    SetHighlightLastColumn(value *bool)()
    SetLegacyId(value *string)()
    SetName(value *string)()
    SetRows(value []WorkbookTableRowable)()
    SetShowBandedColumns(value *bool)()
    SetShowBandedRows(value *bool)()
    SetShowFilterButton(value *bool)()
    SetShowHeaders(value *bool)()
    SetShowTotals(value *bool)()
    SetSort(value WorkbookTableSortable)()
    SetStyle(value *string)()
    SetWorksheet(value WorkbookWorksheetable)()
}
