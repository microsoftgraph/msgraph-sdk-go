package addkey

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// KeyCredentialPostRequestBody provides operations to call the addKey method.
type KeyCredentialPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The keyCredential property
    keyCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable
    // The passwordCredential property
    passwordCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable
    // The proof property
    proof *string
}
// NewKeyCredentialPostRequestBody instantiates a new KeyCredentialPostRequestBody and sets the default values.
func NewKeyCredentialPostRequestBody()(*KeyCredentialPostRequestBody) {
    m := &KeyCredentialPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKeyCredentialPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyCredentialPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyCredentialPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredentialPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyCredentialPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredential(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable))
        }
        return nil
    }
    res["passwordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable))
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
func (m *KeyCredentialPostRequestBody) GetKeyCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable) {
    return m.keyCredential
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *KeyCredentialPostRequestBody) GetPasswordCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable) {
    return m.passwordCredential
}
// GetProof gets the proof property value. The proof property
func (m *KeyCredentialPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *KeyCredentialPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *KeyCredentialPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeyCredential sets the keyCredential property value. The keyCredential property
func (m *KeyCredentialPostRequestBody) SetKeyCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable)() {
    m.keyCredential = value
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *KeyCredentialPostRequestBody) SetPasswordCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable)() {
    m.passwordCredential = value
}
// SetProof sets the proof property value. The proof property
func (m *KeyCredentialPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
