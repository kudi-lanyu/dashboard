import {stateName as chromeStateName} from '../../chrome/state';
import {breadcrumbsConfig} from "../../common/components/breadcrumbs/service";

import {stateName} from './state';
import {stateUrl} from './state';
import {AtlastopController} from './controller';

export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    url: stateUrl,
    parent: chromeStateName,
    resolve:{
      'atlastop':atlasTopResource,
    },
    data:{
      [breadcrumbsConfig]:{
        'label': i18n.MSG_BREADCRUMBS_ATLASTOP_LABEL,
      },
    },
    views:{
      '':{
        controller:AtlastopController,
        controllerAs:'$ctrl',
        templateUrl:'overview/overview.html',
      }
    },
  });
}

export function atlasTopResource(kdAtlasTopResource,kdDataSelectService) {
  console.log("atlasTopResource");
  console.log(kdAtlasTopResource);
  // let paginationQuery = kdDataSelectService.getDefaultResourceQuery();
  // console.log(kdAtlasTopResource.get().$promise);
  return kdAtlasTopResource.get().$promise;
}

const i18n = {
  MSG_BREADCRUMBS_ATLASTOP_LABEL: goog.getMsg('Atlastop'),
};
