package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-poll-api/internal/flusher"
	"github.com/ozoncp/ocp-poll-api/internal/mocks"
	"github.com/ozoncp/ocp-poll-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		controller *gomock.Controller
		repo *mocks.MockRepo
		f flusher.Flusher
		polls []models.Poll
		chunkSize int
	)
	BeforeEach(func() {
		controller = gomock.NewController(GinkgoT())
		repo = mocks.NewMockRepo(controller)
		polls = []models.Poll {
			0: {0,0,0,0},
			1: {1,1,1,1},
			2: {2,2,2,2},
			3: {3,1,2,3},
			4: {4,0,0,0},
		}
	})
	AfterEach(func() {
		controller.Finish()
	})
	Context("Flush", func() {
		When("Correct test", func() {
			BeforeEach(func() {
				chunkSize = 3
				repo.EXPECT().AddEntities(gomock.Any()).Return(nil).AnyTimes()
				f = flusher.NewFlusher(chunkSize, repo)
			})
			It("returnsNil", func() {
				result, err := f.Flush(polls)
				Expect(result).Should(BeNil())
				Expect(err).Should(BeNil())
			})
		})
		When("Error - size = nil", func() {
			BeforeEach(func() {
				chunkSize = 0
				repo.EXPECT().AddEntities(gomock.Any()).Return(nil).AnyTimes()
				f = flusher.NewFlusher(chunkSize, repo)
			})
			It("error exists", func() {
				result, err := f.Flush(polls)
				Expect(result).Should(BeNil())
				Expect(err).Should(Not(BeNil()))
			})
		})
		When("Error - entities array is nil", func() {
			BeforeEach(func() {
				chunkSize = 3
				repo.EXPECT().AddEntities(gomock.Any()).Return(nil).AnyTimes()
				f = flusher.NewFlusher(chunkSize, repo)
			})
			It("error exists", func() {
				result, err := f.Flush(nil)
				Expect(result).Should(BeNil())
				Expect(err).Should(Not(BeNil()))
			})
		})
	})
})
