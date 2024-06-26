# Created by Isaque Veras on 03/15/24.
# Copyright © 2024 Isaque Veras. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

version := $(shell git tag | sort -V | tail -n 1)

build:
	go build -v -o ./bin/ -ldflags "-X github.com/isaqueveras/juaz/version.Version=${version}"
