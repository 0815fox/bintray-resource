package bintray_test

import (
	"github.com/hashicorp/go-version"
	"github.com/jamiemonserrate/bintray-resource/bintray"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Package", func() {
	Context(".VersionsSince", func() {
		It("Returns all versions after the provided version", func() {
			bintrayPackage := bintray.Package{RawVersions: []string{"3.3.3", "2.2.2", "1.1.1"}}

			version1, _ := version.NewVersion("1.1.1")
			expectedVersion2, _ := version.NewVersion("2.2.2")
			expectedVersion3, _ := version.NewVersion("3.3.3")
			Expect(bintrayPackage.VersionsSince(version1)).To(Equal([]*version.Version{expectedVersion3, expectedVersion2}))
		})
	})
})
