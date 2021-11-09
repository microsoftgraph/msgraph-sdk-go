package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/ediscovery"
)

// 
type Case_escaped struct {
    Entity
    // The user who closed the case.
    closedBy *IdentitySet;
    // The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    closedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Returns a list of case custodian objects for this case.  Nullable.
    custodians []Custodian;
    // The case description.
    description *string;
    // The case name.
    displayName *string;
    // The external case number for customer reference.
    externalId *string;
    // The last user who modified the entity.
    lastModifiedBy *IdentitySet;
    // The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Returns a list of case legalHold objects for this case.  Nullable.
    legalHolds []LegalHold;
    // Returns a list of case noncustodialDataSource objects for this case.  Nullable.
    noncustodialDataSources []NoncustodialDataSource;
    // Returns a list of case operation objects for this case. Nullable.
    operations []CaseOperation;
    // Returns a list of reviewSet objects in the case. Read-only. Nullable.
    reviewSets []ReviewSet;
    // 
    settings *CaseSettings;
    // Returns a list of sourceCollection objects associated with this case.
    sourceCollections []SourceCollection;
    // The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
    status *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.CaseStatus;
    // Returns a list of tag objects associated to this case.
    tags []Tag;
}
// Instantiates a new case_escaped and sets the default values.
func NewCase_escaped()(*Case_escaped) {
    m := &Case_escaped{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the closedBy property value. The user who closed the case.
func (m *Case_escaped) GetClosedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.closedBy
    }
}
// Gets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetClosedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.closedDateTime
    }
}
// Gets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
func (m *Case_escaped) GetCustodians()([]Custodian) {
    if m == nil {
        return nil
    } else {
        return m.custodians
    }
}
// Gets the description property value. The case description.
func (m *Case_escaped) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The case name.
func (m *Case_escaped) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the externalId property value. The external case number for customer reference.
func (m *Case_escaped) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the lastModifiedBy property value. The last user who modified the entity.
func (m *Case_escaped) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Case_escaped) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
func (m *Case_escaped) GetLegalHolds()([]LegalHold) {
    if m == nil {
        return nil
    } else {
        return m.legalHolds
    }
}
// Gets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
func (m *Case_escaped) GetNoncustodialDataSources()([]NoncustodialDataSource) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialDataSources
    }
}
// Gets the operations property value. Returns a list of case operation objects for this case. Nullable.
func (m *Case_escaped) GetOperations()([]CaseOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
func (m *Case_escaped) GetReviewSets()([]ReviewSet) {
    if m == nil {
        return nil
    } else {
        return m.reviewSets
    }
}
// Gets the settings property value. 
func (m *Case_escaped) GetSettings()(*CaseSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
func (m *Case_escaped) GetSourceCollections()([]SourceCollection) {
    if m == nil {
        return nil
    } else {
        return m.sourceCollections
    }
}
// Gets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
func (m *Case_escaped) GetStatus()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.CaseStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the tags property value. Returns a list of tag objects associated to this case.
func (m *Case_escaped) GetTags()([]Tag) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// The deserialization information for the current model
func (m *Case_escaped) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["closedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetClosedBy(val.(*IdentitySet))
        return nil
    }
    res["closedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetClosedDateTime(val)
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
    res["custodians"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustodian() })
        if err != nil {
            return err
        }
        res := make([]Custodian, len(val))
        for i, v := range val {
            res[i] = *(v.(*Custodian))
        }
        m.SetCustodians(res)
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
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
    res["legalHolds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLegalHold() })
        if err != nil {
            return err
        }
        res := make([]LegalHold, len(val))
        for i, v := range val {
            res[i] = *(v.(*LegalHold))
        }
        m.SetLegalHolds(res)
        return nil
    }
    res["noncustodialDataSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNoncustodialDataSource() })
        if err != nil {
            return err
        }
        res := make([]NoncustodialDataSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*NoncustodialDataSource))
        }
        m.SetNoncustodialDataSources(res)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseOperation() })
        if err != nil {
            return err
        }
        res := make([]CaseOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*CaseOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["reviewSets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReviewSet() })
        if err != nil {
            return err
        }
        res := make([]ReviewSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*ReviewSet))
        }
        m.SetReviewSets(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCaseSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*CaseSettings))
        return nil
    }
    res["sourceCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSourceCollection() })
        if err != nil {
            return err
        }
        res := make([]SourceCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*SourceCollection))
        }
        m.SetSourceCollections(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseCaseStatus)
        if err != nil {
            return err
        }
        cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.CaseStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        res := make([]Tag, len(val))
        for i, v := range val {
            res[i] = *(v.(*Tag))
        }
        m.SetTags(res)
        return nil
    }
    return res
}
func (m *Case_escaped) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Case_escaped) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("closedBy", m.GetClosedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closedDateTime", m.GetClosedDateTime())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustodians()))
        for i, v := range m.GetCustodians() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("custodians", cast)
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLegalHolds()))
        for i, v := range m.GetLegalHolds() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("legalHolds", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNoncustodialDataSources()))
        for i, v := range m.GetNoncustodialDataSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialDataSources", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReviewSets()))
        for i, v := range m.GetReviewSets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reviewSets", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSourceCollections()))
        for i, v := range m.GetSourceCollections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sourceCollections", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tags", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the closedBy property value. The user who closed the case.
// Parameters:
//  - value : Value to set for the closedBy property.
func (m *Case_escaped) SetClosedBy(value *IdentitySet)() {
    m.closedBy = value
}
// Sets the closedDateTime property value. The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the closedDateTime property.
func (m *Case_escaped) SetClosedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closedDateTime = value
}
// Sets the createdDateTime property value. The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Case_escaped) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the custodians property value. Returns a list of case custodian objects for this case.  Nullable.
// Parameters:
//  - value : Value to set for the custodians property.
func (m *Case_escaped) SetCustodians(value []Custodian)() {
    m.custodians = value
}
// Sets the description property value. The case description.
// Parameters:
//  - value : Value to set for the description property.
func (m *Case_escaped) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The case name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Case_escaped) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the externalId property value. The external case number for customer reference.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *Case_escaped) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the lastModifiedBy property value. The last user who modified the entity.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *Case_escaped) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. The latest date and time when the case was modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Case_escaped) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the legalHolds property value. Returns a list of case legalHold objects for this case.  Nullable.
// Parameters:
//  - value : Value to set for the legalHolds property.
func (m *Case_escaped) SetLegalHolds(value []LegalHold)() {
    m.legalHolds = value
}
// Sets the noncustodialDataSources property value. Returns a list of case noncustodialDataSource objects for this case.  Nullable.
// Parameters:
//  - value : Value to set for the noncustodialDataSources property.
func (m *Case_escaped) SetNoncustodialDataSources(value []NoncustodialDataSource)() {
    m.noncustodialDataSources = value
}
// Sets the operations property value. Returns a list of case operation objects for this case. Nullable.
// Parameters:
//  - value : Value to set for the operations property.
func (m *Case_escaped) SetOperations(value []CaseOperation)() {
    m.operations = value
}
// Sets the reviewSets property value. Returns a list of reviewSet objects in the case. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the reviewSets property.
func (m *Case_escaped) SetReviewSets(value []ReviewSet)() {
    m.reviewSets = value
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *Case_escaped) SetSettings(value *CaseSettings)() {
    m.settings = value
}
// Sets the sourceCollections property value. Returns a list of sourceCollection objects associated with this case.
// Parameters:
//  - value : Value to set for the sourceCollections property.
func (m *Case_escaped) SetSourceCollections(value []SourceCollection)() {
    m.sourceCollections = value
}
// Sets the status property value. The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For details, see the following table.
// Parameters:
//  - value : Value to set for the status property.
func (m *Case_escaped) SetStatus(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.CaseStatus)() {
    m.status = value
}
// Sets the tags property value. Returns a list of tag objects associated to this case.
// Parameters:
//  - value : Value to set for the tags property.
func (m *Case_escaped) SetTags(value []Tag)() {
    m.tags = value
}
