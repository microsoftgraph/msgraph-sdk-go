package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingQuestionAssignment provides operations to manage the solutionsRoot singleton.
type BookingQuestionAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ID of the custom question.
    isRequired *bool;
    // Indicates whether it is mandatory to answer the custom question.
    questionId *string;
}
// NewBookingQuestionAssignment instantiates a new bookingQuestionAssignment and sets the default values.
func NewBookingQuestionAssignment()(*BookingQuestionAssignment) {
    m := &BookingQuestionAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingQuestionAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingQuestionAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingQuestionAssignment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingQuestionAssignment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingQuestionAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["questionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuestionId(val)
        }
        return nil
    }
    return res
}
// GetIsRequired gets the isRequired property value. The ID of the custom question.
func (m *BookingQuestionAssignment) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// GetQuestionId gets the questionId property value. Indicates whether it is mandatory to answer the custom question.
func (m *BookingQuestionAssignment) GetQuestionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.questionId
    }
}
func (m *BookingQuestionAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingQuestionAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("questionId", m.GetQuestionId())
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
func (m *BookingQuestionAssignment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsRequired sets the isRequired property value. The ID of the custom question.
func (m *BookingQuestionAssignment) SetIsRequired(value *bool)() {
    if m != nil {
        m.isRequired = value
    }
}
// SetQuestionId sets the questionId property value. Indicates whether it is mandatory to answer the custom question.
func (m *BookingQuestionAssignment) SetQuestionId(value *string)() {
    if m != nil {
        m.questionId = value
    }
}
