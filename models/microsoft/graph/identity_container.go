package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityContainer provides operations to manage the identityContainer singleton.
type IdentityContainer struct {
    Entity
    // Represents entry point for API connectors.
    apiConnectors []IdentityApiConnectorable;
    // Represents entry point for B2X/self-service sign-up identity userflows.
    b2xUserFlows []B2xIdentityUserFlowable;
    // the entry point for the Conditional Access (CA) object model.
    conditionalAccess ConditionalAccessRootable;
    // Represents entry point for identity provider base.
    identityProviders []IdentityProviderBaseable;
    // Represents entry point for identity userflow attributes.
    userFlowAttributes []IdentityUserFlowAttributeable;
}
// NewIdentityContainer instantiates a new identityContainer and sets the default values.
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityContainerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIdentityContainer(), nil
}
// GetApiConnectors gets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) GetApiConnectors()([]IdentityApiConnectorable) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectors
    }
}
// GetB2xUserFlows gets the b2xUserFlows property value. Represents entry point for B2X/self-service sign-up identity userflows.
func (m *IdentityContainer) GetB2xUserFlows()([]B2xIdentityUserFlowable) {
    if m == nil {
        return nil
    } else {
        return m.b2xUserFlows
    }
}
// GetConditionalAccess gets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) GetConditionalAccess()(ConditionalAccessRootable) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccess
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["apiConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityApiConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityApiConnectorable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityApiConnectorable)
            }
            m.SetApiConnectors(res)
        }
        return nil
    }
    res["b2xUserFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateB2xIdentityUserFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]B2xIdentityUserFlowable, len(val))
            for i, v := range val {
                res[i] = v.(B2xIdentityUserFlowable)
            }
            m.SetB2xUserFlows(res)
        }
        return nil
    }
    res["conditionalAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccess(val.(ConditionalAccessRootable))
        }
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderBaseable)
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["userFlowAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityUserFlowAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityUserFlowAttributeable)
            }
            m.SetUserFlowAttributes(res)
        }
        return nil
    }
    return res
}
// GetIdentityProviders gets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) GetIdentityProviders()([]IdentityProviderBaseable) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
// GetUserFlowAttributes gets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) GetUserFlowAttributes()([]IdentityUserFlowAttributeable) {
    if m == nil {
        return nil
    } else {
        return m.userFlowAttributes
    }
}
func (m *IdentityContainer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentityContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApiConnectors() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApiConnectors()))
        for i, v := range m.GetApiConnectors() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apiConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetB2xUserFlows() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetB2xUserFlows()))
        for i, v := range m.GetB2xUserFlows() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("b2xUserFlows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccess", m.GetConditionalAccess())
        if err != nil {
            return err
        }
    }
    if m.GetIdentityProviders() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("identityProviders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlowAttributes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlowAttributes()))
        for i, v := range m.GetUserFlowAttributes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userFlowAttributes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiConnectors sets the apiConnectors property value. Represents entry point for API connectors.
func (m *IdentityContainer) SetApiConnectors(value []IdentityApiConnectorable)() {
    if m != nil {
        m.apiConnectors = value
    }
}
// SetB2xUserFlows sets the b2xUserFlows property value. Represents entry point for B2X/self-service sign-up identity userflows.
func (m *IdentityContainer) SetB2xUserFlows(value []B2xIdentityUserFlowable)() {
    if m != nil {
        m.b2xUserFlows = value
    }
}
// SetConditionalAccess sets the conditionalAccess property value. the entry point for the Conditional Access (CA) object model.
func (m *IdentityContainer) SetConditionalAccess(value ConditionalAccessRootable)() {
    if m != nil {
        m.conditionalAccess = value
    }
}
// SetIdentityProviders sets the identityProviders property value. Represents entry point for identity provider base.
func (m *IdentityContainer) SetIdentityProviders(value []IdentityProviderBaseable)() {
    if m != nil {
        m.identityProviders = value
    }
}
// SetUserFlowAttributes sets the userFlowAttributes property value. Represents entry point for identity userflow attributes.
func (m *IdentityContainer) SetUserFlowAttributes(value []IdentityUserFlowAttributeable)() {
    if m != nil {
        m.userFlowAttributes = value
    }
}
