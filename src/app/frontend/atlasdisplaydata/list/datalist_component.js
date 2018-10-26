export class AtlasDataCardListController {
  constructor(){
    this.atlasDataList;
    this.atlasDataListResource;
  }

  // function
}

export const atlasDataCardListComponent = {
  transclude:{
    'header':'?kdHeader',// define it by yourself
    'zerostate':'?kdEmptyListText',
  },
  controller:AtlasDataCardListController,
  bindings:{
    'atlasDataList':'<',
    'atlasDataListResource':'<',
  },
  templateUrl: 'atlasdisplaydata/list/datalist.html',
};
