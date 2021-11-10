package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RedundancyDetectionSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether email threading and near duplicate detection are enabled.
    isEnabled *bool;
    // See Minimum/maximum number of words to learn more.
    maxWords *int32;
    // See Minimum/maximum number of words to learn more.
    minWords *int32;
    // See Document and email similarity threshold to learn more.
    similarityThreshold *int32;
}
// Instantiates a new redundancyDetectionSettings and sets the default values.
func NewRedundancyDetectionSettings()(*RedundancyDetectionSettings) {
    m := &RedundancyDetectionSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedundancyDetectionSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isEnabled property value. Indicates whether email threading and near duplicate detection are enabled.
func (m *RedundancyDetectionSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the maxWords property value. See Minimum/maximum number of words to learn more.
func (m *RedundancyDetectionSettings) GetMaxWords()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxWords
    }
}
// Gets the minWords property value. See Minimum/maximum number of words to learn more.
func (m *RedundancyDetectionSettings) GetMinWords()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minWords
    }
}
// Gets the similarityThreshold property value. See Document and email similarity threshold to learn more.
func (m *RedundancyDetectionSettings) GetSimilarityThreshold()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.similarityThreshold
    }
}
// The deserialization information for the current model
func (m *RedundancyDetectionSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["maxWords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxWords(val)
        }
        return nil
    }
    res["minWords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinWords(val)
        }
        return nil
    }
    res["similarityThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimilarityThreshold(val)
        }
        return nil
    }
    return res
}
func (m *RedundancyDetectionSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RedundancyDetectionSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxWords", m.GetMaxWords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minWords", m.GetMinWords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("similarityThreshold", m.GetSimilarityThreshold())
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
func (m *RedundancyDetectionSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isEnabled property value. Indicates whether email threading and near duplicate detection are enabled.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *RedundancyDetectionSettings) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the maxWords property value. See Minimum/maximum number of words to learn more.
// Parameters:
//  - value : Value to set for the maxWords property.
func (m *RedundancyDetectionSettings) SetMaxWords(value *int32)() {
    m.maxWords = value
}
// Sets the minWords property value. See Minimum/maximum number of words to learn more.
// Parameters:
//  - value : Value to set for the minWords property.
func (m *RedundancyDetectionSettings) SetMinWords(value *int32)() {
    m.minWords = value
}
// Sets the similarityThreshold property value. See Document and email similarity threshold to learn more.
// Parameters:
//  - value : Value to set for the similarityThreshold property.
func (m *RedundancyDetectionSettings) SetSimilarityThreshold(value *int32)() {
    m.similarityThreshold = value
}
