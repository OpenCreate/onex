// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ServiceAccountTokenProjectionApplyConfiguration represents an declarative configuration of the ServiceAccountTokenProjection type for use
// with apply.
type ServiceAccountTokenProjectionApplyConfiguration struct {
	Audience          *string `json:"audience,omitempty"`
	ExpirationSeconds *int64  `json:"expirationSeconds,omitempty"`
	Path              *string `json:"path,omitempty"`
}

// ServiceAccountTokenProjectionApplyConfiguration constructs an declarative configuration of the ServiceAccountTokenProjection type for use with
// apply.
func ServiceAccountTokenProjection() *ServiceAccountTokenProjectionApplyConfiguration {
	return &ServiceAccountTokenProjectionApplyConfiguration{}
}

// WithAudience sets the Audience field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Audience field is set to the value of the last call.
func (b *ServiceAccountTokenProjectionApplyConfiguration) WithAudience(value string) *ServiceAccountTokenProjectionApplyConfiguration {
	b.Audience = &value
	return b
}

// WithExpirationSeconds sets the ExpirationSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExpirationSeconds field is set to the value of the last call.
func (b *ServiceAccountTokenProjectionApplyConfiguration) WithExpirationSeconds(value int64) *ServiceAccountTokenProjectionApplyConfiguration {
	b.ExpirationSeconds = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *ServiceAccountTokenProjectionApplyConfiguration) WithPath(value string) *ServiceAccountTokenProjectionApplyConfiguration {
	b.Path = &value
	return b
}
