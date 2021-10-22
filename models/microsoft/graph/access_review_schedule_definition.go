package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewScheduleDefinition struct {
    Entity
    createdBy *UserIdentity;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    descriptionForAdmins *string;
    descriptionForReviewers *string;
    displayName *string;
    fallbackReviewers []AccessReviewReviewerScope;
    instanceEnumerationScope *AccessReviewScope;
    instances []AccessReviewInstance;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    reviewers []AccessReviewReviewerScope;
    scope *AccessReviewScope;
    settings *AccessReviewScheduleSettings;
    status *string;
}
func NewAccessReviewScheduleDefinition()(*AccessReviewScheduleDefinition) {
    m := &AccessReviewScheduleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessReviewScheduleDefinition) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *AccessReviewScheduleDefinition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *AccessReviewScheduleDefinition) GetDescriptionForAdmins()(*string) {
    if m == nil {
        return nil
    } else {
        return m.descriptionForAdmins
    }
}
func (m *AccessReviewScheduleDefinition) GetDescriptionForReviewers()(*string) {
    if m == nil {
        return nil
    } else {
        return m.descriptionForReviewers
    }
}
func (m *AccessReviewScheduleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AccessReviewScheduleDefinition) GetFallbackReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.fallbackReviewers
    }
}
func (m *AccessReviewScheduleDefinition) GetInstanceEnumerationScope()(*AccessReviewScope) {
    if m == nil {
        return nil
    } else {
        return m.instanceEnumerationScope
    }
}
func (m *AccessReviewScheduleDefinition) GetInstances()([]AccessReviewInstance) {
    if m == nil {
        return nil
    } else {
        return m.instances
    }
}
func (m *AccessReviewScheduleDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *AccessReviewScheduleDefinition) GetReviewers()([]AccessReviewReviewerScope) {
    if m == nil {
        return nil
    } else {
        return m.reviewers
    }
}
func (m *AccessReviewScheduleDefinition) GetScope()(*AccessReviewScope) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
func (m *AccessReviewScheduleDefinition) GetSettings()(*AccessReviewScheduleSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *AccessReviewScheduleDefinition) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *AccessReviewScheduleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*UserIdentity))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["descriptionForAdmins"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescriptionForAdmins(val)
        return nil
    }
    res["descriptionForReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescriptionForReviewers(val)
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
    res["fallbackReviewers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewReviewerScope() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewReviewerScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewReviewerScope))
        }
        m.SetFallbackReviewers(res)
        return nil
    }
    res["instanceEnumerationScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScope() })
        if err != nil {
            return err
        }
        m.SetInstanceEnumerationScope(val.(*AccessReviewScope))
        return nil
    }
    res["instances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewInstance() })
        if err != nil {
            return err
        }
        res := make([]AccessReviewInstance, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessReviewInstance))
        }
        m.SetInstances(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
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
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScope() })
        if err != nil {
            return err
        }
        m.SetScope(val.(*AccessReviewScope))
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewScheduleSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*AccessReviewScheduleSettings))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *AccessReviewScheduleDefinition) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewScheduleDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForAdmins", m.GetDescriptionForAdmins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForReviewers", m.GetDescriptionForReviewers())
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFallbackReviewers()))
        for i, v := range m.GetFallbackReviewers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("instanceEnumerationScope", m.GetInstanceEnumerationScope())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AccessReviewScheduleDefinition) SetCreatedBy(value *UserIdentity)() {
    m.createdBy = value
}
func (m *AccessReviewScheduleDefinition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *AccessReviewScheduleDefinition) SetDescriptionForAdmins(value *string)() {
    m.descriptionForAdmins = value
}
func (m *AccessReviewScheduleDefinition) SetDescriptionForReviewers(value *string)() {
    m.descriptionForReviewers = value
}
func (m *AccessReviewScheduleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AccessReviewScheduleDefinition) SetFallbackReviewers(value []AccessReviewReviewerScope)() {
    m.fallbackReviewers = value
}
func (m *AccessReviewScheduleDefinition) SetInstanceEnumerationScope(value *AccessReviewScope)() {
    m.instanceEnumerationScope = value
}
func (m *AccessReviewScheduleDefinition) SetInstances(value []AccessReviewInstance)() {
    m.instances = value
}
func (m *AccessReviewScheduleDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *AccessReviewScheduleDefinition) SetReviewers(value []AccessReviewReviewerScope)() {
    m.reviewers = value
}
func (m *AccessReviewScheduleDefinition) SetScope(value *AccessReviewScope)() {
    m.scope = value
}
func (m *AccessReviewScheduleDefinition) SetSettings(value *AccessReviewScheduleSettings)() {
    m.settings = value
}
func (m *AccessReviewScheduleDefinition) SetStatus(value *string)() {
    m.status = value
}
