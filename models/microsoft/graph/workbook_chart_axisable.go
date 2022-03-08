package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxisable 
type WorkbookChartAxisable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFormat()(WorkbookChartAxisFormatable)
    GetMajorGridlines()(WorkbookChartGridlinesable)
    GetMajorUnit()(Jsonable)
    GetMaximum()(Jsonable)
    GetMinimum()(Jsonable)
    GetMinorGridlines()(WorkbookChartGridlinesable)
    GetMinorUnit()(Jsonable)
    GetTitle()(WorkbookChartAxisTitleable)
    SetFormat(value WorkbookChartAxisFormatable)()
    SetMajorGridlines(value WorkbookChartGridlinesable)()
    SetMajorUnit(value Jsonable)()
    SetMaximum(value Jsonable)()
    SetMinimum(value Jsonable)()
    SetMinorGridlines(value WorkbookChartGridlinesable)()
    SetMinorUnit(value Jsonable)()
    SetTitle(value WorkbookChartAxisTitleable)()
}
