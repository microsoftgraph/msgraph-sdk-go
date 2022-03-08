package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchResponse provides operations to call the query method.
type SearchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A collection of search results.
    hitsContainers []SearchHitsContainerable;
    // Provides information related to spelling corrections in the alteration response.
    queryAlterationResponse AlterationResponseable;
    // A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
    resultTemplates ResultTemplateDictionaryable;
    // Contains the search terms sent in the initial search query.
    searchTerms []string;
}
// NewSearchResponse instantiates a new searchResponse and sets the default values.
func NewSearchResponse()(*SearchResponse) {
    m := &SearchResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSearchResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hitsContainers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSearchHitsContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchHitsContainerable, len(val))
            for i, v := range val {
                res[i] = v.(SearchHitsContainerable)
            }
            m.SetHitsContainers(res)
        }
        return nil
    }
    res["queryAlterationResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlterationResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationResponse(val.(AlterationResponseable))
        }
        return nil
    }
    res["resultTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultTemplateDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplates(val.(ResultTemplateDictionaryable))
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
// GetHitsContainers gets the hitsContainers property value. A collection of search results.
func (m *SearchResponse) GetHitsContainers()([]SearchHitsContainerable) {
    if m == nil {
        return nil
    } else {
        return m.hitsContainers
    }
}
// GetQueryAlterationResponse gets the queryAlterationResponse property value. Provides information related to spelling corrections in the alteration response.
func (m *SearchResponse) GetQueryAlterationResponse()(AlterationResponseable) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationResponse
    }
}
// GetResultTemplates gets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResponse) GetResultTemplates()(ResultTemplateDictionaryable) {
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
func (m *SearchResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHitsContainers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHitsContainers()))
        for i, v := range m.GetHitsContainers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *SearchResponse) SetHitsContainers(value []SearchHitsContainerable)() {
    if m != nil {
        m.hitsContainers = value
    }
}
// SetQueryAlterationResponse sets the queryAlterationResponse property value. Provides information related to spelling corrections in the alteration response.
func (m *SearchResponse) SetQueryAlterationResponse(value AlterationResponseable)() {
    if m != nil {
        m.queryAlterationResponse = value
    }
}
// SetResultTemplates sets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResponse) SetResultTemplates(value ResultTemplateDictionaryable)() {
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
