package tests

import (
	"github.com/dencat/id3golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByteToIntSynchsafe(t *testing.T) {
	asrt := assert.New(t)
	// {0,0,16,30} -> 2078
	sizeInt := id3golang.ByteToIntSynchsafe([]byte{0, 0, 16, 30})
	asrt.Equal(2078, sizeInt)
}

func TestIntToByteSynchsafe(t *testing.T) {
	asrt := assert.New(t)
	// 2078 -> {0,0,16,30}
	sizeByte := id3golang.IntToByteSynchsafe(2078)
	asrt.Equal([]byte{0, 0, 16, 30}, sizeByte)
}
