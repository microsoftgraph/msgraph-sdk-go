package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody provides operations to call the move method.
type MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The DestinationId property
    destinationId *string
}
// NewMeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody instantiates a new MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody and sets the default values.
func NewMeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody()(*MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) {
    m := &MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDestinationId gets the destinationId property value. The DestinationId property
func (m *MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) GetDestinationId()(*string) {
    return m.destinationId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationId", m.GetDestinationId())
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
func (m *MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDestinationId sets the destinationId property value. The DestinationId property
func (m *MeMailFoldersItemChildFoldersItemMessagesItemMovePostRequestBody) SetDestinationId(value *string)() {
    m.destinationId = value
}
