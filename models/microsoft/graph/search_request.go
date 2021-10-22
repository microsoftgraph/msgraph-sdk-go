package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchRequest struct {
    additionalData map[string]interface{};
    contentSources []string;
    enableTopResults *bool;
    entityTypes []EntityType;
    fields []string;
    from *int32;
    query *SearchQuery;
    size *int32;
}
func NewSearchRequest()(*SearchRequest) {
    m := &SearchRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchRequest) GetContentSources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.contentSources
    }
}
func (m *SearchRequest) GetEnableTopResults()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableTopResults
    }
}
func (m *SearchRequest) GetEntityTypes()([]EntityType) {
    if m == nil {
        return nil
    } else {
        return m.entityTypes
    }
}
func (m *SearchRequest) GetFields()([]string) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
func (m *SearchRequest) GetFrom()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
func (m *SearchRequest) GetQuery()(*SearchQuery) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
func (m *SearchRequest) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *SearchRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
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
            res[i] = v.(string)
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
    return res
}
func (m *SearchRequest) IsNil()(bool) {
    return m == nil
}
func (m *SearchRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SearchRequest) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchRequest) SetContentSources(value []string)() {
    m.contentSources = value
}
func (m *SearchRequest) SetEnableTopResults(value *bool)() {
    m.enableTopResults = value
}
func (m *SearchRequest) SetEntityTypes(value []EntityType)() {
    m.entityTypes = value
}
func (m *SearchRequest) SetFields(value []string)() {
    m.fields = value
}
func (m *SearchRequest) SetFrom(value *int32)() {
    m.from = value
}
func (m *SearchRequest) SetQuery(value *SearchQuery)() {
    m.query = value
}
func (m *SearchRequest) SetSize(value *int32)() {
    m.size = value
}
