package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxisTitleable 
type WorkbookChartAxisTitleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFormat()(WorkbookChartAxisTitleFormatable)
    GetText()(*string)
    GetVisible()(*bool)
    SetFormat(value WorkbookChartAxisTitleFormatable)()
    SetText(value *string)()
    SetVisible(value *bool)()
}
