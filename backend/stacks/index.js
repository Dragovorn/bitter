import V1Stack from "./V1Stack";
import StorageStack from "./StorageStack";
import * as iam from "@aws-cdk/aws-iam";

const EMAIL_ADDRESS = "shinobu@bitter.social"

export default function main(app) {
    const email_policy = new iam.PolicyStatement({
        actions: [
            "ses:SendEmail",
            "ses:SendRawEmail"
        ],
        resources: [ "*" ],
        conditions: {
            "StringEquals": {
                "ses:FromAddress": EMAIL_ADDRESS
            }
        }
    })

    // Set default runtime for all functions
    app.setDefaultFunctionProps({
        runtime: "go1.x"
    });

    const storage = new StorageStack(app, "persist");
    new V1Stack(app, "v1", {
        users: storage.users_table,
        email_address: EMAIL_ADDRESS,
        email: email_policy
    });
}
