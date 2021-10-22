package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IdentityContainer struct {
    Entity
    apiConnectors []IdentityApiConnector;
    b2xUserFlows []B2xIdentityUserFlow;
    conditionalAccess *ConditionalAccessRoot;
    identityProviders []IdentityProviderBase;
    userFlowAttributes []IdentityUserFlowAttribute;
}
func NewIdentityContainer()(*IdentityContainer) {
    m := &IdentityContainer{
        Entity: *NewEntity(),
    }
    return m
}
func (m *IdentityContainer) GetApiConnectors()([]IdentityApiConnector) {
    if m == nil {
        return nil
    } else {
        return m.apiConnectors
    }
}
func (m *IdentityContainer) GetB2xUserFlows()([]B2xIdentityUserFlow) {
    if m == nil {
        return nil
    } else {
        return m.b2xUserFlows
    }
}
func (m *IdentityContainer) GetConditionalAccess()(*ConditionalAccessRoot) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccess
    }
}
func (m *IdentityContainer) GetIdentityProviders()([]IdentityProviderBase) {
    if m == nil {
        return nil
    } else {
        return m.identityProviders
    }
}
func (m *IdentityContainer) GetUserFlowAttributes()([]IdentityUserFlowAttribute) {
    if m == nil {
        return nil
    } else {
        return m.userFlowAttributes
    }
}
func (m *IdentityContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["apiConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityApiConnector() })
        if err != nil {
            return err
        }
        res := make([]IdentityApiConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityApiConnector))
        }
        m.SetApiConnectors(res)
        return nil
    }
    res["b2xUserFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewB2xIdentityUserFlow() })
        if err != nil {
            return err
        }
        res := make([]B2xIdentityUserFlow, len(val))
        for i, v := range val {
            res[i] = *(v.(*B2xIdentityUserFlow))
        }
        m.SetB2xUserFlows(res)
        return nil
    }
    res["conditionalAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessRoot() })
        if err != nil {
            return err
        }
        m.SetConditionalAccess(val.(*ConditionalAccessRoot))
        return nil
    }
    res["identityProviders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityProviderBase() })
        if err != nil {
            return err
        }
        res := make([]IdentityProviderBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityProviderBase))
        }
        m.SetIdentityProviders(res)
        return nil
    }
    res["userFlowAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttribute() })
        if err != nil {
            return err
        }
        res := make([]IdentityUserFlowAttribute, len(val))
        for i, v := range val {
            res[i] = *(v.(*IdentityUserFlowAttribute))
        }
        m.SetUserFlowAttributes(res)
        return nil
    }
    return res
}
func (m *IdentityContainer) IsNil()(bool) {
    return m == nil
}
func (m *IdentityContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApiConnectors()))
        for i, v := range m.GetApiConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("apiConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetB2xUserFlows()))
        for i, v := range m.GetB2xUserFlows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserFlowAttributes()))
        for i, v := range m.GetUserFlowAttributes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userFlowAttributes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *IdentityContainer) SetApiConnectors(value []IdentityApiConnector)() {
    m.apiConnectors = value
}
func (m *IdentityContainer) SetB2xUserFlows(value []B2xIdentityUserFlow)() {
    m.b2xUserFlows = value
}
func (m *IdentityContainer) SetConditionalAccess(value *ConditionalAccessRoot)() {
    m.conditionalAccess = value
}
func (m *IdentityContainer) SetIdentityProviders(value []IdentityProviderBase)() {
    m.identityProviders = value
}
func (m *IdentityContainer) SetUserFlowAttributes(value []IdentityUserFlowAttribute)() {
    m.userFlowAttributes = value
}
