import * as sst from "@serverless-stack/resources";

export default class V1Stack extends sst.Stack {
    api;

    constructor(scope, id, props) {
        super(scope, id, props);

        const { users_table } = props;

        // Create a HTTP API
        this.api = new sst.Api(this, "Api", {
            defaultFunctionProps: {
                environment: {
                    USERNAME_INDEX: "usernameIndex",
                    EMAIL_ADDRESS: "shinobu@bitter.social",
                    USERS_TABLE: users_table.tableName
                },
            },
            routes: {
                "POST /v1/users/register": "src/v1/endpoint/users/register.go",
            },
        });

        this.api.attachPermissions([users_table])

        // Show the endpoint in the output
        this.addOutputs({
            "ApiEndpoint": this.api.url,
        });
    }
}
