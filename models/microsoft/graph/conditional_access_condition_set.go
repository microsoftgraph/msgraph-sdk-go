package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessConditionSet struct {
    additionalData map[string]interface{};
    applications *ConditionalAccessApplications;
    clientAppTypes []ConditionalAccessClientApp;
    locations *ConditionalAccessLocations;
    platforms *ConditionalAccessPlatforms;
    signInRiskLevels []RiskLevel;
    userRiskLevels []RiskLevel;
    users *ConditionalAccessUsers;
}
func NewConditionalAccessConditionSet()(*ConditionalAccessConditionSet) {
    m := &ConditionalAccessConditionSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessConditionSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessConditionSet) GetApplications()(*ConditionalAccessApplications) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
func (m *ConditionalAccessConditionSet) GetClientAppTypes()([]ConditionalAccessClientApp) {
    if m == nil {
        return nil
    } else {
        return m.clientAppTypes
    }
}
func (m *ConditionalAccessConditionSet) GetLocations()(*ConditionalAccessLocations) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
func (m *ConditionalAccessConditionSet) GetPlatforms()(*ConditionalAccessPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
func (m *ConditionalAccessConditionSet) GetSignInRiskLevels()([]RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.signInRiskLevels
    }
}
func (m *ConditionalAccessConditionSet) GetUserRiskLevels()([]RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.userRiskLevels
    }
}
func (m *ConditionalAccessConditionSet) GetUsers()(*ConditionalAccessUsers) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
func (m *ConditionalAccessConditionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessApplications() })
        if err != nil {
            return err
        }
        m.SetApplications(val.(*ConditionalAccessApplications))
        return nil
    }
    res["clientAppTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessClientApp)
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessClientApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessClientApp))
        }
        m.SetClientAppTypes(res)
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessLocations() })
        if err != nil {
            return err
        }
        m.SetLocations(val.(*ConditionalAccessLocations))
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPlatforms() })
        if err != nil {
            return err
        }
        m.SetPlatforms(val.(*ConditionalAccessPlatforms))
        return nil
    }
    res["signInRiskLevels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        res := make([]RiskLevel, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskLevel))
        }
        m.SetSignInRiskLevels(res)
        return nil
    }
    res["userRiskLevels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        res := make([]RiskLevel, len(val))
        for i, v := range val {
            res[i] = *(v.(*RiskLevel))
        }
        m.SetUserRiskLevels(res)
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessUsers() })
        if err != nil {
            return err
        }
        m.SetUsers(val.(*ConditionalAccessUsers))
        return nil
    }
    return res
}
func (m *ConditionalAccessConditionSet) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessConditionSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applications", m.GetApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("clientAppTypes", SerializeConditionalAccessClientApp(m.GetClientAppTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locations", m.GetLocations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("platforms", m.GetPlatforms())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("signInRiskLevels", SerializeRiskLevel(m.GetSignInRiskLevels()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("userRiskLevels", SerializeRiskLevel(m.GetUserRiskLevels()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("users", m.GetUsers())
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
func (m *ConditionalAccessConditionSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessConditionSet) SetApplications(value *ConditionalAccessApplications)() {
    m.applications = value
}
func (m *ConditionalAccessConditionSet) SetClientAppTypes(value []ConditionalAccessClientApp)() {
    m.clientAppTypes = value
}
func (m *ConditionalAccessConditionSet) SetLocations(value *ConditionalAccessLocations)() {
    m.locations = value
}
func (m *ConditionalAccessConditionSet) SetPlatforms(value *ConditionalAccessPlatforms)() {
    m.platforms = value
}
func (m *ConditionalAccessConditionSet) SetSignInRiskLevels(value []RiskLevel)() {
    m.signInRiskLevels = value
}
func (m *ConditionalAccessConditionSet) SetUserRiskLevels(value []RiskLevel)() {
    m.userRiskLevels = value
}
func (m *ConditionalAccessConditionSet) SetUsers(value *ConditionalAccessUsers)() {
    m.users = value
}
