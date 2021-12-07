package getmailtips

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// GetMailTips 
type GetMailTips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Mail tips for automatic reply if it has been set up by the recipient.
    automaticReplies *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AutomaticRepliesMailTips;
    // A custom mail tip that can be set on the recipient's mailbox.
    customMailTip *string;
    // Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
    deliveryRestricted *bool;
    // The email address of the recipient to get mailtips for.
    emailAddress *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress;
    // Errors that occur during the getMailTips action.
    error *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsError;
    // The number of external members if the recipient is a distribution list.
    externalMemberCount *int32;
    // Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
    isModerated *bool;
    // The mailbox full status of the recipient.
    mailboxFull *bool;
    // The maximum message size that has been configured for the recipient's organization or mailbox.
    maxMessageSize *int32;
    // The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
    recipientScope *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RecipientScopeType;
    // Recipients suggested based on previous contexts where they appear in the same message.
    recipientSuggestions []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient;
    // The number of members if the recipient is a distribution list.
    totalMemberCount *int32;
}
// NewGetMailTips instantiates a new getMailTips and sets the default values.
func NewGetMailTips()(*GetMailTips) {
    m := &GetMailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetMailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAutomaticReplies gets the automaticReplies property value. Mail tips for automatic reply if it has been set up by the recipient.
func (m *GetMailTips) GetAutomaticReplies()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AutomaticRepliesMailTips) {
    if m == nil {
        return nil
    } else {
        return m.automaticReplies
    }
}
// GetCustomMailTip gets the customMailTip property value. A custom mail tip that can be set on the recipient's mailbox.
func (m *GetMailTips) GetCustomMailTip()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customMailTip
    }
}
// GetDeliveryRestricted gets the deliveryRestricted property value. Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
func (m *GetMailTips) GetDeliveryRestricted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deliveryRestricted
    }
}
// GetEmailAddress gets the emailAddress property value. The email address of the recipient to get mailtips for.
func (m *GetMailTips) GetEmailAddress()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetError gets the error property value. Errors that occur during the getMailTips action.
func (m *GetMailTips) GetError()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetExternalMemberCount gets the externalMemberCount property value. The number of external members if the recipient is a distribution list.
func (m *GetMailTips) GetExternalMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.externalMemberCount
    }
}
// GetIsModerated gets the isModerated property value. Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
func (m *GetMailTips) GetIsModerated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isModerated
    }
}
// GetMailboxFull gets the mailboxFull property value. The mailbox full status of the recipient.
func (m *GetMailTips) GetMailboxFull()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailboxFull
    }
}
// GetMaxMessageSize gets the maxMessageSize property value. The maximum message size that has been configured for the recipient's organization or mailbox.
func (m *GetMailTips) GetMaxMessageSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxMessageSize
    }
}
// GetRecipientScope gets the recipientScope property value. The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
func (m *GetMailTips) GetRecipientScope()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RecipientScopeType) {
    if m == nil {
        return nil
    } else {
        return m.recipientScope
    }
}
// GetRecipientSuggestions gets the recipientSuggestions property value. Recipients suggested based on previous contexts where they appear in the same message.
func (m *GetMailTips) GetRecipientSuggestions()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.recipientSuggestions
    }
}
// GetTotalMemberCount gets the totalMemberCount property value. The number of members if the recipient is a distribution list.
func (m *GetMailTips) GetTotalMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalMemberCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetMailTips) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["automaticReplies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAutomaticRepliesMailTips() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticReplies(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AutomaticRepliesMailTips))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress))
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMailTipsError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsError))
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
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseRecipientScopeType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RecipientScopeType)
            m.SetRecipientScope(&cast)
        }
        return nil
    }
    res["recipientSuggestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient))
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
func (m *GetMailTips) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetMailTips) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetRecipientScope().String()
        err := writer.WriteStringValue("recipientScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecipientSuggestions()))
        for i, v := range m.GetRecipientSuggestions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *GetMailTips) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAutomaticReplies sets the automaticReplies property value. Mail tips for automatic reply if it has been set up by the recipient.
func (m *GetMailTips) SetAutomaticReplies(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AutomaticRepliesMailTips)() {
    if m != nil {
        m.automaticReplies = value
    }
}
// SetCustomMailTip sets the customMailTip property value. A custom mail tip that can be set on the recipient's mailbox.
func (m *GetMailTips) SetCustomMailTip(value *string)() {
    if m != nil {
        m.customMailTip = value
    }
}
// SetDeliveryRestricted sets the deliveryRestricted property value. Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
func (m *GetMailTips) SetDeliveryRestricted(value *bool)() {
    if m != nil {
        m.deliveryRestricted = value
    }
}
// SetEmailAddress sets the emailAddress property value. The email address of the recipient to get mailtips for.
func (m *GetMailTips) SetEmailAddress(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EmailAddress)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetError sets the error property value. Errors that occur during the getMailTips action.
func (m *GetMailTips) SetError(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MailTipsError)() {
    if m != nil {
        m.error = value
    }
}
// SetExternalMemberCount sets the externalMemberCount property value. The number of external members if the recipient is a distribution list.
func (m *GetMailTips) SetExternalMemberCount(value *int32)() {
    if m != nil {
        m.externalMemberCount = value
    }
}
// SetIsModerated sets the isModerated property value. Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
func (m *GetMailTips) SetIsModerated(value *bool)() {
    if m != nil {
        m.isModerated = value
    }
}
// SetMailboxFull sets the mailboxFull property value. The mailbox full status of the recipient.
func (m *GetMailTips) SetMailboxFull(value *bool)() {
    if m != nil {
        m.mailboxFull = value
    }
}
// SetMaxMessageSize sets the maxMessageSize property value. The maximum message size that has been configured for the recipient's organization or mailbox.
func (m *GetMailTips) SetMaxMessageSize(value *int32)() {
    if m != nil {
        m.maxMessageSize = value
    }
}
// SetRecipientScope sets the recipientScope property value. The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
func (m *GetMailTips) SetRecipientScope(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.RecipientScopeType)() {
    if m != nil {
        m.recipientScope = value
    }
}
// SetRecipientSuggestions sets the recipientSuggestions property value. Recipients suggested based on previous contexts where they appear in the same message.
func (m *GetMailTips) SetRecipientSuggestions(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient)() {
    if m != nil {
        m.recipientSuggestions = value
    }
}
// SetTotalMemberCount sets the totalMemberCount property value. The number of members if the recipient is a distribution list.
func (m *GetMailTips) SetTotalMemberCount(value *int32)() {
    if m != nil {
        m.totalMemberCount = value
    }
}
