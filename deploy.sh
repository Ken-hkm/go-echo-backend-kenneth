#!/bin/bash

# Variables - adjust these as needed.
LAMBDA_FUNCTION_NAME="go-backend-kenneth"
GO_OS="linux"
GO_ARCH="arm64"
OUTPUT_BINARY="bootstrap"  # Lambda expects a binary called 'bootstrap'
ZIP_FILE="deploy.zip"
BUILD_DIR="cmd/lambda"
REST_API_ID="3f0tv6ipo2"  # Replace with your API Gateway REST API ID
DEPLOYMENT_STAGE="dev"  # Change to your preferred stage name

echo "🚀 Building Go binary for AWS Lambda..."

# Navigate to the build directory.
pushd $BUILD_DIR > /dev/null

# Build the binary for Lambda.
GOOS=$GO_OS GOARCH=$GO_ARCH go build -o $OUTPUT_BINARY main.go
if [ $? -ne 0 ]; then
  echo "❌ Build failed in $BUILD_DIR!"
  popd > /dev/null
  exit 1
fi
echo "✅ Build successful in $BUILD_DIR!"

# Move the binary to the root directory.
mv $OUTPUT_BINARY ../../$OUTPUT_BINARY
popd > /dev/null

echo "📦 Zipping deployment package..."
zip -q $ZIP_FILE $OUTPUT_BINARY
if [ $? -ne 0 ]; then
  echo "❌ Failed to create zip package!"
  exit 1
fi

echo "🔄 Updating AWS Lambda function..."
aws lambda update-function-code --function-name $LAMBDA_FUNCTION_NAME --zip-file fileb://$ZIP_FILE --output text
if [ $? -ne 0 ]; then
  echo "❌ Deployment failed!"
  exit 1
fi

echo "🌍 Deploying API Gateway stage..."
aws apigateway create-deployment --rest-api-id $REST_API_ID --stage-name $DEPLOYMENT_STAGE
if [ $? -ne 0 ]; then
  echo "❌ API Gateway deployment failed!"
  exit 1
fi

echo "🎉 Deployment successful! API is live at:"
echo "https://$REST_API_ID.execute-api.$(aws configure get region).amazonaws.com/$DEPLOYMENT_STAGE"