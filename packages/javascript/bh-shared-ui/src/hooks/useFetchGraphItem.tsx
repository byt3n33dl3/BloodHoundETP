// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { RequestOptions } from 'js-client-library';
import { entityInformationEndpoints } from '../utils/content';
import { apiClient } from '../utils/api';
import { getNodeByDatabaseIdCypher } from '../utils/entityInfoDisplay';
import { validateNodeType } from '../hooks/useSearch/useSearch';
import { useQuery } from 'react-query';

type GraphItemData = { cacheId: string; objectId: string; nodeType: string; databaseId?: string };
type GraphItemExport = {
    graphItemProperties: { objectid: string; [k: string]: any };
    informationAvailable: boolean;
    isLoading: boolean;
    isError: boolean;
};

export const useFetchGraphItem: (param: GraphItemData) => GraphItemExport = ({
    cacheId,
    objectId,
    nodeType,
    databaseId,
}) => {
    const requestDetails: {
        endpoint: (
            params: string,
            options?: RequestOptions,
            includeProperties?: boolean
        ) => Promise<Record<string, any>>;
        param: string;
    } = {
        endpoint: async function () {
            return {};
        },
        param: '',
    };

    const validatedKind = validateNodeType(nodeType);

    if (validatedKind) {
        requestDetails.endpoint = entityInformationEndpoints[validatedKind];
        requestDetails.param = objectId;
    } else if (databaseId) {
        requestDetails.endpoint = apiClient.cypherSearch;
        requestDetails.param = getNodeByDatabaseIdCypher(databaseId);
    }

    const informationAvailable = !!validatedKind || !!databaseId;

    const {
        data: graphItemProperties,
        isLoading,
        isError,
    } = useQuery(
        [cacheId, nodeType, objectId],
        ({ signal }) =>
            requestDetails.endpoint(requestDetails.param, { signal }, true).then((res) => {
                if (validatedKind) return res.data.data.props;
                else if (databaseId) return Object.values(res.data.data.nodes as Record<string, any>)[0].properties;
                else return {};
            }),
        {
            refetchOnWindowFocus: false,
            retry: false,
            enabled: informationAvailable,
        }
    );
    return {
        graphItemProperties,
        informationAvailable,
        isLoading,
        isError,
    };
};