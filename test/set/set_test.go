package set_test

import (
	adt "TDD-in-GoLang/AbstractDataTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var s *adt.Set
var _ = Describe("Set", func() {
	BeforeEach(func() {
		s = adt.NewSet()
	})
	Describe("Emptiness", func() {
		Context("When the set does not contain items", func() {
			It("Should be empty", func() {
				Expect(s.IsEmpty()).To(BeTrue())
			})
		})

		Context("When the set contains items", func() {
			It("should not be empty", func() {
				s.Add("White")
				Expect(s.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func() {
		Context("As items are added", func() {
			It("Should return an increasing size", func() {
				By("Initially the set size being 0", func() {
					Expect(s.IsEmpty()).To(BeTrue())
				})
				By("Adding a first item", func() {
					s.Add("Red")
					Expect(s.Size()).To(Equal(1))
				})
				By("Adding a second item", func() {

					s.Add("Green")
					Expect(s.Size()).To(Equal(2))
				})
			})
		})
	})

	Describe("Contains", func() {
		Context("When red has not been added", func() {
			It("should not contain red", func() {
				Expect(s.Contains("Red")).To(BeFalse())
			})
		})

		Context("When red has been added", func() {
			It("Should contains red", func() {
				s.Add("Red")
				Expect(s.Contains("Red")).To(BeTrue())
			})
		})
	})
})
