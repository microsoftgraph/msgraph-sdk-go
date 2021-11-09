package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/ediscovery"
)

// 
type SourceCollection struct {
    Entity
    // Adds an additional source to the sourceCollection.
    additionalSources []DataSource;
    // Adds the results of the sourceCollection to the specified reviewSet.
    addToReviewSetOperation *AddToReviewSetOperation;
    // The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
    contentQuery *string;
    // The user who created the sourceCollection.
    createdBy *IdentitySet;
    // The date and time the sourceCollection was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Custodian sources that are included in the sourceCollection.
    custodianSources []DataSource;
    // When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
    dataSourceScopes *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceScopes;
    // The description of the sourceCollection.
    description *string;
    // The display name of the sourceCollection.
    displayName *string;
    // The last estimate operation associated with the sourceCollection.
    lastEstimateStatisticsOperation *EstimateStatisticsOperation;
    // The last user who modified the sourceCollection.
    lastModifiedBy *IdentitySet;
    // The last date and time the sourceCollection was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // noncustodialDataSource sources that are included in the sourceCollection
    noncustodialSources []NoncustodialDataSource;
}
// Instantiates a new sourceCollection and sets the default values.
func NewSourceCollection()(*SourceCollection) {
    m := &SourceCollection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the additionalSources property value. Adds an additional source to the sourceCollection.
func (m *SourceCollection) GetAdditionalSources()([]DataSource) {
    if m == nil {
        return nil
    } else {
        return m.additionalSources
    }
}
// Gets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
func (m *SourceCollection) GetAddToReviewSetOperation()(*AddToReviewSetOperation) {
    if m == nil {
        return nil
    } else {
        return m.addToReviewSetOperation
    }
}
// Gets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *SourceCollection) GetContentQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentQuery
    }
}
// Gets the createdBy property value. The user who created the sourceCollection.
func (m *SourceCollection) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The date and time the sourceCollection was created.
func (m *SourceCollection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the custodianSources property value. Custodian sources that are included in the sourceCollection.
func (m *SourceCollection) GetCustodianSources()([]DataSource) {
    if m == nil {
        return nil
    } else {
        return m.custodianSources
    }
}
// Gets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
func (m *SourceCollection) GetDataSourceScopes()(*i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceScopes) {
    if m == nil {
        return nil
    } else {
        return m.dataSourceScopes
    }
}
// Gets the description property value. The description of the sourceCollection.
func (m *SourceCollection) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name of the sourceCollection.
func (m *SourceCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
func (m *SourceCollection) GetLastEstimateStatisticsOperation()(*EstimateStatisticsOperation) {
    if m == nil {
        return nil
    } else {
        return m.lastEstimateStatisticsOperation
    }
}
// Gets the lastModifiedBy property value. The last user who modified the sourceCollection.
func (m *SourceCollection) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// Gets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
func (m *SourceCollection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
func (m *SourceCollection) GetNoncustodialSources()([]NoncustodialDataSource) {
    if m == nil {
        return nil
    } else {
        return m.noncustodialSources
    }
}
// The deserialization information for the current model
func (m *SourceCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSource() })
        if err != nil {
            return err
        }
        res := make([]DataSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*DataSource))
        }
        m.SetAdditionalSources(res)
        return nil
    }
    res["addToReviewSetOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAddToReviewSetOperation() })
        if err != nil {
            return err
        }
        m.SetAddToReviewSetOperation(val.(*AddToReviewSetOperation))
        return nil
    }
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
    res["custodianSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSource() })
        if err != nil {
            return err
        }
        res := make([]DataSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*DataSource))
        }
        m.SetCustodianSources(res)
        return nil
    }
    res["dataSourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.ParseDataSourceScopes)
        if err != nil {
            return err
        }
        cast := val.(i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceScopes)
        m.SetDataSourceScopes(&cast)
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
    res["lastEstimateStatisticsOperation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEstimateStatisticsOperation() })
        if err != nil {
            return err
        }
        m.SetLastEstimateStatisticsOperation(val.(*EstimateStatisticsOperation))
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
    res["noncustodialSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNoncustodialDataSource() })
        if err != nil {
            return err
        }
        res := make([]NoncustodialDataSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*NoncustodialDataSource))
        }
        m.SetNoncustodialSources(res)
        return nil
    }
    return res
}
func (m *SourceCollection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SourceCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalSources()))
        for i, v := range m.GetAdditionalSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("additionalSources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("addToReviewSetOperation", m.GetAddToReviewSetOperation())
        if err != nil {
            return err
        }
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustodianSources()))
        for i, v := range m.GetCustodianSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("custodianSources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDataSourceScopes() != nil {
        cast := m.GetDataSourceScopes().String()
        err = writer.WriteStringValue("dataSourceScopes", &cast)
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
        err = writer.WriteObjectValue("lastEstimateStatisticsOperation", m.GetLastEstimateStatisticsOperation())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNoncustodialSources()))
        for i, v := range m.GetNoncustodialSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("noncustodialSources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the additionalSources property value. Adds an additional source to the sourceCollection.
// Parameters:
//  - value : Value to set for the additionalSources property.
func (m *SourceCollection) SetAdditionalSources(value []DataSource)() {
    m.additionalSources = value
}
// Sets the addToReviewSetOperation property value. Adds the results of the sourceCollection to the specified reviewSet.
// Parameters:
//  - value : Value to set for the addToReviewSetOperation property.
func (m *SourceCollection) SetAddToReviewSetOperation(value *AddToReviewSetOperation)() {
    m.addToReviewSetOperation = value
}
// Sets the contentQuery property value. The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for Content Search and eDiscovery. You can refine searches by using fields paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
// Parameters:
//  - value : Value to set for the contentQuery property.
func (m *SourceCollection) SetContentQuery(value *string)() {
    m.contentQuery = value
}
// Sets the createdBy property value. The user who created the sourceCollection.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *SourceCollection) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The date and time the sourceCollection was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *SourceCollection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the custodianSources property value. Custodian sources that are included in the sourceCollection.
// Parameters:
//  - value : Value to set for the custodianSources property.
func (m *SourceCollection) SetCustodianSources(value []DataSource)() {
    m.custodianSources = value
}
// Sets the dataSourceScopes property value. When specified, the collection will span across a service for an entire workload. Possible values are: none, allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
// Parameters:
//  - value : Value to set for the dataSourceScopes property.
func (m *SourceCollection) SetDataSourceScopes(value *i2f6e0000f37d36d482a1a2be3e8024654d1a2bafb61d18de02f698fc19f7d94d.DataSourceScopes)() {
    m.dataSourceScopes = value
}
// Sets the description property value. The description of the sourceCollection.
// Parameters:
//  - value : Value to set for the description property.
func (m *SourceCollection) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name of the sourceCollection.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SourceCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastEstimateStatisticsOperation property value. The last estimate operation associated with the sourceCollection.
// Parameters:
//  - value : Value to set for the lastEstimateStatisticsOperation property.
func (m *SourceCollection) SetLastEstimateStatisticsOperation(value *EstimateStatisticsOperation)() {
    m.lastEstimateStatisticsOperation = value
}
// Sets the lastModifiedBy property value. The last user who modified the sourceCollection.
// Parameters:
//  - value : Value to set for the lastModifiedBy property.
func (m *SourceCollection) SetLastModifiedBy(value *IdentitySet)() {
    m.lastModifiedBy = value
}
// Sets the lastModifiedDateTime property value. The last date and time the sourceCollection was modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *SourceCollection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the noncustodialSources property value. noncustodialDataSource sources that are included in the sourceCollection
// Parameters:
//  - value : Value to set for the noncustodialSources property.
func (m *SourceCollection) SetNoncustodialSources(value []NoncustodialDataSource)() {
    m.noncustodialSources = value
}
