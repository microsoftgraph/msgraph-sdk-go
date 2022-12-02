package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemGetMailTipsPostRequestBody provides operations to call the getMailTips method.
type UsersItemGetMailTipsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The EmailAddresses property
    emailAddresses []string
    // The MailTipsOptions property
    mailTipsOptions *MailTipsType
}
// NewUsersItemGetMailTipsPostRequestBody instantiates a new UsersItemGetMailTipsPostRequestBody and sets the default values.
func NewUsersItemGetMailTipsPostRequestBody()(*UsersItemGetMailTipsPostRequestBody) {
    m := &UsersItemGetMailTipsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemGetMailTipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemGetMailTipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemGetMailTipsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemGetMailTipsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEmailAddresses gets the emailAddresses property value. The EmailAddresses property
func (m *UsersItemGetMailTipsPostRequestBody) GetEmailAddresses()([]string) {
    return m.emailAddresses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemGetMailTipsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["mailTipsOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMailTipsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailTipsOptions(val.(*MailTipsType))
        }
        return nil
    }
    return res
}
// GetMailTipsOptions gets the mailTipsOptions property value. The MailTipsOptions property
func (m *UsersItemGetMailTipsPostRequestBody) GetMailTipsOptions()(*MailTipsType) {
    return m.mailTipsOptions
}
// Serialize serializes information the current object
func (m *UsersItemGetMailTipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("emailAddresses", m.GetEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetMailTipsOptions() != nil {
        cast := (*m.GetMailTipsOptions()).String()
        err := writer.WriteStringValue("mailTipsOptions", &cast)
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
func (m *UsersItemGetMailTipsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEmailAddresses sets the emailAddresses property value. The EmailAddresses property
func (m *UsersItemGetMailTipsPostRequestBody) SetEmailAddresses(value []string)() {
    m.emailAddresses = value
}
// SetMailTipsOptions sets the mailTipsOptions property value. The MailTipsOptions property
func (m *UsersItemGetMailTipsPostRequestBody) SetMailTipsOptions(value *MailTipsType)() {
    m.mailTipsOptions = value
}
