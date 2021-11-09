package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/ediscovery"
)

// 
type LegalHold struct {
    Entity
    // KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
    contentQuery *string;
    // The user who created the legal hold.
    createdBy *IdentitySet;
    // The date and time the legal hold was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The legal hold description.
    description *string;
    // The display name of the legal hold.
    displayName *string;
    // Lists any errors that happened while placing the hold.
    errors []string;
    // Indicates whether the hold is enabled and actively holding content.
    isEnabled *bool;
    // the user who last modified the legal hold.
    lastModifiedBy *IdentitySet;
    // The date and time the legal hold was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Data source entity for SharePoint sites associated with the legal hold.
    siteSources []SiteSource;
    // The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
    status *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.LegalHoldStatus;
    // Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
    userSources []UserSource;
}
// Instantiates a new legalHold and sets the default values.
func NewLegalHold()(*LegalHold) {
    m := &LegalHold{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
func (m *LegalHold) GetContentQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentQuery
    }
}
// Gets the createdBy property value. The user who created the legal hold.
func (m *LegalHold) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The date and time the legal hold was created.
func (m *LegalHold) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The legal hold description.
func (m *LegalHold) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the legal hold.
func (m *LegalHold) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the errors property value. Lists any errors that happened while placing the hold.
func (m *LegalHold) GetErrors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.errors
    }
}
// Gets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
func (m *LegalHold) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the lastModifiedBy property value. the user who last modified the legal hold.
func (m *LegalHold) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
func (m *LegalHold) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
func (m *LegalHold) GetSiteSources()([]SiteSource) {
    if m == nil {
        return nil
    } else {
        return m.siteSources
    }
}
// Gets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
func (m *LegalHold) GetStatus()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.LegalHoldStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
func (m *LegalHold) GetUserSources()([]UserSource) {
    if m == nil {
        return nil
    } else {
        return m.userSources
    }
}
// The deserialization information for the current model
func (m *LegalHold) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentQuery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentQuery(val)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*IdentitySet))
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["errors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetErrors(res)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetLastModifiedBy(val.(*IdentitySet))
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseLegalHoldStatus)
        if err != nil {
            return err
        }
        cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.LegalHoldStatus)
        m.SetStatus(&cast)
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
func (m *LegalHold) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LegalHold) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentQuery", m.GetContentQuery())
        if err != nil {
            return err
        }
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
        err = writer.WriteCollectionOfStringValues("errors", m.GetErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
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
// Sets the contentQuery property value. KQL query that specifies content to be held in the specified locations. To learn more, see Keyword queries and search conditions for Content Search and eDiscovery.  To hold all content in the specified locations, leave contentQuery blank.
// Parameters:
//  - value : Value to set for the contentQuery property.
func (m *LegalHold) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// Sets the createdBy property value. The user who created the legal hold.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *LegalHold) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The date and time the legal hold was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *LegalHold) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The legal hold description.
// Parameters:
//  - value : Value to set for the description property.
func (m *LegalHold) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the legal hold.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *LegalHold) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the errors property value. Lists any errors that happened while placing the hold.
// Parameters:
//  - value : Value to set for the errors property.
func (m *LegalHold) SetErrors(value []string)() {
    m.errors = value
}
// Sets the isEnabled property value. Indicates whether the hold is enabled and actively holding content.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *LegalHold) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the lastModifiedBy property value. the user who last modified the legal hold.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *LegalHold) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. The date and time the legal hold was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *LegalHold) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the siteSources property value. Data source entity for SharePoint sites associated with the legal hold.
// Parameters:
//  - value : Value to set for the siteSources property.
func (m *LegalHold) SetSiteSources(value []SiteSource)() {
    m.siteSources = value
}
// Sets the status property value. The status of the legal hold. Possible values are: Pending, Error, Success, UnknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *LegalHold) SetStatus(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.LegalHoldStatus)() {
    m.status = value
}
// Sets the userSources property value. Data source entity for a the legal hold. This is the container for a mailbox and OneDrive for Business site.
// Parameters:
//  - value : Value to set for the userSources property.
func (m *LegalHold) SetUserSources(value []UserSource)() {
    m.userSources = value
}
