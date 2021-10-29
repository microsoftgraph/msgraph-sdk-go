package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintJobConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
    collate *bool;
    // The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
    colorMode *PrintColorMode;
    // The number of copies that should be printed. Read-only.
    copies *int32;
    // The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
    dpi *int32;
    // The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
    duplexMode *PrintDuplexMode;
    // The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
    feedOrientation *PrinterFeedOrientation;
    // Finishing processes to use when printing.
    finishings []PrintFinishing;
    // 
    fitPdfToPage *bool;
    // The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
    inputBin *string;
    // The margin settings to use when printing.
    margin *PrintMargin;
    // The media size to use when printing. Supports standard size names for ISO and ANSI media sizes. Valid values listed in the printerCapabilities topic.
    mediaSize *string;
    // The default media (such as paper) type to print the document on.
    mediaType *string;
    // The direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
    multipageLayout *PrintMultipageLayout;
    // The orientation setting the printer should use when printing the job. Valid values are described in the following table.
    orientation *PrintOrientation;
    // The output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
    outputBin *string;
    // The page ranges to print. Read-only.
    pageRanges []IntegerRange;
    // The number of document pages to print on each sheet.
    pagesPerSheet *int32;
    // The print quality to use when printing the job. Valid values are described in the table below. Read-only.
    quality *PrintQuality;
    // Specifies how the printer should scale the document data to fit the requested media. Valid values are described in the following table.
    scaling *PrintScaling;
}
// Instantiates a new printJobConfiguration and sets the default values.
func NewPrintJobConfiguration()(*PrintJobConfiguration) {
    m := &PrintJobConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintJobConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the collate property value. Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
func (m *PrintJobConfiguration) GetCollate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collate
    }
}
// Gets the colorMode property value. The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) GetColorMode()(*PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorMode
    }
}
// Gets the copies property value. The number of copies that should be printed. Read-only.
func (m *PrintJobConfiguration) GetCopies()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.copies
    }
}
// Gets the dpi property value. The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
func (m *PrintJobConfiguration) GetDpi()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dpi
    }
}
// Gets the duplexMode property value. The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) GetDuplexMode()(*PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexMode
    }
}
// Gets the feedOrientation property value. The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
func (m *PrintJobConfiguration) GetFeedOrientation()(*PrinterFeedOrientation) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientation
    }
}
// Gets the finishings property value. Finishing processes to use when printing.
func (m *PrintJobConfiguration) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// Gets the fitPdfToPage property value. 
func (m *PrintJobConfiguration) GetFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fitPdfToPage
    }
}
// Gets the inputBin property value. The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
func (m *PrintJobConfiguration) GetInputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inputBin
    }
}
// Gets the margin property value. The margin settings to use when printing.
func (m *PrintJobConfiguration) GetMargin()(*PrintMargin) {
    if m == nil {
        return nil
    } else {
        return m.margin
    }
}
// Gets the mediaSize property value. The media size to use when printing. Supports standard size names for ISO and ANSI media sizes. Valid values listed in the printerCapabilities topic.
func (m *PrintJobConfiguration) GetMediaSize()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSize
    }
}
// Gets the mediaType property value. The default media (such as paper) type to print the document on.
func (m *PrintJobConfiguration) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// Gets the multipageLayout property value. The direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
func (m *PrintJobConfiguration) GetMultipageLayout()(*PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayout
    }
}
// Gets the orientation property value. The orientation setting the printer should use when printing the job. Valid values are described in the following table.
func (m *PrintJobConfiguration) GetOrientation()(*PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
// Gets the outputBin property value. The output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
func (m *PrintJobConfiguration) GetOutputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputBin
    }
}
// Gets the pageRanges property value. The page ranges to print. Read-only.
func (m *PrintJobConfiguration) GetPageRanges()([]IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.pageRanges
    }
}
// Gets the pagesPerSheet property value. The number of document pages to print on each sheet.
func (m *PrintJobConfiguration) GetPagesPerSheet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// Gets the quality property value. The print quality to use when printing the job. Valid values are described in the table below. Read-only.
func (m *PrintJobConfiguration) GetQuality()(*PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
// Gets the scaling property value. Specifies how the printer should scale the document data to fit the requested media. Valid values are described in the following table.
func (m *PrintJobConfiguration) GetScaling()(*PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scaling
    }
}
// The deserialization information for the current model
func (m *PrintJobConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["collate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCollate(val)
        return nil
    }
    res["colorMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        cast := val.(PrintColorMode)
        m.SetColorMode(&cast)
        return nil
    }
    res["copies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCopies(val)
        return nil
    }
    res["dpi"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDpi(val)
        return nil
    }
    res["duplexMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        cast := val.(PrintDuplexMode)
        m.SetDuplexMode(&cast)
        return nil
    }
    res["feedOrientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        cast := val.(PrinterFeedOrientation)
        m.SetFeedOrientation(&cast)
        return nil
    }
    res["finishings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintFinishing)
        if err != nil {
            return err
        }
        res := make([]PrintFinishing, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintFinishing))
        }
        m.SetFinishings(res)
        return nil
    }
    res["fitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFitPdfToPage(val)
        return nil
    }
    res["inputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInputBin(val)
        return nil
    }
    res["margin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintMargin() })
        if err != nil {
            return err
        }
        m.SetMargin(val.(*PrintMargin))
        return nil
    }
    res["mediaSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaSize(val)
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaType(val)
        return nil
    }
    res["multipageLayout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        cast := val.(PrintMultipageLayout)
        m.SetMultipageLayout(&cast)
        return nil
    }
    res["orientation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOrientation)
        if err != nil {
            return err
        }
        cast := val.(PrintOrientation)
        m.SetOrientation(&cast)
        return nil
    }
    res["outputBin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOutputBin(val)
        return nil
    }
    res["pageRanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        res := make([]IntegerRange, len(val))
        for i, v := range val {
            res[i] = *(v.(*IntegerRange))
        }
        m.SetPageRanges(res)
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPagesPerSheet(val)
        return nil
    }
    res["quality"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintQuality)
        if err != nil {
            return err
        }
        cast := val.(PrintQuality)
        m.SetQuality(&cast)
        return nil
    }
    res["scaling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintScaling)
        if err != nil {
            return err
        }
        cast := val.(PrintScaling)
        m.SetScaling(&cast)
        return nil
    }
    return res
}
func (m *PrintJobConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintJobConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("collate", m.GetCollate())
        if err != nil {
            return err
        }
    }
    if m.GetColorMode() != nil {
        cast := m.GetColorMode().String()
        err := writer.WriteStringValue("colorMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("copies", m.GetCopies())
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
    if m.GetFeedOrientation() != nil {
        cast := m.GetFeedOrientation().String()
        err := writer.WriteStringValue("feedOrientation", &cast)
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
        err := writer.WriteObjectValue("margin", m.GetMargin())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPageRanges()))
        for i, v := range m.GetPageRanges() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("pageRanges", cast)
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
func (m *PrintJobConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the collate property value. Whether the printer should collate pages wehen printing multiple copies of a multi-page document.
// Parameters:
//  - value : Value to set for the collate property.
func (m *PrintJobConfiguration) SetCollate(value *bool)() {
    m.collate = value
}
// Sets the colorMode property value. The color mode the printer should use to print the job. Valid values are described in the table below. Read-only.
// Parameters:
//  - value : Value to set for the colorMode property.
func (m *PrintJobConfiguration) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
// Sets the copies property value. The number of copies that should be printed. Read-only.
// Parameters:
//  - value : Value to set for the copies property.
func (m *PrintJobConfiguration) SetCopies(value *int32)() {
    m.copies = value
}
// Sets the dpi property value. The resolution to use when printing the job, expressed in dots per inch (DPI). Read-only.
// Parameters:
//  - value : Value to set for the dpi property.
func (m *PrintJobConfiguration) SetDpi(value *int32)() {
    m.dpi = value
}
// Sets the duplexMode property value. The duplex mode the printer should use when printing the job. Valid values are described in the table below. Read-only.
// Parameters:
//  - value : Value to set for the duplexMode property.
func (m *PrintJobConfiguration) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
// Sets the feedOrientation property value. The orientation to use when feeding media into the printer. Valid values are described in the following table. Read-only.
// Parameters:
//  - value : Value to set for the feedOrientation property.
func (m *PrintJobConfiguration) SetFeedOrientation(value *PrinterFeedOrientation)() {
    m.feedOrientation = value
}
// Sets the finishings property value. Finishing processes to use when printing.
// Parameters:
//  - value : Value to set for the finishings property.
func (m *PrintJobConfiguration) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// Sets the fitPdfToPage property value. 
// Parameters:
//  - value : Value to set for the fitPdfToPage property.
func (m *PrintJobConfiguration) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
// Sets the inputBin property value. The input bin (tray) to use when printing. See the printer's capabilities for a list of supported input bins.
// Parameters:
//  - value : Value to set for the inputBin property.
func (m *PrintJobConfiguration) SetInputBin(value *string)() {
    m.inputBin = value
}
// Sets the margin property value. The margin settings to use when printing.
// Parameters:
//  - value : Value to set for the margin property.
func (m *PrintJobConfiguration) SetMargin(value *PrintMargin)() {
    m.margin = value
}
// Sets the mediaSize property value. The media size to use when printing. Supports standard size names for ISO and ANSI media sizes. Valid values listed in the printerCapabilities topic.
// Parameters:
//  - value : Value to set for the mediaSize property.
func (m *PrintJobConfiguration) SetMediaSize(value *string)() {
    m.mediaSize = value
}
// Sets the mediaType property value. The default media (such as paper) type to print the document on.
// Parameters:
//  - value : Value to set for the mediaType property.
func (m *PrintJobConfiguration) SetMediaType(value *string)() {
    m.mediaType = value
}
// Sets the multipageLayout property value. The direction to lay out pages when multiple pages are being printed per sheet. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the multipageLayout property.
func (m *PrintJobConfiguration) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
// Sets the orientation property value. The orientation setting the printer should use when printing the job. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the orientation property.
func (m *PrintJobConfiguration) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
// Sets the outputBin property value. The output bin to place completed prints into. See the printer's capabilities for a list of supported output bins.
// Parameters:
//  - value : Value to set for the outputBin property.
func (m *PrintJobConfiguration) SetOutputBin(value *string)() {
    m.outputBin = value
}
// Sets the pageRanges property value. The page ranges to print. Read-only.
// Parameters:
//  - value : Value to set for the pageRanges property.
func (m *PrintJobConfiguration) SetPageRanges(value []IntegerRange)() {
    m.pageRanges = value
}
// Sets the pagesPerSheet property value. The number of document pages to print on each sheet.
// Parameters:
//  - value : Value to set for the pagesPerSheet property.
func (m *PrintJobConfiguration) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
// Sets the quality property value. The print quality to use when printing the job. Valid values are described in the table below. Read-only.
// Parameters:
//  - value : Value to set for the quality property.
func (m *PrintJobConfiguration) SetQuality(value *PrintQuality)() {
    m.quality = value
}
// Sets the scaling property value. Specifies how the printer should scale the document data to fit the requested media. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the scaling property.
func (m *PrintJobConfiguration) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
