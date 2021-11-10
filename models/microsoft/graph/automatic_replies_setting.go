package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AutomaticRepliesSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
    externalAudience *ExternalAudienceScope;
    // The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
    externalReplyMessage *string;
    // The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
    internalReplyMessage *string;
    // The date and time that automatic replies are set to end, if Status is set to Scheduled.
    scheduledEndDateTime *DateTimeTimeZone;
    // The date and time that automatic replies are set to begin, if Status is set to Scheduled.
    scheduledStartDateTime *DateTimeTimeZone;
    // Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
    status *AutomaticRepliesStatus;
}
// Instantiates a new automaticRepliesSetting and sets the default values.
func NewAutomaticRepliesSetting()(*AutomaticRepliesSetting) {
    m := &AutomaticRepliesSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomaticRepliesSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the externalAudience property value. The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
func (m *AutomaticRepliesSetting) GetExternalAudience()(*ExternalAudienceScope) {
    if m == nil {
        return nil
    } else {
        return m.externalAudience
    }
}
// Gets the externalReplyMessage property value. The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) GetExternalReplyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalReplyMessage
    }
}
// Gets the internalReplyMessage property value. The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) GetInternalReplyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalReplyMessage
    }
}
// Gets the scheduledEndDateTime property value. The date and time that automatic replies are set to end, if Status is set to Scheduled.
func (m *AutomaticRepliesSetting) GetScheduledEndDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledEndDateTime
    }
}
// Gets the scheduledStartDateTime property value. The date and time that automatic replies are set to begin, if Status is set to Scheduled.
func (m *AutomaticRepliesSetting) GetScheduledStartDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledStartDateTime
    }
}
// Gets the status property value. Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
func (m *AutomaticRepliesSetting) GetStatus()(*AutomaticRepliesStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *AutomaticRepliesSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["externalAudience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAudienceScope)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ExternalAudienceScope)
            m.SetExternalAudience(&cast)
        }
        return nil
    }
    res["externalReplyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalReplyMessage(val)
        }
        return nil
    }
    res["internalReplyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalReplyMessage(val)
        }
        return nil
    }
    res["scheduledEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledEndDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["scheduledStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledStartDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutomaticRepliesStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AutomaticRepliesStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *AutomaticRepliesSetting) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AutomaticRepliesSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetExternalAudience() != nil {
        cast := m.GetExternalAudience().String()
        err := writer.WriteStringValue("externalAudience", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalReplyMessage", m.GetExternalReplyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalReplyMessage", m.GetInternalReplyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledEndDateTime", m.GetScheduledEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledStartDateTime", m.GetScheduledStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *AutomaticRepliesSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the externalAudience property value. The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
// Parameters:
//  - value : Value to set for the externalAudience property.
func (m *AutomaticRepliesSetting) SetExternalAudience(value *ExternalAudienceScope)() {
    m.externalAudience = value
}
// Sets the externalReplyMessage property value. The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
// Parameters:
//  - value : Value to set for the externalReplyMessage property.
func (m *AutomaticRepliesSetting) SetExternalReplyMessage(value *string)() {
    m.externalReplyMessage = value
}
// Sets the internalReplyMessage property value. The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
// Parameters:
//  - value : Value to set for the internalReplyMessage property.
func (m *AutomaticRepliesSetting) SetInternalReplyMessage(value *string)() {
    m.internalReplyMessage = value
}
// Sets the scheduledEndDateTime property value. The date and time that automatic replies are set to end, if Status is set to Scheduled.
// Parameters:
//  - value : Value to set for the scheduledEndDateTime property.
func (m *AutomaticRepliesSetting) SetScheduledEndDateTime(value *DateTimeTimeZone)() {
    m.scheduledEndDateTime = value
}
// Sets the scheduledStartDateTime property value. The date and time that automatic replies are set to begin, if Status is set to Scheduled.
// Parameters:
//  - value : Value to set for the scheduledStartDateTime property.
func (m *AutomaticRepliesSetting) SetScheduledStartDateTime(value *DateTimeTimeZone)() {
    m.scheduledStartDateTime = value
}
// Sets the status property value. Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
// Parameters:
//  - value : Value to set for the status property.
func (m *AutomaticRepliesSetting) SetStatus(value *AutomaticRepliesStatus)() {
    m.status = value
}
