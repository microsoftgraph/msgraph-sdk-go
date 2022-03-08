package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookWorksheetable 
type WorkbookWorksheetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCharts()([]WorkbookChartable)
    GetName()(*string)
    GetNames()([]WorkbookNamedItemable)
    GetPivotTables()([]WorkbookPivotTableable)
    GetPosition()(*int32)
    GetProtection()(WorkbookWorksheetProtectionable)
    GetTables()([]WorkbookTableable)
    GetVisibility()(*string)
    SetCharts(value []WorkbookChartable)()
    SetName(value *string)()
    SetNames(value []WorkbookNamedItemable)()
    SetPivotTables(value []WorkbookPivotTableable)()
    SetPosition(value *int32)()
    SetProtection(value WorkbookWorksheetProtectionable)()
    SetTables(value []WorkbookTableable)()
    SetVisibility(value *string)()
}
