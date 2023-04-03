// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a contact from a contact list.
func (c *Client) GetContact(ctx context.Context, params *GetContactInput, optFns ...func(*Options)) (*GetContactOutput, error) {
	if params == nil {
		params = &GetContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContact", params, optFns, c.addOperationGetContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContactInput struct {

	// The name of the contact list to which the contact belongs.
	//
	// This member is required.
	ContactListName *string

	// The contact's email addres.
	//
	// This member is required.
	EmailAddress *string

	noSmithyDocumentSerde
}

type GetContactOutput struct {

	// The attribute data attached to a contact.
	AttributesData *string

	// The name of the contact list to which the contact belongs.
	ContactListName *string

	// A timestamp noting when the contact was created.
	CreatedTimestamp *time.Time

	// The contact's email addres.
	EmailAddress *string

	// A timestamp noting the last time the contact's information was updated.
	LastUpdatedTimestamp *time.Time

	// The default topic preferences applied to the contact.
	TopicDefaultPreferences []types.TopicPreference

	// The contact's preference for being opted-in to or opted-out of a topic.>
	TopicPreferences []types.TopicPreference

	// A boolean value status noting if the contact is unsubscribed from all contact
	// list topics.
	UnsubscribeAll bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContact(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetContact",
	}
}
