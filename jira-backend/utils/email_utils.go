package utils

import (
	"fmt"
	md "jira-backend/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var emailSession *ses.SES

func InitializeEmailSession() error {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(Vconfig.GetString("email.aws-region")),
		Credentials: credentials.NewStaticCredentials(Vconfig.GetString("email.access_key"), Vconfig.GetString("email.secret_key"), ""),
	})

	emailSession = ses.New(session)
	return err
}

func prepareHTML(projectName string, encryptedString string) string {
	htmlBody := "<h3>Join Jira-Clone </h3>" +
		"You have been invited to project: %s <br/>" +
		"To join the project click on this: <a href='https://localhost:4200/joinproject?invite=%s'>link</a>"
	return fmt.Sprintf(htmlBody, projectName, encryptedString)
}

func SendEmail(project md.Project, emailId string) (string, error) {
	encryptedInvite, err := EncryptEmailInvite(emailId, project.ProjectId)

	if err != nil {
		return "", err
	}

	sesEmailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(emailId)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(prepareHTML(project.ProjectName, encryptedInvite))},
			},
			Subject: &ses.Content{
				Data: aws.String(fmt.Sprintf("%s: Project Invite On Jira-Clone", project.ProjectName)),
			},
		},
		Source: aws.String(Vconfig.GetString("email.email_id")),
		ReplyToAddresses: []*string{
			aws.String("se.gators23@gmail.com"),
		},
	}

	status, err := emailSession.SendEmail(sesEmailInput)

	fmt.Println(status)

	return encryptedInvite, err
}
