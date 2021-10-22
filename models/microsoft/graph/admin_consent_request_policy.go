package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AdminConsentRequestPolicy struct {
    Entity
    isEnabled *bool;
    notifyReviewers *bool;
    remindersEnabled *bool;
    requestDurationInDays *int32;
    reviewers []AccessReviewReviewerScope;
    version *int32;
}
func NewAdminConsentRequestPolicy()(*AdminConsentRequestPolicy) {
    m := &AdminConsentRequestPolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AdminConsentRequestPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
func (m *AdminConsentRequestPolicy) GetNotifyReviewers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyReviewers
    }
}
func (m *AdminConsentRequestPolicy) GetRemindersEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remindersEnabled
    }
}
func (m *AdminConsentRequestPolicy) GetRequestDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestDurationInDays
    }
}
func (m *AdminConsentRequestPolicy) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
func (m *AdminConsentRequestPolicy) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
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
func (m *AdminConsentRequestPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
func (m *AdminConsentRequestPolicy) SetNotifyReviewers(value *bool)() {
    m.notifyReviewers = value
}
func (m *AdminConsentRequestPolicy) SetRemindersEnabled(value *bool)() {
    m.remindersEnabled = value
}
func (m *AdminConsentRequestPolicy) SetRequestDurationInDays(value *int32)() {
    m.requestDurationInDays = value
}
func (m *AdminConsentRequestPolicy) SetReviewers(value []AccessReviewReviewerScope)() {
    m.reviewers = value
}
func (m *AdminConsentRequestPolicy) SetVersion(value *int32)() {
    m.version = value
}
