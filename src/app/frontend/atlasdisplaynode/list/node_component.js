export class AtlasNodeCardController {
  constructor($state, kdAtlasNodeInfoResource, kdDataSelectService) {
    this.node;
    this.state_ = $state;
    this.kdAtlasNodeInfoResource_ = kdAtlasNodeInfoResource;
    this.kdDataSelectService_ = kdDataSelectService;

    this.nodeInfo = {};

    this.addr = '';
  }

  /** @export */
  $onInit() {
    console.log("atlas node componet! ");
    // console.log(this.node.objectMeta.name);

    let query = this.kdDataSelectService_.getDefaultResourceQuery('', this.node.objectMeta.name);
    this.kdAtlasNodeInfoResource_.get(query,
      (response) => {
        this.nodeInfo = response;
        console.log(this.nodeInfo);
      },
      (err) => {
        console.log("resource get err: ", err);
      });
    
    console.log("may not content will be display.");

  }


  /**
   * Returns true if node is in ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInReadyState() {
    return this.node.ready === 'True';
  }

  /**
   * Returns true if node is in non-ready state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInNotReadyState() {
    return this.node.ready === 'False';
  }

  /**
   * Returns true if node is in unknown state, false otherwise.
   * @return {boolean}
   * @export
   */
  isInUnknownState() {
    return this.node.ready === 'Unknown';
  }
}

export const atlasNodeCardComponent = {
  bindings: {
    'node': '=',
  },
  controller: AtlasNodeCardController,
  templateUrl: 'atlasdisplaynode/list/node.html',
};
