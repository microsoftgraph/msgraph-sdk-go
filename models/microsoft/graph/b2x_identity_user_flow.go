package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// B2xIdentityUserFlow provides operations to manage the identityContainer singleton.
type B2xIdentityUserFlow struct {
    IdentityUserFlow
    // Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
    apiConnectorConfiguration UserFlowApiConnectorConfigurationable;
    // The identity providers included in the user flow.
    identityProviders []IdentityProviderable;
    // The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
    languages []UserFlowLanguageConfigurationable;
    // The user attribute assignments included in the user flow.
    userAttributeAssignments []IdentityUserFlowAttributeAssignmentable;
    // 
    userFlowIdentityProviders []IdentityProviderBaseable;
}
// NewB2xIdentityUserFlow instantiates a new b2xIdentityUserFlow and sets the default values.
func NewB2xIdentityUserFlow()(*B2xIdentityUserFlow) {
    m := &B2xIdentityUserFlow{
        IdentityUserFlow: *NewIdentityUserFlow(),
    }
    return m
}
// CreateB2xIdentityUserFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2xIdentityUserFlowFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewB2xIdentityUserFlow(), nil
}
// GetApiConnectorConfiguration gets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2xIdentityUserFlow) GetApiConnectorConfiguration()(UserFlowApiConnectorConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectorConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2xIdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentityUserFlow.GetFieldDeserializers()
    res["apiConnectorConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFlowApiConnectorConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiConnectorConfiguration(val.(UserFlowApiConnectorConfigurationable))
        }
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderable)
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["languages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFlowLanguageConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguageConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(UserFlowLanguageConfigurationable)
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["userAttributeAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityUserFlowAttributeAssignmentable)
            }
            m.SetUserAttributeAssignments(res)
        }
        return nil
    }
    res["userFlowIdentityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderBaseable)
            }
            m.SetUserFlowIdentityProviders(res)
        }
        return nil
    }
    return res
}
// GetIdentityProviders gets the identityProviders property value. The identity providers included in the user flow.
func (m *B2xIdentityUserFlow) GetIdentityProviders()([]IdentityProviderable) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
// GetLanguages gets the languages property value. The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *B2xIdentityUserFlow) GetLanguages()([]UserFlowLanguageConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
// GetUserAttributeAssignments gets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2xIdentityUserFlow) GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.userAttributeAssignments
    }
}
// GetUserFlowIdentityProviders gets the userFlowIdentityProviders property value. 
func (m *B2xIdentityUserFlow) GetUserFlowIdentityProviders()([]IdentityProviderBaseable) {
    if m == nil {
        return nil
    } else {
        return m.userFlowIdentityProviders
    }
}
func (m *B2xIdentityUserFlow) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *B2xIdentityUserFlow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentityUserFlow.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("apiConnectorConfiguration", m.GetApiConnectorConfiguration())
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
    if m.GetLanguages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAttributeAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserAttributeAssignments()))
        for i, v := range m.GetUserAttributeAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserFlowIdentityProviders() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlowIdentityProviders()))
        for i, v := range m.GetUserFlowIdentityProviders() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userFlowIdentityProviders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiConnectorConfiguration sets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2xIdentityUserFlow) SetApiConnectorConfiguration(value UserFlowApiConnectorConfigurationable)() {
    if m != nil {
        m.apiConnectorConfiguration = value
    }
}
// SetIdentityProviders sets the identityProviders property value. The identity providers included in the user flow.
func (m *B2xIdentityUserFlow) SetIdentityProviders(value []IdentityProviderable)() {
    if m != nil {
        m.identityProviders = value
    }
}
// SetLanguages sets the languages property value. The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *B2xIdentityUserFlow) SetLanguages(value []UserFlowLanguageConfigurationable)() {
    if m != nil {
        m.languages = value
    }
}
// SetUserAttributeAssignments sets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2xIdentityUserFlow) SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignmentable)() {
    if m != nil {
        m.userAttributeAssignments = value
    }
}
// SetUserFlowIdentityProviders sets the userFlowIdentityProviders property value. 
func (m *B2xIdentityUserFlow) SetUserFlowIdentityProviders(value []IdentityProviderBaseable)() {
    if m != nil {
        m.userFlowIdentityProviders = value
    }
}
