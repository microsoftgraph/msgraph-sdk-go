package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdminConsentRequestPolicy 
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
// NewAdminConsentRequestPolicy instantiates a new adminConsentRequestPolicy and sets the default values.
func NewAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    m := &AdminConsentRequestPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetIsEnabled gets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetNotifyReviewers gets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) GetNotifyReviewers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyReviewers
    }
}
// GetRemindersEnabled gets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) GetRemindersEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remindersEnabled
    }
}
// GetRequestDurationInDays gets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) GetRequestDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestDurationInDays
    }
}
// GetReviewers gets the reviewers property value. The list of reviewers for the admin consent. Required.
func (m *AdminConsentRequestPolicy) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
// GetVersion gets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsentRequestPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["notifyReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotifyReviewers(val)
        }
        return nil
    }
    res["remindersEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindersEnabled(val)
        }
        return nil
    }
    res["requestDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDurationInDays(val)
        }
        return nil
    }
    res["reviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewReviewerScope))
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *AdminConsentRequestPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetIsEnabled sets the isEnabled property value. Specifies whether the admin consent request feature is enabled or disabled. Required.
func (m *AdminConsentRequestPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetNotifyReviewers sets the notifyReviewers property value. Specifies whether reviewers will receive notifications. Required.
func (m *AdminConsentRequestPolicy) SetNotifyReviewers(value *bool)() {
    m.notifyReviewers = value
}
// SetRemindersEnabled sets the remindersEnabled property value. Specifies whether reviewers will receive reminder emails. Required.
func (m *AdminConsentRequestPolicy) SetRemindersEnabled(value *bool)() {
    m.remindersEnabled = value
}
// SetRequestDurationInDays sets the requestDurationInDays property value. Specifies the duration the request is active before it automatically expires if no decision is applied.
func (m *AdminConsentRequestPolicy) SetRequestDurationInDays(value *int32)() {
    m.requestDurationInDays = value
}
// SetReviewers sets the reviewers property value. The list of reviewers for the admin consent. Required.
func (m *AdminConsentRequestPolicy) SetReviewers(value []AccessReviewReviewerScope)() {
    m.reviewers = value
}
// SetVersion sets the version property value. Specifies the version of this policy. When the policy is updated, this version is updated. Read-only.
func (m *AdminConsentRequestPolicy) SetVersion(value *int32)() {
    m.version = value
}
