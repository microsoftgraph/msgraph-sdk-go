package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserFlowApiConnectorConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    postAttributeCollection *IdentityApiConnector;
    // 
    postFederationSignup *IdentityApiConnector;
}
// Instantiates a new userFlowApiConnectorConfiguration and sets the default values.
func NewUserFlowApiConnectorConfiguration()(*UserFlowApiConnectorConfiguration) {
    m := &UserFlowApiConnectorConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserFlowApiConnectorConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the postAttributeCollection property value. 
func (m *UserFlowApiConnectorConfiguration) GetPostAttributeCollection()(*IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.postAttributeCollection
    }
}
// Gets the postFederationSignup property value. 
func (m *UserFlowApiConnectorConfiguration) GetPostFederationSignup()(*IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.postFederationSignup
    }
}
// The deserialization information for the current model
func (m *UserFlowApiConnectorConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["postAttributeCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostAttributeCollection(val.(*IdentityApiConnector))
        }
        return nil
    }
    res["postFederationSignup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostFederationSignup(val.(*IdentityApiConnector))
        }
        return nil
    }
    return res
}
func (m *UserFlowApiConnectorConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserFlowApiConnectorConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the postAttributeCollection property value. 
// Parameters:
//  - value : Value to set for the postAttributeCollection property.
func (m *UserFlowApiConnectorConfiguration) SetPostAttributeCollection(value *IdentityApiConnector)() {
    m.postAttributeCollection = value
}
// Sets the postFederationSignup property value. 
// Parameters:
//  - value : Value to set for the postFederationSignup property.
func (m *UserFlowApiConnectorConfiguration) SetPostFederationSignup(value *IdentityApiConnector)() {
    m.postFederationSignup = value
}
