// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {stateName as chromeStateName} from '../chrome/state';
import {breadcrumbsConfig} from '../common/components/breadcrumbs/service';

import {AtlasTasksController} from './controller';
import {stateName} from './state';
import {stateUrl} from './state';

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    url: stateUrl,
    parent: chromeStateName,
    resolve: {
      'cluster': resolveResource,
    },
    data: {
      [breadcrumbsConfig]: {
        'label': i18n.MSG_BREADCRUMBS_ATLASTASKS_LABEL,
      },
    },
    views: {
      '': {
        controller: AtlasTasksController,
        controllerAs: '$ctrl',
        templateUrl: 'atlastasks/atlastasks.html',
      },
    },
  });
}

/**
 * @param {!angular.$resource} kdAtlasTasksResource
 * @param {!./../common/dataselect/service.DataSelectService} kdDataSelectService
 * @return {!angular.$q.Promise}
 * @ngInject
 */
export function resolveResource(kdAtlasTasksResource, kdDataSelectService) {
  let paginationQuery = kdDataSelectService.getDefaultResourceQuery();
  return kdAtlasTasksResource.get(paginationQuery).$promise;
}

const i18n = {
  /** @type {string} @desc Label 'Cluster' that appears as a breadcrumbs on the action bar. */
  MSG_BREADCRUMBS_ATLASTASKS_LABEL: goog.getMsg('Atlastasks'),
};
