#!/usr/bin/env bash

# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${REPO_ROOT}/hack/ensure-go.sh"
source "${REPO_ROOT}/hack/ensure-kind.sh"
source "${REPO_ROOT}/hack/ensure-kubectl.sh"
cd "${REPO_ROOT}" || exit 1

echo "Integration tests do not exist for v1alpha1. This script only exists to allow presubmits to pass on v1alpha1 branches."
bazel_status="${?}"
python hack/coalesce.py
exit "${bazel_status}"
