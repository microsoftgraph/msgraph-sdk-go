package uploadclientcertificate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UploadClientCertificateRequestBody struct {
    additionalData map[string]interface{};
    password *string;
    pkcs12Value *string;
}
func NewUploadClientCertificateRequestBody()(*UploadClientCertificateRequestBody) {
    m := &UploadClientCertificateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UploadClientCertificateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UploadClientCertificateRequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
func (m *UploadClientCertificateRequestBody) GetPkcs12Value()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pkcs12Value
    }
}
func (m *UploadClientCertificateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPassword(val)
        return nil
    }
    res["pkcs12Value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPkcs12Value(val)
        return nil
    }
    return res
}
func (m *UploadClientCertificateRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *UploadClientCertificateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pkcs12Value", m.GetPkcs12Value())
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
func (m *UploadClientCertificateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UploadClientCertificateRequestBody) SetPassword(value *string)() {
    m.password = value
}
func (m *UploadClientCertificateRequestBody) SetPkcs12Value(value *string)() {
    m.pkcs12Value = value
}
