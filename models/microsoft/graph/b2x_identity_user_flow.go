package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// b2xIdentityUserFlow 
type B2xIdentityUserFlow struct {
    IdentityUserFlow
    // Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
    apiConnectorConfiguration *UserFlowApiConnectorConfiguration;
    // The identity providers included in the user flow.
    identityProviders []IdentityProvider;
    // The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
    languages []UserFlowLanguageConfiguration;
    // The user attribute assignments included in the user flow.
    userAttributeAssignments []IdentityUserFlowAttributeAssignment;
    // 
    userFlowIdentityProviders []IdentityProviderBase;
}
// NewB2xIdentityUserFlow instantiates a new b2xIdentityUserFlow and sets the default values.
func NewB2xIdentityUserFlow()(*B2xIdentityUserFlow) {
    m := &B2xIdentityUserFlow{
        IdentityUserFlow: *NewIdentityUserFlow(),
    }
    return m
}
// GetApiConnectorConfiguration gets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2xIdentityUserFlow) GetApiConnectorConfiguration()(*UserFlowApiConnectorConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectorConfiguration
    }
}
// GetIdentityProviders gets the identityProviders property value. The identity providers included in the user flow.
func (m *B2xIdentityUserFlow) GetIdentityProviders()([]IdentityProvider) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
// GetLanguages gets the languages property value. The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *B2xIdentityUserFlow) GetLanguages()([]UserFlowLanguageConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
// GetUserAttributeAssignments gets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2xIdentityUserFlow) GetUserAttributeAssignments()([]IdentityUserFlowAttributeAssignment) {
    if m == nil {
        return nil
    } else {
        return m.userAttributeAssignments
    }
}
// GetUserFlowIdentityProviders gets the userFlowIdentityProviders property value. 
func (m *B2xIdentityUserFlow) GetUserFlowIdentityProviders()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.userFlowIdentityProviders
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2xIdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentityUserFlow.GetFieldDeserializers()
    res["apiConnectorConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowApiConnectorConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiConnectorConfiguration(val.(*UserFlowApiConnectorConfiguration))
        }
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProvider() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProvider, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityProvider))
            }
            m.SetIdentityProviders(res)
        }
        return nil
    }
    res["languages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguageConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguageConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserFlowLanguageConfiguration))
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["userAttributeAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttributeAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityUserFlowAttributeAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityUserFlowAttributeAssignment))
            }
            m.SetUserAttributeAssignments(res)
        }
        return nil
    }
    res["userFlowIdentityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentityProviderBase))
            }
            m.SetUserFlowIdentityProviders(res)
        }
        return nil
    }
    return res
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdentityProviders()))
        for i, v := range m.GetIdentityProviders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("identityProviders", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserAttributeAssignments()))
        for i, v := range m.GetUserAttributeAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlowIdentityProviders()))
        for i, v := range m.GetUserFlowIdentityProviders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userFlowIdentityProviders", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiConnectorConfiguration sets the apiConnectorConfiguration property value. Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration.
func (m *B2xIdentityUserFlow) SetApiConnectorConfiguration(value *UserFlowApiConnectorConfiguration)() {
    m.apiConnectorConfiguration = value
}
// SetIdentityProviders sets the identityProviders property value. The identity providers included in the user flow.
func (m *B2xIdentityUserFlow) SetIdentityProviders(value []IdentityProvider)() {
    m.identityProviders = value
}
// SetLanguages sets the languages property value. The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows.
func (m *B2xIdentityUserFlow) SetLanguages(value []UserFlowLanguageConfiguration)() {
    m.languages = value
}
// SetUserAttributeAssignments sets the userAttributeAssignments property value. The user attribute assignments included in the user flow.
func (m *B2xIdentityUserFlow) SetUserAttributeAssignments(value []IdentityUserFlowAttributeAssignment)() {
    m.userAttributeAssignments = value
}
// SetUserFlowIdentityProviders sets the userFlowIdentityProviders property value. 
func (m *B2xIdentityUserFlow) SetUserFlowIdentityProviders(value []IdentityProviderBase)() {
    m.userFlowIdentityProviders = value
}
