package typed

import (
	"testing"

	"go.uber.org/mock/gomock"
)

// TestValidInterface assesses whether or not the generated mock is valid
func TestValidInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := NewMockSource(ctrl)
	s.EXPECT().Method().Return("")

	CallForeignMethod(s)
}
