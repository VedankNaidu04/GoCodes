// ********* example of a GO program that makes an HTTP GET request to an API
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL of the API endpoint
	apiUrl := "https://api.example.com/data"
	// Send GET request
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error making GET request: %s\n", err)
		return
	}
	defer response.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %5\n", err)
		return
	}
	// Print the response
	fmt.Println(string(body))
}

// ************ GO SDK for aws instance "https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ec2-example-manage-instances.html"
// Start and stop instances using GO
package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)
func main() {
    // Load session from shared config
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create new EC2 client
    svc := ec2.New(sess)

	if os.Args[1] == "START" {
        input := &ec2.StartInstancesInput{
            InstanceIds: []*string{
                aws.String(os.Args[2]),
            },
            DryRun: aws.Bool(true),
        }
        result, err := svc.StartInstances(input)
        awsErr, ok := err.(awserr.Error)

        if ok && awsErr.Code() == "DryRunOperation" {
            // Let's now set dry run to be false. This will allow us to start the instances
            input.DryRun = aws.Bool(false)
            result, err = svc.StartInstances(input)
            if err != nil {
                fmt.Println("Error", err)
            } else {
                fmt.Println("Success", result.StartingInstances)
            }
        } else { // This could be due to a lack of permissions
            fmt.Println("Error", err)
        }
    } else if os.Args[1] == "STOP" { // Turn instances off
        input := &ec2.StopInstancesInput{ //struct 
            InstanceIds: []*string{
                aws.String(os.Args[2]),
            },
            DryRun: aws.Bool(true),
        }
        result, err := svc.StopInstances(input)
        awsErr, ok := err.(awserr.Error)
        if ok && awsErr.Code() == "DryRunOperation" {
            input.DryRun = aws.Bool(false)
            result, err = svc.StopInstances(input)
            if err != nil {
                fmt.Println("Error", err)
            } else {
                fmt.Println("Success", result.StoppingInstances)
            }
        } else {
            fmt.Println("Error", err)
        }
    }
}
