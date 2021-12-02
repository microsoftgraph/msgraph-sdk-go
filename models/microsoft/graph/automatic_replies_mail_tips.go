package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AutomaticRepliesMailTips 
type AutomaticRepliesMailTips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The automatic reply message.
    message *string;
    // The language that the automatic reply message is in.
    messageLanguage *LocaleInfo;
    // The date and time that automatic replies are set to end.
    scheduledEndTime *DateTimeTimeZone;
    // The date and time that automatic replies are set to begin.
    scheduledStartTime *DateTimeTimeZone;
}
// NewAutomaticRepliesMailTips instantiates a new automaticRepliesMailTips and sets the default values.
func NewAutomaticRepliesMailTips()(*AutomaticRepliesMailTips) {
    m := &AutomaticRepliesMailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomaticRepliesMailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
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
func (m *AutomaticRepliesMailTips) GetMessageLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.messageLanguage
    }
}
// GetScheduledEndTime gets the scheduledEndTime property value. The date and time that automatic replies are set to end.
func (m *AutomaticRepliesMailTips) GetScheduledEndTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledEndTime
    }
}
// GetScheduledStartTime gets the scheduledStartTime property value. The date and time that automatic replies are set to begin.
func (m *AutomaticRepliesMailTips) GetScheduledStartTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledStartTime
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageLanguage(val.(*LocaleInfo))
        }
        return nil
    }
    res["scheduledEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledEndTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["scheduledStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledStartTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    return res
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
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
func (m *AutomaticRepliesMailTips) SetMessageLanguage(value *LocaleInfo)() {
    if m != nil {
        m.messageLanguage = value
    }
}
// SetScheduledEndTime sets the scheduledEndTime property value. The date and time that automatic replies are set to end.
func (m *AutomaticRepliesMailTips) SetScheduledEndTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.scheduledEndTime = value
    }
}
// SetScheduledStartTime sets the scheduledStartTime property value. The date and time that automatic replies are set to begin.
func (m *AutomaticRepliesMailTips) SetScheduledStartTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.scheduledStartTime = value
    }
}
