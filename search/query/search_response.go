package query

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// SearchResponse 
type SearchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A collection of search results.
    hitsContainers []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SearchHitsContainer;
    // Provides details of query alteration response for spelling correction.
    queryAlterationResponse *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AlterationResponse;
    // A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
    resultTemplates *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResultTemplateDictionary;
    // Contains the search terms sent in the initial search query.
    searchTerms []string;
}
// NewSearchResponse instantiates a new SearchResponse and sets the default values.
func NewSearchResponse()(*SearchResponse) {
    m := &SearchResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHitsContainers gets the hitsContainers property value. A collection of search results.
func (m *SearchResponse) GetHitsContainers()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SearchHitsContainer) {
    if m == nil {
        return nil
    } else {
        return m.hitsContainers
    }
}
// GetQueryAlterationResponse gets the queryAlterationResponse property value. Provides details of query alteration response for spelling correction.
func (m *SearchResponse) GetQueryAlterationResponse()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AlterationResponse) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationResponse
    }
}
// GetResultTemplates gets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResponse) GetResultTemplates()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResultTemplateDictionary) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplates
    }
}
// GetSearchTerms gets the searchTerms property value. Contains the search terms sent in the initial search query.
func (m *SearchResponse) GetSearchTerms()([]string) {
    if m == nil {
        return nil
    } else {
        return m.searchTerms
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hitsContainers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSearchHitsContainer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SearchHitsContainer, len(val))
            for i, v := range val {
                res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SearchHitsContainer))
            }
            m.SetHitsContainers(res)
        }
        return nil
    }
    res["queryAlterationResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAlterationResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationResponse(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AlterationResponse))
        }
        return nil
    }
    res["resultTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewResultTemplateDictionary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplates(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResultTemplateDictionary))
        }
        return nil
    }
    res["searchTerms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSearchTerms(res)
        }
        return nil
    }
    return res
}
func (m *SearchResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHitsContainers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHitsContainers()))
        for i, v := range m.GetHitsContainers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("hitsContainers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("queryAlterationResponse", m.GetQueryAlterationResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resultTemplates", m.GetResultTemplates())
        if err != nil {
            return err
        }
    }
    if m.GetSearchTerms() != nil {
        err := writer.WriteCollectionOfStringValues("searchTerms", m.GetSearchTerms())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHitsContainers sets the hitsContainers property value. A collection of search results.
func (m *SearchResponse) SetHitsContainers(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SearchHitsContainer)() {
    if m != nil {
        m.hitsContainers = value
    }
}
// SetQueryAlterationResponse sets the queryAlterationResponse property value. Provides details of query alteration response for spelling correction.
func (m *SearchResponse) SetQueryAlterationResponse(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.AlterationResponse)() {
    if m != nil {
        m.queryAlterationResponse = value
    }
}
// SetResultTemplates sets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResponse) SetResultTemplates(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResultTemplateDictionary)() {
    if m != nil {
        m.resultTemplates = value
    }
}
// SetSearchTerms sets the searchTerms property value. Contains the search terms sent in the initial search query.
func (m *SearchResponse) SetSearchTerms(value []string)() {
    if m != nil {
        m.searchTerms = value
    }
}
