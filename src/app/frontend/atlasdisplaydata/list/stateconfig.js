// atlas
import {stateName as atlasdisplayStateName} from '../../atlasdisplay/state';
// dashboard
import {breadcrumbsConfig} from "../../common/components/breadcrumbs/service";

//
import {AtlasDataListController} from "./controller";
import {stateName as parentState, stateUrl} from '../state';

const i18n = {
  MSG_BREADCRUMBS_ATLASDISPLAYDATA_LABEL: goog.getMsg('Atlasdisplaydata')
};

/**
 * Config state object for the example list view.
 * @type {!ui.router.StateConfig}
 */
export const config = {
  url: stateUrl,
  parent: parentState, //
  //resolve function will inject to controller.<string, function>
  resolve: {
    'dataList': resolveAtlasDataList,
  },
  data: {
    [breadcrumbsConfig]: {
      'label': i18n.MSG_BREADCRUMBS_ATLASDISPLAYDATA_LABEL,
      'parent': atlasdisplayStateName,
    },
  },
  views: {
    '': {
      controller: AtlasDataListController,
      controllerAs: '$ctrl',
      templateUrl: 'atlasdisplaydata/list/list.html',
    },
  },
};


function resolveAtlasDataList(kdAtlasDataListResource, kdDataSelectService) {
  console.log("atlasdatalistmodule: resolveAtlasDataList");
  let query = kdDataSelectService.getDefaultResourceQuery();
  // console.log(kdAtlasDataListResource.get(query).$promise.ListMeta.totalItems);
  return kdAtlasDataListResource.get(query).$promise;
}


