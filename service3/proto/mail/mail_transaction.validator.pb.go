// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mail_transaction.proto

package mail

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *MailTransaction) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.SubscriptionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SubscriptionId", fmt.Errorf(`value '%v' must not be an empty string`, this.SubscriptionId))
	}
	if this.UserId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must not be an empty string`, this.UserId))
	}
	if this.ReportedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ReportedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.ReportedBy))
	}
	return nil
}
