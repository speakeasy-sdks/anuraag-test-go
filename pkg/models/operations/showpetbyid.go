// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"test-petstore/pkg/models/shared"
)

type ShowPetByIDRequest struct {
	// The id of the pet to retrieve
	PetID string `pathParam:"style=simple,explode=false,name=petId"`
}

type ShowPetByIDResponse struct {
	ContentType string
	// unexpected error
	Error *shared.Error
	// Expected response to a valid request
	Pet         *shared.Pet
	StatusCode  int
	RawResponse *http.Response
}
