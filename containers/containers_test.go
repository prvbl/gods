// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// All data structures must implement the container structure

package containers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/emirpasic/gods/utils"
)

// For testing purposes
type ContainerTest[V comparable] struct {
	values []V
}

func (container ContainerTest[V]) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest[V]) Size() int {
	return len(container.values)
}

func (container ContainerTest[V]) Clear() {
	container.values = []V{}
}

func (container ContainerTest[V]) Values() []V {
	return container.values
}

func (container ContainerTest[V]) String() string {
	str := "ContainerTest\n"
	var values []string
	for _, value := range container.values {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest[int]{}
	GetSortedValues[int](container, utils.IntComparator)
	container.values = []int{5, 1, 3, 2, 4}
	values := GetSortedValues[int](container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainerTest[string]{}
	GetSortedValues[string](container, utils.StringComparator)
	container.values = []string{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues[string](container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

// For testing purposes
type StringsContainerTest struct {
	values []string
}

func (container StringsContainerTest) Empty() bool {
	return len(container.values) == 0
}

func (container StringsContainerTest) Size() int {
	return len(container.values)
}

func (container StringsContainerTest) Clear() {
	container.values = []string{}
}

func (container StringsContainerTest) Values() []string {
	return container.values
}

func (container StringsContainerTest) String() string {
	str := "StringsContainerTest\n"
	var values []string
	for _, value := range container.values {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func TestGetSortedValuesStringsFromStringsContainer(t *testing.T) {
	var stringsContainer = StringsContainerTest{}
	var container Container[string] = stringsContainer

	GetSortedValues(container, utils.StringComparator)
	stringsContainer.values = []string{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues(container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}
