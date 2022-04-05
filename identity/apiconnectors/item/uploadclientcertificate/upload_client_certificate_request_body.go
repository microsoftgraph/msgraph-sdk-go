package uploadclientcertificate

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UploadClientCertificateRequestBody provides operations to call the uploadClientCertificate method.
type UploadClientCertificateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The password property
    password *string;
    // The pkcs12Value property
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
func CreateUploadClientCertificateRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
func (m *UploadClientCertificateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["pkcs12Value"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetPassword gets the password property value. The password property
func (m *UploadClientCertificateRequestBody) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
// GetPkcs12Value gets the pkcs12Value property value. The pkcs12Value property
func (m *UploadClientCertificateRequestBody) GetPkcs12Value()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pkcs12Value
    }
}
// Serialize serializes information the current object
func (m *UploadClientCertificateRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetPassword sets the password property value. The password property
func (m *UploadClientCertificateRequestBody) SetPassword(value *string)() {
    if m != nil {
        m.password = value
    }
}
// SetPkcs12Value sets the pkcs12Value property value. The pkcs12Value property
func (m *UploadClientCertificateRequestBody) SetPkcs12Value(value *string)() {
    if m != nil {
        m.pkcs12Value = value
    }
}
