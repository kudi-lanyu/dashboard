

export class AtlasFakeListController{
  // construct
  /**
   *
   */
  constructor(fakeList, kdAtlasFakeListResource){
    this.atlasFakeList = fakeList;
    this.atlasFakeListResource = kdAtlasFakeListResource;
  }

  // controller function
  /**
   * @return {boolean}
   * @export
   */
  shouldShowZeroState() {
    /** @type {number} */
    // let resourcesLength = this.listexample.ListSpec.nums;
    // return resourcesLength === 0;
    return true;
  }

}

