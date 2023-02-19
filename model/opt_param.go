package model

const (
	ParamIpAddress = "ipAddress"
	ParamDomain    = "domain"
	ParamEmail     = "email"

	ParamEscapedUnicode = "escapedUnicode"
)

var optParamsByPriority = []string{ParamEmail, ParamDomain, ParamIpAddress}

type OptParams map[string]interface{}

func (p OptParams) GetHighPriorityParam() *OptParam {
	for _, k := range optParamsByPriority {
		if v, ok := p[k]; ok {
			return NewOptParam(k, v)
		}
	}

	return nil
}

type OptParam struct {
	Key   string
	Value interface{}
}

func NewOptParam(key string, value interface{}) *OptParam {
	return &OptParam{
		Key:   key,
		Value: value,
	}
}

func NewParamIPAddress(value string) *OptParam {
	return &OptParam{
		Key:   ParamIpAddress,
		Value: value,
	}
}

func NewParamDomain(value string) *OptParam {
	return &OptParam{
		Key:   ParamDomain,
		Value: value,
	}
}

func NewParamEmail(value string) *OptParam {
	return &OptParam{
		Key:   ParamEmail,
		Value: value,
	}
}

func NewParamEscapedUnicode(value uint8) *OptParam {
	return &OptParam{
		Key:   ParamEscapedUnicode,
		Value: value,
	}
}
