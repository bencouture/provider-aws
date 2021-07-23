/*
Copyright 2021 The Crossplane Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ec2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/crossplane/provider-aws/apis/ec2/manualv1alpha1"
)

const (
	// InstanceNotFound is the code that is returned by ec2 when the given InstanceID is not valid
	InstanceNotFound = "InvalidInstanceID.NotFound"
)

// InstanceClient is the external client used for VPC Custom Resourc
type InstanceClient interface {
	RunInstancesRequest(*ec2.RunInstancesInput) ec2.RunInstancesRequest
	TerminateInstancesRequest(*ec2.TerminateInstancesInput) ec2.TerminateInstancesRequest
	DescribeInstancesRequest(*ec2.DescribeInstancesInput) ec2.DescribeInstancesRequest
	DescribeInstanceAttributeRequest(*ec2.DescribeInstanceAttributeInput) ec2.DescribeInstanceAttributeRequest
	ModifyInstanceAttributeRequest(*ec2.ModifyInstanceAttributeInput) ec2.ModifyInstanceAttributeRequest
	// CreateTagsRequest(*ec2.CreateTagsInput) ec2.CreateTagsRequest
	// ModifyVpcTenancyRequest(*ec2.ModifyVpcTenancyInput) ec2.ModifyVpcTenancyRequest
}

// NewInstanceClient returns a new client using AWS credentials as JSON encoded data.
func NewInstanceClient(cfg aws.Config) InstanceClient {
	return ec2.New(cfg)
}

// IsInstanceNotFoundErr returns true if the error is because the item doesn't exist
func IsInstanceNotFoundErr(err error) bool {
	if awsErr, ok := err.(awserr.Error); ok {
		if awsErr.Code() == InstanceNotFound {
			return true
		}
	}

	return false
}

// IsInstanceUpToDate returns true if there is no update-able difference between desired
// and observed state of the resource.
func IsInstanceUpToDate(spec manualv1alpha1.InstanceParameters, instance ec2.Instance, attributes ec2.DescribeInstanceAttributeOutput) bool {
	// if aws.StringValue(spec.InstanceTenancy) != string(vpc.InstanceTenancy) {
	// 	return false
	// }

	// if aws.BoolValue(spec.EnableDNSHostNames) != aws.BoolValue(attributes.EnableDnsHostnames.Value) ||
	// 	aws.BoolValue(spec.EnableDNSSupport) != aws.BoolValue(attributes.EnableDnsSupport.Value) {
	// 	return false
	// }

	// return manualv1alpha1.CompareTags(spec.Tags, vpc.Tags)
	return true
}

// GenerateInstanceObservation is used to produce manualv1alpha1.InstanceObservation from
// ec2.Instance.
func GenerateInstanceObservation(vpc ec2.Instance) manualv1alpha1.InstanceObservation {
	o := manualv1alpha1.InstanceObservation{
		// IsDefault:     aws.BoolValue(vpc.IsDefault),
		// DHCPOptionsID: aws.StringValue(vpc.DhcpOptionsId),
		// OwnerID:       aws.StringValue(vpc.OwnerId),
		// VPCState:      string(vpc.State),
	}

	// if len(vpc.CidrBlockAssociationSet) > 0 {
	// 	o.CIDRBlockAssociationSet = make([]v1beta1.VPCCIDRBlockAssociation, len(vpc.CidrBlockAssociationSet))
	// 	for i, v := range vpc.CidrBlockAssociationSet {
	// 		o.CIDRBlockAssociationSet[i] = v1beta1.VPCCIDRBlockAssociation{
	// 			AssociationID: aws.StringValue(v.AssociationId),
	// 			CIDRBlock:     aws.StringValue(v.CidrBlock),
	// 		}
	// 		o.CIDRBlockAssociationSet[i].CIDRBlockState = v1beta1.VPCCIDRBlockState{
	// 			State:         string(v.CidrBlockState.State),
	// 			StatusMessage: aws.StringValue(v.CidrBlockState.StatusMessage),
	// 		}
	// 	}
	// }

	// if len(vpc.Ipv6CidrBlockAssociationSet) > 0 {
	// 	o.IPv6CIDRBlockAssociationSet = make([]v1beta1.VPCIPv6CidrBlockAssociation, len(vpc.Ipv6CidrBlockAssociationSet))
	// 	for i, v := range vpc.Ipv6CidrBlockAssociationSet {
	// 		o.IPv6CIDRBlockAssociationSet[i] = v1beta1.VPCIPv6CidrBlockAssociation{
	// 			AssociationID:      aws.StringValue(v.AssociationId),
	// 			IPv6CIDRBlock:      aws.StringValue(v.Ipv6CidrBlock),
	// 			IPv6Pool:           aws.StringValue(v.Ipv6Pool),
	// 			NetworkBorderGroup: aws.StringValue(v.NetworkBorderGroup),
	// 		}
	// 		o.IPv6CIDRBlockAssociationSet[i].IPv6CIDRBlockState = v1beta1.VPCCIDRBlockState{
	// 			State:         string(v.Ipv6CidrBlockState.State),
	// 			StatusMessage: raws.StringValue(v.Ipv6CidrBlockState.StatusMessage),
	// 		}
	// 	}
	// }

	return o
}

// LateInitializeInstance fills the empty fields in *manualv1alpha1.InstanceParameters with
// the values seen in ec2.Instance and ec2.DescribeInstanceAttributeOutput.
func LateInitializeInstance(in *manualv1alpha1.InstanceParameters, v *ec2.Instance, attributes *ec2.DescribeInstanceAttributeOutput) { // nolint:gocyclo
	if v == nil {
		return
	}

	// in.CIDRBlock = awsclients.LateInitializeString(in.CIDRBlock, v.CidrBlock)
	// in.InstanceTenancy = awsclients.LateInitializeStringPtr(in.InstanceTenancy, aws.String(string(v.InstanceTenancy)))
	// if attributes.EnableDnsHostnames != nil {
	// 	in.EnableDNSHostNames = awsclients.LateInitializeBoolPtr(in.EnableDNSHostNames, attributes.EnableDnsHostnames.Value)
	// }
	// if attributes.EnableDnsHostnames != nil {
	// 	in.EnableDNSSupport = awsclients.LateInitializeBoolPtr(in.EnableDNSSupport, attributes.EnableDnsSupport.Value)
	// }
}

// GenerateEC2Monitoring converts internal RunInstancesMonitoringEnabled into ec2.RunInstancesMonitoringEnabled
func GenerateEC2Monitoring(m *manualv1alpha1.RunInstancesMonitoringEnabled) *ec2.RunInstancesMonitoringEnabled {
	if m != nil {
		var res ec2.RunInstancesMonitoringEnabled
		res.Enabled = m.Enabled
		return &res
	}
	return nil
}

// TransformTagSpecifications takes a slice of TagSpecifications, converts it to a slice
// of ec2.TagSpecification and lastly injects the special instance name awsec2.TagSpecification
func TransformTagSpecifications(mdgName string, tagSpecs []manualv1alpha1.TagSpecification) []ec2.TagSpecification {
	ec2Specs := generateEC2TagSpecifications(tagSpecs)
	return injectInstanceNameTagSpecification(mdgName, ec2Specs)
}

// generateEC2TagSpecifications takes a slice of TagSpecifications and converts it to a
// slice of ec2.TagSpecification
func generateEC2TagSpecifications(tagSpecs []manualv1alpha1.TagSpecification) []ec2.TagSpecification {
	res := make([]ec2.TagSpecification, len(tagSpecs))
	for i, ts := range tagSpecs {
		res[i] = ec2.TagSpecification{
			ResourceType: ec2.ResourceType(*ts.ResourceType),
		}

		tags := make([]ec2.Tag, len(ts.Tags))
		for i, t := range ts.Tags {
			tags[i] = ec2.Tag{
				Key:   aws.String(t.Key),
				Value: aws.String(t.Value),
			}
		}

		res[i].Tags = tags
	}
	return res
}

// injectInstanceNameTagSpecification will inject a special TagSpecification of the following
// shape into the give slice, if it does not yet exist:
// resourceType: instance,
// tags:
//   - key: Name
//     value: <name>
// The resulting behavior is that the specified instance name in metadata.Name is reflected
// in the AWS console.
// Note: If the a TagSpecification of resourceType `instance` with key `Name` was supplied,
// this method does not modify the supplied TagSpecification.
func injectInstanceNameTagSpecification(name string, tagSpecs []ec2.TagSpecification) []ec2.TagSpecification {
	instanceNameKey := "Name"

	specialTagSpec := ec2.TagSpecification{
		ResourceType: "instance",
		Tags: []ec2.Tag{
			{
				Key:   aws.String(instanceNameKey), // if the 'N' isn't capitalized AWS treats this as a general tag
				Value: aws.String(name),
			},
		},
	}

	foundSpecial := false
	for _, ts := range tagSpecs {
		if ts.ResourceType == "instance" {
			for _, tag := range ts.Tags {
				if *tag.Key == instanceNameKey {
					foundSpecial = true
					break
				}
			}
		}
	}

	if !foundSpecial {
		tagSpecs = append(tagSpecs, specialTagSpec)
	}

	return tagSpecs
}
