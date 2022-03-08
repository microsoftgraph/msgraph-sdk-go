package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintSettings provides operations to manage the print singleton.
type PrintSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether document conversion is enabled for the tenant. If document conversion is enabled, Universal Print service will automatically convert documents into a format compatible with the printer (xps to pdf) when needed.
    documentConversionEnabled *bool;
}
// NewPrintSettings instantiates a new printSettings and sets the default values.
func NewPrintSettings()(*PrintSettings) {
    m := &PrintSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDocumentConversionEnabled gets the documentConversionEnabled property value. Specifies whether document conversion is enabled for the tenant. If document conversion is enabled, Universal Print service will automatically convert documents into a format compatible with the printer (xps to pdf) when needed.
func (m *PrintSettings) GetDocumentConversionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.documentConversionEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["documentConversionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentConversionEnabled(val)
        }
        return nil
    }
    return res
}
func (m *PrintSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("documentConversionEnabled", m.GetDocumentConversionEnabled())
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
func (m *PrintSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDocumentConversionEnabled sets the documentConversionEnabled property value. Specifies whether document conversion is enabled for the tenant. If document conversion is enabled, Universal Print service will automatically convert documents into a format compatible with the printer (xps to pdf) when needed.
func (m *PrintSettings) SetDocumentConversionEnabled(value *bool)() {
    if m != nil {
        m.documentConversionEnabled = value
    }
}
