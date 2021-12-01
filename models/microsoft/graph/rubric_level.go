package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RubricLevel 
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
// NewRubricLevel instantiates a new rubricLevel and sets the default values.
func NewRubricLevel()(*RubricLevel) {
    m := &RubricLevel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricLevel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. The description of this rubric level.
func (m *RubricLevel) GetDescription()(*EducationItemBody) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of this rubric level.
func (m *RubricLevel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGrading gets the grading property value. Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
func (m *RubricLevel) GetGrading()(*EducationAssignmentGradeType) {
    if m == nil {
        return nil
    } else {
        return m.grading
    }
}
// GetLevelId gets the levelId property value. The ID of this resource.
func (m *RubricLevel) GetLevelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.levelId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricLevel) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. The description of this rubric level.
func (m *RubricLevel) SetDescription(value *EducationItemBody)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name of this rubric level.
func (m *RubricLevel) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGrading sets the grading property value. Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
func (m *RubricLevel) SetGrading(value *EducationAssignmentGradeType)() {
    if m != nil {
        m.grading = value
    }
}
// SetLevelId sets the levelId property value. The ID of this resource.
func (m *RubricLevel) SetLevelId(value *string)() {
    if m != nil {
        m.levelId = value
    }
}
