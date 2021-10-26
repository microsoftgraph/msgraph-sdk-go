package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AverageComparativeScore struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Average score within specified basis.
    averageScore *float64;
    // Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
    basis *string;
}
// Instantiates a new averageComparativeScore and sets the default values.
func NewAverageComparativeScore()(*AverageComparativeScore) {
    m := &AverageComparativeScore{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AverageComparativeScore) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the averageScore property value. Average score within specified basis.
func (m *AverageComparativeScore) GetAverageScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageScore
    }
}
// Gets the basis property value. Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
func (m *AverageComparativeScore) GetBasis()(*string) {
    if m == nil {
        return nil
    } else {
        return m.basis
    }
}
// The deserialization information for the current model
func (m *AverageComparativeScore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["averageScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetAverageScore(val)
        return nil
    }
    res["basis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBasis(val)
        return nil
    }
    return res
}
func (m *AverageComparativeScore) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AverageComparativeScore) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("averageScore", m.GetAverageScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("basis", m.GetBasis())
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
func (m *AverageComparativeScore) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the averageScore property value. Average score within specified basis.
// Parameters:
//  - value : Value to set for the averageScore property.
func (m *AverageComparativeScore) SetAverageScore(value *float64)() {
    m.averageScore = value
}
// Sets the basis property value. Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
// Parameters:
//  - value : Value to set for the basis property.
func (m *AverageComparativeScore) SetBasis(value *string)() {
    m.basis = value
}
