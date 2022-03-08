package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AverageComparativeScore provides operations to manage the security singleton.
type AverageComparativeScore struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Average score within specified basis.
    averageScore *float64;
    // Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
    basis *string;
}
// NewAverageComparativeScore instantiates a new averageComparativeScore and sets the default values.
func NewAverageComparativeScore()(*AverageComparativeScore) {
    m := &AverageComparativeScore{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAverageComparativeScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAverageComparativeScoreFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAverageComparativeScore(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AverageComparativeScore) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAverageScore gets the averageScore property value. Average score within specified basis.
func (m *AverageComparativeScore) GetAverageScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageScore
    }
}
// GetBasis gets the basis property value. Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
func (m *AverageComparativeScore) GetBasis()(*string) {
    if m == nil {
        return nil
    } else {
        return m.basis
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AverageComparativeScore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["averageScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageScore(val)
        }
        return nil
    }
    res["basis"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasis(val)
        }
        return nil
    }
    return res
}
func (m *AverageComparativeScore) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AverageComparativeScore) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAverageScore sets the averageScore property value. Average score within specified basis.
func (m *AverageComparativeScore) SetAverageScore(value *float64)() {
    if m != nil {
        m.averageScore = value
    }
}
// SetBasis sets the basis property value. Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
func (m *AverageComparativeScore) SetBasis(value *string)() {
    if m != nil {
        m.basis = value
    }
}
