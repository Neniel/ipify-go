package model

const (
	ParamIpAddress = "ipAddress"
	ParamDomain    = "domain"
	ParamEmail     = "email"

	ParamEscapedUnicode = "escapedUnicode"
)

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

var optParamsByPriority = []string{ParamEmail, ParamDomain, ParamIpAddress}

type OptParams map[string]interface{}

func NewOptParams() *OptParams {
	return &OptParams{}
}

func (p OptParams) Add(key string, value interface{}) *OptParams {
	p[key] = value
	return &p
}

func (p OptParams) GetHighPriorityParam() *OptParam {
	for _, k := range optParamsByPriority {
		if v, ok := p[k]; ok {
			return NewOptParam(k, v)
		}
	}

	return nil
}

func (p OptParams) GetParam(key string) *OptParam {
	if v, ok := p[key]; ok {
		return NewOptParam(key, v)
	}

	return nil
}
