package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamFunSettings 
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
// NewTeamFunSettings instantiates a new teamFunSettings and sets the default values.
func NewTeamFunSettings()(*TeamFunSettings) {
    m := &TeamFunSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamFunSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamFunSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamFunSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamFunSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowCustomMemes gets the allowCustomMemes property value. If set to true, enables users to include custom memes.
func (m *TeamFunSettings) GetAllowCustomMemes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCustomMemes
    }
}
// GetAllowGiphy gets the allowGiphy property value. If set to true, enables Giphy use.
func (m *TeamFunSettings) GetAllowGiphy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowGiphy
    }
}
// GetAllowStickersAndMemes gets the allowStickersAndMemes property value. If set to true, enables users to include stickers and memes.
func (m *TeamFunSettings) GetAllowStickersAndMemes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowStickersAndMemes
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetGiphyContentRating(val.(*GiphyRatingType))
        }
        return nil
    }
    return res
}
// GetGiphyContentRating gets the giphyContentRating property value. Giphy content rating. Possible values are: moderate, strict.
func (m *TeamFunSettings) GetGiphyContentRating()(*GiphyRatingType) {
    if m == nil {
        return nil
    } else {
        return m.giphyContentRating
    }
}
// Serialize serializes information the current object
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
        cast := (*m.GetGiphyContentRating()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamFunSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowCustomMemes sets the allowCustomMemes property value. If set to true, enables users to include custom memes.
func (m *TeamFunSettings) SetAllowCustomMemes(value *bool)() {
    if m != nil {
        m.allowCustomMemes = value
    }
}
// SetAllowGiphy sets the allowGiphy property value. If set to true, enables Giphy use.
func (m *TeamFunSettings) SetAllowGiphy(value *bool)() {
    if m != nil {
        m.allowGiphy = value
    }
}
// SetAllowStickersAndMemes sets the allowStickersAndMemes property value. If set to true, enables users to include stickers and memes.
func (m *TeamFunSettings) SetAllowStickersAndMemes(value *bool)() {
    if m != nil {
        m.allowStickersAndMemes = value
    }
}
// SetGiphyContentRating sets the giphyContentRating property value. Giphy content rating. Possible values are: moderate, strict.
func (m *TeamFunSettings) SetGiphyContentRating(value *GiphyRatingType)() {
    if m != nil {
        m.giphyContentRating = value
    }
}
