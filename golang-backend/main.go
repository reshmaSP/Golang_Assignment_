
package main

import (
	
	"fmt"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

)

func getInstance(c *gin.Context) {
	// AWS credentials
	viper.SetConfigFile("config.toml")
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("Error reading config file:", err)
        return
    }
	accessKey := viper.GetString("ACCESS_KEY")
    secretKey:= viper.GetString("SECRET_KEY")
	// accessKey := "AKIAXJC476V42U5UUEPX"
	// secretKey := "l5uSK1cPWNcY8hGffQbpWJanqrqP2xosa9fd4Wu9"
	region := "us-east-1" // Replace with your desired AWS region

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create an EC2 service client
	svc := ec2.New(sess)

	// Get the instance ID from the query parameters
	instanceID := c.Query("id")
	fmt.Println(" this is %s", instanceID)
	if instanceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instance ID is required"})
		return
	}

	// Make a request to DescribeInstances for the specific instance ID
	resp, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()

	// Use the /describe-instance route and pass the instance ID as a query parameter
	r.GET("/describe-instance", cors.Default(), getInstance)

	// Run the server
	r.Run("localhost:8080")
}
