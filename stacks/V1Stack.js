import * as sst from "@serverless-stack/resources";

export default class V1Stack extends sst.Stack {
    api;

    constructor(scope, id, props) {
        super(scope, id, props);

        const {
            hosted_zone,
            api_url,
            base_url,
            users,
            validation,
            email,
            email_address,
            username_index,
            user_id_index,
        } = props;

        // Create a HTTP API
        this.api = new sst.Api(this, "Api", {
            customDomain: {
                domainName: api_url,
                hostedZone: hosted_zone,
                path: "v1",
            },
            defaultFunctionProps: {
                environment: {
                    API_URL: api_url,
                    BASE_URL: base_url,
                    USERNAME_INDEX: username_index,
                    USER_ID_INDEX: user_id_index,
                    EMAIL_ADDRESS: email_address,
                    USERS_TABLE: users.tableName,
                    VALIDATION_TABLE: validation.tableName,
                },
            },
            routes: {
                "GET /users/{uid}/validate": {
                    function: {
                        handler: "src/v1/endpoint/users/validate/validate.go",
                        permissions: [users, email, validation]
                    }
                },
                "POST /users/register": {
                    function: {
                        handler: "src/v1/endpoint/users/register/register.go",
                        permissions: [users, email, validation],
                    },
                },
            },
        });

        // Show the endpoint in the output
        this.addOutputs({
            "ApiEndpoint": this.api.url,
        });
    }
}
