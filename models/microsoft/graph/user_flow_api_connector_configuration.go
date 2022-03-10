package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserFlowApiConnectorConfiguration provides operations to manage the identityContainer singleton.
type UserFlowApiConnectorConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    postAttributeCollection IdentityApiConnectorable;
    // 
    postFederationSignup IdentityApiConnectorable;
}
// NewUserFlowApiConnectorConfiguration instantiates a new userFlowApiConnectorConfiguration and sets the default values.
func NewUserFlowApiConnectorConfiguration()(*UserFlowApiConnectorConfiguration) {
    m := &UserFlowApiConnectorConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserFlowApiConnectorConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserFlowApiConnectorConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserFlowApiConnectorConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserFlowApiConnectorConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserFlowApiConnectorConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["postAttributeCollection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityApiConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostAttributeCollection(val.(IdentityApiConnectorable))
        }
        return nil
    }
    res["postFederationSignup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityApiConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostFederationSignup(val.(IdentityApiConnectorable))
        }
        return nil
    }
    return res
}
// GetPostAttributeCollection gets the postAttributeCollection property value. 
func (m *UserFlowApiConnectorConfiguration) GetPostAttributeCollection()(IdentityApiConnectorable) {
    if m == nil {
        return nil
    } else {
        return m.postAttributeCollection
    }
}
// GetPostFederationSignup gets the postFederationSignup property value. 
func (m *UserFlowApiConnectorConfiguration) GetPostFederationSignup()(IdentityApiConnectorable) {
    if m == nil {
        return nil
    } else {
        return m.postFederationSignup
    }
}
func (m *UserFlowApiConnectorConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserFlowApiConnectorConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPostAttributeCollection sets the postAttributeCollection property value. 
func (m *UserFlowApiConnectorConfiguration) SetPostAttributeCollection(value IdentityApiConnectorable)() {
    if m != nil {
        m.postAttributeCollection = value
    }
}
// SetPostFederationSignup sets the postFederationSignup property value. 
func (m *UserFlowApiConnectorConfiguration) SetPostFederationSignup(value IdentityApiConnectorable)() {
    if m != nil {
        m.postFederationSignup = value
    }
}
