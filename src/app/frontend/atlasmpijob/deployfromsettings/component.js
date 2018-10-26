import DeployLabel from "../../deploy/deployfromsettings/deploylabel/deploylabel";
import showNamespaceDialog from "../../deploy/deployfromsettings/createnamespace/dialog";
import {stateName as overview} from "../../overview/state";
import {uniqueNameValidationKey} from "../../deploy/deployfromsettings/uniquename_directive";
import showCreateSecretDialog from "../../deploy/deployfromsettings/createsecret/dialog";


/** @final */
class DeployFromSettingsController {
  /**
   * @param {!angular.$log} $log
   * @param {!angular.$resource} $resource
   * @param {!md.$dialog} $mdDialog
   * @param {!../../chrome/state.StateParams} $stateParams
   * @param {!../../common/history/service.HistoryService} kdHistoryService
   * @param {!../../common/namespace/service.NamespaceService} kdNamespaceService
   * @param {!../service.AtlasMpiJobService} kdAtlasMpiJobService
   * @ngInject
   */
  constructor($log, $resource, $mdDialog, $stateParams, kdHistoryService, kdNamespaceService,
              kdAtlasMpiJobService) {
    /** @export {!angular.FormController} */
    this.form;

    /** @private {!../../common/namespace/service.NamespaceService} */
    this.kdNamespaceService_ = kdNamespaceService;

    /** @private {boolean} */
    this.showMoreOptions_ = false;

    /** mpijob args*/

    /** @export {string} */
    this.name = '';

    /** @export {string} */
    this.image = '';

    /** @export {?number} */
    this.gpus = '';

    /** @export {string} */
    this.command = '';

    // optional
    /** @export {string} */
    this.cpu = '';

    /** @export {string} */
    this.memory = '';

    // for common optional

    /** @export {!Array<!backendApi.EnvironmentVariable>} */
    this.envs = [];

    /** @export {string} */
    this.workingDir = '';

    /** @export {?number} */
    this.retry = '';

    this.dataset = [];

    this.dataDirs = [];

    /** tensorboard optional */

    /** @export {boolean} */
    this.useTensorboard = false;

    /** @export {string} */
    this.trainingLogdir = '';

    /** @export {boolean} */
    this.isLocalLogging = false;

    /** @export {string} */
    this.hostLogPath = '';

    /** synccode optional */

    /** @export {string} */
    this.syncMode = '';

    /** @export {string} */
    this.syncSource = '';

    /** @export {string} */
    this.syncImage = '';

    /** @export {string} */
    this.syncGitProjectName = '';

    /**
     * Checks that a name begins and ends with a lowercase letter
     * and contains nothing but lowercase letters and hyphens ("-")
     * (leading and trailing spaces are ignored by default)
     * @export {!RegExp}
     */
    this.namePattern = new RegExp('^[a-z]([-a-z0-9]*[a-z0-9])?$');

    /** @export {string} */
    this.nameMaxLength = '24';

    /** @private {!angular.$resource} */
    this.resource_ = $resource;

    /** @private {!angular.$log} */
    this.log_ = $log;

    /** @private {!md.$dialog} */
    this.mdDialog_ = $mdDialog;

    /** @private {!../../chrome/state.StateParams} */
    this.stateParams_ = $stateParams;

    /** @private {!../../common/history/service.HistoryService} */
    this.historyService_ = kdHistoryService;


    /** @export {string} */
    this.namespace = this.stateParams_.namespace || '';

    /** @private {!../service.DeployService} */
    this.deployService_ = kdAtlasMpiJobService;

  }

  /** @export */
  $onInit() {
    let namespacesResource = this.resource_('api/v1/namespace');
    namespacesResource.get(
      (namespaces) => {
        this.namespaces = namespaces.namespaces.map((n) => n.objectMeta.name);
        this.namespace = !this.kdNamespaceService_.areMultipleNamespacesSelected() ?
          this.stateParams_.namespace || this.namespaces[0] :
          this.namespaces[0];
      },
      (err) => {
        this.log_.log(`Error during getting namespaces: ${err}`);
      });


  }


  /**
   * Returns true when the deploy action should be enabled.
   * @return {boolean}
   * @export
   */
  isDeployDisabled() {
    return !this.form.$valid || this.deployService_.isDeployDisabled();
  }

  /**
   * @export
   */
  cancel() {
    this.historyService_.back(overview);
  }

  /**
   * @export
   */
  deploy() {
    /** @type {!backendApi.MpiJobSpec} */
    let spec = {
      // necessary
      name:this.name,
      image: this.image,
      gpus: angular.isNumber(this.gpus) ? this.gpus : null,
      command: this.command,

      // optional
      cpu: angular.isNumber(this.cpu) ? this.cpu : null,
      memory: this.memory,

      // for common optional
      envs: this.envs,
      workingDir: this.workingDir,
      retry: this.retry,

      // complete in atlas-dashboard rel1
      dataSet:this.dataset,
      dataDirs:this.dataDirs,

      // tensorboard
      useTensorboard:this.useTensorboard,
      trainingLogdir:this.trainingLogdir,
      isLocalLogging:this.isLocalLogging,
      hostLogPath:this.hostLogPath,

      // SyncCode
      syncMode:this.syncMode,
      syncSource:this.syncSource,
      syncImage:this.syncImage,
      syncGitProjectName:this.syncGitProjectName,
    };

    this.deployService_.deploy(spec);
  }


  /**
   * Returns true when name input should show error. This overrides default behavior to show name
   * uniqueness errors even in the middle of typing.
   *
   * @return {boolean}
   * @export
   */
  isNameError() {
    /** @type {!angular.NgModelController} */
    let name = this.form['name'];

    return name.$error[uniqueNameValidationKey] ||
      (name.$invalid && (name.$touched || this.form.$submitted));
  }

  /**
   * Converts array of DeployLabel to array of backend api label.
   *
   * @param {!Array<!DeployLabel>} labels
   * @return {!Array<!backendApi.Label>}
   * @private
   */
  toBackendApiLabels_(labels) {
    // Omit labels with empty key/value
    /** @type {!Array<!DeployLabel>} */
    let apiLabels = labels.filter((label) => {
      return label.key.length !== 0 && label.value().length !== 0;
    });

    // Transform to array of backend api labels
    return apiLabels.map((label) => {
      return label.toBackendApi();
    });
  }

  isUseTensorboardEnabled(){
    return this.useTensorboard;
  }

  /**
   * Returns true when the given port mapping is filled by the user, i.e., is not empty.
   *
   * @param {!backendApi.PortMapping} portMapping
   * @return {boolean}
   * @private
   */
  isPortMappingFilled_(portMapping) {
    return !!portMapping.port && !!portMapping.targetPort;
  }

  /**
   * @param {!backendApi.EnvironmentVariable} variable
   * @return {boolean}
   * @private
   */
  isVariableFilled_(variable) {
    return !!variable.name;
  }

  /**
   * @return {string}
   * @private
   */
  getName_() {
    return this.name;
  }

  /**
   * Returns true if more options have been enabled and should be shown, false otherwise.
   *
   * @return {boolean}
   * @export
   */
  isMoreOptionsEnabled() {
    return this.showMoreOptions_;
  }

  /**
   * Shows or hides more options.
   * @export
   */
  switchMoreOptions() {
    this.showMoreOptions_ = !this.showMoreOptions_;
  }
}

/**
 * @return {!angular.Component}
 */
export const deployFromSettingsComponent = {
  controller: DeployFromSettingsController,
  templateUrl: 'atlasmpijob/deployfromsettings/deployfromsettings.html',
};
