package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterCapabilities 
type PrinterCapabilities struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A list of supported bottom margins(in microns) for the printer.
    bottomMargins []int32
    // True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
    collation *bool
    // The color modes supported by the printer. Valid values are described in the following table.
    colorModes []string
    // A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
    contentTypes []string
    // The range of copies per job supported by the printer.
    copiesPerJob IntegerRangeable
    // The list of print resolutions in DPI that are supported by the printer.
    dpis []int32
    // The list of duplex modes that are supported by the printer. Valid values are described in the following table.
    duplexModes []string
    // The list of feed orientations that are supported by the printer.
    feedOrientations []string
    // Finishing processes the printer supports for a printed document.
    finishings []string
    // Supported input bins for the printer.
    inputBins []string
    // True if color printing is supported by the printer; false otherwise. Read-only.
    isColorPrintingSupported *bool
    // True if the printer supports printing by page ranges; false otherwise.
    isPageRangeSupported *bool
    // A list of supported left margins(in microns) for the printer.
    leftMargins []int32
    // The media (i.e., paper) colors supported by the printer.
    mediaColors []string
    // The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
    mediaSizes []string
    // The media types supported by the printer.
    mediaTypes []string
    // The presentation directions supported by the printer. Supported values are described in the following table.
    multipageLayouts []string
    // The OdataType property
    odataType *string
    // The print orientations supported by the printer. Valid values are described in the following table.
    orientations []string
    // The printer's supported output bins (trays).
    outputBins []string
    // Supported number of Input Pages to impose upon a single Impression.
    pagesPerSheet []int32
    // The print qualities supported by the printer.
    qualities []string
    // A list of supported right margins(in microns) for the printer.
    rightMargins []int32
    // Supported print scalings.
    scalings []string
    // True if the printer supports scaling PDF pages to match the print media size; false otherwise.
    supportsFitPdfToPage *bool
    // A list of supported top margins(in microns) for the printer.
    topMargins []int32
}
// NewPrinterCapabilities instantiates a new printerCapabilities and sets the default values.
func NewPrinterCapabilities()(*PrinterCapabilities) {
    m := &PrinterCapabilities{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.printerCapabilities";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePrinterCapabilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterCapabilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterCapabilities(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterCapabilities) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBottomMargins gets the bottomMargins property value. A list of supported bottom margins(in microns) for the printer.
func (m *PrinterCapabilities) GetBottomMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.bottomMargins
    }
}
// GetCollation gets the collation property value. True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
func (m *PrinterCapabilities) GetCollation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.collation
    }
}
// GetColorModes gets the colorModes property value. The color modes supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetColorModes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.colorModes
    }
}
// GetContentTypes gets the contentTypes property value. A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
func (m *PrinterCapabilities) GetContentTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentTypes
    }
}
// GetCopiesPerJob gets the copiesPerJob property value. The range of copies per job supported by the printer.
func (m *PrinterCapabilities) GetCopiesPerJob()(IntegerRangeable) {
    if m == nil {
        return nil
    } else {
        return m.copiesPerJob
    }
}
// GetDpis gets the dpis property value. The list of print resolutions in DPI that are supported by the printer.
func (m *PrinterCapabilities) GetDpis()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.dpis
    }
}
// GetDuplexModes gets the duplexModes property value. The list of duplex modes that are supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetDuplexModes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.duplexModes
    }
}
// GetFeedOrientations gets the feedOrientations property value. The list of feed orientations that are supported by the printer.
func (m *PrinterCapabilities) GetFeedOrientations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.feedOrientations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterCapabilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bottomMargins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["collation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCollation(val)
        }
        return nil
    }
    res["colorModes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetColorModes(res)
        }
        return nil
    }
    res["contentTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["copiesPerJob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntegerRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopiesPerJob(val.(IntegerRangeable))
        }
        return nil
    }
    res["dpis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["duplexModes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDuplexModes(res)
        }
        return nil
    }
    res["feedOrientations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFeedOrientations(res)
        }
        return nil
    }
    res["finishings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFinishings(res)
        }
        return nil
    }
    res["inputBins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["isColorPrintingSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsColorPrintingSupported(val)
        }
        return nil
    }
    res["isPageRangeSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPageRangeSupported(val)
        }
        return nil
    }
    res["leftMargins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["mediaColors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["mediaSizes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["mediaTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["multipageLayouts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMultipageLayouts(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["orientations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrientations(res)
        }
        return nil
    }
    res["outputBins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["pagesPerSheet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["qualities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetQualities(res)
        }
        return nil
    }
    res["rightMargins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["scalings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetScalings(res)
        }
        return nil
    }
    res["supportsFitPdfToPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsFitPdfToPage(val)
        }
        return nil
    }
    res["topMargins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFinishings gets the finishings property value. Finishing processes the printer supports for a printed document.
func (m *PrinterCapabilities) GetFinishings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
// GetInputBins gets the inputBins property value. Supported input bins for the printer.
func (m *PrinterCapabilities) GetInputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.inputBins
    }
}
// GetIsColorPrintingSupported gets the isColorPrintingSupported property value. True if color printing is supported by the printer; false otherwise. Read-only.
func (m *PrinterCapabilities) GetIsColorPrintingSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isColorPrintingSupported
    }
}
// GetIsPageRangeSupported gets the isPageRangeSupported property value. True if the printer supports printing by page ranges; false otherwise.
func (m *PrinterCapabilities) GetIsPageRangeSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPageRangeSupported
    }
}
// GetLeftMargins gets the leftMargins property value. A list of supported left margins(in microns) for the printer.
func (m *PrinterCapabilities) GetLeftMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.leftMargins
    }
}
// GetMediaColors gets the mediaColors property value. The media (i.e., paper) colors supported by the printer.
func (m *PrinterCapabilities) GetMediaColors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaColors
    }
}
// GetMediaSizes gets the mediaSizes property value. The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
func (m *PrinterCapabilities) GetMediaSizes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSizes
    }
}
// GetMediaTypes gets the mediaTypes property value. The media types supported by the printer.
func (m *PrinterCapabilities) GetMediaTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mediaTypes
    }
}
// GetMultipageLayouts gets the multipageLayouts property value. The presentation directions supported by the printer. Supported values are described in the following table.
func (m *PrinterCapabilities) GetMultipageLayouts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayouts
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterCapabilities) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetOrientations gets the orientations property value. The print orientations supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) GetOrientations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orientations
    }
}
// GetOutputBins gets the outputBins property value. The printer's supported output bins (trays).
func (m *PrinterCapabilities) GetOutputBins()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outputBins
    }
}
// GetPagesPerSheet gets the pagesPerSheet property value. Supported number of Input Pages to impose upon a single Impression.
func (m *PrinterCapabilities) GetPagesPerSheet()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
// GetQualities gets the qualities property value. The print qualities supported by the printer.
func (m *PrinterCapabilities) GetQualities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.qualities
    }
}
// GetRightMargins gets the rightMargins property value. A list of supported right margins(in microns) for the printer.
func (m *PrinterCapabilities) GetRightMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.rightMargins
    }
}
// GetScalings gets the scalings property value. Supported print scalings.
func (m *PrinterCapabilities) GetScalings()([]string) {
    if m == nil {
        return nil
    } else {
        return m.scalings
    }
}
// GetSupportsFitPdfToPage gets the supportsFitPdfToPage property value. True if the printer supports scaling PDF pages to match the print media size; false otherwise.
func (m *PrinterCapabilities) GetSupportsFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsFitPdfToPage
    }
}
// GetTopMargins gets the topMargins property value. A list of supported top margins(in microns) for the printer.
func (m *PrinterCapabilities) GetTopMargins()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.topMargins
    }
}
// Serialize serializes information the current object
func (m *PrinterCapabilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBottomMargins() != nil {
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
    if m.GetColorModes() != nil {
        err := writer.WriteCollectionOfStringValues("colorModes", m.GetColorModes())
        if err != nil {
            return err
        }
    }
    if m.GetContentTypes() != nil {
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
    if m.GetDpis() != nil {
        err := writer.WriteCollectionOfInt32Values("dpis", m.GetDpis())
        if err != nil {
            return err
        }
    }
    if m.GetDuplexModes() != nil {
        err := writer.WriteCollectionOfStringValues("duplexModes", m.GetDuplexModes())
        if err != nil {
            return err
        }
    }
    if m.GetFeedOrientations() != nil {
        err := writer.WriteCollectionOfStringValues("feedOrientations", m.GetFeedOrientations())
        if err != nil {
            return err
        }
    }
    if m.GetFinishings() != nil {
        err := writer.WriteCollectionOfStringValues("finishings", m.GetFinishings())
        if err != nil {
            return err
        }
    }
    if m.GetInputBins() != nil {
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
    if m.GetLeftMargins() != nil {
        err := writer.WriteCollectionOfInt32Values("leftMargins", m.GetLeftMargins())
        if err != nil {
            return err
        }
    }
    if m.GetMediaColors() != nil {
        err := writer.WriteCollectionOfStringValues("mediaColors", m.GetMediaColors())
        if err != nil {
            return err
        }
    }
    if m.GetMediaSizes() != nil {
        err := writer.WriteCollectionOfStringValues("mediaSizes", m.GetMediaSizes())
        if err != nil {
            return err
        }
    }
    if m.GetMediaTypes() != nil {
        err := writer.WriteCollectionOfStringValues("mediaTypes", m.GetMediaTypes())
        if err != nil {
            return err
        }
    }
    if m.GetMultipageLayouts() != nil {
        err := writer.WriteCollectionOfStringValues("multipageLayouts", m.GetMultipageLayouts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOrientations() != nil {
        err := writer.WriteCollectionOfStringValues("orientations", m.GetOrientations())
        if err != nil {
            return err
        }
    }
    if m.GetOutputBins() != nil {
        err := writer.WriteCollectionOfStringValues("outputBins", m.GetOutputBins())
        if err != nil {
            return err
        }
    }
    if m.GetPagesPerSheet() != nil {
        err := writer.WriteCollectionOfInt32Values("pagesPerSheet", m.GetPagesPerSheet())
        if err != nil {
            return err
        }
    }
    if m.GetQualities() != nil {
        err := writer.WriteCollectionOfStringValues("qualities", m.GetQualities())
        if err != nil {
            return err
        }
    }
    if m.GetRightMargins() != nil {
        err := writer.WriteCollectionOfInt32Values("rightMargins", m.GetRightMargins())
        if err != nil {
            return err
        }
    }
    if m.GetScalings() != nil {
        err := writer.WriteCollectionOfStringValues("scalings", m.GetScalings())
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
    if m.GetTopMargins() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterCapabilities) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBottomMargins sets the bottomMargins property value. A list of supported bottom margins(in microns) for the printer.
func (m *PrinterCapabilities) SetBottomMargins(value []int32)() {
    if m != nil {
        m.bottomMargins = value
    }
}
// SetCollation sets the collation property value. True if the printer supports collating when printing muliple copies of a multi-page document; false otherwise.
func (m *PrinterCapabilities) SetCollation(value *bool)() {
    if m != nil {
        m.collation = value
    }
}
// SetColorModes sets the colorModes property value. The color modes supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) SetColorModes(value []string)() {
    if m != nil {
        m.colorModes = value
    }
}
// SetContentTypes sets the contentTypes property value. A list of supported content (MIME) types that the printer supports. It is not guaranteed that the Universal Print service supports printing all of these MIME types.
func (m *PrinterCapabilities) SetContentTypes(value []string)() {
    if m != nil {
        m.contentTypes = value
    }
}
// SetCopiesPerJob sets the copiesPerJob property value. The range of copies per job supported by the printer.
func (m *PrinterCapabilities) SetCopiesPerJob(value IntegerRangeable)() {
    if m != nil {
        m.copiesPerJob = value
    }
}
// SetDpis sets the dpis property value. The list of print resolutions in DPI that are supported by the printer.
func (m *PrinterCapabilities) SetDpis(value []int32)() {
    if m != nil {
        m.dpis = value
    }
}
// SetDuplexModes sets the duplexModes property value. The list of duplex modes that are supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) SetDuplexModes(value []string)() {
    if m != nil {
        m.duplexModes = value
    }
}
// SetFeedOrientations sets the feedOrientations property value. The list of feed orientations that are supported by the printer.
func (m *PrinterCapabilities) SetFeedOrientations(value []string)() {
    if m != nil {
        m.feedOrientations = value
    }
}
// SetFinishings sets the finishings property value. Finishing processes the printer supports for a printed document.
func (m *PrinterCapabilities) SetFinishings(value []string)() {
    if m != nil {
        m.finishings = value
    }
}
// SetInputBins sets the inputBins property value. Supported input bins for the printer.
func (m *PrinterCapabilities) SetInputBins(value []string)() {
    if m != nil {
        m.inputBins = value
    }
}
// SetIsColorPrintingSupported sets the isColorPrintingSupported property value. True if color printing is supported by the printer; false otherwise. Read-only.
func (m *PrinterCapabilities) SetIsColorPrintingSupported(value *bool)() {
    if m != nil {
        m.isColorPrintingSupported = value
    }
}
// SetIsPageRangeSupported sets the isPageRangeSupported property value. True if the printer supports printing by page ranges; false otherwise.
func (m *PrinterCapabilities) SetIsPageRangeSupported(value *bool)() {
    if m != nil {
        m.isPageRangeSupported = value
    }
}
// SetLeftMargins sets the leftMargins property value. A list of supported left margins(in microns) for the printer.
func (m *PrinterCapabilities) SetLeftMargins(value []int32)() {
    if m != nil {
        m.leftMargins = value
    }
}
// SetMediaColors sets the mediaColors property value. The media (i.e., paper) colors supported by the printer.
func (m *PrinterCapabilities) SetMediaColors(value []string)() {
    if m != nil {
        m.mediaColors = value
    }
}
// SetMediaSizes sets the mediaSizes property value. The media sizes supported by the printer. Supports standard size names for ISO and ANSI media sizes. Valid values are in the following table.
func (m *PrinterCapabilities) SetMediaSizes(value []string)() {
    if m != nil {
        m.mediaSizes = value
    }
}
// SetMediaTypes sets the mediaTypes property value. The media types supported by the printer.
func (m *PrinterCapabilities) SetMediaTypes(value []string)() {
    if m != nil {
        m.mediaTypes = value
    }
}
// SetMultipageLayouts sets the multipageLayouts property value. The presentation directions supported by the printer. Supported values are described in the following table.
func (m *PrinterCapabilities) SetMultipageLayouts(value []string)() {
    if m != nil {
        m.multipageLayouts = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterCapabilities) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetOrientations sets the orientations property value. The print orientations supported by the printer. Valid values are described in the following table.
func (m *PrinterCapabilities) SetOrientations(value []string)() {
    if m != nil {
        m.orientations = value
    }
}
// SetOutputBins sets the outputBins property value. The printer's supported output bins (trays).
func (m *PrinterCapabilities) SetOutputBins(value []string)() {
    if m != nil {
        m.outputBins = value
    }
}
// SetPagesPerSheet sets the pagesPerSheet property value. Supported number of Input Pages to impose upon a single Impression.
func (m *PrinterCapabilities) SetPagesPerSheet(value []int32)() {
    if m != nil {
        m.pagesPerSheet = value
    }
}
// SetQualities sets the qualities property value. The print qualities supported by the printer.
func (m *PrinterCapabilities) SetQualities(value []string)() {
    if m != nil {
        m.qualities = value
    }
}
// SetRightMargins sets the rightMargins property value. A list of supported right margins(in microns) for the printer.
func (m *PrinterCapabilities) SetRightMargins(value []int32)() {
    if m != nil {
        m.rightMargins = value
    }
}
// SetScalings sets the scalings property value. Supported print scalings.
func (m *PrinterCapabilities) SetScalings(value []string)() {
    if m != nil {
        m.scalings = value
    }
}
// SetSupportsFitPdfToPage sets the supportsFitPdfToPage property value. True if the printer supports scaling PDF pages to match the print media size; false otherwise.
func (m *PrinterCapabilities) SetSupportsFitPdfToPage(value *bool)() {
    if m != nil {
        m.supportsFitPdfToPage = value
    }
}
// SetTopMargins sets the topMargins property value. A list of supported top margins(in microns) for the printer.
func (m *PrinterCapabilities) SetTopMargins(value []int32)() {
    if m != nil {
        m.topMargins = value
    }
}
