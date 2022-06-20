package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationRubricOutcome 
type EducationRubricOutcome struct {
    EducationOutcome
    // A copy of the rubricQualityFeedback property that is made when the grade is released to the student.
    publishedRubricQualityFeedback []RubricQualityFeedbackModelable
    // A copy of the rubricQualitySelectedLevels property that is made when the grade is released to the student.
    publishedRubricQualitySelectedLevels []RubricQualitySelectedColumnModelable
    // A collection of specific feedback for each quality of this rubric.
    rubricQualityFeedback []RubricQualityFeedbackModelable
    // The level that the teacher has selected for each quality while grading this assignment.
    rubricQualitySelectedLevels []RubricQualitySelectedColumnModelable
}
// NewEducationRubricOutcome instantiates a new EducationRubricOutcome and sets the default values.
func NewEducationRubricOutcome()(*EducationRubricOutcome) {
    m := &EducationRubricOutcome{
        EducationOutcome: *NewEducationOutcome(),
    }
    return m
}
// CreateEducationRubricOutcomeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationRubricOutcomeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationRubricOutcome(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationRubricOutcome) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationOutcome.GetFieldDeserializers()
    res["publishedRubricQualityFeedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricQualityFeedbackModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricQualityFeedbackModelable, len(val))
            for i, v := range val {
                res[i] = v.(RubricQualityFeedbackModelable)
            }
            m.SetPublishedRubricQualityFeedback(res)
        }
        return nil
    }
    res["publishedRubricQualitySelectedLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricQualitySelectedColumnModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricQualitySelectedColumnModelable, len(val))
            for i, v := range val {
                res[i] = v.(RubricQualitySelectedColumnModelable)
            }
            m.SetPublishedRubricQualitySelectedLevels(res)
        }
        return nil
    }
    res["rubricQualityFeedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricQualityFeedbackModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricQualityFeedbackModelable, len(val))
            for i, v := range val {
                res[i] = v.(RubricQualityFeedbackModelable)
            }
            m.SetRubricQualityFeedback(res)
        }
        return nil
    }
    res["rubricQualitySelectedLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricQualitySelectedColumnModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricQualitySelectedColumnModelable, len(val))
            for i, v := range val {
                res[i] = v.(RubricQualitySelectedColumnModelable)
            }
            m.SetRubricQualitySelectedLevels(res)
        }
        return nil
    }
    return res
}
// GetPublishedRubricQualityFeedback gets the publishedRubricQualityFeedback property value. A copy of the rubricQualityFeedback property that is made when the grade is released to the student.
func (m *EducationRubricOutcome) GetPublishedRubricQualityFeedback()([]RubricQualityFeedbackModelable) {
    if m == nil {
        return nil
    } else {
        return m.publishedRubricQualityFeedback
    }
}
// GetPublishedRubricQualitySelectedLevels gets the publishedRubricQualitySelectedLevels property value. A copy of the rubricQualitySelectedLevels property that is made when the grade is released to the student.
func (m *EducationRubricOutcome) GetPublishedRubricQualitySelectedLevels()([]RubricQualitySelectedColumnModelable) {
    if m == nil {
        return nil
    } else {
        return m.publishedRubricQualitySelectedLevels
    }
}
// GetRubricQualityFeedback gets the rubricQualityFeedback property value. A collection of specific feedback for each quality of this rubric.
func (m *EducationRubricOutcome) GetRubricQualityFeedback()([]RubricQualityFeedbackModelable) {
    if m == nil {
        return nil
    } else {
        return m.rubricQualityFeedback
    }
}
// GetRubricQualitySelectedLevels gets the rubricQualitySelectedLevels property value. The level that the teacher has selected for each quality while grading this assignment.
func (m *EducationRubricOutcome) GetRubricQualitySelectedLevels()([]RubricQualitySelectedColumnModelable) {
    if m == nil {
        return nil
    } else {
        return m.rubricQualitySelectedLevels
    }
}
// Serialize serializes information the current object
func (m *EducationRubricOutcome) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationOutcome.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPublishedRubricQualityFeedback() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublishedRubricQualityFeedback()))
        for i, v := range m.GetPublishedRubricQualityFeedback() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publishedRubricQualityFeedback", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublishedRubricQualitySelectedLevels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublishedRubricQualitySelectedLevels()))
        for i, v := range m.GetPublishedRubricQualitySelectedLevels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publishedRubricQualitySelectedLevels", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRubricQualityFeedback() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRubricQualityFeedback()))
        for i, v := range m.GetRubricQualityFeedback() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rubricQualityFeedback", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRubricQualitySelectedLevels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRubricQualitySelectedLevels()))
        for i, v := range m.GetRubricQualitySelectedLevels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rubricQualitySelectedLevels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPublishedRubricQualityFeedback sets the publishedRubricQualityFeedback property value. A copy of the rubricQualityFeedback property that is made when the grade is released to the student.
func (m *EducationRubricOutcome) SetPublishedRubricQualityFeedback(value []RubricQualityFeedbackModelable)() {
    if m != nil {
        m.publishedRubricQualityFeedback = value
    }
}
// SetPublishedRubricQualitySelectedLevels sets the publishedRubricQualitySelectedLevels property value. A copy of the rubricQualitySelectedLevels property that is made when the grade is released to the student.
func (m *EducationRubricOutcome) SetPublishedRubricQualitySelectedLevels(value []RubricQualitySelectedColumnModelable)() {
    if m != nil {
        m.publishedRubricQualitySelectedLevels = value
    }
}
// SetRubricQualityFeedback sets the rubricQualityFeedback property value. A collection of specific feedback for each quality of this rubric.
func (m *EducationRubricOutcome) SetRubricQualityFeedback(value []RubricQualityFeedbackModelable)() {
    if m != nil {
        m.rubricQualityFeedback = value
    }
}
// SetRubricQualitySelectedLevels sets the rubricQualitySelectedLevels property value. The level that the teacher has selected for each quality while grading this assignment.
func (m *EducationRubricOutcome) SetRubricQualitySelectedLevels(value []RubricQualitySelectedColumnModelable)() {
    if m != nil {
        m.rubricQualitySelectedLevels = value
    }
}
