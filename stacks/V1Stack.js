import * as sst from "@serverless-stack/resources";

export default class V1Stack extends sst.Stack {
    api;

    constructor(scope, id, props) {
        super(scope, id, props);

        const {
            hosted_zone,
            api_url,
            users,
            validation,
            email,
            username_index,
            user_id_index,
        } = props;

        this.setDefaultFunctionProps({
            environment: {
                USERNAME_INDEX: username_index,
                USER_ID_INDEX: user_id_index,
                USERS_TABLE: users.tableName,
                VALIDATION_TABLE: validation.tableName,
            }
        })

        // Create an HTTP API
        this.api = new sst.Api(this, "Api", {
            customDomain: {
                domainName: api_url,
                hostedZone: hosted_zone,
                path: "v1",
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
