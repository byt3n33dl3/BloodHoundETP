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

import EntityInfoPanel from '../Explore/EntityInfo/EntityInfoPanel';
import { DropdownOption, GroupManagementContent } from 'bh-shared-ui';
import { SelectedNode } from 'src/ducks/entityinfo/types';
import { useState } from 'react';
import DataSelector from '../QA/DataSelector';
import { AssetGroup, AssetGroupMember } from 'js-client-library';
import { GraphNodeTypes } from 'src/ducks/graph/types';
import { faGem } from '@fortawesome/free-solid-svg-icons';
import { useSelector } from 'react-redux';
import { Domain } from 'src/ducks/global/types';
import { setSelectedNode } from 'src/ducks/entityinfo/actions';
import { useNavigate } from 'react-router-dom';
import { ROUTE_EXPLORE } from 'src/ducks/global/routes';
import { setSearchValue, startSearchSelected } from 'src/ducks/searchbar/actions';
import { PRIMARY_SEARCH, SEARCH_TYPE_EXACT } from 'src/ducks/searchbar/types';
import { TIER_ZERO_LABEL, TIER_ZERO_TAG } from 'src/constants';
import { useAppDispatch } from 'src/store';

const GroupManagement = () => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const globalDomain: Domain = useSelector((state: any) => state.global.options.domain);

    // Kept out of the shared UI due to diff between GraphNodeTypes across apps
    const [openNode, setOpenNode] = useState<SelectedNode | null>(null);

    const handleClickMember = (member: AssetGroupMember) => {
        setOpenNode({
            id: member.object_id,
            type: member.primary_kind as GraphNodeTypes,
            name: member.name,
        });
    };

    const handleShowNodeInExplore = () => {
        if (openNode) {
            const searchNode = {
                objectid: openNode.id,
                label: openNode.name,
                ...openNode,
            };
            dispatch(setSearchValue(searchNode, PRIMARY_SEARCH, SEARCH_TYPE_EXACT));
            dispatch(startSearchSelected(PRIMARY_SEARCH));
            dispatch(setSelectedNode(openNode));

            navigate(ROUTE_EXPLORE);
        }
    };

    // Handle tier zero case
    const mapAssetGroups = (assetGroups: AssetGroup[]): DropdownOption[] => {
        return assetGroups.map((assetGroup) => {
            const isTierZero = assetGroup.tag === TIER_ZERO_TAG;
            return {
                key: assetGroup.id,
                value: isTierZero ? TIER_ZERO_LABEL : assetGroup.name,
                icon: isTierZero ? faGem : undefined,
            };
        });
    };

    return (
        <GroupManagementContent
            globalDomain={globalDomain}
            showExplorePageLink={!!openNode}
            tierZeroLabel={TIER_ZERO_LABEL}
            tierZeroTag={TIER_ZERO_TAG}
            // Both these components should eventually be moved into the shared UI library
            entityPanelComponent={<EntityInfoPanel selectedNode={openNode} />}
            generateDomainSelectorComponent={(props) => <DataSelector {...props} />}
            onShowNodeInExplore={handleShowNodeInExplore}
            onClickMember={handleClickMember}
            mapAssetGroups={mapAssetGroups}
        />
    );
};

export default GroupManagement;
