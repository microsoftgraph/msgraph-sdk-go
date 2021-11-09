package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Custodian struct {
    DataSourceContainer
    // Date and time the custodian acknowledged a hold notification.
    acknowledgedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identifies whether a custodian's sources were placed on hold during creation.
    applyHoldToSources *bool;
    // Email address of the custodian.
    email *string;
    // Data source entity for SharePoint sites associated with the custodian.
    siteSources []SiteSource;
    // Data source entity for groups associated with the custodian.
    unifiedGroupSources []UnifiedGroupSource;
    // Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
    userSources []UserSource;
}
// Instantiates a new custodian and sets the default values.
func NewCustodian()(*Custodian) {
    m := &Custodian{
        DataSourceContainer: *NewDataSourceContainer(),
    }
    return m
}
// Gets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
func (m *Custodian) GetAcknowledgedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.acknowledgedDateTime
    }
}
// Gets the applyHoldToSources property value. Identifies whether a custodian's sources were placed on hold during creation.
func (m *Custodian) GetApplyHoldToSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyHoldToSources
    }
}
// Gets the email property value. Email address of the custodian.
func (m *Custodian) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the siteSources property value. Data source entity for SharePoint sites associated with the custodian.
func (m *Custodian) GetSiteSources()([]SiteSource) {
    if m == nil {
        return nil
    } else {
        return m.siteSources
    }
}
// Gets the unifiedGroupSources property value. Data source entity for groups associated with the custodian.
func (m *Custodian) GetUnifiedGroupSources()([]UnifiedGroupSource) {
    if m == nil {
        return nil
    } else {
        return m.unifiedGroupSources
    }
}
// Gets the userSources property value. Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
func (m *Custodian) GetUserSources()([]UserSource) {
    if m == nil {
        return nil
    } else {
        return m.userSources
    }
}
// The deserialization information for the current model
func (m *Custodian) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DataSourceContainer.GetFieldDeserializers()
    res["acknowledgedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAcknowledgedDateTime(val)
        return nil
    }
    res["applyHoldToSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApplyHoldToSources(val)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["siteSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSiteSource() })
        if err != nil {
            return err
        }
        res := make([]SiteSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*SiteSource))
        }
        m.SetSiteSources(res)
        return nil
    }
    res["unifiedGroupSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedGroupSource() })
        if err != nil {
            return err
        }
        res := make([]UnifiedGroupSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedGroupSource))
        }
        m.SetUnifiedGroupSources(res)
        return nil
    }
    res["userSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSource() })
        if err != nil {
            return err
        }
        res := make([]UserSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSource))
        }
        m.SetUserSources(res)
        return nil
    }
    return res
}
func (m *Custodian) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Custodian) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DataSourceContainer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("acknowledgedDateTime", m.GetAcknowledgedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applyHoldToSources", m.GetApplyHoldToSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSiteSources()))
        for i, v := range m.GetSiteSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("siteSources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUnifiedGroupSources()))
        for i, v := range m.GetUnifiedGroupSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("unifiedGroupSources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserSources()))
        for i, v := range m.GetUserSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the acknowledgedDateTime property value. Date and time the custodian acknowledged a hold notification.
// Parameters:
//  - value : Value to set for the acknowledgedDateTime property.
func (m *Custodian) SetAcknowledgedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.acknowledgedDateTime = value
}
// Sets the applyHoldToSources property value. Identifies whether a custodian's sources were placed on hold during creation.
// Parameters:
//  - value : Value to set for the applyHoldToSources property.
func (m *Custodian) SetApplyHoldToSources(value *bool)() {
    m.applyHoldToSources = value
}
// Sets the email property value. Email address of the custodian.
// Parameters:
//  - value : Value to set for the email property.
func (m *Custodian) SetEmail(value *string)() {
    m.email = value
}
// Sets the siteSources property value. Data source entity for SharePoint sites associated with the custodian.
// Parameters:
//  - value : Value to set for the siteSources property.
func (m *Custodian) SetSiteSources(value []SiteSource)() {
    m.siteSources = value
}
// Sets the unifiedGroupSources property value. Data source entity for groups associated with the custodian.
// Parameters:
//  - value : Value to set for the unifiedGroupSources property.
func (m *Custodian) SetUnifiedGroupSources(value []UnifiedGroupSource)() {
    m.unifiedGroupSources = value
}
// Sets the userSources property value. Data source entity for a the custodian. This is the container for a custodian's mailbox and OneDrive for Business site.
// Parameters:
//  - value : Value to set for the userSources property.
func (m *Custodian) SetUserSources(value []UserSource)() {
    m.userSources = value
}
