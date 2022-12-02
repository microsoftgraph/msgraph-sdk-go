package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityApiConnectorsItemUploadClientCertificatePostRequestBody provides operations to call the uploadClientCertificate method.
type IdentityApiConnectorsItemUploadClientCertificatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The password property
    password *string
    // The pkcs12Value property
    pkcs12Value *string
}
// NewIdentityApiConnectorsItemUploadClientCertificatePostRequestBody instantiates a new IdentityApiConnectorsItemUploadClientCertificatePostRequestBody and sets the default values.
func NewIdentityApiConnectorsItemUploadClientCertificatePostRequestBody()(*IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) {
    m := &IdentityApiConnectorsItemUploadClientCertificatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityApiConnectorsItemUploadClientCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityApiConnectorsItemUploadClientCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityApiConnectorsItemUploadClientCertificatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["pkcs12Value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) GetPassword()(*string) {
    return m.password
}
// GetPkcs12Value gets the pkcs12Value property value. The pkcs12Value property
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) GetPkcs12Value()(*string) {
    return m.pkcs12Value
}
// Serialize serializes information the current object
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPassword sets the password property value. The password property
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) SetPassword(value *string)() {
    m.password = value
}
// SetPkcs12Value sets the pkcs12Value property value. The pkcs12Value property
func (m *IdentityApiConnectorsItemUploadClientCertificatePostRequestBody) SetPkcs12Value(value *string)() {
    m.pkcs12Value = value
}
