package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2SpotFleet AWS CloudFormation Resource (AWS::EC2::SpotFleet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html
type AWSEC2SpotFleet struct {

	// SpotFleetRequestConfigData AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-spotfleet.html#cfn-ec2-spotfleet-spotfleetrequestconfigdata
	SpotFleetRequestConfigData *AWSEC2SpotFleet_SpotFleetRequestConfigData `json:"SpotFleetRequestConfigData,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SpotFleet) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2SpotFleet) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2SpotFleet
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2SpotFleet) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2SpotFleet
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSEC2SpotFleet(*res.Properties)
	}

	return nil
}

// GetAllAWSEC2SpotFleetResources retrieves all AWSEC2SpotFleet items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SpotFleetResources() map[string]AWSEC2SpotFleet {
	results := map[string]AWSEC2SpotFleet{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2SpotFleet:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SpotFleet" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2SpotFleet
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2SpotFleetWithName retrieves all AWSEC2SpotFleet items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SpotFleetWithName(name string) (AWSEC2SpotFleet, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2SpotFleet:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SpotFleet" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2SpotFleet
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2SpotFleet{}, errors.New("resource not found")
}
