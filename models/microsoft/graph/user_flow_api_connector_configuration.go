package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserFlowApiConnectorConfiguration struct {
    additionalData map[string]interface{};
    postAttributeCollection *IdentityApiConnector;
    postFederationSignup *IdentityApiConnector;
}
func NewUserFlowApiConnectorConfiguration()(*UserFlowApiConnectorConfiguration) {
    m := &UserFlowApiConnectorConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserFlowApiConnectorConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserFlowApiConnectorConfiguration) GetPostAttributeCollection()(*IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.postAttributeCollection
    }
}
func (m *UserFlowApiConnectorConfiguration) GetPostFederationSignup()(*IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.postFederationSignup
    }
}
func (m *UserFlowApiConnectorConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["postAttributeCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        m.SetPostAttributeCollection(val.(*IdentityApiConnector))
        return nil
    }
    res["postFederationSignup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        m.SetPostFederationSignup(val.(*IdentityApiConnector))
        return nil
    }
    return res
}
func (m *UserFlowApiConnectorConfiguration) IsNil()(bool) {
    return m == nil
}
func (m *UserFlowApiConnectorConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("postAttributeCollection", m.GetPostAttributeCollection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("postFederationSignup", m.GetPostFederationSignup())
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
func (m *UserFlowApiConnectorConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserFlowApiConnectorConfiguration) SetPostAttributeCollection(value *IdentityApiConnector)() {
    m.postAttributeCollection = value
}
func (m *UserFlowApiConnectorConfiguration) SetPostFederationSignup(value *IdentityApiConnector)() {
    m.postFederationSignup = value
}
