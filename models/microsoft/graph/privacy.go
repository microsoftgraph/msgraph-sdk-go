package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Privacy 
type Privacy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    subjectRightsRequests []SubjectRightsRequest;
}
// NewPrivacy instantiates a new Privacy and sets the default values.
func NewPrivacy()(*Privacy) {
    m := &Privacy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Privacy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSubjectRightsRequests gets the subjectRightsRequests property value. 
func (m *Privacy) GetSubjectRightsRequests()([]SubjectRightsRequest) {
    if m == nil {
        return nil
    } else {
        return m.subjectRightsRequests
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Privacy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["subjectRightsRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSubjectRightsRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectRightsRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*SubjectRightsRequest))
            }
            m.SetSubjectRightsRequests(res)
        }
        return nil
    }
    return res
}
func (m *Privacy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Privacy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSubjectRightsRequests()))
        for i, v := range m.GetSubjectRightsRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Privacy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSubjectRightsRequests sets the subjectRightsRequests property value. 
func (m *Privacy) SetSubjectRightsRequests(value []SubjectRightsRequest)() {
    if m != nil {
        m.subjectRightsRequests = value
    }
}
