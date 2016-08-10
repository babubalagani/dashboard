// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

// Options for GenericDataSelect which takes []GenericDataCell and returns selected data.
// Can be extended to include any kind of selection - for example filtering.
// Currently included only Pagination and Sort options.
type DataSelectQuery struct {
	PaginationQuery *PaginationQuery
	SortQuery       *SortQuery
	//	Filter     *FilterQuery
}
// SortQuery holds options for sort functionality of data select.
type SortQuery struct {
	SortByList []SortBy

}

// SortBy holds the name of the property that should be sorted and whether order should be ascending or descending.
type SortBy struct {
	Property  PropertyName
	Ascending bool
}

// NoSort is as option for no sort.
var NoSort = &SortQuery{
	SortByList: []SortBy{},
}

// NoDataSelect is an option for no data select (same data will be returned).
var NoDataSelect = NewDataSelectQuery(NoPagination, NoSort)


// NewDataSelectQuery creates DataSelectQuery object from simpler data select queries.
func NewDataSelectQuery(paginationQuery *PaginationQuery, sortQuery *SortQuery) (*DataSelectQuery) {
	return &DataSelectQuery{
		PaginationQuery: paginationQuery,
		SortQuery:       sortQuery,
	}
}

// NewSortQuery takes raw sort options list and returns SortQuery object. For example:
// ["a", "parameter1", "d", "parameter2"] - means that the data should be sorted by
// parameter1 (ascending) and later - for results that return equal under parameter 1 sort - by parameter2 (descending)
func NewSortQuery(sortByListRaw []string) (*SortQuery) {
	if sortByListRaw == nil || len(sortByListRaw)%2 == 1{
		// Empty sort list or invalid (odd) length
		return NoSort
	}
	sortByList := []SortBy{}
	for i:=0; i+1 < len(sortByListRaw); i+=2 {
		// parse order option
		var ascending bool
		orderOption := sortByListRaw[i]
		if orderOption == "a" {
			ascending = true
		} else if orderOption == "d" {
			ascending = false
		} else {
			//  Invalid order option. Only ascending (a), descending (d) options are supported
			return NoSort
		}

		// parse property name
		propertyName := sortByListRaw[i+1]
		sortBy := SortBy{
			Property: PropertyName(propertyName),
			Ascending: ascending,
		}
		// Add to the sort options.
		sortByList = append(sortByList, sortBy)
	}
	return &SortQuery{
		SortByList: sortByList,
	}
}
