// Copyright 2021 iWuxc <iwuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	v1 "github.com/iWuxc/api/apiserver/v1"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

type Template interface {
	Push(ctx context.Context, opts metav1.CreateOptions) error
}