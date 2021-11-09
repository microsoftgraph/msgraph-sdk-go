package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrinterCapabilities struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of supported bottom margins(in microns) for the printer.
    bottomMargins []int32;
    // True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
    collation *bool;
    // The color modes supported by the printer. Valid values are described in the following table.
    colorModes []PrintColorMode;
    // A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
    contentTypes []string;
    // The range of copies per job supported by the printer.
    copiesPerJob *IntegerRange;
    // The list of print resolutions in DPI that are supported by the printer.
    dpis []int32;
    // The list of duplex modes that are supported by the printer. Valid values are described in the following table.
    duplexModes []PrintDuplexMode;
    // The list of feed orientations that are supported by the printer.
    feedOrientations []PrinterFeedOrientation;
    // Finishing processes the printer supports for a printed document.
    finishings []PrintFinishing;
    // Supported input bins for the printer.
    inputBins []string;
    // True if color printing is supported by the printer; false otherwise. Read-only.
    isColorPrintingSupported *bool;
    // True if the printer supports printing by page ranges; false otherwise.
    isPageRangeSupported *bool;
    // A list of supported left margins(in microns) for the printer.
    leftMargins []int32;
    // The media (i.e., paper) colors supported by the printer.
    mediaColors []string;
    // The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
    mediaSizes []string;
    // The media types supported by the printer.
    mediaTypes []string;
    // The presentation directions supported by the printer. Supported values are described in the following table.
    multipageLayouts []PrintMultipageLayout;
    // The print orientations supported by the printer. Valid values are described in the following table.
    orientations []PrintOrientation;
    // The printer's supported output bins (trays).
    outputBins []string;
    // Supported number of Input Pages to impose upon a single Impression.
    pagesPerSheet []int32;
    // The print qualities supported by the printer.
    qualities []PrintQuality;
    // A list of supported right margins(in microns) for the printer.
    rightMargins []int32;
    // Supported print scalings.
    scalings []PrintScaling;
    // True if the printer supports scaling PDF pages to match the print media size; false otherwise.
    supportsFitPdfToPage *bool;
    // A list of supported top margins(in microns) for the printer.
    topMargins []int32;
}
// Instantiates a new printerCapabilities and sets the default values.
func NewPrinterCapabilities()(*PrinterCapabilities) {
    m := &PrinterCapabilities{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterCapabilities) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the bottomMargins property value. A list of supported bottom margins(in microns) for the printer.
func (m *PrinterCapabilities) GetBottomMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.bottomMargins
    }
}
// Gets the collation property value. True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
func (m *PrinterCapabilities) GetCollation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collation
    }
}
// Gets the colorModes property value. The color modes supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetColorModes()([]PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorModes
    }
}
// Gets the contentTypes property value. A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
func (m *PrinterCapabilities) GetContentTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentTypes
    }
}
// Gets the copiesPerJob property value. The range of copies per job supported by the printer.
func (m *PrinterCapabilities) GetCopiesPerJob()(*IntegerRange) {
    if m == nil {
        return nil
    } else {
        return m.copiesPerJob
    }
}
// Gets the dpis property value. The list of print resolutions in DPI that are supported by the printer.
func (m *PrinterCapabilities) GetDpis()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.dpis
    }
}
// Gets the duplexModes property value. The list of duplex modes that are supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetDuplexModes()([]PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexModes
    }
}
// Gets the feedOrientations property value. The list of feed orientations that are supported by the printer.
func (m *PrinterCapabilities) GetFeedOrientations()([]PrinterFeedOrientation) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientations
    }
}
// Gets the finishings property value. Finishing processes the printer supports for a printed document.
func (m *PrinterCapabilities) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// Gets the inputBins property value. Supported input bins for the printer.
func (m *PrinterCapabilities) GetInputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inputBins
    }
}
// Gets the isColorPrintingSupported property value. True if color printing is supported by the printer; false otherwise. Read-only.
func (m *PrinterCapabilities) GetIsColorPrintingSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isColorPrintingSupported
    }
}
// Gets the isPageRangeSupported property value. True if the printer supports printing by page ranges; false otherwise.
func (m *PrinterCapabilities) GetIsPageRangeSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPageRangeSupported
    }
}
// Gets the leftMargins property value. A list of supported left margins(in microns) for the printer.
func (m *PrinterCapabilities) GetLeftMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.leftMargins
    }
}
// Gets the mediaColors property value. The media (i.e., paper) colors supported by the printer.
func (m *PrinterCapabilities) GetMediaColors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaColors
    }
}
// Gets the mediaSizes property value. The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
func (m *PrinterCapabilities) GetMediaSizes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSizes
    }
}
// Gets the mediaTypes property value. The media types supported by the printer.
func (m *PrinterCapabilities) GetMediaTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaTypes
    }
}
// Gets the multipageLayouts property value. The presentation directions supported by the printer. Supported values are described in the following table.
func (m *PrinterCapabilities) GetMultipageLayouts()([]PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayouts
    }
}
// Gets the orientations property value. The print orientations supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetOrientations()([]PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientations
    }
}
// Gets the outputBins property value. The printer's supported output bins (trays).
func (m *PrinterCapabilities) GetOutputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outputBins
    }
}
// Gets the pagesPerSheet property value. Supported number of Input Pages to impose upon a single Impression.
func (m *PrinterCapabilities) GetPagesPerSheet()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// Gets the qualities property value. The print qualities supported by the printer.
func (m *PrinterCapabilities) GetQualities()([]PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.qualities
    }
}
// Gets the rightMargins property value. A list of supported right margins(in microns) for the printer.
func (m *PrinterCapabilities) GetRightMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.rightMargins
    }
}
// Gets the scalings property value. Supported print scalings.
func (m *PrinterCapabilities) GetScalings()([]PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scalings
    }
}
// Gets the supportsFitPdfToPage property value. True if the printer supports scaling PDF pages to match the print media size; false otherwise.
func (m *PrinterCapabilities) GetSupportsFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsFitPdfToPage
    }
}
// Gets the topMargins property value. A list of supported top margins(in microns) for the printer.
func (m *PrinterCapabilities) GetTopMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.topMargins
    }
}
// The deserialization information for the current model
func (m *PrinterCapabilities) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bottomMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetBottomMargins(res)
        }
        return nil
    }
    res["collation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollation(val)
        }
        return nil
    }
    res["colorModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintColorMode)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintColorMode, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintColorMode))
            }
            m.SetColorModes(res)
        }
        return nil
    }
    res["contentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContentTypes(res)
        }
        return nil
    }
    res["copiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntegerRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPerJob(val.(*IntegerRange))
        }
        return nil
    }
    res["dpis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetDpis(res)
        }
        return nil
    }
    res["duplexModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintDuplexMode)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintDuplexMode, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintDuplexMode))
            }
            m.SetDuplexModes(res)
        }
        return nil
    }
    res["feedOrientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrinterFeedOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterFeedOrientation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrinterFeedOrientation))
            }
            m.SetFeedOrientations(res)
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
    res["inputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInputBins(res)
        }
        return nil
    }
    res["isColorPrintingSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsColorPrintingSupported(val)
        }
        return nil
    }
    res["isPageRangeSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPageRangeSupported(val)
        }
        return nil
    }
    res["leftMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetLeftMargins(res)
        }
        return nil
    }
    res["mediaColors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMediaColors(res)
        }
        return nil
    }
    res["mediaSizes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMediaSizes(res)
        }
        return nil
    }
    res["mediaTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMediaTypes(res)
        }
        return nil
    }
    res["multipageLayouts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintMultipageLayout)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintMultipageLayout, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintMultipageLayout))
            }
            m.SetMultipageLayouts(res)
        }
        return nil
    }
    res["orientations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintOrientation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintOrientation))
            }
            m.SetOrientations(res)
        }
        return nil
    }
    res["outputBins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOutputBins(res)
        }
        return nil
    }
    res["pagesPerSheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetPagesPerSheet(res)
        }
        return nil
    }
    res["qualities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintQuality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintQuality, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintQuality))
            }
            m.SetQualities(res)
        }
        return nil
    }
    res["rightMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetRightMargins(res)
        }
        return nil
    }
    res["scalings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParsePrintScaling)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintScaling, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintScaling))
            }
            m.SetScalings(res)
        }
        return nil
    }
    res["supportsFitPdfToPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsFitPdfToPage(val)
        }
        return nil
    }
    res["topMargins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetTopMargins(res)
        }
        return nil
    }
    return res
}
func (m *PrinterCapabilities) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrinterCapabilities) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfInt32Values("bottomMargins", m.GetBottomMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("collation", m.GetCollation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("colorModes", SerializePrintColorMode(m.GetColorModes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("contentTypes", m.GetContentTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("copiesPerJob", m.GetCopiesPerJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("dpis", m.GetDpis())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("duplexModes", SerializePrintDuplexMode(m.GetDuplexModes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("feedOrientations", SerializePrinterFeedOrientation(m.GetFeedOrientations()))
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
        err := writer.WriteCollectionOfStringValues("inputBins", m.GetInputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isColorPrintingSupported", m.GetIsColorPrintingSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPageRangeSupported", m.GetIsPageRangeSupported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("leftMargins", m.GetLeftMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("mediaColors", m.GetMediaColors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("mediaSizes", m.GetMediaSizes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("mediaTypes", m.GetMediaTypes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("multipageLayouts", SerializePrintMultipageLayout(m.GetMultipageLayouts()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("orientations", SerializePrintOrientation(m.GetOrientations()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("outputBins", m.GetOutputBins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("qualities", SerializePrintQuality(m.GetQualities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("rightMargins", m.GetRightMargins())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("scalings", SerializePrintScaling(m.GetScalings()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("supportsFitPdfToPage", m.GetSupportsFitPdfToPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("topMargins", m.GetTopMargins())
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
func (m *PrinterCapabilities) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the bottomMargins property value. A list of supported bottom margins(in microns) for the printer.
// Parameters:
//  - value : Value to set for the bottomMargins property.
func (m *PrinterCapabilities) SetBottomMargins(value []int32)() {
    m.bottomMargins = value
}
// Sets the collation property value. True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
// Parameters:
//  - value : Value to set for the collation property.
func (m *PrinterCapabilities) SetCollation(value *bool)() {
    m.collation = value
}
// Sets the colorModes property value. The color modes supported by the printer. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the colorModes property.
func (m *PrinterCapabilities) SetColorModes(value []PrintColorMode)() {
    m.colorModes = value
}
// Sets the contentTypes property value. A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
// Parameters:
//  - value : Value to set for the contentTypes property.
func (m *PrinterCapabilities) SetContentTypes(value []string)() {
    m.contentTypes = value
}
// Sets the copiesPerJob property value. The range of copies per job supported by the printer.
// Parameters:
//  - value : Value to set for the copiesPerJob property.
func (m *PrinterCapabilities) SetCopiesPerJob(value *IntegerRange)() {
    m.copiesPerJob = value
}
// Sets the dpis property value. The list of print resolutions in DPI that are supported by the printer.
// Parameters:
//  - value : Value to set for the dpis property.
func (m *PrinterCapabilities) SetDpis(value []int32)() {
    m.dpis = value
}
// Sets the duplexModes property value. The list of duplex modes that are supported by the printer. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the duplexModes property.
func (m *PrinterCapabilities) SetDuplexModes(value []PrintDuplexMode)() {
    m.duplexModes = value
}
// Sets the feedOrientations property value. The list of feed orientations that are supported by the printer.
// Parameters:
//  - value : Value to set for the feedOrientations property.
func (m *PrinterCapabilities) SetFeedOrientations(value []PrinterFeedOrientation)() {
    m.feedOrientations = value
}
// Sets the finishings property value. Finishing processes the printer supports for a printed document.
// Parameters:
//  - value : Value to set for the finishings property.
func (m *PrinterCapabilities) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
// Sets the inputBins property value. Supported input bins for the printer.
// Parameters:
//  - value : Value to set for the inputBins property.
func (m *PrinterCapabilities) SetInputBins(value []string)() {
    m.inputBins = value
}
// Sets the isColorPrintingSupported property value. True if color printing is supported by the printer; false otherwise. Read-only.
// Parameters:
//  - value : Value to set for the isColorPrintingSupported property.
func (m *PrinterCapabilities) SetIsColorPrintingSupported(value *bool)() {
    m.isColorPrintingSupported = value
}
// Sets the isPageRangeSupported property value. True if the printer supports printing by page ranges; false otherwise.
// Parameters:
//  - value : Value to set for the isPageRangeSupported property.
func (m *PrinterCapabilities) SetIsPageRangeSupported(value *bool)() {
    m.isPageRangeSupported = value
}
// Sets the leftMargins property value. A list of supported left margins(in microns) for the printer.
// Parameters:
//  - value : Value to set for the leftMargins property.
func (m *PrinterCapabilities) SetLeftMargins(value []int32)() {
    m.leftMargins = value
}
// Sets the mediaColors property value. The media (i.e., paper) colors supported by the printer.
// Parameters:
//  - value : Value to set for the mediaColors property.
func (m *PrinterCapabilities) SetMediaColors(value []string)() {
    m.mediaColors = value
}
// Sets the mediaSizes property value. The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
// Parameters:
//  - value : Value to set for the mediaSizes property.
func (m *PrinterCapabilities) SetMediaSizes(value []string)() {
    m.mediaSizes = value
}
// Sets the mediaTypes property value. The media types supported by the printer.
// Parameters:
//  - value : Value to set for the mediaTypes property.
func (m *PrinterCapabilities) SetMediaTypes(value []string)() {
    m.mediaTypes = value
}
// Sets the multipageLayouts property value. The presentation directions supported by the printer. Supported values are described in the following table.
// Parameters:
//  - value : Value to set for the multipageLayouts property.
func (m *PrinterCapabilities) SetMultipageLayouts(value []PrintMultipageLayout)() {
    m.multipageLayouts = value
}
// Sets the orientations property value. The print orientations supported by the printer. Valid values are described in the following table.
// Parameters:
//  - value : Value to set for the orientations property.
func (m *PrinterCapabilities) SetOrientations(value []PrintOrientation)() {
    m.orientations = value
}
// Sets the outputBins property value. The printer's supported output bins (trays).
// Parameters:
//  - value : Value to set for the outputBins property.
func (m *PrinterCapabilities) SetOutputBins(value []string)() {
    m.outputBins = value
}
// Sets the pagesPerSheet property value. Supported number of Input Pages to impose upon a single Impression.
// Parameters:
//  - value : Value to set for the pagesPerSheet property.
func (m *PrinterCapabilities) SetPagesPerSheet(value []int32)() {
    m.pagesPerSheet = value
}
// Sets the qualities property value. The print qualities supported by the printer.
// Parameters:
//  - value : Value to set for the qualities property.
func (m *PrinterCapabilities) SetQualities(value []PrintQuality)() {
    m.qualities = value
}
// Sets the rightMargins property value. A list of supported right margins(in microns) for the printer.
// Parameters:
//  - value : Value to set for the rightMargins property.
func (m *PrinterCapabilities) SetRightMargins(value []int32)() {
    m.rightMargins = value
}
// Sets the scalings property value. Supported print scalings.
// Parameters:
//  - value : Value to set for the scalings property.
func (m *PrinterCapabilities) SetScalings(value []PrintScaling)() {
    m.scalings = value
}
// Sets the supportsFitPdfToPage property value. True if the printer supports scaling PDF pages to match the print media size; false otherwise.
// Parameters:
//  - value : Value to set for the supportsFitPdfToPage property.
func (m *PrinterCapabilities) SetSupportsFitPdfToPage(value *bool)() {
    m.supportsFitPdfToPage = value
}
// Sets the topMargins property value. A list of supported top margins(in microns) for the printer.
// Parameters:
//  - value : Value to set for the topMargins property.
func (m *PrinterCapabilities) SetTopMargins(value []int32)() {
    m.topMargins = value
}
