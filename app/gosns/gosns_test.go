package gosns

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"net/url"
	"io/ioutil"
	"strconv"
	"encoding/json"
)

func TestListTopicshandler_POST_NoTopics(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTopics)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "<Topics></Topics>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateTopicshandler_POST_CreateTopics(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{}
	form.Add("Action", "CreateTopic")
	form.Add("Name", "UnitTestTopic1")
	req.PostForm = form

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTopic)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "UnitTestTopic1"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPublishhandler_POST_SendMessage(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{}
	form.Add("TopicArn", "arn:aws:sns:local:000000000000:UnitTestTopic1")
	form.Add("Message", "TestMessage1")
	req.PostForm = form

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Publish)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "<MessageId>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSubscribehandler_POST_Success(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{}
	form.Add("TopicArn", "arn:aws:sns:local:000000000000:UnitTestTopic1")
	form.Add("Protocol", "sqs")
	form.Add("Endpoint", "http://localhost:4100/queue/noqueue1")
	req.PostForm = form

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Subscribe)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "</SubscriptionArn>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPublish_No_Queue_Error_handler_POST_Success(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{}
	form.Add("TopicArn", "arn:aws:sns:local:000000000000:UnitTestTopic1")
	form.Add("Message", "TestMessage1")
	req.PostForm = form

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Publish)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "<MessageId>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteTopichandler_POST_Success(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{}
	form.Add("TopicArn", "arn:aws:sns:local:000000000000:UnitTestTopic1")
	form.Add("Message", "TestMessage1")
	req.PostForm = form

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTopic)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "</DeleteTopicResponse>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	// Check the response body is what we expect.
	expected = "</ResponseMetadata>"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
<<<<<<< HEAD:gosns/gosns_test.go

/*
 * https://golang.org/pkg/net/http/httptest/#example_Server
 */
func TestPublish_POST_http_endpoint(t *testing.T) {

	const TopicName string = "UnitTestTopic2"
	const TopicArn string = "arn:aws:sns:local:000000000000:" + TopicName

	/*
	 * Create a mock-http endpoint to respond to our published messages.
	 * TODO This does not handle "SubscribeURL"
	 */
	var request http.Request
	var body []byte
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture our Request for later analysis
		request = *r
		// Dump the body while we have access to it
		body, _ = ioutil.ReadAll(r.Body)
		r.Body.Close()

		// Create a response
		w.Header().Set("Content-Type", "text/html")
	}))
	defer ts.Close()

	/*
	 * Create a topic
	 */
	{
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		form := url.Values{}
		form.Add("Action", "CreateTopic")
		form.Add("Name", TopicName)
		req.PostForm = form


		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateTopic)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	/*
	 * Subscribe to the topic with our mock-http endpoint
	 */
	{
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		form := url.Values{}
		form.Add("TopicArn", TopicArn)
		form.Add("Protocol", "http")
		form.Add("Endpoint", ts.URL)
		req.PostForm = form

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Subscribe)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	/*
	 * Next publish a message
	 */
	{
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		form := url.Values{}
		form.Add("TopicArn", TopicArn)
		form.Add("Message", "TestMessage1")
		form.Add("Subject", "TestSubject1")
		req.PostForm = form

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Publish)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	/*
	 * Assert the results from our request
	 */
	if request.Method != http.MethodPost {
		t.Errorf("request method was unexpected: got %s want %s",
				request.Method, "POST")
	}

	/*
	 * Assert our body is properly formatted
	 */
	if (len(body) == 0)	 {
		t.Fatal("request body is empty")
	}

	content_length, _ := strconv.Atoi(request.Header.Get("Content-Length"))
	if (content_length != len(body)) {
		t.Errorf("request body length (%v) does not match Content-Length (%v)",
				len(body),
				content_length)
	}

	for _, f := range []string{"x-amz-sns-message-type",
				"x-amz-sns-message-id",
				"x-amz-sns-topic-arn"} {
		field := request.Header.Get(f)
		if field == "" {
			t.Errorf("request header %v is missing and/or empty", f)
		}
	}

	// This should be JSON
	jbody := map[string]string{}
	err := json.Unmarshal(body, &jbody)
	if err != nil {
		t.Fatal(err)
	}

	if jbody["Type"] != "Notification" {
		t.Errorf("request body Type does not match.	got: %v  expected: %v", jbody["Type"], "Notification")
	}
	// verify that these keys exist and check that the data is not empty
	for _, f := range []string{"MessageId",
				"TopicArn",
				"Subject",
				"Message",
				"Timestamp",
				"SignatureVersion",
				"Signature",
				"SigningCertURL"} {
		if jbody[f] == "" {
			t.Errorf("request body %v is missing and/or empty", f)
		}
	}

}

=======
>>>>>>> upstream/master:app/gosns/gosns_test.go
