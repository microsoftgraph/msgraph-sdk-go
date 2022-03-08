package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ReportRootable 
type ReportRootable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDailyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetDailyPrintUsageByUser()([]PrintUsageByUserable)
    GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinterable)
    GetMonthlyPrintUsageByUser()([]PrintUsageByUserable)
    SetDailyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetDailyPrintUsageByUser(value []PrintUsageByUserable)()
    SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinterable)()
    SetMonthlyPrintUsageByUser(value []PrintUsageByUserable)()
}
