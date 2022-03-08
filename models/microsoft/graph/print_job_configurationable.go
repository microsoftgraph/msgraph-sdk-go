package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintJobConfigurationable 
type PrintJobConfigurationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCollate()(*bool)
    GetColorMode()(*PrintColorMode)
    GetCopies()(*int32)
    GetDpi()(*int32)
    GetDuplexMode()(*PrintDuplexMode)
    GetFeedOrientation()(*PrinterFeedOrientation)
    GetFinishings()([]PrintFinishing)
    GetFitPdfToPage()(*bool)
    GetInputBin()(*string)
    GetMargin()(PrintMarginable)
    GetMediaSize()(*string)
    GetMediaType()(*string)
    GetMultipageLayout()(*PrintMultipageLayout)
    GetOrientation()(*PrintOrientation)
    GetOutputBin()(*string)
    GetPageRanges()([]IntegerRangeable)
    GetPagesPerSheet()(*int32)
    GetQuality()(*PrintQuality)
    GetScaling()(*PrintScaling)
    SetCollate(value *bool)()
    SetColorMode(value *PrintColorMode)()
    SetCopies(value *int32)()
    SetDpi(value *int32)()
    SetDuplexMode(value *PrintDuplexMode)()
    SetFeedOrientation(value *PrinterFeedOrientation)()
    SetFinishings(value []PrintFinishing)()
    SetFitPdfToPage(value *bool)()
    SetInputBin(value *string)()
    SetMargin(value PrintMarginable)()
    SetMediaSize(value *string)()
    SetMediaType(value *string)()
    SetMultipageLayout(value *PrintMultipageLayout)()
    SetOrientation(value *PrintOrientation)()
    SetOutputBin(value *string)()
    SetPageRanges(value []IntegerRangeable)()
    SetPagesPerSheet(value *int32)()
    SetQuality(value *PrintQuality)()
    SetScaling(value *PrintScaling)()
}
