package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TermsAndConditionsAcceptanceStatus struct {
    Entity
    // DateTime when the terms were last accepted by the user.
    acceptedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Most recent version number of the T&C accepted by the user.
    acceptedVersion *int32;
    // Navigation link to the terms and conditions that are assigned.
    termsAndConditions *TermsAndConditions;
    // Display name of the user whose acceptance the entity represents.
    userDisplayName *string;
    // The userPrincipalName of the User that accepted the term.
    userPrincipalName *string;
}
// Instantiates a new termsAndConditionsAcceptanceStatus and sets the default values.
func NewTermsAndConditionsAcceptanceStatus()(*TermsAndConditionsAcceptanceStatus) {
    m := &TermsAndConditionsAcceptanceStatus{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the acceptedDateTime property value. DateTime when the terms were last accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acceptedDateTime
    }
}
// Gets the acceptedVersion property value. Most recent version number of the T&C accepted by the user.
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.acceptedVersion
    }
}
// Gets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
func (m *TermsAndConditionsAcceptanceStatus) GetTermsAndConditions()(*TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
// Gets the userDisplayName property value. Display name of the user whose acceptance the entity represents.
func (m *TermsAndConditionsAcceptanceStatus) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userPrincipalName property value. The userPrincipalName of the User that accepted the term.
func (m *TermsAndConditionsAcceptanceStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *TermsAndConditionsAcceptanceStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAcceptedDateTime(val)
        return nil
    }
    res["acceptedVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAcceptedVersion(val)
        return nil
    }
    res["termsAndConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsAndConditions() })
        if err != nil {
            return err
        }
        m.SetTermsAndConditions(val.(*TermsAndConditions))
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserDisplayName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *TermsAndConditionsAcceptanceStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TermsAndConditionsAcceptanceStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acceptedDateTime", m.GetAcceptedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("acceptedVersion", m.GetAcceptedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("termsAndConditions", m.GetTermsAndConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the acceptedDateTime property value. DateTime when the terms were last accepted by the user.
// Parameters:
//  - value : Value to set for the acceptedDateTime property.
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acceptedDateTime = value
}
// Sets the acceptedVersion property value. Most recent version number of the T&C accepted by the user.
// Parameters:
//  - value : Value to set for the acceptedVersion property.
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedVersion(value *int32)() {
    m.acceptedVersion = value
}
// Sets the termsAndConditions property value. Navigation link to the terms and conditions that are assigned.
// Parameters:
//  - value : Value to set for the termsAndConditions property.
func (m *TermsAndConditionsAcceptanceStatus) SetTermsAndConditions(value *TermsAndConditions)() {
    m.termsAndConditions = value
}
// Sets the userDisplayName property value. Display name of the user whose acceptance the entity represents.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *TermsAndConditionsAcceptanceStatus) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userPrincipalName property value. The userPrincipalName of the User that accepted the term.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *TermsAndConditionsAcceptanceStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
