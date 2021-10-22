package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrivacyProfile struct {
    additionalData map[string]interface{};
    contactEmail *string;
    statementUrl *string;
}
func NewPrivacyProfile()(*PrivacyProfile) {
    m := &PrivacyProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrivacyProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrivacyProfile) GetContactEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactEmail
    }
}
func (m *PrivacyProfile) GetStatementUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statementUrl
    }
}
func (m *PrivacyProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contactEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContactEmail(val)
        return nil
    }
    res["statementUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatementUrl(val)
        return nil
    }
    return res
}
func (m *PrivacyProfile) IsNil()(bool) {
    return m == nil
}
func (m *PrivacyProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contactEmail", m.GetContactEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("statementUrl", m.GetStatementUrl())
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
func (m *PrivacyProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrivacyProfile) SetContactEmail(value *string)() {
    m.contactEmail = value
}
func (m *PrivacyProfile) SetStatementUrl(value *string)() {
    m.statementUrl = value
}
