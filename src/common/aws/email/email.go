package email

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ses"
    "main/src/common"
    "main/src/common/aws/session"
)

const (
    Charset = "UTF-8"
)

var sesClient *ses.SES

func Init() {
    sesClient = ses.New(session.Session())
}

func TextMessage(text string, subject string) *ses.Message {
    return &ses.Message{
        Body: &ses.Body{
            Text: Content(text),
        },
        Subject: Content(subject),
    }
}

func HTMLMessage(html string, subject string) *ses.Message {
    return &ses.Message{
        Body: &ses.Body{
            Html: Content(html),
        },
        Subject: Content(subject),
    }
}

func Message(text string, html string, subject string) *ses.Message {
    return &ses.Message{
        Body: &ses.Body{
            Html: Content(html),
            Text: Content(text),
        },
        Subject: Content(subject),
    }
}

func To(to ...string) *ses.Destination {
    var toAddresses []*string

    for _, str := range to {
        toAddresses = append(toAddresses, aws.String(str))
    }

    return &ses.Destination{
        ToAddresses: toAddresses,
    }
}

func Content(body string) *ses.Content {
    return &ses.Content{
        Charset: aws.String(Charset),
        Data: aws.String(body),
    }
}

func Send(destination *ses.Destination, message *ses.Message) (*ses.SendEmailOutput, error) {
    return SendPreMade(&ses.SendEmailInput{
        Destination: destination,
        Message: message,
        Source: aws.String(common.Constants().Email()),
    })
}

func SendPreMade(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
    return sesClient.SendEmail(input)
}
