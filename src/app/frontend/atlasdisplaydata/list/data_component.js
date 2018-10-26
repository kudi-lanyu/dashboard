import {StateParams} from "../../common/resource/resourcedetail";
// import {stateName} from "../../persistentvolumeclaim/detail/state";


export class AtlasDataCardController {
  constructor($state, $interpolate, kdNamespaceService){
    this.pvc;
    this.state_ = $state;
    this.kdNamespaceService_ = kdNamespaceService;
  }

  /**
   * @return {boolean}
   * @export
   */
  areMultipleNamespacesSelected() {
    console.log("data component are MultipleNamespaceSelected.");
    return this.kdNamespaceService_.areMultipleNamespacesSelected();
  }

  // /**
  //  * @return {string}
  //  * @export
  //  */
  // getPersistentVolumeClaimDetailHref() {
  //   return this.state_.href(
  //     stateName,
  //     new StateParams(
  //       this.pvc.objectMeta.namespace,
  //       this.pvc.objectMeta.name));
  // }


  getDescription(){
    console.log("data component are 1.");

    return this.pvc.objectMeta.annotations["description"]
  }

  getOwner(){
    console.log("data component are 2.");

    return this.pvc.objectMeta.annotations["owner"]
  }

  /**
   * Returns true if persistent volume claim is in bound state, false otherwise.
   * @return {boolean}
   @export
   */
  isBound() {
    return this.pvc.status === 'Bound';
  }

  /**
   * Returns true if persistent volume claim is in pending state, false otherwise.
   * @return {boolean}
   * @export
   */
  isPending() {
    return this.pvc.status === 'Pending';
  }

  /**
   * Returns true if persistent volume claim is in lost state, false otherwise.
   * @return {boolean}
   * @export
   */
  isLost() {
    return this.pvc.status === 'Lost';
  }
}

export const atlasDataCardComponent = {
  bindings: {
    'pvc': '=',
  },
  controller: AtlasDataCardController,
  templateUrl: 'atlasdisplaydata/list/data.html',
};
