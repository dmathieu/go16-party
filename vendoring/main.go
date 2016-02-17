package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

	var params *s3.ListBucketsInput
	resp, err := svc.ListBuckets(params)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var bucketNames []string
	for _, bucket := range resp.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}

	content := fmt.Sprintf("You have %d buckets: %s", len(resp.Buckets), strings.Join(bucketNames, ", "))
	w.Write([]byte(content))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8000", r)
}
