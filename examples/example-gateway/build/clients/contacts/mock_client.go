// Code generated by mockery v1.0.0

// +build ignore

package contactsClient

import contacts "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/contacts/contacts"
import context "context"
import mock "github.com/stretchr/testify/mock"
import zanzibar "github.com/uber/zanzibar/runtime"

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// HTTPClient provides a mock function with given fields:
func (_m *MockClient) HTTPClient() *zanzibar.HTTPClient {
	ret := _m.Called()

	var r0 *zanzibar.HTTPClient
	if rf, ok := ret.Get(0).(func() *zanzibar.HTTPClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zanzibar.HTTPClient)
		}
	}

	return r0
}

// SaveContacts provides a mock function with given fields: ctx, reqHeaders, args
func (_m *MockClient) SaveContacts(ctx context.Context, reqHeaders map[string]string, args *contacts.SaveContactsRequest) (*contacts.SaveContactsResponse, map[string]string, error) {
	ret := _m.Called(ctx, reqHeaders, args)

	var r0 *contacts.SaveContactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string, *contacts.SaveContactsRequest) *contacts.SaveContactsResponse); ok {
		r0 = rf(ctx, reqHeaders, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contacts.SaveContactsResponse)
		}
	}

	var r1 map[string]string
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string, *contacts.SaveContactsRequest) map[string]string); ok {
		r1 = rf(ctx, reqHeaders, args)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, map[string]string, *contacts.SaveContactsRequest) error); ok {
		r2 = rf(ctx, reqHeaders, args)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TestURLURL provides a mock function with given fields: ctx, reqHeaders
func (_m *MockClient) TestURLURL(ctx context.Context, reqHeaders map[string]string) (string, map[string]string, error) {
	ret := _m.Called(ctx, reqHeaders)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) string); ok {
		r0 = rf(ctx, reqHeaders)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 map[string]string
	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) map[string]string); ok {
		r1 = rf(ctx, reqHeaders)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, map[string]string) error); ok {
		r2 = rf(ctx, reqHeaders)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
