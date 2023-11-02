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

import { Menu, MenuItem } from '@mui/material';
import { useNotifications } from 'bh-shared-ui';
import { FC, useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { destinationNodeSuggested, sourceNodeSuggested } from 'src/ducks/searchbar/actions';
import { AppState, useAppDispatch } from 'src/store';

const ContextMenu: FC<{ anchorPosition: { x: number; y: number } }> = ({ anchorPosition }) => {
    const dispatch = useAppDispatch();
    const [open, setOpen] = useState(false);

    const selectedNode = useSelector((state: AppState) => state.entityinfo.selectedNode);

    useEffect(() => {
        if (anchorPosition) {
            setOpen(true);
        } else {
            setOpen(false);
        }
    }, [anchorPosition]);

    const handleClick = () => {
        setOpen(false);
    };

    const handleSetStartingNode = () => {
        if (selectedNode) {
            dispatch(
                sourceNodeSuggested({
                    name: selectedNode.name,
                    objectid: selectedNode.id,
                    type: selectedNode.type,
                })
            );
        }
    };

    const handleSetEndingNode = () => {
        if (selectedNode) {
            dispatch(
                destinationNodeSuggested({
                    name: selectedNode.name,
                    objectid: selectedNode.id,
                    type: selectedNode.type,
                })
            );
        }
    };

    return (
        <Menu
            open={open}
            anchorPosition={{ left: anchorPosition?.x || 0, top: anchorPosition?.y || 0 }}
            anchorReference='anchorPosition'
            onClick={handleClick}>
            <MenuItem onClick={handleSetStartingNode}>Set as starting node</MenuItem>
            <MenuItem onClick={handleSetEndingNode}>Set as ending node</MenuItem>
            <CopyMenuItem />
        </Menu>
    );
};

const CopyMenuItem = () => {
    const { addNotification } = useNotifications();

    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
    const open = Boolean(anchorEl);

    const handleMenuOpen = (event: React.MouseEvent<HTMLLIElement>) => {
        // stop propagation so that parent menu click event one level up doesn't fire
        event.stopPropagation();
        setAnchorEl(event.currentTarget);
    };

    const handleCopyMenuClick = () => {
        setAnchorEl(null);
    };

    const selectedNode = useSelector((state: AppState) => state.entityinfo.selectedNode);

    const handleDisplayName = () => {
        if (selectedNode) {
            navigator.clipboard.writeText(selectedNode.name);
            addNotification(`Display name copied to clipboard`, 'copyToClipboard');
        }
    };

    const handleObjectId = () => {
        if (selectedNode) {
            navigator.clipboard.writeText(selectedNode.id);
            addNotification(`Object ID name copied to clipboard`, 'copyToClipboard');
        }
    };

    const handleCypher = () => {
        if (selectedNode) {
            const cypher = `MATCH (n:${selectedNode.type}) WHERE n.objectid = '${selectedNode.id}' RETURN n`;
            navigator.clipboard.writeText(cypher);
            addNotification(`Cypher copied to clipboard`, 'copyToClipboard');
        }
    };

    return (
        <>
            <MenuItem onClick={handleMenuOpen}>Copy</MenuItem>
            <Menu
                open={open}
                anchorEl={anchorEl}
                anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                }}
                transformOrigin={{
                    vertical: 'top',
                    horizontal: 'left',
                }}
                onClick={handleCopyMenuClick}>
                <MenuItem onClick={handleDisplayName}>Display Name</MenuItem>
                <MenuItem onClick={handleObjectId}>Object ID</MenuItem>
                <MenuItem onClick={handleCypher}>Cypher</MenuItem>
            </Menu>
        </>
    );
};

export default ContextMenu;
