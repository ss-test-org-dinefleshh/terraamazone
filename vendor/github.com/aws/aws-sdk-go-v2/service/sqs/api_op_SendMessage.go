// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendMessageInput struct {
	_ struct{} `type:"structure"`

	// The length of time, in seconds, for which to delay a specific message. Valid
	// values: 0 to 900. Maximum: 15 minutes. Messages with a positive DelaySeconds
	// value become available for processing after the delay period is finished.
	// If you don't specify a value, the default value for the queue applies.
	//
	// When you set FifoQueue, you can't set DelaySeconds per message. You can set
	// this parameter only on a queue level.
	DelaySeconds *int64 `type:"integer"`

	// Each message attribute consists of a Name, Type, and Value. For more information,
	// see Amazon SQS Message Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageAttributes map[string]MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The message to send. The maximum string size is 256 KB.
	//
	// A message can include only XML, JSON, and unformatted text. The following
	// Unicode characters are allowed:
	//
	// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
	//
	// Any characters not included in this list will be rejected. For more information,
	// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
	//
	// MessageBody is a required field
	MessageBody *string `type:"string" required:"true"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of sent messages. If a message with a particular
	// MessageDeduplicationId is sent successfully, any messages sent with the same
	// MessageDeduplicationId are accepted successfully but aren't delivered during
	// the 5-minute deduplication interval. For more information, see Exactly-Once
	// Processing (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//    * Every message must have a unique MessageDeduplicationId, You may provide
	//    a MessageDeduplicationId explicitly. If you aren't able to provide a MessageDeduplicationId
	//    and you enable ContentBasedDeduplication for your queue, Amazon SQS uses
	//    a SHA-256 hash to generate the MessageDeduplicationId using the body of
	//    the message (but not the attributes of the message). If you don't provide
	//    a MessageDeduplicationId and the queue doesn't have ContentBasedDeduplication
	//    set, the action fails with an error. If the queue has ContentBasedDeduplication
	//    set, your MessageDeduplicationId overrides the generated one.
	//
	//    * When ContentBasedDeduplication is in effect, messages with identical
	//    content sent within the deduplication interval are treated as duplicates
	//    and only one copy of the message is delivered.
	//
	//    * If you send one message with ContentBasedDeduplication enabled and then
	//    another message with a MessageDeduplicationId that is the same as the
	//    one generated for the first MessageDeduplicationId, the two messages are
	//    treated as duplicates and only one copy of the message is delivered.
	//
	// The MessageDeduplicationId is available to the consumer of the message (this
	// can be useful for troubleshooting delivery issues).
	//
	// If a message is sent successfully but the acknowledgement is lost and the
	// message is resent with the same MessageDeduplicationId after the deduplication
	// interval, Amazon SQS can't detect duplicate messages.
	//
	// Amazon SQS continues to keep track of the message deduplication ID even after
	// the message is received and deleted.
	//
	// The length of MessageDeduplicationId is 128 characters. MessageDeduplicationId
	// can contain alphanumeric characters (a-z, A-Z, 0-9) and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageDeduplicationId, see Using the MessageDeduplicationId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageDeduplicationId *string `type:"string"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The tag that specifies that a message belongs to a specific message group.
	// Messages that belong to the same message group are processed in a FIFO manner
	// (however, messages in different message groups might be processed out of
	// order). To interleave multiple ordered streams within a single queue, use
	// MessageGroupId values (for example, session data for multiple users). In
	// this scenario, multiple consumers can process the queue, but the session
	// data of each user is processed in a FIFO fashion.
	//
	//    * You must associate a non-empty MessageGroupId with a message. If you
	//    don't provide a MessageGroupId, the action fails.
	//
	//    * ReceiveMessage might return messages with multiple MessageGroupId values.
	//    For each MessageGroupId, the messages are sorted by time sent. The caller
	//    can't specify a MessageGroupId.
	//
	// The length of MessageGroupId is 128 characters. Valid values: alphanumeric
	// characters and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageGroupId, see Using the MessageGroupId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagegroupid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// MessageGroupId is required for FIFO queues. You can't use it for Standard
	// queues.
	MessageGroupId *string `type:"string"`

	// The message system attribute to send. Each message system attribute consists
	// of a Name, Type, and Value.
	//
	//    * Currently, the only supported message system attribute is AWSTraceHeader.
	//    Its type must be String and its value must be a correctly formatted AWS
	//    X-Ray trace string.
	//
	//    * The size of a message system attribute doesn't count towards the total
	//    size of a message.
	MessageSystemAttributes map[string]MessageSystemAttributeValue `locationName:"MessageSystemAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The URL of the Amazon SQS queue to which a message is sent.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SendMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendMessageInput"}

	if s.MessageBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("MessageBody"))
	}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}
	if s.MessageAttributes != nil {
		for i, v := range s.MessageAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.MessageSystemAttributes != nil {
		for i, v := range s.MessageSystemAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageSystemAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The MD5OfMessageBody and MessageId elements.
type SendMessageOutput struct {
	_ struct{} `type:"structure"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageAttributes *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageBody *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message system attribute string. You
	// can use this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest.
	MD5OfMessageSystemAttributes *string `type:"string"`

	// An attribute containing the MessageId of the message sent to the queue. For
	// more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageId *string `type:"string"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The large, non-consecutive number that Amazon SQS assigns to each message.
	//
	// The length of SequenceNumber is 128 bits. SequenceNumber continues to increase
	// for a particular MessageGroupId.
	SequenceNumber *string `type:"string"`
}

// String returns the string representation
func (s SendMessageOutput) String() string {
	return awsutil.Prettify(s)
}

const opSendMessage = "SendMessage"

// SendMessageRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Delivers a message to the specified queue.
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed:
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// Any characters not included in this list will be rejected. For more information,
// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
//
//    // Example sending a request using SendMessageRequest.
//    req := client.SendMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SendMessage
func (c *Client) SendMessageRequest(input *SendMessageInput) SendMessageRequest {
	op := &aws.Operation{
		Name:       opSendMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendMessageInput{}
	}

	req := c.newRequest(op, input, &SendMessageOutput{})
	return SendMessageRequest{Request: req, Input: input, Copy: c.SendMessageRequest}
}

// SendMessageRequest is the request type for the
// SendMessage API operation.
type SendMessageRequest struct {
	*aws.Request
	Input *SendMessageInput
	Copy  func(*SendMessageInput) SendMessageRequest
}

// Send marshals and sends the SendMessage API request.
func (r SendMessageRequest) Send(ctx context.Context) (*SendMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SendMessageResponse{
		SendMessageOutput: r.Request.Data.(*SendMessageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SendMessageResponse is the response type for the
// SendMessage API operation.
type SendMessageResponse struct {
	*SendMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SendMessage request.
func (r *SendMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}