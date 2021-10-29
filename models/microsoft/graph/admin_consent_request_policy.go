package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AdminConsentRequestPolicy struct {
    Entity
    // Specifies whether the admin consent request feature is enabled or disabled. Required.
    isEnabled *bool;
    // Specifies whether reviewers will receive notifications. Required.
    notifyReviewers *bool;
    // Specifies whether reviewers will receive reminder emails. Required.
    remindersEnabled *bool;
    // Specifies the duration the request is active before it automatically expires if no decision is applied.
    requestDurationInDays *int32;
    // The list of reviewers for the admin consent. Required.
    reviewers []AccessReviewReviewerScope;
    // Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
    version *int32;
}
// Instantiates a new adminConsentRequestPolicy and sets the default values.
func NewAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    m := &AdminConsentRequestPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) GetNotifyReviewers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyReviewers
    }
}
// Gets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) GetRemindersEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remindersEnabled
    }
}
// Gets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) GetRequestDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestDurationInDays
    }
}
// Gets the reviewers property value. The list of reviewers for the admin consent. Required.
func (m *AdminConsentRequestPolicy) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// Gets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *AdminConsentRequestPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["notifyReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNotifyReviewers(val)
        return nil
    }
    res["remindersEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRemindersEnabled(val)
        return nil
    }
    res["requestDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRequestDurationInDays(val)
        return nil
    }
    res["reviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewReviewerScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewReviewerScope))
        }
        m.SetReviewers(res)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *AdminConsentRequestPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AdminConsentRequestPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notifyReviewers", m.GetNotifyReviewers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("remindersEnabled", m.GetRemindersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("requestDurationInDays", m.GetRequestDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *AdminConsentRequestPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
// Parameters:
//  - value : Value to set for the notifyReviewers property.
func (m *AdminConsentRequestPolicy) SetNotifyReviewers(value *bool)() {
    m.notifyReviewers = value
}
// Sets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
// Parameters:
//  - value : Value to set for the remindersEnabled property.
func (m *AdminConsentRequestPolicy) SetRemindersEnabled(value *bool)() {
    m.remindersEnabled = value
}
// Sets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
// Parameters:
//  - value : Value to set for the requestDurationInDays property.
func (m *AdminConsentRequestPolicy) SetRequestDurationInDays(value *int32)() {
    m.requestDurationInDays = value
}
// Sets the reviewers property value. The list of reviewers for the admin consent. Required.
// Parameters:
//  - value : Value to set for the reviewers property.
func (m *AdminConsentRequestPolicy) SetReviewers(value []AccessReviewReviewerScope)() {
    m.reviewers = value
}
// Sets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
// Parameters:
//  - value : Value to set for the version property.
func (m *AdminConsentRequestPolicy) SetVersion(value *int32)() {
    m.version = value
}
