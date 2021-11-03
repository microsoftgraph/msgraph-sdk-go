package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format '{field}:/'{aggregationFilterToken}/''. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format '{field}:or(/'{aggregationFilterToken1}/',/'{aggregationFilterToken2}/')'. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses.
    aggregationFilters []string;
    // Specifies aggregations (also known as refiners) to be returned alongside search results. Optional.
    aggregations []AggregationOption;
    // Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType=externalItem. Optional.
    contentSources []string;
    // This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType=message. Optional.
    enableTopResults *bool;
    // One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem. See known limitations for those combinations of two or more entity types that are supported in the same search request. Required.
    entityTypes []EntityType;
    // Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional.
    fields []string;
    // Specifies the offset for the search results. Offset 0 returns the very first result. Optional.
    from *int32;
    // 
    query *SearchQuery;
    // The size of the page to be retrieved. Optional.
    size *int32;
    // Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional.
    sortProperties []SortProperty;
}
// Instantiates a new searchRequest and sets the default values.
func NewSearchRequest()(*SearchRequest) {
    m := &SearchRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the aggregationFilters property value. Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format '{field}:/'{aggregationFilterToken}/''. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format '{field}:or(/'{aggregationFilterToken1}/',/'{aggregationFilterToken2}/')'. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses.
func (m *SearchRequest) GetAggregationFilters()([]string) {
    if m == nil {
        return nil
    } else {
        return m.aggregationFilters
    }
}
// Gets the aggregations property value. Specifies aggregations (also known as refiners) to be returned alongside search results. Optional.
func (m *SearchRequest) GetAggregations()([]AggregationOption) {
    if m == nil {
        return nil
    } else {
        return m.aggregations
    }
}
// Gets the contentSources property value. Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType=externalItem. Optional.
func (m *SearchRequest) GetContentSources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentSources
    }
}
// Gets the enableTopResults property value. This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType=message. Optional.
func (m *SearchRequest) GetEnableTopResults()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableTopResults
    }
}
// Gets the entityTypes property value. One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem. See known limitations for those combinations of two or more entity types that are supported in the same search request. Required.
func (m *SearchRequest) GetEntityTypes()([]EntityType) {
    if m == nil {
        return nil
    } else {
        return m.entityTypes
    }
}
// Gets the fields property value. Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional.
func (m *SearchRequest) GetFields()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// Gets the from property value. Specifies the offset for the search results. Offset 0 returns the very first result. Optional.
func (m *SearchRequest) GetFrom()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// Gets the query property value. 
func (m *SearchRequest) GetQuery()(*SearchQuery) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
// Gets the size property value. The size of the page to be retrieved. Optional.
func (m *SearchRequest) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Gets the sortProperties property value. Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional.
func (m *SearchRequest) GetSortProperties()([]SortProperty) {
    if m == nil {
        return nil
    } else {
        return m.sortProperties
    }
}
// The deserialization information for the current model
func (m *SearchRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aggregationFilters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetAggregationFilters(res)
        return nil
    }
    res["aggregations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAggregationOption() })
        if err != nil {
            return err
        }
        res := make([]AggregationOption, len(val))
        for i, v := range val {
            res[i] = *(v.(*AggregationOption))
        }
        m.SetAggregations(res)
        return nil
    }
    res["contentSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetContentSources(res)
        return nil
    }
    res["enableTopResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnableTopResults(val)
        return nil
    }
    res["entityTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseEntityType)
        if err != nil {
            return err
        }
        res := make([]EntityType, len(val))
        for i, v := range val {
            res[i] = *(v.(*EntityType))
        }
        m.SetEntityTypes(res)
        return nil
    }
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetFields(res)
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFrom(val)
        return nil
    }
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchQuery() })
        if err != nil {
            return err
        }
        m.SetQuery(val.(*SearchQuery))
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    res["sortProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSortProperty() })
        if err != nil {
            return err
        }
        res := make([]SortProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*SortProperty))
        }
        m.SetSortProperties(res)
        return nil
    }
    return res
}
func (m *SearchRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SearchRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("aggregationFilters", m.GetAggregationFilters())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAggregations()))
        for i, v := range m.GetAggregations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("aggregations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("contentSources", m.GetContentSources())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTopResults", m.GetEnableTopResults())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("entityTypes", SerializeEntityType(m.GetEntityTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSortProperties()))
        for i, v := range m.GetSortProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sortProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SearchRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the aggregationFilters property value. Contains one or more filters to obtain search results aggregated and filtered to a specific value of a field. Optional.Build this filter based on a prior search that aggregates by the same field. From the response of the prior search, identify the searchBucket that filters results to the specific value of the field, use the string in its aggregationFilterToken property, and build an aggregation filter string in the format '{field}:/'{aggregationFilterToken}/''. If multiple values for the same field need to be provided, use the strings in its aggregationFilterToken property and build an aggregation filter string in the format '{field}:or(/'{aggregationFilterToken1}/',/'{aggregationFilterToken2}/')'. For example, searching and aggregating drive items by file type returns a searchBucket for the file type docx in the response. You can conveniently use the aggregationFilterToken returned for this searchBucket in a subsequent search query and filter matches down to drive items of the docx file type. Example 1 and example 2 show the actual requests and responses.
// Parameters:
//  - value : Value to set for the aggregationFilters property.
func (m *SearchRequest) SetAggregationFilters(value []string)() {
    m.aggregationFilters = value
}
// Sets the aggregations property value. Specifies aggregations (also known as refiners) to be returned alongside search results. Optional.
// Parameters:
//  - value : Value to set for the aggregations property.
func (m *SearchRequest) SetAggregations(value []AggregationOption)() {
    m.aggregations = value
}
// Sets the contentSources property value. Contains the connection to be targeted. Respects the following format : /external/connections/connectionid where connectionid is the ConnectionId defined in the Connectors Administration.  Note: contentSource is only applicable when entityType=externalItem. Optional.
// Parameters:
//  - value : Value to set for the contentSources property.
func (m *SearchRequest) SetContentSources(value []string)() {
    m.contentSources = value
}
// Sets the enableTopResults property value. This triggers hybrid sort for messages: the first 3 messages are the most relevant. This property is only applicable to entityType=message. Optional.
// Parameters:
//  - value : Value to set for the enableTopResults property.
func (m *SearchRequest) SetEnableTopResults(value *bool)() {
    m.enableTopResults = value
}
// Sets the entityTypes property value. One or more types of resources expected in the response. Possible values are: list, site, listItem, message, event, drive, driveItem, person, externalItem. See known limitations for those combinations of two or more entity types that are supported in the same search request. Required.
// Parameters:
//  - value : Value to set for the entityTypes property.
func (m *SearchRequest) SetEntityTypes(value []EntityType)() {
    m.entityTypes = value
}
// Sets the fields property value. Contains the fields to be returned for each resource object specified in entityTypes, allowing customization of the fields returned by default otherwise, including additional fields such as custom managed properties from SharePoint and OneDrive, or custom fields in externalItem from content that Microsoft Graph connectors bring in. The fields property can be using the semantic labels applied to properties. For example, if a property is label as title, you can retrieve it using the following syntax : label_title.Optional.
// Parameters:
//  - value : Value to set for the fields property.
func (m *SearchRequest) SetFields(value []string)() {
    m.fields = value
}
// Sets the from property value. Specifies the offset for the search results. Offset 0 returns the very first result. Optional.
// Parameters:
//  - value : Value to set for the from property.
func (m *SearchRequest) SetFrom(value *int32)() {
    m.from = value
}
// Sets the query property value. 
// Parameters:
//  - value : Value to set for the query property.
func (m *SearchRequest) SetQuery(value *SearchQuery)() {
    m.query = value
}
// Sets the size property value. The size of the page to be retrieved. Optional.
// Parameters:
//  - value : Value to set for the size property.
func (m *SearchRequest) SetSize(value *int32)() {
    m.size = value
}
// Sets the sortProperties property value. Contains the ordered collection of fields and direction to sort results. There can be at most 5 sort properties in the collection. Optional.
// Parameters:
//  - value : Value to set for the sortProperties property.
func (m *SearchRequest) SetSortProperties(value []SortProperty)() {
    m.sortProperties = value
}
