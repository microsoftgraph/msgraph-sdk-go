package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemPresenceClearPresencePostRequestBody provides operations to call the clearPresence method.
type UsersItemPresenceClearPresencePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The sessionId property
    sessionId *string
}
// NewUsersItemPresenceClearPresencePostRequestBody instantiates a new UsersItemPresenceClearPresencePostRequestBody and sets the default values.
func NewUsersItemPresenceClearPresencePostRequestBody()(*UsersItemPresenceClearPresencePostRequestBody) {
    m := &UsersItemPresenceClearPresencePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemPresenceClearPresencePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemPresenceClearPresencePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemPresenceClearPresencePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemPresenceClearPresencePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemPresenceClearPresencePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sessionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    return res
}
// GetSessionId gets the sessionId property value. The sessionId property
func (m *UsersItemPresenceClearPresencePostRequestBody) GetSessionId()(*string) {
    return m.sessionId
}
// Serialize serializes information the current object
func (m *UsersItemPresenceClearPresencePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
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
func (m *UsersItemPresenceClearPresencePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *UsersItemPresenceClearPresencePostRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
