// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIListRunsResponse api list runs response
// swagger:model apiListRunsResponse
type APIListRunsResponse struct {

	// next page token
	NextPageToken string `json:"next_page_token,omitempty"`

	// run details
	RunDetails []*APIRunDetail `json:"run_details"`

	// runs
	Runs []*APIRun `json:"runs"`

	// total size
	TotalSize int32 `json:"total_size,omitempty"`
}

// Validate validates this api list runs response
func (m *APIListRunsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIListRunsResponse) validateRunDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.RunDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.RunDetails); i++ {
		if swag.IsZero(m.RunDetails[i]) { // not required
			continue
		}

		if m.RunDetails[i] != nil {
			if err := m.RunDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("run_details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIListRunsResponse) validateRuns(formats strfmt.Registry) error {

	if swag.IsZero(m.Runs) { // not required
		return nil
	}

	for i := 0; i < len(m.Runs); i++ {
		if swag.IsZero(m.Runs[i]) { // not required
			continue
		}

		if m.Runs[i] != nil {
			if err := m.Runs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIListRunsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIListRunsResponse) UnmarshalBinary(b []byte) error {
	var res APIListRunsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
