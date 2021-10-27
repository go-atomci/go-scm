// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"context"
	"fmt"

	"github.com/go-atomci/go-scm/scm"
)

type searchService struct {
	client *wrapper
}

func (s *searchService) FindProjects(ctx context.Context, projectName string) ([]*scm.Repository, *scm.Response, error) {
	path := fmt.Sprintf("api/v4/search?scope=projects&search=%s", encode(projectName))
	out := []*repository{}
	res, err := s.client.do(ctx, "GET", path, nil, out)
	return convertRepositoryList(out), res, err
}
