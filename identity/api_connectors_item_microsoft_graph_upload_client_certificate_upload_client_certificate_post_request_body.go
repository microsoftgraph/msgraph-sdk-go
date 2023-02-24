package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody 
type ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody instantiates a new ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody and sets the default values.
func NewApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody()(*ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) {
    m := &ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) GetPassword()(*string) {
    val, err := m.GetBackingStore().Get("password")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPkcs12Value gets the pkcs12Value property value. The pkcs12Value property
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) GetPkcs12Value()(*string) {
    val, err := m.GetBackingStore().Get("pkcs12Value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetPassword sets the password property value. The password property
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) SetPassword(value *string)() {
    err := m.GetBackingStore().Set("password", value)
    if err != nil {
        panic(err)
    }
}
// SetPkcs12Value sets the pkcs12Value property value. The pkcs12Value property
func (m *ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBody) SetPkcs12Value(value *string)() {
    err := m.GetBackingStore().Set("pkcs12Value", value)
    if err != nil {
        panic(err)
    }
}
// ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBodyable 
type ApiConnectorsItemMicrosoftGraphUploadClientCertificateUploadClientCertificatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetPassword()(*string)
    GetPkcs12Value()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetPassword(value *string)()
    SetPkcs12Value(value *string)()
}
