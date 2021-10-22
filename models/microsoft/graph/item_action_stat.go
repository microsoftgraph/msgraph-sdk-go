package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemActionStat struct {
    actionCount *int32;
    actorCount *int32;
    additionalData map[string]interface{};
}
func NewItemActionStat()(*ItemActionStat) {
    m := &ItemActionStat{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ItemActionStat) GetActionCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.actionCount
    }
}
func (m *ItemActionStat) GetActorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.actorCount
    }
}
func (m *ItemActionStat) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
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
func (m *ItemActionStat) SetActionCount(value *int32)() {
    m.actionCount = value
}
func (m *ItemActionStat) SetActorCount(value *int32)() {
    m.actorCount = value
}
func (m *ItemActionStat) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
