/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

/* eslint-disable import/order, import/no-duplicates, @typescript-eslint/no-unused-vars */

import * as Types from '../../../../__generated__/schema';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type PolicyDetailsVariables = {
  input: Types.GetPolicyInput;
};

export type PolicyDetails = {
  policy?: Types.Maybe<
    Pick<
      Types.PolicyDetails,
      | 'autoRemediationId'
      | 'autoRemediationParameters'
      | 'description'
      | 'displayName'
      | 'enabled'
      | 'suppressions'
      | 'id'
      | 'reference'
      | 'resourceTypes'
      | 'runbook'
      | 'severity'
      | 'tags'
      | 'body'
    > & {
      tests?: Types.Maybe<
        Array<
          Types.Maybe<
            Pick<Types.PolicyUnitTest, 'expectedResult' | 'name' | 'resource' | 'resourceType'>
          >
        >
      >;
    }
  >;
};

export const PolicyDetailsDocument = gql`
  query PolicyDetails($input: GetPolicyInput!) {
    policy(input: $input) {
      autoRemediationId
      autoRemediationParameters
      description
      displayName
      enabled
      suppressions
      id
      reference
      resourceTypes
      runbook
      severity
      tags
      body
      tests {
        expectedResult
        name
        resource
        resourceType
      }
    }
  }
`;

/**
 * __usePolicyDetails__
 *
 * To run a query within a React component, call `usePolicyDetails` and pass it any options that fit your needs.
 * When your component renders, `usePolicyDetails` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = usePolicyDetails({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function usePolicyDetails(
  baseOptions?: ApolloReactHooks.QueryHookOptions<PolicyDetails, PolicyDetailsVariables>
) {
  return ApolloReactHooks.useQuery<PolicyDetails, PolicyDetailsVariables>(
    PolicyDetailsDocument,
    baseOptions
  );
}
export function usePolicyDetailsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<PolicyDetails, PolicyDetailsVariables>
) {
  return ApolloReactHooks.useLazyQuery<PolicyDetails, PolicyDetailsVariables>(
    PolicyDetailsDocument,
    baseOptions
  );
}
export type PolicyDetailsHookResult = ReturnType<typeof usePolicyDetails>;
export type PolicyDetailsLazyQueryHookResult = ReturnType<typeof usePolicyDetailsLazyQuery>;
export type PolicyDetailsQueryResult = ApolloReactCommon.QueryResult<
  PolicyDetails,
  PolicyDetailsVariables
>;
