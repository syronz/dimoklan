package model

import (
	"dimoklan/consts/entity"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// func TestAuth(t *testing.T) {
// 	fmt.Println(">>>>> TestAuth")
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "auth Suite")
// }

var _ = Describe("Books", func() {
	var foxInSocks, lesMis *Auth

	BeforeEach(func() {
		lesMis = &Auth{
			Email: "email1",
		}

		foxInSocks = &Auth{
			Email: "email2",
		}
	})

	Describe("Categorizing books", func() {
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.ToRepo().EntityType).To(Equal(entity.Auth))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.ToRepo().EntityType).To(Equal(entity.Auth))
			})
		})
	})
})
