import { FC } from "react";
import { ListItem, ListItemText } from "@mui/material";
import HighlightedText from "../HighlightedText";
import NodeIcon from "../NodeIcon";

type NodeSearchResult = {
    label: string;
    objectId: string;
    kind: string;
}

const SearchResultItem: FC<{
    item: NodeSearchResult;
    index: number;
    highlightedIndex?: number;
    keyword: string;
    getItemProps: (options: any) => any
}> = ({ item, index, highlightedIndex, keyword, getItemProps }) => {
    return (
        <ListItem
            button
            dense
            selected={highlightedIndex ? highlightedIndex === index : false}
            key={item.objectId}
            data-testid='explore_search_result-list-item'
            {...getItemProps({ item, index })}>
            <ListItemText
                primary={
                    <div
                        style={{
                            width: '100%',
                            display: 'flex',
                            alignItems: 'center',
                        }}>
                        <NodeIcon nodeType={item.kind} />
                        <div
                            style={{
                                flexGrow: 1,
                                marginRight: '1em',
                            }}>
                            <HighlightedText
                                text={item.label || item.objectId}
                                search={keyword}
                            />
                        </div>
                    </div>
                }
                primaryTypographyProps={{
                    style: {
                        whiteSpace: 'nowrap',
                        verticalAlign: 'center',
                    },
                }}
            />
        </ListItem>
    );
}

export default SearchResultItem;