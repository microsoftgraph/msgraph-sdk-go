package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AutomaticRepliesMailTips provides operations to call the getMailTips method.
type AutomaticRepliesMailTips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The automatic reply message.
    message *string;
    // The language that the automatic reply message is in.
    messageLanguage LocaleInfoable;
    // The date and time that automatic replies are set to end.
    scheduledEndTime DateTimeTimeZoneable;
    // The date and time that automatic replies are set to begin.
    scheduledStartTime DateTimeTimeZoneable;
}
// NewAutomaticRepliesMailTips instantiates a new automaticRepliesMailTips and sets the default values.
func NewAutomaticRepliesMailTips()(*AutomaticRepliesMailTips) {
    m := &AutomaticRepliesMailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAutomaticRepliesMailTipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAutomaticRepliesMailTipsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAutomaticRepliesMailTips(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomaticRepliesMailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AutomaticRepliesMailTips) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["messageLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["scheduledEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledEndTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["scheduledStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledStartTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The automatic reply message.
func (m *AutomaticRepliesMailTips) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetMessageLanguage gets the messageLanguage property value. The language that the automatic reply message is in.
func (m *AutomaticRepliesMailTips) GetMessageLanguage()(LocaleInfoable) {
    if m == nil {
        return nil
    } else {
        return m.messageLanguage
    }
}
// GetScheduledEndTime gets the scheduledEndTime property value. The date and time that automatic replies are set to end.
func (m *AutomaticRepliesMailTips) GetScheduledEndTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.scheduledEndTime
    }
}
// GetScheduledStartTime gets the scheduledStartTime property value. The date and time that automatic replies are set to begin.
func (m *AutomaticRepliesMailTips) GetScheduledStartTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.scheduledStartTime
    }
}
func (m *AutomaticRepliesMailTips) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AutomaticRepliesMailTips) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("messageLanguage", m.GetMessageLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledEndTime", m.GetScheduledEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledStartTime", m.GetScheduledStartTime())
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
func (m *AutomaticRepliesMailTips) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessage sets the message property value. The automatic reply message.
func (m *AutomaticRepliesMailTips) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetMessageLanguage sets the messageLanguage property value. The language that the automatic reply message is in.
func (m *AutomaticRepliesMailTips) SetMessageLanguage(value LocaleInfoable)() {
    if m != nil {
        m.messageLanguage = value
    }
}
// SetScheduledEndTime sets the scheduledEndTime property value. The date and time that automatic replies are set to end.
func (m *AutomaticRepliesMailTips) SetScheduledEndTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.scheduledEndTime = value
    }
}
// SetScheduledStartTime sets the scheduledStartTime property value. The date and time that automatic replies are set to begin.
func (m *AutomaticRepliesMailTips) SetScheduledStartTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.scheduledStartTime = value
    }
}
