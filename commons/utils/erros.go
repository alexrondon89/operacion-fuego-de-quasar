package utils

const (
	QueryDynamoError = "One error occurred during query execution!: "
	UpdateDynamoError = "One error occurred during update execution!: "
	UnknownSatellite = "This satellite has not sent any message before "
	EmptyDataBase	 = "there are no satellites with enemy information sent "
	InsufficientInformation ="It´s not possible to get enemy information "
	EmptySatelliteList	= "at least one satellite must be added "
	InvalidInformation = "information from satellite is invalid "
	)

func CheckError(err error) bool {
	return err != nil
}


