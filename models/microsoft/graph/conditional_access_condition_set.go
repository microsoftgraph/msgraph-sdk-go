package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessConditionSet 
type ConditionalAccessConditionSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Applications and user actions included in and excluded from the policy. Required.
    applications *ConditionalAccessApplications;
    // Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
    clientAppTypes []ConditionalAccessClientApp;
    // Devices in the policy.
    devices *ConditionalAccessDevices;
    // Locations included in and excluded from the policy.
    locations *ConditionalAccessLocations;
    // Platforms included in and excluded from the policy.
    platforms *ConditionalAccessPlatforms;
    // Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
    signInRiskLevels []RiskLevel;
    // User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
    userRiskLevels []RiskLevel;
    // 
    users *ConditionalAccessUsers;
}
// NewConditionalAccessConditionSet instantiates a new conditionalAccessConditionSet and sets the default values.
func NewConditionalAccessConditionSet()(*ConditionalAccessConditionSet) {
    m := &ConditionalAccessConditionSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessConditionSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplications gets the applications property value. Applications and user actions included in and excluded from the policy. Required.
func (m *ConditionalAccessConditionSet) GetApplications()(*ConditionalAccessApplications) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
// GetClientAppTypes gets the clientAppTypes property value. Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
func (m *ConditionalAccessConditionSet) GetClientAppTypes()([]ConditionalAccessClientApp) {
    if m == nil {
        return nil
    } else {
        return m.clientAppTypes
    }
}
// GetDevices gets the devices property value. Devices in the policy.
func (m *ConditionalAccessConditionSet) GetDevices()(*ConditionalAccessDevices) {
    if m == nil {
        return nil
    } else {
        return m.devices
    }
}
// GetLocations gets the locations property value. Locations included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) GetLocations()(*ConditionalAccessLocations) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// GetPlatforms gets the platforms property value. Platforms included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) GetPlatforms()(*ConditionalAccessPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// GetSignInRiskLevels gets the signInRiskLevels property value. Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) GetSignInRiskLevels()([]RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.signInRiskLevels
    }
}
// GetUserRiskLevels gets the userRiskLevels property value. User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) GetUserRiskLevels()([]RiskLevel) {
    if m == nil {
        return nil
    } else {
        return m.userRiskLevels
    }
}
// GetUsers gets the users property value. 
func (m *ConditionalAccessConditionSet) GetUsers()(*ConditionalAccessUsers) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessConditionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessApplications() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplications(val.(*ConditionalAccessApplications))
        }
        return nil
    }
    res["clientAppTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessClientApp)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessClientApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessClientApp))
            }
            m.SetClientAppTypes(res)
        }
        return nil
    }
    res["devices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessDevices() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevices(val.(*ConditionalAccessDevices))
        }
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessLocations() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocations(val.(*ConditionalAccessLocations))
        }
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessPlatforms() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(*ConditionalAccessPlatforms))
        }
        return nil
    }
    res["signInRiskLevels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskLevel, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskLevel))
            }
            m.SetSignInRiskLevels(res)
        }
        return nil
    }
    res["userRiskLevels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskLevel, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskLevel))
            }
            m.SetUserRiskLevels(res)
        }
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConditionalAccessUsers() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsers(val.(*ConditionalAccessUsers))
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessConditionSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err := writer.WriteObjectValue("devices", m.GetDevices())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessConditionSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplications sets the applications property value. Applications and user actions included in and excluded from the policy. Required.
func (m *ConditionalAccessConditionSet) SetApplications(value *ConditionalAccessApplications)() {
    if m != nil {
        m.applications = value
    }
}
// SetClientAppTypes sets the clientAppTypes property value. Client application types included in the policy. Possible values are: all, browser, mobileAppsAndDesktopClients, exchangeActiveSync, easSupported, other. Required.
func (m *ConditionalAccessConditionSet) SetClientAppTypes(value []ConditionalAccessClientApp)() {
    if m != nil {
        m.clientAppTypes = value
    }
}
// SetDevices sets the devices property value. Devices in the policy.
func (m *ConditionalAccessConditionSet) SetDevices(value *ConditionalAccessDevices)() {
    if m != nil {
        m.devices = value
    }
}
// SetLocations sets the locations property value. Locations included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) SetLocations(value *ConditionalAccessLocations)() {
    if m != nil {
        m.locations = value
    }
}
// SetPlatforms sets the platforms property value. Platforms included in and excluded from the policy.
func (m *ConditionalAccessConditionSet) SetPlatforms(value *ConditionalAccessPlatforms)() {
    if m != nil {
        m.platforms = value
    }
}
// SetSignInRiskLevels sets the signInRiskLevels property value. Sign-in risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) SetSignInRiskLevels(value []RiskLevel)() {
    if m != nil {
        m.signInRiskLevels = value
    }
}
// SetUserRiskLevels sets the userRiskLevels property value. User risk levels included in the policy. Possible values are: low, medium, high, hidden, none, unknownFutureValue. Required.
func (m *ConditionalAccessConditionSet) SetUserRiskLevels(value []RiskLevel)() {
    if m != nil {
        m.userRiskLevels = value
    }
}
// SetUsers sets the users property value. 
func (m *ConditionalAccessConditionSet) SetUsers(value *ConditionalAccessUsers)() {
    if m != nil {
        m.users = value
    }
}
