package router

import "termcode/internal/domain/model"

type Capability int

const (
	CapStreaming Capability = iota + 1
	CapToolCalling
	CapReasoning
	CapVision
	CapEmbeddings
	CapJSONMode
	CapFunctionCalling
)

type Matcher struct{}

func NewMatcher() *Matcher {
	return &Matcher{}
}

func (m *Matcher) ModelHasCapability(mod *model.Model, cap Capability) bool {
	switch cap {
	case CapStreaming:
		return mod.Capabilities.Streaming
	case CapToolCalling:
		return mod.Capabilities.ToolCalling
	case CapReasoning:
		return mod.Capabilities.Reasoning
	case CapVision:
		return mod.Capabilities.Vision
	case CapEmbeddings:
		return mod.Capabilities.Embeddings
	case CapJSONMode:
		return mod.Capabilities.JSONMode
	case CapFunctionCalling:
		return mod.Capabilities.FunctionCalling
	default:
		return false
	}
}

func (m *Matcher) FilterByCapability(models []*model.Model, cap Capability) []*model.Model {
	var result []*model.Model
	for _, mod := range models {
		if m.ModelHasCapability(mod, cap) {
			result = append(result, mod)
		}
	}
	return result
}
