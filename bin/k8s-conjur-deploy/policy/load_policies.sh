#!/bin/bash

set -eo pipefail

if [ "$CONJUR_APPLIANCE_URL" != "" ]; then
  conjur init -u $CONJUR_APPLIANCE_URL -a $CONJUR_ACCOUNT
fi

# check for unset vars after checking for appliance url
set -u

conjur authn login -u admin -p $CONJUR_ADMIN_PASSWORD

readonly POLICY_DIR="/policy"

# NOTE: generated files are prefixed with the test app namespace to allow for parallel CI
readonly POLICY_FILES=(
  "$POLICY_DIR/users.yml"
  "$POLICY_DIR/generated/$TEST_APP_NAMESPACE_NAME.project-authn.yml"
  "$POLICY_DIR/generated/$TEST_APP_NAMESPACE_NAME.cluster-authn-svc.yml"
  "$POLICY_DIR/generated/$TEST_APP_NAMESPACE_NAME.app-identity.yml"

)

for policy_file in "${POLICY_FILES[@]}"; do
  echo "Loading policy $policy_file..."
  if [[ "$CONJUR_VERSION" == "4" ]]; then
    conjur policy load --as-group security_admin $policy_file
  elif [[ "$CONJUR_VERSION" == "5" ]]; then
    conjur policy load root $policy_file
  fi
done



conjur authn logout
