export class AtlasFakeCardListController {
  constructor(){
    this.atlasFakeList;
    this.atlasFakeListResource;
  }

  // function
}

export const atlasFakeCardListComponent = {
  transclude:{
    'header':'?kdHeader',// define it by yourself
    'zerostate':'?kdEmptyListText',
  },
  controller:AtlasFakeCardListController,
  bindings:{
    'atlasFakeList':'<',
    'atlasFakeListResource':'<',
  },
  templateUrl: 'atlasfakelist/list/fakelist.html',
};
