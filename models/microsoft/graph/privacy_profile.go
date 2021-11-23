package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivacyProfile 
type PrivacyProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A valid smtp email address for the privacy statement contact. Not required.
    contactEmail *string;
    // A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
    statementUrl *string;
}
// NewPrivacyProfile instantiates a new privacyProfile and sets the default values.
func NewPrivacyProfile()(*PrivacyProfile) {
    m := &PrivacyProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivacyProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContactEmail gets the contactEmail property value. A valid smtp email address for the privacy statement contact. Not required.
func (m *PrivacyProfile) GetContactEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactEmail
    }
}
// GetStatementUrl gets the statementUrl property value. A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
func (m *PrivacyProfile) GetStatementUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statementUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivacyProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contactEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactEmail(val)
        }
        return nil
    }
    res["statementUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatementUrl(val)
        }
        return nil
    }
    return res
}
func (m *PrivacyProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrivacyProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContactEmail sets the contactEmail property value. A valid smtp email address for the privacy statement contact. Not required.
func (m *PrivacyProfile) SetContactEmail(value *string)() {
    m.contactEmail = value
}
// SetStatementUrl sets the statementUrl property value. A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company's privacy statement. Not required.
func (m *PrivacyProfile) SetStatementUrl(value *string)() {
    m.statementUrl = value
}
