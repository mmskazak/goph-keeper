package config

import (
	"dario.cat/mergo"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type Student struct {
	name string
	age  int
	iq   string
}

func TestMerGo(t *testing.T) {
	stu1 := &Student{
		name: "stu1",
		age:  20,
		iq:   "",
	}
	stu2 := &Student{
		name: "stu2",
		age:  41,
		iq:   "120",
	}

	err := mergo.Merge(stu1, stu2, mergo.WithOverride)
	require.NoError(t, err)
	assert.Equal(t, "stu2", stu1.name)
	assert.Equal(t, 41, stu1.age)
	assert.Equal(t, "120", stu1.iq)
	fmt.Printf("%+v\n", stu1)
}
