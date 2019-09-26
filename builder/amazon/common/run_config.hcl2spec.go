// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package common

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*RunConfig) HCL2Spec() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"AssociatePublicIpAddress": &hcldec.AttrSpec{Name:"AssociatePublicIpAddress", Type:cty.Bool, Required:false},
		"AvailabilityZone": &hcldec.AttrSpec{Name:"AvailabilityZone", Type:cty.String, Required:false},
		"BlockDurationMinutes": nil,
		"DisableStopInstance": &hcldec.AttrSpec{Name:"DisableStopInstance", Type:cty.Bool, Required:false},
		"EbsOptimized": &hcldec.AttrSpec{Name:"EbsOptimized", Type:cty.Bool, Required:false},
		"EnableT2Unlimited": &hcldec.AttrSpec{Name:"EnableT2Unlimited", Type:cty.Bool, Required:false},
		"IamInstanceProfile": &hcldec.AttrSpec{Name:"IamInstanceProfile", Type:cty.String, Required:false},
		"InstanceInitiatedShutdownBehavior": &hcldec.AttrSpec{Name:"InstanceInitiatedShutdownBehavior", Type:cty.String, Required:false},
		"InstanceType": &hcldec.AttrSpec{Name:"InstanceType", Type:cty.String, Required:false},
		"SecurityGroupFilter": nil,
		"RunTags": nil,
		"SecurityGroupId": &hcldec.AttrSpec{Name:"SecurityGroupId", Type:cty.String, Required:false},
		"SecurityGroupIds": &hcldec.AttrSpec{Name:"SecurityGroupIds", Type:cty.List(cty.String), Required:false},
		"SourceAmi": &hcldec.AttrSpec{Name:"SourceAmi", Type:cty.String, Required:false},
		"SourceAmiFilter": nil,
		"SpotInstanceTypes": &hcldec.AttrSpec{Name:"SpotInstanceTypes", Type:cty.List(cty.String), Required:false},
		"SpotPrice": &hcldec.AttrSpec{Name:"SpotPrice", Type:cty.String, Required:false},
		"SpotPriceAutoProduct": &hcldec.AttrSpec{Name:"SpotPriceAutoProduct", Type:cty.String, Required:false},
		"SpotTags": nil,
		"SubnetFilter": nil,
		"SubnetId": &hcldec.AttrSpec{Name:"SubnetId", Type:cty.String, Required:false},
		"TemporaryKeyPairName": &hcldec.AttrSpec{Name:"TemporaryKeyPairName", Type:cty.String, Required:false},
		"TemporarySGSourceCidrs": &hcldec.AttrSpec{Name:"TemporarySGSourceCidrs", Type:cty.List(cty.String), Required:false},
		"UserData": &hcldec.AttrSpec{Name:"UserData", Type:cty.String, Required:false},
		"UserDataFile": &hcldec.AttrSpec{Name:"UserDataFile", Type:cty.String, Required:false},
		"VpcFilter": nil,
		"VpcId": &hcldec.AttrSpec{Name:"VpcId", Type:cty.String, Required:false},
		"WindowsPasswordTimeout": &hcldec.AttrSpec{Name:"WindowsPasswordTimeout", Type:cty.String, Required:false},
		"Comm": nil,
		"SSHInterface": &hcldec.AttrSpec{Name:"SSHInterface", Type:cty.String, Required:false},
	}
	return hcldec.ObjectSpec(s)
}

func (*AmiFilterOptions) HCL2Spec() hcldec.ObjectSpec {
	s := map[string]hcldec.Spec{
		"Filters": nil,
		"Owners": nil,
		"MostRecent": &hcldec.AttrSpec{Name:"MostRecent", Type:cty.Bool, Required:false},
	}
	return hcldec.ObjectSpec(s)
}