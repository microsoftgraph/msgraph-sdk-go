package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MailTips provides operations to call the getMailTips method.
type MailTips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Mail tips for automatic reply if it has been set up by the recipient.
    automaticReplies AutomaticRepliesMailTipsable;
    // A custom mail tip that can be set on the recipient's mailbox.
    customMailTip *string;
    // Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
    deliveryRestricted *bool;
    // The email address of the recipient to get mailtips for.
    emailAddress EmailAddressable;
    // Errors that occur during the getMailTips action.
    error MailTipsErrorable;
    // The number of external members if the recipient is a distribution list.
    externalMemberCount *int32;
    // Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
    isModerated *bool;
    // The mailbox full status of the recipient.
    mailboxFull *bool;
    // The maximum message size that has been configured for the recipient's organization or mailbox.
    maxMessageSize *int32;
    // The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
    recipientScope *RecipientScopeType;
    // Recipients suggested based on previous contexts where they appear in the same message.
    recipientSuggestions []Recipientable;
    // The number of members if the recipient is a distribution list.
    totalMemberCount *int32;
}
// NewMailTips instantiates a new mailTips and sets the default values.
func NewMailTips()(*MailTips) {
    m := &MailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMailTipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailTipsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMailTips(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAutomaticReplies gets the automaticReplies property value. Mail tips for automatic reply if it has been set up by the recipient.
func (m *MailTips) GetAutomaticReplies()(AutomaticRepliesMailTipsable) {
    if m == nil {
        return nil
    } else {
        return m.automaticReplies
    }
}
// GetCustomMailTip gets the customMailTip property value. A custom mail tip that can be set on the recipient's mailbox.
func (m *MailTips) GetCustomMailTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customMailTip
    }
}
// GetDeliveryRestricted gets the deliveryRestricted property value. Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
func (m *MailTips) GetDeliveryRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deliveryRestricted
    }
}
// GetEmailAddress gets the emailAddress property value. The email address of the recipient to get mailtips for.
func (m *MailTips) GetEmailAddress()(EmailAddressable) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetError gets the error property value. Errors that occur during the getMailTips action.
func (m *MailTips) GetError()(MailTipsErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetExternalMemberCount gets the externalMemberCount property value. The number of external members if the recipient is a distribution list.
func (m *MailTips) GetExternalMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.externalMemberCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailTips) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["automaticReplies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutomaticRepliesMailTipsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticReplies(val.(AutomaticRepliesMailTipsable))
        }
        return nil
    }
    res["customMailTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomMailTip(val)
        }
        return nil
    }
    res["deliveryRestricted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryRestricted(val)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val.(EmailAddressable))
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateMailTipsErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(MailTipsErrorable))
        }
        return nil
    }
    res["externalMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalMemberCount(val)
        }
        return nil
    }
    res["isModerated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsModerated(val)
        }
        return nil
    }
    res["mailboxFull"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxFull(val)
        }
        return nil
    }
    res["maxMessageSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxMessageSize(val)
        }
        return nil
    }
    res["recipientScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecipientScopeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientScope(val.(*RecipientScopeType))
        }
        return nil
    }
    res["recipientSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetRecipientSuggestions(res)
        }
        return nil
    }
    res["totalMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMemberCount(val)
        }
        return nil
    }
    return res
}
// GetIsModerated gets the isModerated property value. Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
func (m *MailTips) GetIsModerated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isModerated
    }
}
// GetMailboxFull gets the mailboxFull property value. The mailbox full status of the recipient.
func (m *MailTips) GetMailboxFull()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailboxFull
    }
}
// GetMaxMessageSize gets the maxMessageSize property value. The maximum message size that has been configured for the recipient's organization or mailbox.
func (m *MailTips) GetMaxMessageSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxMessageSize
    }
}
// GetRecipientScope gets the recipientScope property value. The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
func (m *MailTips) GetRecipientScope()(*RecipientScopeType) {
    if m == nil {
        return nil
    } else {
        return m.recipientScope
    }
}
// GetRecipientSuggestions gets the recipientSuggestions property value. Recipients suggested based on previous contexts where they appear in the same message.
func (m *MailTips) GetRecipientSuggestions()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.recipientSuggestions
    }
}
// GetTotalMemberCount gets the totalMemberCount property value. The number of members if the recipient is a distribution list.
func (m *MailTips) GetTotalMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalMemberCount
    }
}
func (m *MailTips) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MailTips) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automaticReplies", m.GetAutomaticReplies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customMailTip", m.GetCustomMailTip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deliveryRestricted", m.GetDeliveryRestricted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("externalMemberCount", m.GetExternalMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isModerated", m.GetIsModerated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailboxFull", m.GetMailboxFull())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxMessageSize", m.GetMaxMessageSize())
        if err != nil {
            return err
        }
    }
    if m.GetRecipientScope() != nil {
        cast := (*m.GetRecipientScope()).String()
        err := writer.WriteStringValue("recipientScope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecipientSuggestions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecipientSuggestions()))
        for i, v := range m.GetRecipientSuggestions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipientSuggestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalMemberCount", m.GetTotalMemberCount())
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
func (m *MailTips) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAutomaticReplies sets the automaticReplies property value. Mail tips for automatic reply if it has been set up by the recipient.
func (m *MailTips) SetAutomaticReplies(value AutomaticRepliesMailTipsable)() {
    if m != nil {
        m.automaticReplies = value
    }
}
// SetCustomMailTip sets the customMailTip property value. A custom mail tip that can be set on the recipient's mailbox.
func (m *MailTips) SetCustomMailTip(value *string)() {
    if m != nil {
        m.customMailTip = value
    }
}
// SetDeliveryRestricted sets the deliveryRestricted property value. Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
func (m *MailTips) SetDeliveryRestricted(value *bool)() {
    if m != nil {
        m.deliveryRestricted = value
    }
}
// SetEmailAddress sets the emailAddress property value. The email address of the recipient to get mailtips for.
func (m *MailTips) SetEmailAddress(value EmailAddressable)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetError sets the error property value. Errors that occur during the getMailTips action.
func (m *MailTips) SetError(value MailTipsErrorable)() {
    if m != nil {
        m.error = value
    }
}
// SetExternalMemberCount sets the externalMemberCount property value. The number of external members if the recipient is a distribution list.
func (m *MailTips) SetExternalMemberCount(value *int32)() {
    if m != nil {
        m.externalMemberCount = value
    }
}
// SetIsModerated sets the isModerated property value. Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
func (m *MailTips) SetIsModerated(value *bool)() {
    if m != nil {
        m.isModerated = value
    }
}
// SetMailboxFull sets the mailboxFull property value. The mailbox full status of the recipient.
func (m *MailTips) SetMailboxFull(value *bool)() {
    if m != nil {
        m.mailboxFull = value
    }
}
// SetMaxMessageSize sets the maxMessageSize property value. The maximum message size that has been configured for the recipient's organization or mailbox.
func (m *MailTips) SetMaxMessageSize(value *int32)() {
    if m != nil {
        m.maxMessageSize = value
    }
}
// SetRecipientScope sets the recipientScope property value. The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
func (m *MailTips) SetRecipientScope(value *RecipientScopeType)() {
    if m != nil {
        m.recipientScope = value
    }
}
// SetRecipientSuggestions sets the recipientSuggestions property value. Recipients suggested based on previous contexts where they appear in the same message.
func (m *MailTips) SetRecipientSuggestions(value []Recipientable)() {
    if m != nil {
        m.recipientSuggestions = value
    }
}
// SetTotalMemberCount sets the totalMemberCount property value. The number of members if the recipient is a distribution list.
func (m *MailTips) SetTotalMemberCount(value *int32)() {
    if m != nil {
        m.totalMemberCount = value
    }
}
