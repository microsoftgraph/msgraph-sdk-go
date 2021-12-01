package batchrecorddecisions

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BatchRecordDecisionsRequestBody 
type BatchRecordDecisionsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    decision *string;
    // 
    justification *string;
    // 
    principalId *string;
    // 
    resourceId *string;
}
// NewBatchRecordDecisionsRequestBody instantiates a new batchRecordDecisionsRequestBody and sets the default values.
func NewBatchRecordDecisionsRequestBody()(*BatchRecordDecisionsRequestBody) {
    m := &BatchRecordDecisionsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BatchRecordDecisionsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecision gets the decision property value. 
func (m *BatchRecordDecisionsRequestBody) GetDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decision
    }
}
// GetJustification gets the justification property value. 
func (m *BatchRecordDecisionsRequestBody) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetPrincipalId gets the principalId property value. 
func (m *BatchRecordDecisionsRequestBody) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// GetResourceId gets the resourceId property value. 
func (m *BatchRecordDecisionsRequestBody) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BatchRecordDecisionsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["decision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecision(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["principalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    return res
}
func (m *BatchRecordDecisionsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BatchRecordDecisionsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
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
func (m *BatchRecordDecisionsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecision sets the decision property value. 
func (m *BatchRecordDecisionsRequestBody) SetDecision(value *string)() {
    if m != nil {
        m.decision = value
    }
}
// SetJustification sets the justification property value. 
func (m *BatchRecordDecisionsRequestBody) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetPrincipalId sets the principalId property value. 
func (m *BatchRecordDecisionsRequestBody) SetPrincipalId(value *string)() {
    if m != nil {
        m.principalId = value
    }
}
// SetResourceId sets the resourceId property value. 
func (m *BatchRecordDecisionsRequestBody) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
