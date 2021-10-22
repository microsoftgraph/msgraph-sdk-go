package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ScoredEmailAddress struct {
    additionalData map[string]interface{};
    address *string;
    itemId *string;
    relevanceScore *float64;
    selectionLikelihood *SelectionLikelihoodInfo;
}
func NewScoredEmailAddress()(*ScoredEmailAddress) {
    m := &ScoredEmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ScoredEmailAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ScoredEmailAddress) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
func (m *ScoredEmailAddress) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
func (m *ScoredEmailAddress) GetRelevanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.relevanceScore
    }
}
func (m *ScoredEmailAddress) GetSelectionLikelihood()(*SelectionLikelihoodInfo) {
    if m == nil {
        return nil
    } else {
        return m.selectionLikelihood
    }
}
func (m *ScoredEmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddress(val)
        return nil
    }
    res["itemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetItemId(val)
        return nil
    }
    res["relevanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetRelevanceScore(val)
        return nil
    }
    res["selectionLikelihood"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSelectionLikelihoodInfo)
        if err != nil {
            return err
        }
        cast := val.(SelectionLikelihoodInfo)
        m.SetSelectionLikelihood(&cast)
        return nil
    }
    return res
}
func (m *ScoredEmailAddress) IsNil()(bool) {
    return m == nil
}
func (m *ScoredEmailAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("relevanceScore", m.GetRelevanceScore())
        if err != nil {
            return err
        }
    }
    if m.GetSelectionLikelihood() != nil {
        cast := m.GetSelectionLikelihood().String()
        err := writer.WriteStringValue("selectionLikelihood", &cast)
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
func (m *ScoredEmailAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ScoredEmailAddress) SetAddress(value *string)() {
    m.address = value
}
func (m *ScoredEmailAddress) SetItemId(value *string)() {
    m.itemId = value
}
func (m *ScoredEmailAddress) SetRelevanceScore(value *float64)() {
    m.relevanceScore = value
}
func (m *ScoredEmailAddress) SetSelectionLikelihood(value *SelectionLikelihoodInfo)() {
    m.selectionLikelihood = value
}
