import V1Stack from "./V1Stack";
import StorageStack from "./StorageStack";
import FrontendStack from "./FrontendStack";
import * as iam from "@aws-cdk/aws-iam";

export default function main(app) {
    const HOSTED_ZONE = process.env.HOSTED_ZONE;
    const API_URL = process.env.API_URL;
    const BASE_URL = process.env.BASE_URL;
    const EMAIL_ADDRESS = process.env.EMAIL_ADDRESS;
    const REDIRECT_WWW = process.env.REDIRECT_WWW;

    const EMAIL_POLICY = new iam.PolicyStatement({
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
    });

    // Set default runtime for all functions
    app.setDefaultFunctionProps({
        runtime: "go1.x"
    });

    const storage = new StorageStack(app, "data");

    new V1Stack(app, "v1", {
        hosted_zone: HOSTED_ZONE,
        api_url: API_URL,
        base_url: BASE_URL,
        username_index: storage.username_index,
        user_id_index: storage.user_id_index,
        users: storage.users_table,
        email_address: EMAIL_ADDRESS,
        email: EMAIL_POLICY,
        validation: storage.validation_table,
    });

    // Please make sure that frontend always gets deployed last

    new FrontendStack(app, "frontend", {
        hosted_zone: HOSTED_ZONE,
        domain_name: BASE_URL,
        redirect_www: REDIRECT_WWW,
    })
}
