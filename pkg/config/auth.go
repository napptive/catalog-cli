/**
 * Copyright 2021 Napptive
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package config

import (
	"context"
	"google.golang.org/grpc/metadata"
	"time"
)

// AuthorizationHeader with the key name for the authorization payload.
const AuthorizationHeader = "authorization"

// AgentHeader with the key name for the agent payload.
const AgentHeader = "agent"

// AgentValue with the value for the agent payload.
const AgentValue = "catalog"

// VersionHeader with the key name for the version payload.
const VersionHeader = "version"

// ContextTimeout with the default timeout for Napptive catalog operations.
const ContextTimeout = 5 * time.Minute

// TokenConfig with a struct to load the JWT token
type TokenConfig struct {
	// Token with the JWT.
	Token string
	// RefreshToken contains a JWT that can be used to renew the active JWT.
	Refresh string
}

// AuthToken with a struct to manage and send the token in the context metadata
type AuthToken struct {
	// AuthEnable with a flag if the authentication is enabled or not
	AuthEnable bool
	// Token with the token to send
	Token      string
	// Version with the version sent in the metadata
	Version    string
}

// NewContextHelper creates a ContextHelper with a given configuration.
func NewAuthToken(cfg *Config) *AuthToken {
	return &AuthToken{
		Version:    cfg.Version,
		Token:      cfg.TokenConfig.Token,
		AuthEnable: cfg.AuthEnable,
	}
}

// GetContext returns a context depending if the metadata is enabled or not
func (a *AuthToken) GetContext() (context.Context, context.CancelFunc) {
		md := metadata.New(map[string]string{AgentHeader: AgentValue, VersionHeader: a.Version})
	if a.AuthEnable {
		md = metadata.New(map[string]string{AuthorizationHeader: a.Token, AgentHeader: AgentValue, VersionHeader: a.Version})
	}
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return context.WithTimeout(ctx, ContextTimeout)
}
