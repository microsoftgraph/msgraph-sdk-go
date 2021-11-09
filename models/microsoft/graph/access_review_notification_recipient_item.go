package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessReviewNotificationRecipientItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Determines the recipient of the notification email.
    notificationRecipientScope *AccessReviewNotificationRecipientScope;
    // Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients.
    notificationTemplateType *string;
}
// Instantiates a new accessReviewNotificationRecipientItem and sets the default values.
func NewAccessReviewNotificationRecipientItem()(*AccessReviewNotificationRecipientItem) {
    m := &AccessReviewNotificationRecipientItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewNotificationRecipientItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the notificationRecipientScope property value. Determines the recipient of the notification email.
func (m *AccessReviewNotificationRecipientItem) GetNotificationRecipientScope()(*AccessReviewNotificationRecipientScope) {
    if m == nil {
        return nil
    } else {
        return m.notificationRecipientScope
    }
}
// Gets the notificationTemplateType property value. Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients.
func (m *AccessReviewNotificationRecipientItem) GetNotificationTemplateType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateType
    }
}
// The deserialization information for the current model
func (m *AccessReviewNotificationRecipientItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notificationRecipientScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewNotificationRecipientScope() })
        if err != nil {
            return err
        }
        m.SetNotificationRecipientScope(val.(*AccessReviewNotificationRecipientScope))
        return nil
    }
    res["notificationTemplateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationTemplateType(val)
        return nil
    }
    return res
}
func (m *AccessReviewNotificationRecipientItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessReviewNotificationRecipientItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("notificationRecipientScope", m.GetNotificationRecipientScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTemplateType", m.GetNotificationTemplateType())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AccessReviewNotificationRecipientItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the notificationRecipientScope property value. Determines the recipient of the notification email.
// Parameters:
//  - value : Value to set for the notificationRecipientScope property.
func (m *AccessReviewNotificationRecipientItem) SetNotificationRecipientScope(value *AccessReviewNotificationRecipientScope)() {
    m.notificationRecipientScope = value
}
// Sets the notificationTemplateType property value. Indicates the type of access review email to be sent. Supported template type is CompletedAdditionalRecipients, which sends review completion notifications to the recipients.
// Parameters:
//  - value : Value to set for the notificationTemplateType property.
func (m *AccessReviewNotificationRecipientItem) SetNotificationTemplateType(value *string)() {
    m.notificationTemplateType = value
}
