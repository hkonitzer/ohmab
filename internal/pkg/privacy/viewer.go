// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package privacy

import (
	"context"
	"fmt"
	"hynie.de/ohmab/ent/user"
)

// Role for viewer actions.
type Role int

// List of roles.
const ( // see also verifyuserrole.go
	_ Role = 1 << iota
	Admin
	Owner
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool // If viewer is admin.
	Owner() bool
	HasScopes() bool
	HasScope(scope string) bool
	GetUserID() string
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	Role   Role     // Attached roles.
	Scopes []string // Attached businesses for this role
	Claims map[string]string
}

func AdminRoleAsString() string {
	return fmt.Sprint(Admin)
}

func OwnerRoleAsString() string {
	return fmt.Sprint(Owner)
}

func ViewerRoleAsString() string {
	return fmt.Sprint(View)
}

func ViewerRole() int {
	return int(View)
}

func (v *UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

func (v *UserViewer) Owner() bool {
	return v.Role&Owner != 0
}

func (v *UserViewer) HasScopes() bool {
	return len(v.Scopes) > 0
}

func (v *UserViewer) HasScope(scope string) bool {
	if !v.HasScopes() {
		return false
	}
	for _, s := range v.Scopes {
		if s == scope {
			return true
		}
	}
	return false
}

func (v *UserViewer) GetUserID() string {
	if len(v.Claims) == 0 {
		return ""
	}
	return v.Claims["user_"+user.FieldID]
}

func (v *UserViewer) SetUserID(id string) {
	if v.Claims == nil {
		v.Claims = make(map[string]string)
	}
	v.Claims["user_"+user.FieldID] = id
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
