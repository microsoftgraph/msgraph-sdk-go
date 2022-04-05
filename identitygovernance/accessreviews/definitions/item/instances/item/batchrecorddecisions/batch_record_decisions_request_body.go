package batchrecorddecisions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BatchRecordDecisionsRequestBody provides operations to call the batchRecordDecisions method.
type BatchRecordDecisionsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The decision property
    decision *string;
    // The justification property
    justification *string;
    // The principalId property
    principalId *string;
    // The resourceId property
    resourceId *string;
}
// NewBatchRecordDecisionsRequestBody instantiates a new batchRecordDecisionsRequestBody and sets the default values.
func NewBatchRecordDecisionsRequestBody()(*BatchRecordDecisionsRequestBody) {
    m := &BatchRecordDecisionsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBatchRecordDecisionsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBatchRecordDecisionsRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBatchRecordDecisionsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BatchRecordDecisionsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDecision gets the decision property value. The decision property
func (m *BatchRecordDecisionsRequestBody) GetDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decision
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BatchRecordDecisionsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decision"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecision(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["principalId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["resourceId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetJustification gets the justification property value. The justification property
func (m *BatchRecordDecisionsRequestBody) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetPrincipalId gets the principalId property value. The principalId property
func (m *BatchRecordDecisionsRequestBody) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// GetResourceId gets the resourceId property value. The resourceId property
func (m *BatchRecordDecisionsRequestBody) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// Serialize serializes information the current object
func (m *BatchRecordDecisionsRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BatchRecordDecisionsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDecision sets the decision property value. The decision property
func (m *BatchRecordDecisionsRequestBody) SetDecision(value *string)() {
    if m != nil {
        m.decision = value
    }
}
// SetJustification sets the justification property value. The justification property
func (m *BatchRecordDecisionsRequestBody) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetPrincipalId sets the principalId property value. The principalId property
func (m *BatchRecordDecisionsRequestBody) SetPrincipalId(value *string)() {
    if m != nil {
        m.principalId = value
    }
}
// SetResourceId sets the resourceId property value. The resourceId property
func (m *BatchRecordDecisionsRequestBody) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
