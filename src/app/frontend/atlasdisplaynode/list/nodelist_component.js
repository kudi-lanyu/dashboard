export class AtlasNodeCardListController {
  constructor(){
    this.atlasNodeList;
    this.atlasNodeListResource;
  }
}

export const atlasNodeCardListComponent = {
  transclude:{
    'header':'?kdHeader',// define it by yourself
    'zerostate':'?kdEmptyListText',
  },
  controller:AtlasNodeCardListController,
  bindings:{
    'atlasNodeList':'<',
    'atlasNodeListResource':'<',
  },
  templateUrl: 'atlasdisplaynode/list/nodelist.html',
};
