package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchQuery 
type SearchQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The search query containing the search terms. Required.
    queryString *string;
}
// NewSearchQuery instantiates a new searchQuery and sets the default values.
func NewSearchQuery()(*SearchQuery) {
    m := &SearchQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchQuery) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetQueryString gets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) GetQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryString
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchQuery) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["queryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryString(val)
        }
        return nil
    }
    return res
}
func (m *SearchQuery) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchQuery) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("queryString", m.GetQueryString())
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
func (m *SearchQuery) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQueryString sets the queryString property value. The search query containing the search terms. Required.
func (m *SearchQuery) SetQueryString(value *string)() {
    if m != nil {
        m.queryString = value
    }
}
