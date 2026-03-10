#!/bin/bash

# Workflow validation script for Temporal workflows
# Usage: ./validate-workflow.sh [workflow-name]

set -e

WORKFLOW_NAME=${1:-"basic-workflow"}
WORKFLOW_DIR="backend/workflows"
ACTIVITY_DIR="backend/activities"

echo "🔍 Validating Temporal workflow: $WORKFLOW_NAME"

# Check if workflow file exists
if [ ! -f "$WORKFLOW_DIR/${WORKFLOW_NAME}.go" ]; then
    echo "❌ Workflow file not found: $WORKFLOW_DIR/${WORKFLOW_NAME}.go"
    exit 1
fi

# Check for required imports
echo "📦 Checking imports..."
if ! grep -q "go.temporal.io/sdk/workflow" "$WORKFLOW_DIR/${WORKFLOW_NAME}.go"; then
    echo "❌ Missing workflow import"
    exit 1
fi

# Check workflow function signature
echo "🔧 Checking workflow function signature..."
if ! grep -q "func.*Workflow.*workflow.Context" "$WORKFLOW_DIR/${WORKFLOW_NAME}.go"; then
    echo "❌ Invalid workflow function signature"
    exit 1
fi

# Check for proper error handling
echo "⚠️  Checking error handling..."
if ! grep -q "return.*err" "$WORKFLOW_DIR/${WORKFLOW_NAME}.go"; then
    echo "⚠️  Warning: No error handling found"
fi

# Check for logging
echo "📝 Checking logging..."
if ! grep -q "workflow.GetLogger" "$WORKFLOW_DIR/${WORKFLOW_NAME}.go"; then
    echo "⚠️  Warning: No logging found"
fi

# Validate Go syntax
echo "🐹 Validating Go syntax..."
if ! go fmt "$WORKFLOW_DIR/${WORKFLOW_NAME}.go" > /dev/null; then
    echo "❌ Go syntax error"
    exit 1
fi

# Run tests if they exist
if [ -f "$WORKFLOW_DIR/${WORKFLOW_NAME}_test.go" ]; then
    echo "🧪 Running tests..."
    cd "$WORKFLOW_DIR"
    if ! go test -v -run "$WORKFLOW_NAME"; then
        echo "❌ Tests failed"
        exit 1
    fi
    cd - > /dev/null
else
    echo "⚠️  No tests found"
fi

echo "✅ Workflow validation completed successfully!"
