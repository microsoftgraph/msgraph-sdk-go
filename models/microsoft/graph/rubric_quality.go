package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// rubricQuality 
type RubricQuality struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The collection of criteria for this rubric quality.
    criteria []RubricCriterion;
    // The description of this rubric quality.
    description *EducationItemBody;
    // The name of this rubric quality.
    displayName *string;
    // The ID of this resource.
    qualityId *string;
    // If present, a numerical weight for this quality.  Weights must add up to 100.
    weight *float32;
}
// NewRubricQuality instantiates a new rubricQuality and sets the default values.
func NewRubricQuality()(*RubricQuality) {
    m := &RubricQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCriteria gets the criteria property value. The collection of criteria for this rubric quality.
func (m *RubricQuality) GetCriteria()([]RubricCriterion) {
    if m == nil {
        return nil
    } else {
        return m.criteria
    }
}
// GetDescription gets the description property value. The description of this rubric quality.
func (m *RubricQuality) GetDescription()(*EducationItemBody) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of this rubric quality.
func (m *RubricQuality) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetQualityId gets the qualityId property value. The ID of this resource.
func (m *RubricQuality) GetQualityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qualityId
    }
}
// GetWeight gets the weight property value. If present, a numerical weight for this quality.  Weights must add up to 100.
func (m *RubricQuality) GetWeight()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.weight
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RubricQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["criteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRubricCriterion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricCriterion, len(val))
            for i, v := range val {
                res[i] = *(v.(*RubricCriterion))
            }
            m.SetCriteria(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(*EducationItemBody))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["qualityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityId(val)
        }
        return nil
    }
    res["weight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
func (m *RubricQuality) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RubricQuality) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCriteria()))
        for i, v := range m.GetCriteria() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("criteria", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("qualityId", m.GetQualityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("weight", m.GetWeight())
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
func (m *RubricQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCriteria sets the criteria property value. The collection of criteria for this rubric quality.
func (m *RubricQuality) SetCriteria(value []RubricCriterion)() {
    m.criteria = value
}
// SetDescription sets the description property value. The description of this rubric quality.
func (m *RubricQuality) SetDescription(value *EducationItemBody)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of this rubric quality.
func (m *RubricQuality) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetQualityId sets the qualityId property value. The ID of this resource.
func (m *RubricQuality) SetQualityId(value *string)() {
    m.qualityId = value
}
// SetWeight sets the weight property value. If present, a numerical weight for this quality.  Weights must add up to 100.
func (m *RubricQuality) SetWeight(value *float32)() {
    m.weight = value
}
