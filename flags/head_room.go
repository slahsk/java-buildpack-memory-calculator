/*
 * Copyright 2015-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package flags

import (
	"fmt"
	"strconv"
)

const (
	DefaultHeadRoom = HeadRoom(0)
	FlagHeadRoom    = "head-room"
)

type HeadRoom int

func (h *HeadRoom) Set(s string) error {
	f, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*h = HeadRoom(f)
	return nil
}

func (h *HeadRoom) String() string {
	return strconv.FormatInt(int64(*h), 10)
}

func (h *HeadRoom) Type() string {
	return "int"
}

func (h *HeadRoom) Validate() error {
	if *h < 0 || *h > 100 {
		return fmt.Errorf("--%s must be a valid percentage: %d", FlagHeadRoom, *h)
	}

	return nil
}
