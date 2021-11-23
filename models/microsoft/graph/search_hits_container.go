package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchHitsContainer 
type SearchHitsContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Contains the collection of aggregations computed based on the provided aggregationOption specified in the request.
    aggregations []SearchAggregation;
    // A collection of the search results.
    hits []SearchHit;
    // Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly.
    moreResultsAvailable *bool;
    // The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.
    total *int32;
}
// NewSearchHitsContainer instantiates a new searchHitsContainer and sets the default values.
func NewSearchHitsContainer()(*SearchHitsContainer) {
    m := &SearchHitsContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHitsContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAggregations gets the aggregations property value. Contains the collection of aggregations computed based on the provided aggregationOption specified in the request.
func (m *SearchHitsContainer) GetAggregations()([]SearchAggregation) {
    if m == nil {
        return nil
    } else {
        return m.aggregations
    }
}
// GetHits gets the hits property value. A collection of the search results.
func (m *SearchHitsContainer) GetHits()([]SearchHit) {
    if m == nil {
        return nil
    } else {
        return m.hits
    }
}
// GetMoreResultsAvailable gets the moreResultsAvailable property value. Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly.
func (m *SearchHitsContainer) GetMoreResultsAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.moreResultsAvailable
    }
}
// GetTotal gets the total property value. The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.
func (m *SearchHitsContainer) GetTotal()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchHitsContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aggregations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchAggregation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchAggregation, len(val))
            for i, v := range val {
                res[i] = *(v.(*SearchAggregation))
            }
            m.SetAggregations(res)
        }
        return nil
    }
    res["hits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchHit() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchHit, len(val))
            for i, v := range val {
                res[i] = *(v.(*SearchHit))
            }
            m.SetHits(res)
        }
        return nil
    }
    res["moreResultsAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoreResultsAvailable(val)
        }
        return nil
    }
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
func (m *SearchHitsContainer) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchHitsContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHits()))
        for i, v := range m.GetHits() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("hits", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("moreResultsAvailable", m.GetMoreResultsAvailable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *SearchHitsContainer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAggregations sets the aggregations property value. Contains the collection of aggregations computed based on the provided aggregationOption specified in the request.
func (m *SearchHitsContainer) SetAggregations(value []SearchAggregation)() {
    m.aggregations = value
}
// SetHits sets the hits property value. A collection of the search results.
func (m *SearchHitsContainer) SetHits(value []SearchHit)() {
    m.hits = value
}
// SetMoreResultsAvailable sets the moreResultsAvailable property value. Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly.
func (m *SearchHitsContainer) SetMoreResultsAvailable(value *bool)() {
    m.moreResultsAvailable = value
}
// SetTotal sets the total property value. The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.
func (m *SearchHitsContainer) SetTotal(value *int32)() {
    m.total = value
}
