// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeAffinityApplyConfiguration represents an declarative configuration of the NodeAffinity type for use
// with apply.
type NodeAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  *NodeSelectorApplyConfiguration             `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []PreferredSchedulingTermApplyConfiguration `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// NodeAffinityApplyConfiguration constructs an declarative configuration of the NodeAffinity type for use with
// apply.
func NodeAffinity() *NodeAffinityApplyConfiguration {
	return &NodeAffinityApplyConfiguration{}
}

// WithRequiredDuringSchedulingIgnoredDuringExecution sets the RequiredDuringSchedulingIgnoredDuringExecution field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RequiredDuringSchedulingIgnoredDuringExecution field is set to the value of the last call.
func (b *NodeAffinityApplyConfiguration) WithRequiredDuringSchedulingIgnoredDuringExecution(value *NodeSelectorApplyConfiguration) *NodeAffinityApplyConfiguration {
	b.RequiredDuringSchedulingIgnoredDuringExecution = value
	return b
}

// WithPreferredDuringSchedulingIgnoredDuringExecution adds the given value to the PreferredDuringSchedulingIgnoredDuringExecution field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (b *NodeAffinityApplyConfiguration) WithPreferredDuringSchedulingIgnoredDuringExecution(values ...*PreferredSchedulingTermApplyConfiguration) *NodeAffinityApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPreferredDuringSchedulingIgnoredDuringExecution")
		}
		b.PreferredDuringSchedulingIgnoredDuringExecution = append(b.PreferredDuringSchedulingIgnoredDuringExecution, *values[i])
	}
	return b
}
