package configs

type FileSystemConfig struct {
	PathToReports        string
	VirtualPathToReports string
}

func GetFileSystemConfig() (FileSystemConfig, error) {
	fsConfig := FileSystemConfig{}
	pathToReports, err := getEnvVariable("PATH_TO_REPORTS")
	if err != nil {
		return FileSystemConfig{}, err
	}
	fsConfig.PathToReports = pathToReports
	virtualPathToReports, err := getEnvVariable("VIRTUAL_PATH_TO_REPORTS")
	if err != nil {
		return FileSystemConfig{}, err
	}
	fsConfig.VirtualPathToReports = virtualPathToReports
	return fsConfig, nil
}
