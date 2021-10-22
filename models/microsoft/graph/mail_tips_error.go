package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MailTipsError struct {
    additionalData map[string]interface{};
    code *string;
    message *string;
}
func NewMailTipsError()(*MailTipsError) {
    m := &MailTipsError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MailTipsError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MailTipsError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
func (m *MailTipsError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *MailTipsError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    return res
}
func (m *MailTipsError) IsNil()(bool) {
    return m == nil
}
func (m *MailTipsError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
func (m *MailTipsError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MailTipsError) SetCode(value *string)() {
    m.code = value
}
func (m *MailTipsError) SetMessage(value *string)() {
    m.message = value
}
