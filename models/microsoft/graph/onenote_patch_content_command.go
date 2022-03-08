package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenotePatchContentCommand provides operations to call the onenotePatchContent method.
type OnenotePatchContentCommand struct {
    // The action to perform on the target element. The possible values are: replace, append, delete, insert, or prepend.
    action *OnenotePatchActionType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
    content *string;
    // The location to add the supplied content, relative to the target element. The possible values are: after (default) or before.
    position *OnenotePatchInsertPosition;
    // The element to update. Must be the #<data-id> or the generated <id> of the element, or the body or title keyword.
    target *string;
}
// NewOnenotePatchContentCommand instantiates a new onenotePatchContentCommand and sets the default values.
func NewOnenotePatchContentCommand()(*OnenotePatchContentCommand) {
    m := &OnenotePatchContentCommand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnenotePatchContentCommandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePatchContentCommandFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenotePatchContentCommand(), nil
}
// GetAction gets the action property value. The action to perform on the target element. The possible values are: replace, append, delete, insert, or prepend.
func (m *OnenotePatchContentCommand) GetAction()(*OnenotePatchActionType) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePatchContentCommand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContent gets the content property value. A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
func (m *OnenotePatchContentCommand) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePatchContentCommand) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenotePatchActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*OnenotePatchActionType))
        }
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenotePatchInsertPosition)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val.(*OnenotePatchInsertPosition))
        }
        return nil
    }
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetPosition gets the position property value. The location to add the supplied content, relative to the target element. The possible values are: after (default) or before.
func (m *OnenotePatchContentCommand) GetPosition()(*OnenotePatchInsertPosition) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// GetTarget gets the target property value. The element to update. Must be the #<data-id> or the generated <id> of the element, or the body or title keyword.
func (m *OnenotePatchContentCommand) GetTarget()(*string) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
func (m *OnenotePatchContentCommand) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OnenotePatchContentCommand) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
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
        cast := (*m.GetPosition()).String()
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
// SetAction sets the action property value. The action to perform on the target element. The possible values are: replace, append, delete, insert, or prepend.
func (m *OnenotePatchContentCommand) SetAction(value *OnenotePatchActionType)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePatchContentCommand) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContent sets the content property value. A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
func (m *OnenotePatchContentCommand) SetContent(value *string)() {
    if m != nil {
        m.content = value
    }
}
// SetPosition sets the position property value. The location to add the supplied content, relative to the target element. The possible values are: after (default) or before.
func (m *OnenotePatchContentCommand) SetPosition(value *OnenotePatchInsertPosition)() {
    if m != nil {
        m.position = value
    }
}
// SetTarget sets the target property value. The element to update. Must be the #<data-id> or the generated <id> of the element, or the body or title keyword.
func (m *OnenotePatchContentCommand) SetTarget(value *string)() {
    if m != nil {
        m.target = value
    }
}
