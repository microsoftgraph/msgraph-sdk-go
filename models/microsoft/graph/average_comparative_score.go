package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AverageComparativeScore struct {
    additionalData map[string]interface{};
    averageScore *float64;
    basis *string;
}
func NewAverageComparativeScore()(*AverageComparativeScore) {
    m := &AverageComparativeScore{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AverageComparativeScore) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AverageComparativeScore) GetAverageScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.averageScore
    }
}
func (m *AverageComparativeScore) GetBasis()(*string) {
    if m == nil {
        return nil
    } else {
        return m.basis
    }
}
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
func (m *AverageComparativeScore) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AverageComparativeScore) SetAverageScore(value *float64)() {
    m.averageScore = value
}
func (m *AverageComparativeScore) SetBasis(value *string)() {
    m.basis = value
}
