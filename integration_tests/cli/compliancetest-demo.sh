#!/bin/bash
set -e
source integration_tests/cli/common.sh

# Preparation of Actors

echo "Assign Vendor role to Jack"
result=$(echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "Vendor" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

test_house_account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "Create TestHouse account with address: $test_house_account"
create_account_with_name $test_house_account

test_house_address=$(zblcli keys show "$test_house_account" -a)

result=$(echo "test1234" | zblcli tx authz assign-role $test_house_address "TestHouse" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

second_test_house_account=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 6 | head -n 1)
echo "Create second TestHouse account with address: $second_test_house_account"
create_account_with_name $second_test_house_account

second_test_house_address=$(zblcli keys show "$second_test_house_account" -a)

result=$(echo "test1234" | zblcli tx authz assign-role $second_test_house_address "TestHouse" --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

# Body

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="http://first.place.com"
test_date="2020-01-01T00:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$testing_result" "$test_date"  --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Second Testing Result for Model VID: $vid PID: $pid"
second_testing_result="http://second.place.com"
second_test_date="2020-04-04T10:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$second_testing_result" $second_test_date --from $second_test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliancetest test-result $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
check_response "$result" "\"test_result\": \"$second_testing_result\""
check_response "$result" "\"test_date\": \"$second_test_date\""
echo "$result"

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: $vid PID: $pid"
result=$(echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Add Testing Result for Model VID: $vid PID: $pid"
testing_result="blob string"
test_date="2020-11-24T10:00:00Z"
result=$(echo "test1234" | zblcli tx compliancetest add-test-result $vid $pid "$testing_result" "$test_date" --from $test_house_account --yes)
check_response "$result" "\"success\": true"
echo "$result"

echo "Get Testing Result for Model with VID: ${vid} PID: ${pid}"
result=$(zblcli query compliancetest test-result $vid $pid)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"test_result\": \"$testing_result\""
check_response "$result" "\"test_date\": \"$test_date\""
echo "$result"