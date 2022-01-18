package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConnectedOrganization 
type ConnectedOrganization struct {
    Entity
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the connected organization.
    description *string;
    // The display name of the connected organization. Supports $filter (eq).
    displayName *string;
    // Nullable.
    externalSponsors []DirectoryObject;
    // The identity sources in this connected organization, one of azureActiveDirectoryTenant, domainIdentitySource or externalDomainFederation. Nullable.
    identitySources []IdentitySource;
    // Nullable.
    internalSponsors []DirectoryObject;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The state of a connected organization defines whether assignment policies with requestor scope type AllConfiguredConnectedOrganizationSubjects are applicable or not. Possible values are: configured, proposed.
    state *ConnectedOrganizationState;
}
// NewConnectedOrganization instantiates a new connectedOrganization and sets the default values.
func NewConnectedOrganization()(*ConnectedOrganization) {
    m := &ConnectedOrganization{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ConnectedOrganization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description of the connected organization.
func (m *ConnectedOrganization) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the connected organization. Supports $filter (eq).
func (m *ConnectedOrganization) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalSponsors gets the externalSponsors property value. Nullable.
func (m *ConnectedOrganization) GetExternalSponsors()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.externalSponsors
    }
}
// GetIdentitySources gets the identitySources property value. The identity sources in this connected organization, one of azureActiveDirectoryTenant, domainIdentitySource or externalDomainFederation. Nullable.
func (m *ConnectedOrganization) GetIdentitySources()([]IdentitySource) {
    if m == nil {
        return nil
    } else {
        return m.identitySources
    }
}
// GetInternalSponsors gets the internalSponsors property value. Nullable.
func (m *ConnectedOrganization) GetInternalSponsors()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.internalSponsors
    }
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ConnectedOrganization) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
// GetState gets the state property value. The state of a connected organization defines whether assignment policies with requestor scope type AllConfiguredConnectedOrganizationSubjects are applicable or not. Possible values are: configured, proposed.
func (m *ConnectedOrganization) GetState()(*ConnectedOrganizationState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectedOrganization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["externalSponsors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetExternalSponsors(res)
        }
        return nil
    }
    res["identitySources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentitySource, len(val))
            for i, v := range val {
                res[i] = *(v.(*IdentitySource))
            }
            m.SetIdentitySources(res)
        }
        return nil
    }
    res["internalSponsors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetInternalSponsors(res)
        }
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectedOrganizationState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConnectedOrganizationState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *ConnectedOrganization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConnectedOrganization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExternalSponsors()))
        for i, v := range m.GetExternalSponsors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("externalSponsors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIdentitySources()))
        for i, v := range m.GetIdentitySources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("identitySources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInternalSponsors()))
        for i, v := range m.GetInternalSponsors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("internalSponsors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ConnectedOrganization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description of the connected organization.
func (m *ConnectedOrganization) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the connected organization. Supports $filter (eq).
func (m *ConnectedOrganization) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalSponsors sets the externalSponsors property value. Nullable.
func (m *ConnectedOrganization) SetExternalSponsors(value []DirectoryObject)() {
    if m != nil {
        m.externalSponsors = value
    }
}
// SetIdentitySources sets the identitySources property value. The identity sources in this connected organization, one of azureActiveDirectoryTenant, domainIdentitySource or externalDomainFederation. Nullable.
func (m *ConnectedOrganization) SetIdentitySources(value []IdentitySource)() {
    if m != nil {
        m.identitySources = value
    }
}
// SetInternalSponsors sets the internalSponsors property value. Nullable.
func (m *ConnectedOrganization) SetInternalSponsors(value []DirectoryObject)() {
    if m != nil {
        m.internalSponsors = value
    }
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ConnectedOrganization) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.modifiedDateTime = value
    }
}
// SetState sets the state property value. The state of a connected organization defines whether assignment policies with requestor scope type AllConfiguredConnectedOrganizationSubjects are applicable or not. Possible values are: configured, proposed.
func (m *ConnectedOrganization) SetState(value *ConnectedOrganizationState)() {
    if m != nil {
        m.state = value
    }
}
