package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AutomaticRepliesSetting struct {
    additionalData map[string]interface{};
    externalAudience *ExternalAudienceScope;
    externalReplyMessage *string;
    internalReplyMessage *string;
    scheduledEndDateTime *DateTimeTimeZone;
    scheduledStartDateTime *DateTimeTimeZone;
    status *AutomaticRepliesStatus;
}
func NewAutomaticRepliesSetting()(*AutomaticRepliesSetting) {
    m := &AutomaticRepliesSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AutomaticRepliesSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AutomaticRepliesSetting) GetExternalAudience()(*ExternalAudienceScope) {
    if m == nil {
        return nil
    } else {
        return m.externalAudience
    }
}
func (m *AutomaticRepliesSetting) GetExternalReplyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalReplyMessage
    }
}
func (m *AutomaticRepliesSetting) GetInternalReplyMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalReplyMessage
    }
}
func (m *AutomaticRepliesSetting) GetScheduledEndDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledEndDateTime
    }
}
func (m *AutomaticRepliesSetting) GetScheduledStartDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledStartDateTime
    }
}
func (m *AutomaticRepliesSetting) GetStatus()(*AutomaticRepliesStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *AutomaticRepliesSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["externalAudience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAudienceScope)
        if err != nil {
            return err
        }
        cast := val.(ExternalAudienceScope)
        m.SetExternalAudience(&cast)
        return nil
    }
    res["externalReplyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalReplyMessage(val)
        return nil
    }
    res["internalReplyMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternalReplyMessage(val)
        return nil
    }
    res["scheduledEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetScheduledEndDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["scheduledStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetScheduledStartDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutomaticRepliesStatus)
        if err != nil {
            return err
        }
        cast := val.(AutomaticRepliesStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *AutomaticRepliesSetting) IsNil()(bool) {
    return m == nil
}
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
func (m *AutomaticRepliesSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AutomaticRepliesSetting) SetExternalAudience(value *ExternalAudienceScope)() {
    m.externalAudience = value
}
func (m *AutomaticRepliesSetting) SetExternalReplyMessage(value *string)() {
    m.externalReplyMessage = value
}
func (m *AutomaticRepliesSetting) SetInternalReplyMessage(value *string)() {
    m.internalReplyMessage = value
}
func (m *AutomaticRepliesSetting) SetScheduledEndDateTime(value *DateTimeTimeZone)() {
    m.scheduledEndDateTime = value
}
func (m *AutomaticRepliesSetting) SetScheduledStartDateTime(value *DateTimeTimeZone)() {
    m.scheduledStartDateTime = value
}
func (m *AutomaticRepliesSetting) SetStatus(value *AutomaticRepliesStatus)() {
    m.status = value
}
