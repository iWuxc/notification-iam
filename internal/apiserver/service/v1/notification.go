// Copyright 2021 iWuxc <iwuxc.eng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"
	"github.com/marmotedu/iam/internal/apiserver/store"
)

type NotificationSrv interface {
	Get(ctx context.Context, tpl string)
}

type notificationService struct {
	store store.Factory
}

