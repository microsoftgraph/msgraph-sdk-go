package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintSettings struct {
    additionalData map[string]interface{};
    documentConversionEnabled *bool;
}
func NewPrintSettings()(*PrintSettings) {
    m := &PrintSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrintSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrintSettings) GetDocumentConversionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.documentConversionEnabled
    }
}
func (m *PrintSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["documentConversionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDocumentConversionEnabled(val)
        return nil
    }
    return res
}
func (m *PrintSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *PrintSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrintSettings) SetDocumentConversionEnabled(value *bool)() {
    m.documentConversionEnabled = value
}
