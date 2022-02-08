package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsApp 
type TeamsApp struct {
    Entity
    // The details for each version of the app.
    appDefinitions []TeamsAppDefinition;
    // The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
    displayName *string;
    // The method of distribution for the app. Read-only.
    distributionMethod *TeamsAppDistributionMethod;
    // The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
    externalId *string;
}
// NewTeamsApp instantiates a new teamsApp and sets the default values.
func NewTeamsApp()(*TeamsApp) {
    m := &TeamsApp{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppDefinitions gets the appDefinitions property value. The details for each version of the app.
func (m *TeamsApp) GetAppDefinitions()([]TeamsAppDefinition) {
    if m == nil {
        return nil
    } else {
        return m.appDefinitions
    }
}
// GetDisplayName gets the displayName property value. The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDistributionMethod gets the distributionMethod property value. The method of distribution for the app. Read-only.
func (m *TeamsApp) GetDistributionMethod()(*TeamsAppDistributionMethod) {
    if m == nil {
        return nil
    } else {
        return m.distributionMethod
    }
}
// GetExternalId gets the externalId property value. The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamsAppDefinition))
            }
            m.SetAppDefinitions(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["distributionMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppDistributionMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDistributionMethod(val.(*TeamsAppDistributionMethod))
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    return res
}
func (m *TeamsApp) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamsApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppDefinitions() != nil {
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
        cast := (*m.GetDistributionMethod()).String()
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
// SetAppDefinitions sets the appDefinitions property value. The details for each version of the app.
func (m *TeamsApp) SetAppDefinitions(value []TeamsAppDefinition)() {
    if m != nil {
        m.appDefinitions = value
    }
}
// SetDisplayName sets the displayName property value. The name of the catalog app provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDistributionMethod sets the distributionMethod property value. The method of distribution for the app. Read-only.
func (m *TeamsApp) SetDistributionMethod(value *TeamsAppDistributionMethod)() {
    if m != nil {
        m.distributionMethod = value
    }
}
// SetExternalId sets the externalId property value. The ID of the catalog provided by the app developer in the Microsoft Teams zip app package.
func (m *TeamsApp) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
