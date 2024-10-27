// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// CustomResourceConversionApplyConfiguration represents an declarative configuration of the CustomResourceConversion type for use
// with apply.
type CustomResourceConversionApplyConfiguration struct {
	Strategy *v1.ConversionStrategyType           `json:"strategy,omitempty"`
	Webhook  *WebhookConversionApplyConfiguration `json:"webhook,omitempty"`
}

// CustomResourceConversionApplyConfiguration constructs an declarative configuration of the CustomResourceConversion type for use with
// apply.
func CustomResourceConversion() *CustomResourceConversionApplyConfiguration {
	return &CustomResourceConversionApplyConfiguration{}
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *CustomResourceConversionApplyConfiguration) WithStrategy(value v1.ConversionStrategyType) *CustomResourceConversionApplyConfiguration {
	b.Strategy = &value
	return b
}

// WithWebhook sets the Webhook field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Webhook field is set to the value of the last call.
func (b *CustomResourceConversionApplyConfiguration) WithWebhook(value *WebhookConversionApplyConfiguration) *CustomResourceConversionApplyConfiguration {
	b.Webhook = value
	return b
}