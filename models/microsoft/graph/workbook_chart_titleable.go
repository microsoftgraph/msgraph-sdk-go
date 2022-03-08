package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartTitleable 
type WorkbookChartTitleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetFormat()(WorkbookChartTitleFormatable)
    GetOverlay()(*bool)
    GetText()(*string)
    GetVisible()(*bool)
    SetFormat(value WorkbookChartTitleFormatable)()
    SetOverlay(value *bool)()
    SetText(value *string)()
    SetVisible(value *bool)()
}
