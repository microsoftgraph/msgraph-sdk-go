package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InvitedUserMessageInfo 
type InvitedUserMessageInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Additional recipients the invitation message should be sent to. Currently only 1 additional recipient is supported.
    ccRecipients []Recipientable;
    // Customized message body you want to send if you don't want the default message.
    customizedMessageBody *string;
    // The language you want to send the default message in. If the customizedMessageBody is specified, this property is ignored, and the message is sent using the customizedMessageBody. The language format should be in ISO 639. The default is en-US.
    messageLanguage *string;
}
// NewInvitedUserMessageInfo instantiates a new invitedUserMessageInfo and sets the default values.
func NewInvitedUserMessageInfo()(*InvitedUserMessageInfo) {
    m := &InvitedUserMessageInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInvitedUserMessageInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInvitedUserMessageInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInvitedUserMessageInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InvitedUserMessageInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCcRecipients gets the ccRecipients property value. Additional recipients the invitation message should be sent to. Currently only 1 additional recipient is supported.
func (m *InvitedUserMessageInfo) GetCcRecipients()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.ccRecipients
    }
}
// GetCustomizedMessageBody gets the customizedMessageBody property value. Customized message body you want to send if you don't want the default message.
func (m *InvitedUserMessageInfo) GetCustomizedMessageBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedMessageBody
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InvitedUserMessageInfo) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ccRecipients"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetCcRecipients(res)
        }
        return nil
    }
    res["customizedMessageBody"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedMessageBody(val)
        }
        return nil
    }
    res["messageLanguage"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageLanguage(val)
        }
        return nil
    }
    return res
}
// GetMessageLanguage gets the messageLanguage property value. The language you want to send the default message in. If the customizedMessageBody is specified, this property is ignored, and the message is sent using the customizedMessageBody. The language format should be in ISO 639. The default is en-US.
func (m *InvitedUserMessageInfo) GetMessageLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageLanguage
    }
}
// Serialize serializes information the current object
func (m *InvitedUserMessageInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCcRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("ccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customizedMessageBody", m.GetCustomizedMessageBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageLanguage", m.GetMessageLanguage())
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
func (m *InvitedUserMessageInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCcRecipients sets the ccRecipients property value. Additional recipients the invitation message should be sent to. Currently only 1 additional recipient is supported.
func (m *InvitedUserMessageInfo) SetCcRecipients(value []Recipientable)() {
    if m != nil {
        m.ccRecipients = value
    }
}
// SetCustomizedMessageBody sets the customizedMessageBody property value. Customized message body you want to send if you don't want the default message.
func (m *InvitedUserMessageInfo) SetCustomizedMessageBody(value *string)() {
    if m != nil {
        m.customizedMessageBody = value
    }
}
// SetMessageLanguage sets the messageLanguage property value. The language you want to send the default message in. If the customizedMessageBody is specified, this property is ignored, and the message is sent using the customizedMessageBody. The language format should be in ISO 639. The default is en-US.
func (m *InvitedUserMessageInfo) SetMessageLanguage(value *string)() {
    if m != nil {
        m.messageLanguage = value
    }
}
