package atlasjob

type (
	UseTensorboard   = bool
	TensorboardImage = string
	TrainingLogdir   = string
	HostLogPath      = string
	IsLocalLogging   = bool
)

type TensorboardArgs struct {
	UseTensorboard   UseTensorboard   `yaml:"useTensorboard";json:"useTensorboard"`     // --tensorboard
	TensorboardImage TensorboardImage `yaml:"tensorboardImage";json:"tensorboardImage"` // --tensorboardImage
	TrainingLogdir   TrainingLogdir   `yaml:"trainingLogdir";json:"trainingLogdir"`     // --logdir
	HostLogPath      HostLogPath      `yaml:"hostLogPath";json:"hostLogPath"`
	IsLocalLogging   IsLocalLogging   `yaml:"isLocalLogging";json:"isLocalLogging"`
}
