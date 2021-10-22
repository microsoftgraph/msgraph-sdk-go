package sync

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type SyncRequestBody struct {
    additionalData map[string]interface{};
    syncType *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementExchangeConnectorSyncType;
}
func NewSyncRequestBody()(*SyncRequestBody) {
    m := &SyncRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SyncRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SyncRequestBody) GetSyncType()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementExchangeConnectorSyncType) {
    if m == nil {
        return nil
    } else {
        return m.syncType
    }
}
func (m *SyncRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["syncType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseDeviceManagementExchangeConnectorSyncType)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementExchangeConnectorSyncType)
        m.SetSyncType(&cast)
        return nil
    }
    return res
}
func (m *SyncRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *SyncRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetSyncType() != nil {
        cast := m.GetSyncType().String()
        err := writer.WriteStringValue("syncType", &cast)
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
func (m *SyncRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SyncRequestBody) SetSyncType(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagementExchangeConnectorSyncType)() {
    m.syncType = value
}
