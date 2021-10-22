package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SearchResult struct {
    additionalData map[string]interface{};
    onClickTelemetryUrl *string;
}
func NewSearchResult()(*SearchResult) {
    m := &SearchResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SearchResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SearchResult) GetOnClickTelemetryUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onClickTelemetryUrl
    }
}
func (m *SearchResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["onClickTelemetryUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnClickTelemetryUrl(val)
        return nil
    }
    return res
}
func (m *SearchResult) IsNil()(bool) {
    return m == nil
}
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
func (m *SearchResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SearchResult) SetOnClickTelemetryUrl(value *string)() {
    m.onClickTelemetryUrl = value
}
