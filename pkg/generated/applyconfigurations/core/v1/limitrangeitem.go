// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// LimitRangeItemApplyConfiguration represents an declarative configuration of the LimitRangeItem type for use
// with apply.
type LimitRangeItemApplyConfiguration struct {
	Type                 *v1.LimitType    `json:"type,omitempty"`
	Max                  *v1.ResourceList `json:"max,omitempty"`
	Min                  *v1.ResourceList `json:"min,omitempty"`
	Default              *v1.ResourceList `json:"default,omitempty"`
	DefaultRequest       *v1.ResourceList `json:"defaultRequest,omitempty"`
	MaxLimitRequestRatio *v1.ResourceList `json:"maxLimitRequestRatio,omitempty"`
}

// LimitRangeItemApplyConfiguration constructs an declarative configuration of the LimitRangeItem type for use with
// apply.
func LimitRangeItem() *LimitRangeItemApplyConfiguration {
	return &LimitRangeItemApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *LimitRangeItemApplyConfiguration) WithType(value v1.LimitType) *LimitRangeItemApplyConfiguration {
	b.Type = &value
	return b
}

// WithMax sets the Max field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Max field is set to the value of the last call.
func (b *LimitRangeItemApplyConfiguration) WithMax(value v1.ResourceList) *LimitRangeItemApplyConfiguration {
	b.Max = &value
	return b
}

// WithMin sets the Min field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Min field is set to the value of the last call.
func (b *LimitRangeItemApplyConfiguration) WithMin(value v1.ResourceList) *LimitRangeItemApplyConfiguration {
	b.Min = &value
	return b
}

// WithDefault sets the Default field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Default field is set to the value of the last call.
func (b *LimitRangeItemApplyConfiguration) WithDefault(value v1.ResourceList) *LimitRangeItemApplyConfiguration {
	b.Default = &value
	return b
}

// WithDefaultRequest sets the DefaultRequest field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultRequest field is set to the value of the last call.
func (b *LimitRangeItemApplyConfiguration) WithDefaultRequest(value v1.ResourceList) *LimitRangeItemApplyConfiguration {
	b.DefaultRequest = &value
	return b
}

// WithMaxLimitRequestRatio sets the MaxLimitRequestRatio field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxLimitRequestRatio field is set to the value of the last call.
func (b *LimitRangeItemApplyConfiguration) WithMaxLimitRequestRatio(value v1.ResourceList) *LimitRangeItemApplyConfiguration {
	b.MaxLimitRequestRatio = &value
	return b
}
