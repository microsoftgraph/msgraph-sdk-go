package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RubricLevel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The description of this rubric level.
    description *EducationItemBody;
    // The name of this rubric level.
    displayName *string;
    // Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
    grading *EducationAssignmentGradeType;
    // The ID of this resource.
    levelId *string;
}
// Instantiates a new rubricLevel and sets the default values.
func NewRubricLevel()(*RubricLevel) {
    m := &RubricLevel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricLevel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. The description of this rubric level.
func (m *RubricLevel) GetDescription()(*EducationItemBody) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of this rubric level.
func (m *RubricLevel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the grading property value. Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
func (m *RubricLevel) GetGrading()(*EducationAssignmentGradeType) {
    if m == nil {
        return nil
    } else {
        return m.grading
    }
}
// Gets the levelId property value. The ID of this resource.
func (m *RubricLevel) GetLevelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.levelId
    }
}
// The deserialization information for the current model
func (m *RubricLevel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["grading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationAssignmentGradeType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrading(val.(*EducationAssignmentGradeType))
        }
        return nil
    }
    res["levelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevelId(val)
        }
        return nil
    }
    return res
}
func (m *RubricLevel) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RubricLevel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("grading", m.GetGrading())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("levelId", m.GetLevelId())
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
func (m *RubricLevel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. The description of this rubric level.
// Parameters:
//  - value : Value to set for the description property.
func (m *RubricLevel) SetDescription(value *EducationItemBody)() {
    m.description = value
}
// Sets the displayName property value. The name of this rubric level.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *RubricLevel) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the grading property value. Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
// Parameters:
//  - value : Value to set for the grading property.
func (m *RubricLevel) SetGrading(value *EducationAssignmentGradeType)() {
    m.grading = value
}
// Sets the levelId property value. The ID of this resource.
// Parameters:
//  - value : Value to set for the levelId property.
func (m *RubricLevel) SetLevelId(value *string)() {
    m.levelId = value
}
