package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TermsAndConditionsAcceptanceStatus struct {
    Entity
    acceptedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    acceptedVersion *int32;
    termsAndConditions *TermsAndConditions;
    userDisplayName *string;
    userPrincipalName *string;
}
func NewTermsAndConditionsAcceptanceStatus()(*TermsAndConditionsAcceptanceStatus) {
    m := &TermsAndConditionsAcceptanceStatus{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acceptedDateTime
    }
}
func (m *TermsAndConditionsAcceptanceStatus) GetAcceptedVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.acceptedVersion
    }
}
func (m *TermsAndConditionsAcceptanceStatus) GetTermsAndConditions()(*TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
func (m *TermsAndConditionsAcceptanceStatus) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
func (m *TermsAndConditionsAcceptanceStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
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
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acceptedDateTime = value
}
func (m *TermsAndConditionsAcceptanceStatus) SetAcceptedVersion(value *int32)() {
    m.acceptedVersion = value
}
func (m *TermsAndConditionsAcceptanceStatus) SetTermsAndConditions(value *TermsAndConditions)() {
    m.termsAndConditions = value
}
func (m *TermsAndConditionsAcceptanceStatus) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
func (m *TermsAndConditionsAcceptanceStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
