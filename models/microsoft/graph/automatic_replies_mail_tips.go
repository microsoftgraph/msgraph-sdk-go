package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AutomaticRepliesMailTips struct {
    additionalData map[string]interface{};
    message *string;
    messageLanguage *LocaleInfo;
    scheduledEndTime *DateTimeTimeZone;
    scheduledStartTime *DateTimeTimeZone;
}
func NewAutomaticRepliesMailTips()(*AutomaticRepliesMailTips) {
    m := &AutomaticRepliesMailTips{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AutomaticRepliesMailTips) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AutomaticRepliesMailTips) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *AutomaticRepliesMailTips) GetMessageLanguage()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.messageLanguage
    }
}
func (m *AutomaticRepliesMailTips) GetScheduledEndTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledEndTime
    }
}
func (m *AutomaticRepliesMailTips) GetScheduledStartTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.scheduledStartTime
    }
}
func (m *AutomaticRepliesMailTips) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["messageLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        m.SetMessageLanguage(val.(*LocaleInfo))
        return nil
    }
    res["scheduledEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetScheduledEndTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["scheduledStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetScheduledStartTime(val.(*DateTimeTimeZone))
        return nil
    }
    return res
}
func (m *AutomaticRepliesMailTips) IsNil()(bool) {
    return m == nil
}
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
func (m *AutomaticRepliesMailTips) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AutomaticRepliesMailTips) SetMessage(value *string)() {
    m.message = value
}
func (m *AutomaticRepliesMailTips) SetMessageLanguage(value *LocaleInfo)() {
    m.messageLanguage = value
}
func (m *AutomaticRepliesMailTips) SetScheduledEndTime(value *DateTimeTimeZone)() {
    m.scheduledEndTime = value
}
func (m *AutomaticRepliesMailTips) SetScheduledStartTime(value *DateTimeTimeZone)() {
    m.scheduledStartTime = value
}
