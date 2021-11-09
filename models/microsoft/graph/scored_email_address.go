package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new scoredEmailAddress and sets the default values.
func NewScoredEmailAddress()(*ScoredEmailAddress) {
    m := &ScoredEmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScoredEmailAddress) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the address property value. The email address.
func (m *ScoredEmailAddress) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the itemId property value. 
func (m *ScoredEmailAddress) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// Gets the relevanceScore property value. The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
func (m *ScoredEmailAddress) GetRelevanceScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.relevanceScore
    }
}
// Gets the selectionLikelihood property value. 
func (m *ScoredEmailAddress) GetSelectionLikelihood()(*SelectionLikelihoodInfo) {
    if m == nil {
        return nil
    } else {
        return m.selectionLikelihood
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ScoredEmailAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the address property value. The email address.
// Parameters:
//  - value : Value to set for the address property.
func (m *ScoredEmailAddress) SetAddress(value *string)() {
    m.address = value
}
// Sets the itemId property value. 
// Parameters:
//  - value : Value to set for the itemId property.
func (m *ScoredEmailAddress) SetItemId(value *string)() {
    m.itemId = value
}
// Sets the relevanceScore property value. The relevance score of the email address. A relevance score is used as a sort key, in relation to the other returned results. A higher relevance score value corresponds to a more relevant result. Relevance is determined by the user’s communication and collaboration patterns and business relationships.
// Parameters:
//  - value : Value to set for the relevanceScore property.
func (m *ScoredEmailAddress) SetRelevanceScore(value *float64)() {
    m.relevanceScore = value
}
// Sets the selectionLikelihood property value. 
// Parameters:
//  - value : Value to set for the selectionLikelihood property.
func (m *ScoredEmailAddress) SetSelectionLikelihood(value *SelectionLikelihoodInfo)() {
    m.selectionLikelihood = value
}
