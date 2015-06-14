package check

import (
	"github.com/hashicorp/go-version"
	"github.com/jamiemonserrate/bintray-resource"
	"github.com/jamiemonserrate/bintray-resource/bintray"
)

type CheckCommand struct {
	bintrayClient bintray.BintrayClient
}

func NewCheckCommand(bintrayClient bintray.BintrayClient) CheckCommand {
	return CheckCommand{bintrayClient: bintrayClient}
}

func (checkCommand *CheckCommand) Execute(checkRequest CheckRequest) CheckResponse {
	bintrayPackage := checkCommand.bintrayClient.GetPackage(checkRequest.Source.PackageName)

	response := CheckResponse{}
	for _, v := range bintrayPackage.VersionsSince(checkRequest.Version()) {
		response = append(response, bintrayresource.Version{Number: v.String()})
	}
	return response
}

func toVersion(versionString string) *version.Version {
	version, _ := version.NewVersion(versionString)
	return version
}
