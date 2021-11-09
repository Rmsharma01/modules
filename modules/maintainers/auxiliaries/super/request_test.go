/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package super

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Super_Request(t *testing.T) {

	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("Data1")))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, identityID, mutableProperties)

	require.Equal(t, auxiliaryRequest{ClassificationID: classificationID, IdentityID: identityID, MutableProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
