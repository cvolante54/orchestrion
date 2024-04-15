// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package instrument

import (
	"github.com/gorilla/mux"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

func WrapGorillaMuxRouter(r *mux.Router) *muxtrace.Router {
	return muxtrace.WrapRouter(r)
}