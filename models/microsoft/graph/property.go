package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// property 
type Property struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
    aliases []string;
    // Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional.
    isQueryable *bool;
    // Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional.
    isRefinable *bool;
    // Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional.
    isRetrievable *bool;
    // Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional.
    isSearchable *bool;
    // Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional.
    labels []i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label;
    // The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^.  Required.
    name *string;
    // The data type of the property. Possible values are: string, int64, double, dateTime, boolean, stringCollection, int64Collection, doubleCollection, dateTimeCollection, unknownFutureValue.
    type_escaped *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType;
}
// NewProperty instantiates a new property and sets the default values.
func NewProperty()(*Property) {
    m := &Property{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Property) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAliases gets the aliases property value. A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
func (m *Property) GetAliases()([]string) {
    if m == nil {
        return nil
    } else {
        return m.aliases
    }
}
// GetIsQueryable gets the isQueryable property value. Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional.
func (m *Property) GetIsQueryable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isQueryable
    }
}
// GetIsRefinable gets the isRefinable property value. Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional.
func (m *Property) GetIsRefinable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRefinable
    }
}
// GetIsRetrievable gets the isRetrievable property value. Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional.
func (m *Property) GetIsRetrievable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRetrievable
    }
}
// GetIsSearchable gets the isSearchable property value. Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional.
func (m *Property) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetLabels gets the labels property value. Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional.
func (m *Property) GetLabels()([]i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
// GetName gets the name property value. The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^.  Required.
func (m *Property) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType_escaped gets the type_escaped property value. The data type of the property. Possible values are: string, int64, double, dateTime, boolean, stringCollection, int64Collection, doubleCollection, dateTimeCollection, unknownFutureValue.
func (m *Property) GetType_escaped()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Property) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aliases"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAliases(res)
        }
        return nil
    }
    res["isQueryable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsQueryable(val)
        }
        return nil
    }
    res["isRefinable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRefinable(val)
        }
        return nil
    }
    res["isRetrievable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRetrievable(val)
        }
        return nil
    }
    res["isSearchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["labels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseLabel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label, len(val))
            for i, v := range val {
                res[i] = *(v.(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label))
            }
            m.SetLabels(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParsePropertyType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *Property) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Property) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("aliases", m.GetAliases())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQueryable", m.GetIsQueryable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRefinable", m.GetIsRefinable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRetrievable", m.GetIsRetrievable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("labels", i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.SerializeLabel(m.GetLabels()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Property) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAliases sets the aliases property value. A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
func (m *Property) SetAliases(value []string)() {
    m.aliases = value
}
// SetIsQueryable sets the isQueryable property value. Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional.
func (m *Property) SetIsQueryable(value *bool)() {
    m.isQueryable = value
}
// SetIsRefinable sets the isRefinable property value. Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional.
func (m *Property) SetIsRefinable(value *bool)() {
    m.isRefinable = value
}
// SetIsRetrievable sets the isRetrievable property value. Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional.
func (m *Property) SetIsRetrievable(value *bool)() {
    m.isRetrievable = value
}
// SetIsSearchable sets the isSearchable property value. Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional.
func (m *Property) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetLabels sets the labels property value. Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional.
func (m *Property) SetLabels(value []i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.Label)() {
    m.labels = value
}
// SetName sets the name property value. The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^.  Required.
func (m *Property) SetName(value *string)() {
    m.name = value
}
// SetType_escaped sets the type_escaped property value. The data type of the property. Possible values are: string, int64, double, dateTime, boolean, stringCollection, int64Collection, doubleCollection, dateTimeCollection, unknownFutureValue.
func (m *Property) SetType_escaped(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.PropertyType)() {
    m.type_escaped = value
}
