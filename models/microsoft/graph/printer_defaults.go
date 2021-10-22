package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrinterDefaults struct {
    additionalData map[string]interface{};
    colorMode *PrintColorMode;
    contentType *string;
    copiesPerJob *int32;
    dpi *int32;
    duplexMode *PrintDuplexMode;
    finishings []PrintFinishing;
    fitPdfToPage *bool;
    inputBin *string;
    mediaColor *string;
    mediaSize *string;
    mediaType *string;
    multipageLayout *PrintMultipageLayout;
    orientation *PrintOrientation;
    outputBin *string;
    pagesPerSheet *int32;
    quality *PrintQuality;
    scaling *PrintScaling;
}
func NewPrinterDefaults()(*PrinterDefaults) {
    m := &PrinterDefaults{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrinterDefaults) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrinterDefaults) GetColorMode()(*PrintColorMode) {
    if m == nil {
        return nil
    } else {
        return m.colorMode
    }
}
func (m *PrinterDefaults) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *PrinterDefaults) GetCopiesPerJob()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.copiesPerJob
    }
}
func (m *PrinterDefaults) GetDpi()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dpi
    }
}
func (m *PrinterDefaults) GetDuplexMode()(*PrintDuplexMode) {
    if m == nil {
        return nil
    } else {
        return m.duplexMode
    }
}
func (m *PrinterDefaults) GetFinishings()([]PrintFinishing) {
    if m == nil {
        return nil
    } else {
        return m.finishings
    }
}
func (m *PrinterDefaults) GetFitPdfToPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fitPdfToPage
    }
}
func (m *PrinterDefaults) GetInputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inputBin
    }
}
func (m *PrinterDefaults) GetMediaColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaColor
    }
}
func (m *PrinterDefaults) GetMediaSize()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaSize
    }
}
func (m *PrinterDefaults) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
func (m *PrinterDefaults) GetMultipageLayout()(*PrintMultipageLayout) {
    if m == nil {
        return nil
    } else {
        return m.multipageLayout
    }
}
func (m *PrinterDefaults) GetOrientation()(*PrintOrientation) {
    if m == nil {
        return nil
    } else {
        return m.orientation
    }
}
func (m *PrinterDefaults) GetOutputBin()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputBin
    }
}
func (m *PrinterDefaults) GetPagesPerSheet()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pagesPerSheet
    }
}
func (m *PrinterDefaults) GetQuality()(*PrintQuality) {
    if m == nil {
        return nil
    } else {
        return m.quality
    }
}
func (m *PrinterDefaults) GetScaling()(*PrintScaling) {
    if m == nil {
        return nil
    } else {
        return m.scaling
    }
}
func (m *PrinterDefaults) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["colorMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintColorMode)
        if err != nil {
            return err
        }
        cast := val.(PrintColorMode)
        m.SetColorMode(&cast)
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    res["copiesPerJob"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCopiesPerJob(val)
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
    res["mediaColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaColor(val)
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
func (m *PrinterDefaults) IsNil()(bool) {
    return m == nil
}
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
func (m *PrinterDefaults) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrinterDefaults) SetColorMode(value *PrintColorMode)() {
    m.colorMode = value
}
func (m *PrinterDefaults) SetContentType(value *string)() {
    m.contentType = value
}
func (m *PrinterDefaults) SetCopiesPerJob(value *int32)() {
    m.copiesPerJob = value
}
func (m *PrinterDefaults) SetDpi(value *int32)() {
    m.dpi = value
}
func (m *PrinterDefaults) SetDuplexMode(value *PrintDuplexMode)() {
    m.duplexMode = value
}
func (m *PrinterDefaults) SetFinishings(value []PrintFinishing)() {
    m.finishings = value
}
func (m *PrinterDefaults) SetFitPdfToPage(value *bool)() {
    m.fitPdfToPage = value
}
func (m *PrinterDefaults) SetInputBin(value *string)() {
    m.inputBin = value
}
func (m *PrinterDefaults) SetMediaColor(value *string)() {
    m.mediaColor = value
}
func (m *PrinterDefaults) SetMediaSize(value *string)() {
    m.mediaSize = value
}
func (m *PrinterDefaults) SetMediaType(value *string)() {
    m.mediaType = value
}
func (m *PrinterDefaults) SetMultipageLayout(value *PrintMultipageLayout)() {
    m.multipageLayout = value
}
func (m *PrinterDefaults) SetOrientation(value *PrintOrientation)() {
    m.orientation = value
}
func (m *PrinterDefaults) SetOutputBin(value *string)() {
    m.outputBin = value
}
func (m *PrinterDefaults) SetPagesPerSheet(value *int32)() {
    m.pagesPerSheet = value
}
func (m *PrinterDefaults) SetQuality(value *PrintQuality)() {
    m.quality = value
}
func (m *PrinterDefaults) SetScaling(value *PrintScaling)() {
    m.scaling = value
}
