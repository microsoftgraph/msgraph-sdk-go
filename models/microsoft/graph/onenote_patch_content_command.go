package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnenotePatchContentCommand struct {
    action *OnenotePatchActionType;
    additionalData map[string]interface{};
    content *string;
    position *OnenotePatchInsertPosition;
    target *string;
}
func NewOnenotePatchContentCommand()(*OnenotePatchContentCommand) {
    m := &OnenotePatchContentCommand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OnenotePatchContentCommand) GetAction()(*OnenotePatchActionType) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *OnenotePatchContentCommand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OnenotePatchContentCommand) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *OnenotePatchContentCommand) GetPosition()(*OnenotePatchInsertPosition) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
func (m *OnenotePatchContentCommand) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *OnenotePatchContentCommand) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenotePatchActionType)
        if err != nil {
            return err
        }
        cast := val.(OnenotePatchActionType)
        m.SetAction(&cast)
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenotePatchInsertPosition)
        if err != nil {
            return err
        }
        cast := val.(OnenotePatchInsertPosition)
        m.SetPosition(&cast)
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTarget(val)
        return nil
    }
    return res
}
func (m *OnenotePatchContentCommand) IsNil()(bool) {
    return m == nil
}
func (m *OnenotePatchContentCommand) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetPosition() != nil {
        cast := m.GetPosition().String()
        err := writer.WriteStringValue("position", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
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
func (m *OnenotePatchContentCommand) SetAction(value *OnenotePatchActionType)() {
    m.action = value
}
func (m *OnenotePatchContentCommand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OnenotePatchContentCommand) SetContent(value *string)() {
    m.content = value
}
func (m *OnenotePatchContentCommand) SetPosition(value *OnenotePatchInsertPosition)() {
    m.position = value
}
func (m *OnenotePatchContentCommand) SetTarget(value *string)() {
    m.target = value
}
