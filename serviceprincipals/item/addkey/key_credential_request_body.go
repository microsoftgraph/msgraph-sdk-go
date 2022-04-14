package addkey

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// KeyCredentialRequestBody provides operations to call the addKey method.
type KeyCredentialRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The keyCredential property
    keyCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable
    // The passwordCredential property
    passwordCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable
    // The proof property
    proof *string
}
// NewKeyCredentialRequestBody instantiates a new KeyCredentialRequestBody and sets the default values.
func NewKeyCredentialRequestBody()(*KeyCredentialRequestBody) {
    m := &KeyCredentialRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKeyCredentialRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyCredentialRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyCredentialRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyCredentialRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyCredentialRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *KeyCredentialRequestBody) GetKeyCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.keyCredential
    }
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *KeyCredentialRequestBody) GetPasswordCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredential
    }
}
// GetProof gets the proof property value. The proof property
func (m *KeyCredentialRequestBody) GetProof()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proof
    }
}
// Serialize serializes information the current object
func (m *KeyCredentialRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *KeyCredentialRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeyCredential sets the keyCredential property value. The keyCredential property
func (m *KeyCredentialRequestBody) SetKeyCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable)() {
    if m != nil {
        m.keyCredential = value
    }
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *KeyCredentialRequestBody) SetPasswordCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable)() {
    if m != nil {
        m.passwordCredential = value
    }
}
// SetProof sets the proof property value. The proof property
func (m *KeyCredentialRequestBody) SetProof(value *string)() {
    if m != nil {
        m.proof = value
    }
}
