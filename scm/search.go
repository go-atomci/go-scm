// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scm

import (
	"context"
)

type (

	// RepositoryService provides access to repository resources.
	SearchService interface {
		// Find returns a repository by name.
		FindProjects(context.Context, string) ([]*Repository, *Response, error)
	}
)
