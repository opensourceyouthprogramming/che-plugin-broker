//
// Copyright (c) 2019 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package theia

import (
	"reflect"
	"testing"

	"github.com/eclipse/che-plugin-broker/common"
	"github.com/eclipse/che-plugin-broker/model"
)

func TestGenerateSidecarTooling(t *testing.T) {
	type args struct {
		image string
		pj    model.PackageJSON
		rand  common.Random
	}
	tests := []struct {
		name string
		args args
		want *model.ToolingConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSidecarTooling(tt.args.image, tt.args.pj, tt.args.rand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSidecarTooling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containerConfig(t *testing.T) {
	type args struct {
		image string
		rand  common.Random
	}
	tests := []struct {
		name string
		args args
		want *model.Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containerConfig(tt.args.image, tt.args.rand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("containerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addPortToTooling(t *testing.T) {
	type args struct {
		toolingConf *model.ToolingConf
		pj          model.PackageJSON
		rand        common.Random
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addPortToTooling(tt.args.toolingConf, tt.args.pj, tt.args.rand)
		})
	}
}
