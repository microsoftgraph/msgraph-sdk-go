package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalsItemAddKeyPostRequestBody provides operations to call the addKey method.
type ServicePrincipalsItemAddKeyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The keyCredential property
    keyCredential KeyCredentialable
    // The passwordCredential property
    passwordCredential PasswordCredentialable
    // The proof property
    proof *string
}
// NewServicePrincipalsItemAddKeyPostRequestBody instantiates a new ServicePrincipalsItemAddKeyPostRequestBody and sets the default values.
func NewServicePrincipalsItemAddKeyPostRequestBody()(*ServicePrincipalsItemAddKeyPostRequestBody) {
    m := &ServicePrincipalsItemAddKeyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateServicePrincipalsItemAddKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalsItemAddKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalsItemAddKeyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePrincipalsItemAddKeyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalsItemAddKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredential(val.(KeyCredentialable))
        }
        return nil
    }
    res["passwordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(PasswordCredentialable))
        }
        return nil
    }
    res["proof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
// GetKeyCredential gets the keyCredential property value. The keyCredential property
func (m *ServicePrincipalsItemAddKeyPostRequestBody) GetKeyCredential()(KeyCredentialable) {
    return m.keyCredential
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *ServicePrincipalsItemAddKeyPostRequestBody) GetPasswordCredential()(PasswordCredentialable) {
    return m.passwordCredential
}
// GetProof gets the proof property value. The proof property
func (m *ServicePrincipalsItemAddKeyPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *ServicePrincipalsItemAddKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("keyCredential", m.GetKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
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
func (m *ServicePrincipalsItemAddKeyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeyCredential sets the keyCredential property value. The keyCredential property
func (m *ServicePrincipalsItemAddKeyPostRequestBody) SetKeyCredential(value KeyCredentialable)() {
    m.keyCredential = value
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *ServicePrincipalsItemAddKeyPostRequestBody) SetPasswordCredential(value PasswordCredentialable)() {
    m.passwordCredential = value
}
// SetProof sets the proof property value. The proof property
func (m *ServicePrincipalsItemAddKeyPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
