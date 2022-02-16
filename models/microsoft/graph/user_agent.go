package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserAgent 
type UserAgent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identifies the version of application software used by this endpoint.
    applicationVersion *string;
    // User-agent header value reported by this endpoint.
    headerValue *string;
}
// NewUserAgent instantiates a new userAgent and sets the default values.
func NewUserAgent()(*UserAgent) {
    m := &UserAgent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAgent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationVersion gets the applicationVersion property value. Identifies the version of application software used by this endpoint.
func (m *UserAgent) GetApplicationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationVersion
    }
}
// GetHeaderValue gets the headerValue property value. User-agent header value reported by this endpoint.
func (m *UserAgent) GetHeaderValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.headerValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAgent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationVersion(val)
        }
        return nil
    }
    res["headerValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderValue(val)
        }
        return nil
    }
    return res
}
func (m *UserAgent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserAgent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationVersion", m.GetApplicationVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("headerValue", m.GetHeaderValue())
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
func (m *UserAgent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationVersion sets the applicationVersion property value. Identifies the version of application software used by this endpoint.
func (m *UserAgent) SetApplicationVersion(value *string)() {
    if m != nil {
        m.applicationVersion = value
    }
}
// SetHeaderValue sets the headerValue property value. User-agent header value reported by this endpoint.
func (m *UserAgent) SetHeaderValue(value *string)() {
    if m != nil {
        m.headerValue = value
    }
}
