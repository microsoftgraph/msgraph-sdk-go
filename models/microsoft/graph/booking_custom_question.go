package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingCustomQuestion provides operations to manage the solutionsRoot singleton.
type BookingCustomQuestion struct {
    Entity
    // The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
    answerInputType *AnswerInputType;
    // List of possible answer values.
    answerOptions []string;
    // Display name of this entity.
    displayName *string;
}
// NewBookingCustomQuestion instantiates a new bookingCustomQuestion and sets the default values.
func NewBookingCustomQuestion()(*BookingCustomQuestion) {
    m := &BookingCustomQuestion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBookingCustomQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomQuestionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingCustomQuestion(), nil
}
// GetAnswerInputType gets the answerInputType property value. The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
func (m *BookingCustomQuestion) GetAnswerInputType()(*AnswerInputType) {
    if m == nil {
        return nil
    } else {
        return m.answerInputType
    }
}
// GetAnswerOptions gets the answerOptions property value. List of possible answer values.
func (m *BookingCustomQuestion) GetAnswerOptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.answerOptions
    }
}
// GetDisplayName gets the displayName property value. Display name of this entity.
func (m *BookingCustomQuestion) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomQuestion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["answerInputType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerInputType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnswerInputType(val.(*AnswerInputType))
        }
        return nil
    }
    res["answerOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAnswerOptions(res)
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
    return res
}
func (m *BookingCustomQuestion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingCustomQuestion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnswerInputType() != nil {
        cast := (*m.GetAnswerInputType()).String()
        err = writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnswerOptions() != nil {
        err = writer.WriteCollectionOfStringValues("answerOptions", m.GetAnswerOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnswerInputType sets the answerInputType property value. The expected answer type. The possible values are: text, radioButton, unknownFutureValue.
func (m *BookingCustomQuestion) SetAnswerInputType(value *AnswerInputType)() {
    if m != nil {
        m.answerInputType = value
    }
}
// SetAnswerOptions sets the answerOptions property value. List of possible answer values.
func (m *BookingCustomQuestion) SetAnswerOptions(value []string)() {
    if m != nil {
        m.answerOptions = value
    }
}
// SetDisplayName sets the displayName property value. Display name of this entity.
func (m *BookingCustomQuestion) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
