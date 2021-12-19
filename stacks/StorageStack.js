import * as sst from "@serverless-stack/resources";
import {RemovalPolicy} from "@aws-cdk/core";

export default class StorageStack extends sst.Stack {
    users_table;
    validation_table;
    user_content_bucket;
    username_index = "usernameIndex";
    user_id_index = "userIdIndex";

    constructor(scope, id, props) {
        super(scope, id, props);

        this.users_table = new sst.Table(this, "users", {
            fields: {
                uid: sst.TableFieldType.STRING,
                username: sst.TableFieldType.STRING,
            },
            primaryIndex: {
                partitionKey: "uid",
                sortKey: "username",
            },
            globalIndexes: {
                usernameIndex: {
                    partitionKey: "username",
                }
            },
            dynamodbTable: {
                removalPolicy: RemovalPolicy.DESTROY
            }
        });

        this.validation_table = new sst.Table(this, "validation", {
            fields: {
                code: sst.TableFieldType.STRING,
                user_id: sst.TableFieldType.STRING,
            },
            primaryIndex: {
                partitionKey: "code",
                sortKey: "user_id"
            },
            dynamodbTable: {
                removalPolicy: RemovalPolicy.DESTROY
            }
        });

        this.user_content_bucket = new sst.Bucket(this, "user-content", {
            s3Bucket: {
                removalPolicy: RemovalPolicy.DESTROY
            }
        });
    }
}
