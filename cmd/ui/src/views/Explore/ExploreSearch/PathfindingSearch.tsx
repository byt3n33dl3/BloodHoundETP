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

import { Box } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBullseye, faCircle } from '@fortawesome/free-solid-svg-icons';
import { startSearchSelected } from 'src/ducks/searchbar/actions';
import { useDispatch, useSelector } from 'react-redux';
import { useEffect } from 'react';
import { PRIMARY_SEARCH, PATHFINDING_SEARCH } from 'src/ducks/searchbar/types';
import { AppState } from 'src/store';
import EdgeFilter from './EdgeFilter';
import ExploreSearchCombobox from '../ExploreSearchCombobox';
import PathfindingSwapButton from './PathfindingSwapButton';

const PathfindingSearch = () => {
    const dispatch = useDispatch();

    const { primary, secondary } = useSelector((state: AppState) => state.search);

    useEffect(() => {
        if (primary.value && secondary.value) {
            dispatch(startSearchSelected(PATHFINDING_SEARCH));
        } else {
            dispatch(startSearchSelected(PRIMARY_SEARCH));
        }
    }, [primary, secondary, dispatch]);

    return (
        <Box display={'flex'} alignItems={'center'} gap={1}>
            <SourceToBullseyeIcon />

            <Box flexGrow={1} gap={1} display={'flex'} flexDirection={'column'}>
                <ExploreSearchCombobox searchType={PRIMARY_SEARCH} labelText='Start Node' />
                <ExploreSearchCombobox searchType={PATHFINDING_SEARCH} labelText='Destination Node' />
            </Box>

            <PathfindingSwapButton />
            <EdgeFilter />
        </Box>
    );
};

const SourceToBullseyeIcon = () => {
    return (
        <Box display={'flex'} flexDirection={'column'} alignItems={'center'}>
            <FontAwesomeIcon icon={faCircle} size='xs' />
            <Box
                border={'none'}
                borderLeft={'1px dotted black'}
                marginTop={'0.5em'}
                marginBottom={'0.5em'}
                height='1em'></Box>
            <FontAwesomeIcon icon={faBullseye} size='xs' />
        </Box>
    );
};

export default PathfindingSearch;
