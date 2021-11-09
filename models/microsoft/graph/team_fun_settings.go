package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamFunSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If set to true, enables users to include custom memes.
    allowCustomMemes *bool;
    // If set to true, enables Giphy use.
    allowGiphy *bool;
    // If set to true, enables users to include stickers and memes.
    allowStickersAndMemes *bool;
    // Giphy content rating. Possible values are: moderate, strict.
    giphyContentRating *GiphyRatingType;
}
// Instantiates a new teamFunSettings and sets the default values.
func NewTeamFunSettings()(*TeamFunSettings) {
    m := &TeamFunSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamFunSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowCustomMemes property value. If set to true, enables users to include custom memes.
func (m *TeamFunSettings) GetAllowCustomMemes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCustomMemes
    }
}
// Gets the allowGiphy property value. If set to true, enables Giphy use.
func (m *TeamFunSettings) GetAllowGiphy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowGiphy
    }
}
// Gets the allowStickersAndMemes property value. If set to true, enables users to include stickers and memes.
func (m *TeamFunSettings) GetAllowStickersAndMemes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStickersAndMemes
    }
}
// Gets the giphyContentRating property value. Giphy content rating. Possible values are: moderate, strict.
func (m *TeamFunSettings) GetGiphyContentRating()(*GiphyRatingType) {
    if m == nil {
        return nil
    } else {
        return m.giphyContentRating
    }
}
// The deserialization information for the current model
func (m *TeamFunSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowCustomMemes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowCustomMemes(val)
        }
        return nil
    }
    res["allowGiphy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowGiphy(val)
        }
        return nil
    }
    res["allowStickersAndMemes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowStickersAndMemes(val)
        }
        return nil
    }
    res["giphyContentRating"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGiphyRatingType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(GiphyRatingType)
            m.SetGiphyContentRating(&cast)
        }
        return nil
    }
    return res
}
func (m *TeamFunSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TeamFunSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowCustomMemes property value. If set to true, enables users to include custom memes.
// Parameters:
//  - value : Value to set for the allowCustomMemes property.
func (m *TeamFunSettings) SetAllowCustomMemes(value *bool)() {
    m.allowCustomMemes = value
}
// Sets the allowGiphy property value. If set to true, enables Giphy use.
// Parameters:
//  - value : Value to set for the allowGiphy property.
func (m *TeamFunSettings) SetAllowGiphy(value *bool)() {
    m.allowGiphy = value
}
// Sets the allowStickersAndMemes property value. If set to true, enables users to include stickers and memes.
// Parameters:
//  - value : Value to set for the allowStickersAndMemes property.
func (m *TeamFunSettings) SetAllowStickersAndMemes(value *bool)() {
    m.allowStickersAndMemes = value
}
// Sets the giphyContentRating property value. Giphy content rating. Possible values are: moderate, strict.
// Parameters:
//  - value : Value to set for the giphyContentRating property.
func (m *TeamFunSettings) SetGiphyContentRating(value *GiphyRatingType)() {
    m.giphyContentRating = value
}
