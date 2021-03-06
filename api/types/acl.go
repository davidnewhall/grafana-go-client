// Copyright 2018 Augustin Husson
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import "time"

type RoleType string

const (
	RoleViewer RoleType = "Viewer"
	RoleEditor RoleType = "Editor"
	RoleAdmin  RoleType = "Admin"
)

type PermissionTypeAsString string

const (
	PermmissionViewAsString = "View"
	PermissionEditAsString  = "Edit"
)

type PermissionType int

const (
	PermissionView PermissionType = 1 << iota
	PermissionEdit
	PermissionAdmin
)

type FolderOrDashboardPermission struct {
	IsFolder       bool           `json:"isFolder"`
	Inherited      bool           `json:"inherited"`
	OrgID          int64          `json:"-"`
	DashboardID    int64          `json:"dashboardId,omitempty"`
	FolderID       int64          `json:"folderId,omitempty"`
	Created        time.Time      `json:"created"`
	Updated        time.Time      `json:"updated"`
	UserID         int64          `json:"userId"`
	UserLogin      string         `json:"userLogin"`
	UserEmail      string         `json:"userEmail"`
	UserAvatarURL  string         `json:"userAvatarUrl"`
	TeamID         int64          `json:"teamId"`
	TeamEmail      string         `json:"teamEmail"`
	TeamAvatarURL  string         `json:"teamAvatarUrl"`
	Team           string         `json:"team"`
	Role           *RoleType      `json:"role,omitempty"`
	Permission     PermissionType `json:"permission"`
	PermissionName string         `json:"permissionName"`
	UID            string         `json:"uid"`
	Title          string         `json:"title"`
	Slug           string         `json:"slug"`
	URL            string         `json:"url"`
}
