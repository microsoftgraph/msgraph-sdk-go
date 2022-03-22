package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Privacy 
type Privacy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    subjectRightsRequests []SubjectRightsRequestable;
}
// NewPrivacy instantiates a new Privacy and sets the default values.
func NewPrivacy()(*Privacy) {
    m := &Privacy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrivacyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivacyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrivacy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Privacy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Privacy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["subjectRightsRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectRightsRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequestable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectRightsRequestable)
            }
            m.SetSubjectRightsRequests(res)
        }
        return nil
    }
    return res
}
// GetSubjectRightsRequests gets the subjectRightsRequests property value. 
func (m *Privacy) GetSubjectRightsRequests()([]SubjectRightsRequestable) {
    if m == nil {
        return nil
    } else {
        return m.subjectRightsRequests
    }
}
// Serialize serializes information the current object
func (m *Privacy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetSubjectRightsRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubjectRightsRequests()))
        for i, v := range m.GetSubjectRightsRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("subjectRightsRequests", cast)
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
func (m *Privacy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSubjectRightsRequests sets the subjectRightsRequests property value. 
func (m *Privacy) SetSubjectRightsRequests(value []SubjectRightsRequestable)() {
    if m != nil {
        m.subjectRightsRequests = value
    }
}
