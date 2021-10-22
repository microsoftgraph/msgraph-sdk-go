package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchHitsContainer struct {
    additionalData map[string]interface{};
    hits []SearchHit;
    moreResultsAvailable *bool;
    total *int32;
}
func NewSearchHitsContainer()(*SearchHitsContainer) {
    m := &SearchHitsContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchHitsContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchHitsContainer) GetHits()([]SearchHit) {
    if m == nil {
        return nil
    } else {
        return m.hits
    }
}
func (m *SearchHitsContainer) GetMoreResultsAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.moreResultsAvailable
    }
}
func (m *SearchHitsContainer) GetTotal()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
func (m *SearchHitsContainer) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchHit() })
        if err != nil {
            return err
        }
        res := make([]SearchHit, len(val))
        for i, v := range val {
            res[i] = *(v.(*SearchHit))
        }
        m.SetHits(res)
        return nil
    }
    res["moreResultsAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMoreResultsAvailable(val)
        return nil
    }
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotal(val)
        return nil
    }
    return res
}
func (m *SearchHitsContainer) IsNil()(bool) {
    return m == nil
}
func (m *SearchHitsContainer) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *SearchHitsContainer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchHitsContainer) SetHits(value []SearchHit)() {
    m.hits = value
}
func (m *SearchHitsContainer) SetMoreResultsAvailable(value *bool)() {
    m.moreResultsAvailable = value
}
func (m *SearchHitsContainer) SetTotal(value *int32)() {
    m.total = value
}
