package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchHit provides operations to call the query method.
type SearchHit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of the content source which the externalItem is part of .
    contentSource *string;
    // The internal identifier for the item.
    hitId *string;
    // The rank or the order of the result.
    rank *int32;
    // 
    resource Entityable;
    // ID of the result template used to render the search result. This ID must map to a display layout in the resultTemplates dictionary that is also included in the searchResponse.
    resultTemplateId *string;
    // A summary of the result, if a summary is available.
    summary *string;
}
// NewSearchHit instantiates a new searchHit and sets the default values.
func NewSearchHit()(*SearchHit) {
    m := &SearchHit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchHitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchHitFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSearchHit(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentSource gets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) GetContentSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentSource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchHit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentSource(val)
        }
        return nil
    }
    res["hitId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHitId(val)
        }
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(Entityable))
        }
        return nil
    }
    res["resultTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplateId(val)
        }
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
// GetHitId gets the hitId property value. The internal identifier for the item.
func (m *SearchHit) GetHitId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hitId
    }
}
// GetRank gets the rank property value. The rank or the order of the result.
func (m *SearchHit) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// GetResource gets the resource property value. 
func (m *SearchHit) GetResource()(Entityable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResultTemplateId gets the resultTemplateId property value. ID of the result template used to render the search result. This ID must map to a display layout in the resultTemplates dictionary that is also included in the searchResponse.
func (m *SearchHit) GetResultTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultTemplateId
    }
}
// GetSummary gets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
func (m *SearchHit) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchHit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentSource", m.GetContentSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hitId", m.GetHitId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resultTemplateId", m.GetResultTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
func (m *SearchHit) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentSource sets the contentSource property value. The name of the content source which the externalItem is part of .
func (m *SearchHit) SetContentSource(value *string)() {
    if m != nil {
        m.contentSource = value
    }
}
// SetHitId sets the hitId property value. The internal identifier for the item.
func (m *SearchHit) SetHitId(value *string)() {
    if m != nil {
        m.hitId = value
    }
}
// SetRank sets the rank property value. The rank or the order of the result.
func (m *SearchHit) SetRank(value *int32)() {
    if m != nil {
        m.rank = value
    }
}
// SetResource sets the resource property value. 
func (m *SearchHit) SetResource(value Entityable)() {
    if m != nil {
        m.resource = value
    }
}
// SetResultTemplateId sets the resultTemplateId property value. ID of the result template used to render the search result. This ID must map to a display layout in the resultTemplates dictionary that is also included in the searchResponse.
func (m *SearchHit) SetResultTemplateId(value *string)() {
    if m != nil {
        m.resultTemplateId = value
    }
}
// SetSummary sets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) SetSummary(value *string)() {
    if m != nil {
        m.summary = value
    }
}
