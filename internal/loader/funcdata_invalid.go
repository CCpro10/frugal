// +build !go1.15 go1.20

/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package loader

import (
    `github.com/cloudwego/frugal/internal/rt`
)

// triggers a compilation error
const (
    _ = panic("Unsupported Go version. Supported versions are 1.15 ~ 1.19")
)

func registerFunction(_ string, _ uintptr, _ uintptr, _ rt.Frame) {
    panic("Unsupported Go version. Supported versions are 1.15 ~ 1.19")
}
