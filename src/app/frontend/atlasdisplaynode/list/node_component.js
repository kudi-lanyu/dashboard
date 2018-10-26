

export class AtlasNodeCardController {
  constructor($state, kdAtlasNodeInfoResource){
    this.node;
    this.state_ = $state;
    this.kdAtlasNodeInfoResource_ = kdAtlasNodeInfoResource;

    this.allocatedGpu;
    this.totalGpu;
    this.ipAddress;
    this.role;
    this.usage;
  }

  /** @export */
  $onInit() {
    console.log("atlas node componet! ");

    let nodeInfo = this.kdAtlasNodeInfoResource_.get(this.node.objectMeta.name).$promise;
    console.log(nodeInfo);

    this.ipAddress = nodeInfo.ipAddress;
    this.role = nodeInfo.noderole;
    this.allocatedGpu = nodeInfo.nodeAllocatedGpuCount;
    this.totalGpu = nodeInfo.nodeTotalGpuCount;
    this.usage = nodeInfo.usage;

    console.log(this.ipAddress);
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
