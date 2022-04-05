package setpresence

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SetPresenceRequestBody provides operations to call the setPresence method.
type SetPresenceRequestBody struct {
    // The activity property
    activity *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The availability property
    availability *string;
    // The expirationDuration property
    expirationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration;
    // The sessionId property
    sessionId *string;
}
// NewSetPresenceRequestBody instantiates a new setPresenceRequestBody and sets the default values.
func NewSetPresenceRequestBody()(*SetPresenceRequestBody) {
    m := &SetPresenceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSetPresenceRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetPresenceRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetPresenceRequestBody(), nil
}
// GetActivity gets the activity property value. The activity property
func (m *SetPresenceRequestBody) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetPresenceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAvailability gets the availability property value. The availability property
func (m *SetPresenceRequestBody) GetAvailability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
// GetExpirationDuration gets the expirationDuration property value. The expirationDuration property
func (m *SetPresenceRequestBody) GetExpirationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.expirationDuration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetPresenceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activity"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["availability"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailability(val)
        }
        return nil
    }
    res["expirationDuration"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDuration(val)
        }
        return nil
    }
    res["sessionId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *SetPresenceRequestBody) GetSessionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionId
    }
}
// Serialize serializes information the current object
func (m *SetPresenceRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetActivity sets the activity property value. The activity property
func (m *SetPresenceRequestBody) SetActivity(value *string)() {
    if m != nil {
        m.activity = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetPresenceRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAvailability sets the availability property value. The availability property
func (m *SetPresenceRequestBody) SetAvailability(value *string)() {
    if m != nil {
        m.availability = value
    }
}
// SetExpirationDuration sets the expirationDuration property value. The expirationDuration property
func (m *SetPresenceRequestBody) SetExpirationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.expirationDuration = value
    }
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *SetPresenceRequestBody) SetSessionId(value *string)() {
    if m != nil {
        m.sessionId = value
    }
}
