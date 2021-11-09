package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CaseSettings struct {
    Entity
    // The OCR (Optical Character Recognition) settings for the case.
    ocr *OcrSettings;
    // The redundancy (near duplicate and email threading) detection settings for the case.
    redundancyDetection *RedundancyDetectionSettings;
    // The Topic Modeling (Themes) settings for the case.
    topicModeling *TopicModelingSettings;
}
// Instantiates a new caseSettings and sets the default values.
func NewCaseSettings()(*CaseSettings) {
    m := &CaseSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the ocr property value. The OCR (Optical Character Recognition) settings for the case.
func (m *CaseSettings) GetOcr()(*OcrSettings) {
    if m == nil {
        return nil
    } else {
        return m.ocr
    }
}
// Gets the redundancyDetection property value. The redundancy (near duplicate and email threading) detection settings for the case.
func (m *CaseSettings) GetRedundancyDetection()(*RedundancyDetectionSettings) {
    if m == nil {
        return nil
    } else {
        return m.redundancyDetection
    }
}
// Gets the topicModeling property value. The Topic Modeling (Themes) settings for the case.
func (m *CaseSettings) GetTopicModeling()(*TopicModelingSettings) {
    if m == nil {
        return nil
    } else {
        return m.topicModeling
    }
}
// The deserialization information for the current model
func (m *CaseSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ocr"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOcrSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOcr(val.(*OcrSettings))
        }
        return nil
    }
    res["redundancyDetection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRedundancyDetectionSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedundancyDetection(val.(*RedundancyDetectionSettings))
        }
        return nil
    }
    res["topicModeling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTopicModelingSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicModeling(val.(*TopicModelingSettings))
        }
        return nil
    }
    return res
}
func (m *CaseSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CaseSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("ocr", m.GetOcr())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("redundancyDetection", m.GetRedundancyDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("topicModeling", m.GetTopicModeling())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the ocr property value. The OCR (Optical Character Recognition) settings for the case.
// Parameters:
//  - value : Value to set for the ocr property.
func (m *CaseSettings) SetOcr(value *OcrSettings)() {
    m.ocr = value
}
// Sets the redundancyDetection property value. The redundancy (near duplicate and email threading) detection settings for the case.
// Parameters:
//  - value : Value to set for the redundancyDetection property.
func (m *CaseSettings) SetRedundancyDetection(value *RedundancyDetectionSettings)() {
    m.redundancyDetection = value
}
// Sets the topicModeling property value. The Topic Modeling (Themes) settings for the case.
// Parameters:
//  - value : Value to set for the topicModeling property.
func (m *CaseSettings) SetTopicModeling(value *TopicModelingSettings)() {
    m.topicModeling = value
}
