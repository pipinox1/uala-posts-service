import boto3
import os

dynamodb = boto3.client(
    'dynamodb',
    endpoint_url=os.getenv("DYNAMODB_ENDPOINT", "http://localhost:8000"),
    region_name=os.getenv("AWS_REGION", "us-west-2"),
    aws_access_key_id=os.getenv("AWS_ACCESS_KEY_ID", "fake"),
    aws_secret_access_key=os.getenv("AWS_SECRET_ACCESS_KEY", "fake")
)

def create_table():
    try:
        dynamodb.create_table(
            TableName='users-timelines',
            KeySchema=[
                {'AttributeName': 'pk', 'KeyType': 'HASH'},  # Primary key
                {'AttributeName': 'sk', 'KeyType': 'RANGE'}  # Sort key
            ],
            AttributeDefinitions=[
                {'AttributeName': 'pk', 'AttributeType': 'S'},
                {'AttributeName': 'sk', 'AttributeType': 'S'}
            ],
            ProvisionedThroughput={
                'ReadCapacityUnits': 5,
                'WriteCapacityUnits': 5
            }
        )
        print("Table created.")
    except dynamodb.exceptions.ResourceInUseException:
        print("Table already exists.")

create_table()