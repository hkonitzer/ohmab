// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package privacy

import (
	"context"
	"errors"
	"fmt"
	"github.com/hkonitzer/ohmab/ent/user"
	"strconv"
)

// Role for viewer actions.
type Role int

// List of roles.
const ( // see also verifyuserrole.go
	_ Role = 1 << iota
	Admin
	Owner
	Employee
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool // If viewer is admin.
	Owner() bool
	Employee() bool
	SetRole() error
	IsEmpty() bool
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

func EmployeeRoleAsString() string {
	return fmt.Sprint(Employee)
}

func ViewerRoleAsString() string {
	return fmt.Sprint(View)
}

func ViewerRole() int {
	return int(View)
}

func (v *UserViewer) IsEmpty() bool {
	return v.Role == 0
}

func (v *UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

func (v *UserViewer) Owner() bool {
	return v.Role&Owner != 0
}

func (v *UserViewer) Employee() bool {
	return v.Role&Employee != 0
}

func (v *UserViewer) SetRole(role int) error {
	switch role {
	case 16:
		v.Role = View
	case 8:
		v.Role = Employee
	case 4:
		v.Role = Owner
	case 2:
		v.Role = Admin
	default:
		return errors.New("invalid role: " + strconv.Itoa(role))
	}
	return nil
}

func (v *UserViewer) SetRoleAsString(role string) error {
	r, err := strconv.Atoi(role)
	if err != nil {
		return err
	}
	return v.SetRole(r)
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

type contextKey string

var ContextKeyViewer = contextKey("viewer")

func setToken(ctx context.Context, viewer UserViewer) context.Context {
	return context.WithValue(ctx, ContextKeyViewer, viewer)
}

func getToken(ctx context.Context) (UserViewer, bool) {
	tokenStr, ok := ctx.Value(ContextKeyViewer).(UserViewer)
	return tokenStr, ok
}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) (UserViewer, bool) {
	return getToken(ctx)
}

// ToContext sets the Viewer on the given context.
func (v *UserViewer) ToContext(ctx context.Context) context.Context {
	return setToken(ctx, *v)
}

// NewViewerContext returns a copy of parent context with a new constructed Viewer attached with it.
func NewViewerContext(parent context.Context, role Role, scopes []string) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	if scopes == nil {
		scopes = []string{}
	}
	return setToken(parent, UserViewer{
		Role:   role,
		Scopes: scopes,
	})
}
