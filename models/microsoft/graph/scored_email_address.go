package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// scoredEmailAddress 
type ScoredEmailAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The email address.
    address *string;
    // 
    itemId *string;
    // The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
    relevanceScore *float64;
    // 
    selectionLikelihood *SelectionLikelihoodInfo;
}
// NewScoredEmailAddress instantiates a new scoredEmailAddress and sets the default values.
func NewScoredEmailAddress()(*ScoredEmailAddress) {
    m := &ScoredEmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScoredEmailAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. The email address.
func (m *ScoredEmailAddress) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetItemId gets the itemId property value. 
func (m *ScoredEmailAddress) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// GetRelevanceScore gets the relevanceScore property value. The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
func (m *ScoredEmailAddress) GetRelevanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.relevanceScore
    }
}
// GetSelectionLikelihood gets the selectionLikelihood property value. 
func (m *ScoredEmailAddress) GetSelectionLikelihood()(*SelectionLikelihoodInfo) {
    if m == nil {
        return nil
    } else {
        return m.selectionLikelihood
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScoredEmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["itemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["relevanceScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelevanceScore(val)
        }
        return nil
    }
    res["selectionLikelihood"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSelectionLikelihoodInfo)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SelectionLikelihoodInfo)
            m.SetSelectionLikelihood(&cast)
        }
        return nil
    }
    return res
}
func (m *ScoredEmailAddress) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScoredEmailAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddress sets the address property value. The email address.
func (m *ScoredEmailAddress) SetAddress(value *string)() {
    m.address = value
}
// SetItemId sets the itemId property value. 
func (m *ScoredEmailAddress) SetItemId(value *string)() {
    m.itemId = value
}
// SetRelevanceScore sets the relevanceScore property value. The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
func (m *ScoredEmailAddress) SetRelevanceScore(value *float64)() {
    m.relevanceScore = value
}
// SetSelectionLikelihood sets the selectionLikelihood property value. 
func (m *ScoredEmailAddress) SetSelectionLikelihood(value *SelectionLikelihoodInfo)() {
    m.selectionLikelihood = value
}
