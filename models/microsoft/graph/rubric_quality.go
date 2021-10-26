package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new rubricQuality and sets the default values.
func NewRubricQuality()(*RubricQuality) {
    m := &RubricQuality{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricQuality) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the criteria property value. The collection of criteria for this rubric quality.
func (m *RubricQuality) GetCriteria()([]RubricCriterion) {
    if m == nil {
        return nil
    } else {
        return m.criteria
    }
}
// Gets the description property value. The description of this rubric quality.
func (m *RubricQuality) GetDescription()(*EducationItemBody) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of this rubric quality.
func (m *RubricQuality) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the qualityId property value. The ID of this resource.
func (m *RubricQuality) GetQualityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qualityId
    }
}
// Gets the weight property value. If present, a numerical weight for this quality.  Weights must add up to 100.
func (m *RubricQuality) GetWeight()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.weight
    }
}
// The deserialization information for the current model
func (m *RubricQuality) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["criteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRubricCriterion() })
        if err != nil {
            return err
        }
        res := make([]RubricCriterion, len(val))
        for i, v := range val {
            res[i] = *(v.(*RubricCriterion))
        }
        m.SetCriteria(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationItemBody() })
        if err != nil {
            return err
        }
        m.SetDescription(val.(*EducationItemBody))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["qualityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQualityId(val)
        return nil
    }
    res["weight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetWeight(val)
        return nil
    }
    return res
}
func (m *RubricQuality) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RubricQuality) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the criteria property value. The collection of criteria for this rubric quality.
// Parameters:
//  - value : Value to set for the criteria property.
func (m *RubricQuality) SetCriteria(value []RubricCriterion)() {
    m.criteria = value
}
// Sets the description property value. The description of this rubric quality.
// Parameters:
//  - value : Value to set for the description property.
func (m *RubricQuality) SetDescription(value *EducationItemBody)() {
    m.description = value
}
// Sets the displayName property value. The name of this rubric quality.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *RubricQuality) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the qualityId property value. The ID of this resource.
// Parameters:
//  - value : Value to set for the qualityId property.
func (m *RubricQuality) SetQualityId(value *string)() {
    m.qualityId = value
}
// Sets the weight property value. If present, a numerical weight for this quality.  Weights must add up to 100.
// Parameters:
//  - value : Value to set for the weight property.
func (m *RubricQuality) SetWeight(value *float32)() {
    m.weight = value
}
