package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ItemActionStat struct {
    // The number of times the action took place. Read-only.
    actionCount *int32;
    // The number of distinct actors that performed the action. Read-only.
    actorCount *int32;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// Instantiates a new itemActionStat and sets the default values.
func NewItemActionStat()(*ItemActionStat) {
    m := &ItemActionStat{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionCount property value. The number of times the action took place. Read-only.
func (m *ItemActionStat) GetActionCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.actionCount
    }
}
// Gets the actorCount property value. The number of distinct actors that performed the action. Read-only.
func (m *ItemActionStat) GetActorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.actorCount
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionStat) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// The deserialization information for the current model
func (m *ItemActionStat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActionCount(val)
        return nil
    }
    res["actorCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActorCount(val)
        return nil
    }
    return res
}
func (m *ItemActionStat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ItemActionStat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("actionCount", m.GetActionCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("actorCount", m.GetActorCount())
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
// Sets the actionCount property value. The number of times the action took place. Read-only.
// Parameters:
//  - value : Value to set for the actionCount property.
func (m *ItemActionStat) SetActionCount(value *int32)() {
    m.actionCount = value
}
// Sets the actorCount property value. The number of distinct actors that performed the action. Read-only.
// Parameters:
//  - value : Value to set for the actorCount property.
func (m *ItemActionStat) SetActorCount(value *int32)() {
    m.actorCount = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ItemActionStat) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
