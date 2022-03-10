package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterDefaultsable 
type PrinterDefaultsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetColorMode()(*PrintColorMode)
    GetContentType()(*string)
    GetCopiesPerJob()(*int32)
    GetDpi()(*int32)
    GetDuplexMode()(*PrintDuplexMode)
    GetFinishings()([]PrintFinishing)
    GetFitPdfToPage()(*bool)
    GetInputBin()(*string)
    GetMediaColor()(*string)
    GetMediaSize()(*string)
    GetMediaType()(*string)
    GetMultipageLayout()(*PrintMultipageLayout)
    GetOrientation()(*PrintOrientation)
    GetOutputBin()(*string)
    GetPagesPerSheet()(*int32)
    GetQuality()(*PrintQuality)
    GetScaling()(*PrintScaling)
    SetColorMode(value *PrintColorMode)()
    SetContentType(value *string)()
    SetCopiesPerJob(value *int32)()
    SetDpi(value *int32)()
    SetDuplexMode(value *PrintDuplexMode)()
    SetFinishings(value []PrintFinishing)()
    SetFitPdfToPage(value *bool)()
    SetInputBin(value *string)()
    SetMediaColor(value *string)()
    SetMediaSize(value *string)()
    SetMediaType(value *string)()
    SetMultipageLayout(value *PrintMultipageLayout)()
    SetOrientation(value *PrintOrientation)()
    SetOutputBin(value *string)()
    SetPagesPerSheet(value *int32)()
    SetQuality(value *PrintQuality)()
    SetScaling(value *PrintScaling)()
}
