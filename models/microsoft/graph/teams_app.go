package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamsApp struct {
    Entity
    appDefinitions []TeamsAppDefinition;
    displayName *string;
    distributionMethod *TeamsAppDistributionMethod;
    externalId *string;
}
func NewTeamsApp()(*TeamsApp) {
    m := &TeamsApp{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TeamsApp) GetAppDefinitions()([]TeamsAppDefinition) {
    if m == nil {
        return nil
    } else {
        return m.appDefinitions
    }
}
func (m *TeamsApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *TeamsApp) GetDistributionMethod()(*TeamsAppDistributionMethod) {
    if m == nil {
        return nil
    } else {
        return m.distributionMethod
    }
}
func (m *TeamsApp) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *TeamsApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppDefinition() })
        if err != nil {
            return err
        }
        res := make([]TeamsAppDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsAppDefinition))
        }
        m.SetAppDefinitions(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["distributionMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppDistributionMethod)
        if err != nil {
            return err
        }
        cast := val.(TeamsAppDistributionMethod)
        m.SetDistributionMethod(&cast)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    return res
}
func (m *TeamsApp) IsNil()(bool) {
    return m == nil
}
func (m *TeamsApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppDefinitions()))
        for i, v := range m.GetAppDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDistributionMethod() != nil {
        cast := m.GetDistributionMethod().String()
        err = writer.WriteStringValue("distributionMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TeamsApp) SetAppDefinitions(value []TeamsAppDefinition)() {
    m.appDefinitions = value
}
func (m *TeamsApp) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *TeamsApp) SetDistributionMethod(value *TeamsAppDistributionMethod)() {
    m.distributionMethod = value
}
func (m *TeamsApp) SetExternalId(value *string)() {
    m.externalId = value
}
