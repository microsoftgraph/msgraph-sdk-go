package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrinterDefaults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The default color mode to use when printing the document. Valid values are described in the following table.
    colorMode *PrintColorMode;
    // The default content (MIME) type to use when processing documents.
    contentType *string;
    // The default number of copies printed per job.
    copiesPerJob *int32;
    // The default resolution in DPI to use when printing the job.
    dpi *int32;
    // The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
    duplexMode *PrintDuplexMode;
    // The default set of finishings to apply to print jobs. Valid values are described in the following table.
    finishings []PrintFinishing;
    // The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
    fitPdfToPage *bool;
    // The default input bin that serves as the paper source.
    inputBin *string;
    // The default media (such as paper) color to print the document on.
    mediaColor *string;
    // The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
    mediaSize *string;
    // The default media (such as paper) type to print the document on.
    mediaType *string;
    // The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
    multipageLayout *PrintMultipageLayout;
    // The default orientation to use when printing the document. Valid values are described in the following table.
    orientation *PrintOrientation;
    // The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
    outputBin *string;
    // The default number of document pages to print on each sheet.
    pagesPerSheet *int32;
    // The default quality to use when printing the document. Valid values are described in the following table.
    quality *PrintQuality;
    // Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
    scaling *PrintScaling;
}
// Instantiates a new printerDefaults and sets the default values.
func NewPrinterDefaults()(*PrinterDefaults) {
    m := &PrinterDefaults{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterDefaults) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the colorMode property value. The default color mode to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetColorMode()(*PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorMode
    }
}
// Gets the contentType property value. The default content (MIME) type to use when processing documents.
func (m *PrinterDefaults) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// Gets the copiesPerJob property value. The default number of copies printed per job.
func (m *PrinterDefaults) GetCopiesPerJob()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.copiesPerJob
    }
}
// Gets the dpi property value. The default resolution in DPI to use when printing the job.
func (m *PrinterDefaults) GetDpi()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dpi
    }
}
// Gets the duplexMode property value. The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
func (m *PrinterDefaults) GetDuplexMode()(*PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexMode
    }
}
// Gets the finishings property value. The default set of finishings to apply to print jobs. Valid values are described in the following table.
func (m *PrinterDefaults) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// Gets the fitPdfToPage property value. The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
func (m *PrinterDefaults) GetFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fitPdfToPage
    }
}
// Gets the inputBin property value. The default input bin that serves as the paper source.
func (m *PrinterDefaults) GetInputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inputBin
    }
}
// Gets the mediaColor property value. The default media (such as paper) color to print the document on.
func (m *PrinterDefaults) GetMediaColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaColor
    }
}
// Gets the mediaSize property value. The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
func (m *PrinterDefaults) GetMediaSize()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSize
    }
}
// Gets the mediaType property value. The default media (such as paper) type to print the document on.
func (m *PrinterDefaults) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// Gets the multipageLayout property value. The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
func (m *PrinterDefaults) GetMultipageLayout()(*PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayout
    }
}
// Gets the orientation property value. The default orientation to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetOrientation()(*PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// Gets the outputBin property value. The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
func (m *PrinterDefaults) GetOutputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputBin
    }
}
// Gets the pagesPerSheet property value. The default number of document pages to print on each sheet.
func (m *PrinterDefaults) GetPagesPerSheet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// Gets the quality property value. The default quality to use when printing the document. Valid values are described in the following table.
func (m *PrinterDefaults) GetQuality()(*PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// Gets the scaling property value. Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
func (m *PrinterDefaults) GetScaling()(*PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scaling
    }
}
// The deserialization information for the current model
func (m *PrinterDefaults) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["colorMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintColorMode)
            m.SetColorMode(&cast)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["copiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPerJob(val)
        }
        return nil
    }
    res["dpi"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDpi(val)
        }
        return nil
    }
    res["duplexMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintDuplexMode)
            m.SetDuplexMode(&cast)
        }
        return nil
    }
    res["finishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintFinishing, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintFinishing))
            }
            m.SetFinishings(res)
        }
        return nil
    }
    res["fitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFitPdfToPage(val)
        }
        return nil
    }
    res["inputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputBin(val)
        }
        return nil
    }
    res["mediaColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaColor(val)
        }
        return nil
    }
    res["mediaSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSize(val)
        }
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["multipageLayout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintMultipageLayout)
            m.SetMultipageLayout(&cast)
        }
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintOrientation)
            m.SetOrientation(&cast)
        }
        return nil
    }
    res["outputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputBin(val)
        }
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPagesPerSheet(val)
        }
        return nil
    }
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintQuality)
            m.SetQuality(&cast)
        }
        return nil
    }
    res["scaling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintScaling)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PrintScaling)
            m.SetScaling(&cast)
        }
        return nil
    }
    return res
}
func (m *PrinterDefaults) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrinterDefaults) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetColorMode() != nil {
        cast := m.GetColorMode().String()
        err := writer.WriteStringValue("colorMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copiesPerJob", m.GetCopiesPerJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dpi", m.GetDpi())
        if err != nil {
            return err
        }
    }
    if m.GetDuplexMode() != nil {
        cast := m.GetDuplexMode().String()
        err := writer.WriteStringValue("duplexMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("finishings", SerializePrintFinishing(m.GetFinishings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fitPdfToPage", m.GetFitPdfToPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("inputBin", m.GetInputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaColor", m.GetMediaColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaSize", m.GetMediaSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    if m.GetMultipageLayout() != nil {
        cast := m.GetMultipageLayout().String()
        err := writer.WriteStringValue("multipageLayout", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrientation() != nil {
        cast := m.GetOrientation().String()
        err := writer.WriteStringValue("orientation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputBin", m.GetOutputBin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    if m.GetQuality() != nil {
        cast := m.GetQuality().String()
        err := writer.WriteStringValue("quality", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScaling() != nil {
        cast := m.GetScaling().String()
        err := writer.WriteStringValue("scaling", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PrinterDefaults) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the colorMode property value. The default color mode to use when printing the document. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the colorMode property.
func (m *PrinterDefaults) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
// Sets the contentType property value. The default content (MIME) type to use when processing documents.
// Parameters:
//  - value : Value to set for the contentType property.
func (m *PrinterDefaults) SetContentType(value *string)() {
    m.contentType = value
}
// Sets the copiesPerJob property value. The default number of copies printed per job.
// Parameters:
//  - value : Value to set for the copiesPerJob property.
func (m *PrinterDefaults) SetCopiesPerJob(value *int32)() {
    m.copiesPerJob = value
}
// Sets the dpi property value. The default resolution in DPI to use when printing the job.
// Parameters:
//  - value : Value to set for the dpi property.
func (m *PrinterDefaults) SetDpi(value *int32)() {
    m.dpi = value
}
// Sets the duplexMode property value. The default duplex (double-sided) configuration to use when printing a document. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the duplexMode property.
func (m *PrinterDefaults) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
// Sets the finishings property value. The default set of finishings to apply to print jobs. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the finishings property.
func (m *PrinterDefaults) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// Sets the fitPdfToPage property value. The default fitPdfToPage setting. True to fit each page of a PDF document to a physical sheet of media; false to let the printer decide how to lay out impressions.
// Parameters:
//  - value : Value to set for the fitPdfToPage property.
func (m *PrinterDefaults) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
// Sets the inputBin property value. The default input bin that serves as the paper source.
// Parameters:
//  - value : Value to set for the inputBin property.
func (m *PrinterDefaults) SetInputBin(value *string)() {
    m.inputBin = value
}
// Sets the mediaColor property value. The default media (such as paper) color to print the document on.
// Parameters:
//  - value : Value to set for the mediaColor property.
func (m *PrinterDefaults) SetMediaColor(value *string)() {
    m.mediaColor = value
}
// Sets the mediaSize property value. The default media size to use. Supports standard size names for ISO and ANSI media sizes. Valid values are listed in the printerCapabilities topic.
// Parameters:
//  - value : Value to set for the mediaSize property.
func (m *PrinterDefaults) SetMediaSize(value *string)() {
    m.mediaSize = value
}
// Sets the mediaType property value. The default media (such as paper) type to print the document on.
// Parameters:
//  - value : Value to set for the mediaType property.
func (m *PrinterDefaults) SetMediaType(value *string)() {
    m.mediaType = value
}
// Sets the multipageLayout property value. The default direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the multipageLayout property.
func (m *PrinterDefaults) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
// Sets the orientation property value. The default orientation to use when printing the document. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the orientation property.
func (m *PrinterDefaults) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
// Sets the outputBin property value. The default output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
// Parameters:
//  - value : Value to set for the outputBin property.
func (m *PrinterDefaults) SetOutputBin(value *string)() {
    m.outputBin = value
}
// Sets the pagesPerSheet property value. The default number of document pages to print on each sheet.
// Parameters:
//  - value : Value to set for the pagesPerSheet property.
func (m *PrinterDefaults) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
// Sets the quality property value. The default quality to use when printing the document. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the quality property.
func (m *PrinterDefaults) SetQuality(value *PrintQuality)() {
    m.quality = value
}
// Sets the scaling property value. Specifies how the printer scales the document data to fit the requested media. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the scaling property.
func (m *PrinterDefaults) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
