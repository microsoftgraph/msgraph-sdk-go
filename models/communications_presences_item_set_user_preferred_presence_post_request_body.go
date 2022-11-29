package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody provides operations to call the setUserPreferredPresence method.
type CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody struct {
    // The activity property
    activity *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The availability property
    availability *string
    // The expirationDuration property
    expirationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewCommunicationsPresencesItemSetUserPreferredPresencePostRequestBody instantiates a new CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody and sets the default values.
func NewCommunicationsPresencesItemSetUserPreferredPresencePostRequestBody()(*CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) {
    m := &CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsPresencesItemSetUserPreferredPresencePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsPresencesItemSetUserPreferredPresencePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsPresencesItemSetUserPreferredPresencePostRequestBody(), nil
}
// GetActivity gets the activity property value. The activity property
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) GetActivity()(*string) {
    return m.activity
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAvailability gets the availability property value. The availability property
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) GetAvailability()(*string) {
    return m.availability
}
// GetExpirationDuration gets the expirationDuration property value. The expirationDuration property
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) GetExpirationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.expirationDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivity)
    res["availability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAvailability)
    res["expirationDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetExpirationDuration)
    return res
}
// Serialize serializes information the current object
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("expirationDuration", m.GetExpirationDuration())
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
// SetActivity sets the activity property value. The activity property
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) SetActivity(value *string)() {
    m.activity = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAvailability sets the availability property value. The availability property
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) SetAvailability(value *string)() {
    m.availability = value
}
// SetExpirationDuration sets the expirationDuration property value. The expirationDuration property
func (m *CommunicationsPresencesItemSetUserPreferredPresencePostRequestBody) SetExpirationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.expirationDuration = value
}
