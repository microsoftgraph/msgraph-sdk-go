package uploadclientcertificate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UploadClientCertificateRequestBody provides operations to call the uploadClientCertificate method.
type UploadClientCertificateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    password *string;
    // 
    pkcs12Value *string;
}
// NewUploadClientCertificateRequestBody instantiates a new uploadClientCertificateRequestBody and sets the default values.
func NewUploadClientCertificateRequestBody()(*UploadClientCertificateRequestBody) {
    m := &UploadClientCertificateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUploadClientCertificateRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUploadClientCertificateRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUploadClientCertificateRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadClientCertificateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UploadClientCertificateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["pkcs12Value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPkcs12Value(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. 
func (m *UploadClientCertificateRequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// GetPkcs12Value gets the pkcs12Value property value. 
func (m *UploadClientCertificateRequestBody) GetPkcs12Value()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pkcs12Value
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UploadClientCertificateRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPassword sets the password property value. 
func (m *UploadClientCertificateRequestBody) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
// SetPkcs12Value sets the pkcs12Value property value. 
func (m *UploadClientCertificateRequestBody) SetPkcs12Value(value *string)() {
    if m != nil {
        m.pkcs12Value = value
    }
}
