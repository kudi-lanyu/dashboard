import atlasFakeListModule, {atlasFakeListResource} from "../atlasfakelist/module";

export class AtlasdisplayController{

  constructor(atlasdisplay,cluster,atlasfakelist,kdAtlasFakeListResource,kdAtlasDisplayResource,kdNamespaceListResource, kdNodeListResource, kdPersistentVolumeListResource,
              kdRoleListResource, kdStorageClassListResource){
    this.atlasdisplay = atlasdisplay;
    this.kdAtlasDisplayResource = kdAtlasDisplayResource;
    this.cluster = cluster;
    this.atlasfakelist = atlasfakelist;
    this.kdAtlasFakeListResource = kdAtlasFakeListResource;

    // module resource
    /** @export {!angular.Resource} */
    this.kdNamespaceListResource = kdNamespaceListResource;

    /** @export {!angular.Resource} */
    this.kdNodeListResource = kdNodeListResource;

    /** @export {!angular.Resource} */
    this.kdPersistentVolumeListResource = kdPersistentVolumeListResource;

    /** @export {!angular.Resource} */
    this.kdRoleListResource = kdRoleListResource;

    /** @export {!angular.Resource} */
    this.kdStorageClassListResource = kdStorageClassListResource;
  }

  // function use in templatized html file
  /**
   * @return {boolean}
   * @export
   */
  shouldShowZeroState() {
    /** @type {number} */
    // let resourcesLength = this.cluster.nodeList.listMeta.totalItems +
    //   this.cluster.namespaceList.listMeta.totalItems +
    //   this.cluster.persistentVolumeList.listMeta.totalItems +
    //   this.cluster.roleList.listMeta.totalItems +
    //   this.cluster.storageClassList.listMeta.totalItems;
    //

    return true;
  }


}
