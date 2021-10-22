package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintCertificateSigningRequest struct {
    additionalData map[string]interface{};
    content *string;
    transportKey *string;
}
func NewPrintCertificateSigningRequest()(*PrintCertificateSigningRequest) {
    m := &PrintCertificateSigningRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrintCertificateSigningRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrintCertificateSigningRequest) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *PrintCertificateSigningRequest) GetTransportKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.transportKey
    }
}
func (m *PrintCertificateSigningRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["transportKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTransportKey(val)
        return nil
    }
    return res
}
func (m *PrintCertificateSigningRequest) IsNil()(bool) {
    return m == nil
}
func (m *PrintCertificateSigningRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("transportKey", m.GetTransportKey())
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
func (m *PrintCertificateSigningRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrintCertificateSigningRequest) SetContent(value *string)() {
    m.content = value
}
func (m *PrintCertificateSigningRequest) SetTransportKey(value *string)() {
    m.transportKey = value
}
