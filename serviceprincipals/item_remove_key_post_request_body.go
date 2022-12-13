package serviceprincipals

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemRemoveKeyPostRequestBody provides operations to call the removeKey method.
type ItemRemoveKeyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The keyId property
    keyId *UUID
    // The proof property
    proof *string
}
// NewItemRemoveKeyPostRequestBody instantiates a new ItemRemoveKeyPostRequestBody and sets the default values.
func NewItemRemoveKeyPostRequestBody()(*ItemRemoveKeyPostRequestBody) {
    m := &ItemRemoveKeyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemRemoveKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemRemoveKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRemoveKeyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemRemoveKeyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemRemoveKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
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
// GetKeyId gets the keyId property value. The keyId property
func (m *ItemRemoveKeyPostRequestBody) GetKeyId()(*UUID) {
    return m.keyId
}
// GetProof gets the proof property value. The proof property
func (m *ItemRemoveKeyPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *ItemRemoveKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("keyId", m.GetKeyId())
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
func (m *ItemRemoveKeyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeyId sets the keyId property value. The keyId property
func (m *ItemRemoveKeyPostRequestBody) SetKeyId(value *UUID)() {
    m.keyId = value
}
// SetProof sets the proof property value. The proof property
func (m *ItemRemoveKeyPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
