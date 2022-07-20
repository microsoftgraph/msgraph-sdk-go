package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanContainer 
type PlannerPlanContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier of the resource that contains the plan.
    containerId *string
    // The OdataType property
    odataType *string
    // The type property
    type_escaped *PlannerContainerType
    // The full canonical URL of the container.
    url *string
}
// NewPlannerPlanContainer instantiates a new plannerPlanContainer and sets the default values.
func NewPlannerPlanContainer()(*PlannerPlanContainer) {
    m := &PlannerPlanContainer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.plannerPlanContainer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerPlanContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanContainer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerPlanContainer) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContainerId gets the containerId property value. The identifier of the resource that contains the plan.
func (m *PlannerPlanContainer) GetContainerId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["containerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerContainerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*PlannerContainerType))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerPlanContainer) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetType gets the type property value. The type property
func (m *PlannerPlanContainer) GetType()(*PlannerContainerType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUrl gets the url property value. The full canonical URL of the container.
func (m *PlannerPlanContainer) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// Serialize serializes information the current object
func (m *PlannerPlanContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerId", m.GetContainerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *PlannerPlanContainer) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContainerId sets the containerId property value. The identifier of the resource that contains the plan.
func (m *PlannerPlanContainer) SetContainerId(value *string)() {
    if m != nil {
        m.containerId = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerPlanContainer) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetType sets the type property value. The type property
func (m *PlannerPlanContainer) SetType(value *PlannerContainerType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUrl sets the url property value. The full canonical URL of the container.
func (m *PlannerPlanContainer) SetUrl(value *string)() {
    if m != nil {
        m.url = value
    }
}
