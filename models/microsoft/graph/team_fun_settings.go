package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamFunSettings struct {
    additionalData map[string]interface{};
    allowCustomMemes *bool;
    allowGiphy *bool;
    allowStickersAndMemes *bool;
    giphyContentRating *GiphyRatingType;
}
func NewTeamFunSettings()(*TeamFunSettings) {
    m := &TeamFunSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeamFunSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeamFunSettings) GetAllowCustomMemes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCustomMemes
    }
}
func (m *TeamFunSettings) GetAllowGiphy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowGiphy
    }
}
func (m *TeamFunSettings) GetAllowStickersAndMemes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStickersAndMemes
    }
}
func (m *TeamFunSettings) GetGiphyContentRating()(*GiphyRatingType) {
    if m == nil {
        return nil
    } else {
        return m.giphyContentRating
    }
}
func (m *TeamFunSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowCustomMemes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowCustomMemes(val)
        return nil
    }
    res["allowGiphy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowGiphy(val)
        return nil
    }
    res["allowStickersAndMemes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowStickersAndMemes(val)
        return nil
    }
    res["giphyContentRating"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGiphyRatingType)
        if err != nil {
            return err
        }
        cast := val.(GiphyRatingType)
        m.SetGiphyContentRating(&cast)
        return nil
    }
    return res
}
func (m *TeamFunSettings) IsNil()(bool) {
    return m == nil
}
func (m *TeamFunSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowCustomMemes", m.GetAllowCustomMemes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowGiphy", m.GetAllowGiphy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowStickersAndMemes", m.GetAllowStickersAndMemes())
        if err != nil {
            return err
        }
    }
    if m.GetGiphyContentRating() != nil {
        cast := m.GetGiphyContentRating().String()
        err := writer.WriteStringValue("giphyContentRating", &cast)
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
func (m *TeamFunSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeamFunSettings) SetAllowCustomMemes(value *bool)() {
    m.allowCustomMemes = value
}
func (m *TeamFunSettings) SetAllowGiphy(value *bool)() {
    m.allowGiphy = value
}
func (m *TeamFunSettings) SetAllowStickersAndMemes(value *bool)() {
    m.allowStickersAndMemes = value
}
func (m *TeamFunSettings) SetGiphyContentRating(value *GiphyRatingType)() {
    m.giphyContentRating = value
}
