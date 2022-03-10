package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActionStat provides operations to manage the educationRoot singleton.
type ItemActionStat struct {
    // The number of times the action took place. Read-only.
    actionCount *int32;
    // The number of distinct actors that performed the action. Read-only.
    actorCount *int32;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewItemActionStat instantiates a new itemActionStat and sets the default values.
func NewItemActionStat()(*ItemActionStat) {
    m := &ItemActionStat{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemActionStatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActionStatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewItemActionStat(), nil
}
// GetActionCount gets the actionCount property value. The number of times the action took place. Read-only.
func (m *ItemActionStat) GetActionCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.actionCount
    }
}
// GetActorCount gets the actorCount property value. The number of distinct actors that performed the action. Read-only.
func (m *ItemActionStat) GetActorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.actorCount
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionStat) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActionStat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionCount(val)
        }
        return nil
    }
    res["actorCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActorCount(val)
        }
        return nil
    }
    return res
}
func (m *ItemActionStat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActionCount sets the actionCount property value. The number of times the action took place. Read-only.
func (m *ItemActionStat) SetActionCount(value *int32)() {
    if m != nil {
        m.actionCount = value
    }
}
// SetActorCount sets the actorCount property value. The number of distinct actors that performed the action. Read-only.
func (m *ItemActionStat) SetActorCount(value *int32)() {
    if m != nil {
        m.actorCount = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemActionStat) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
