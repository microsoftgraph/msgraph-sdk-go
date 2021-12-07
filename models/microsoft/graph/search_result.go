package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchResult 
type SearchResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if the user interacts with this item to improve the quality of results.
    onClickTelemetryUrl *string;
}
// NewSearchResult instantiates a new searchResult and sets the default values.
func NewSearchResult()(*SearchResult) {
    m := &SearchResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOnClickTelemetryUrl gets the onClickTelemetryUrl property value. A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if the user interacts with this item to improve the quality of results.
func (m *SearchResult) GetOnClickTelemetryUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onClickTelemetryUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["onClickTelemetryUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnClickTelemetryUrl(val)
        }
        return nil
    }
    return res
}
func (m *SearchResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("onClickTelemetryUrl", m.GetOnClickTelemetryUrl())
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
func (m *SearchResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOnClickTelemetryUrl sets the onClickTelemetryUrl property value. A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if the user interacts with this item to improve the quality of results.
func (m *SearchResult) SetOnClickTelemetryUrl(value *string)() {
    if m != nil {
        m.onClickTelemetryUrl = value
    }
}
