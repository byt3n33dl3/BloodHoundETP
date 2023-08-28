import { findAllByTestId, fireEvent, render, within } from "@testing-library/react";
import SearchCurrentNodes, { NO_RESULTS_TEXT } from "./SearchCurrentNodes";
import { GraphNodes } from "./types";

const nodes: GraphNodes = {
    '1': {
        label: 'computer_node',
        kind: 'Computer',
        objectId: '001',
        isTierZero: false,
        lastSeen: ''
    },
    '2': {
        label: 'user_node',
        kind: 'User',
        objectId: '002',
        isTierZero: false,
        lastSeen: ''
    },
    '3': {
        label: 'group_node',
        kind: 'Group',
        objectId: '003',
        isTierZero: false,
        lastSeen: ''
    }
}

const RESULT_ID = 'explore_search_result-list-item';

describe("SearchCurrentNodes", () => {

    const setup = () => {
        const testOnSelect = vi.fn();
        const screen = render(<SearchCurrentNodes currentNodes={nodes} onSelect={testOnSelect} />);

        const input = screen.getByRole('textbox');
        const resultList = screen.getByTestId('current-results-list');

        const setInputValue = (value: string) => fireEvent.change(input, { target: { value } });

        return {
            testOnSelect,
            ...screen,
            input,
            resultList,
            setInputValue
        }
    }
    
    it('displays an autofocused text input', () => {
        const { input } = setup();

        expect(input).toBeInTheDocument();
        expect(input).toHaveFocus();
    });

    it('displays appropriate text when there are no matches for search term', () => {
        const { setInputValue, queryByText } = setup();

        setInputValue('specterops');
        expect(queryByText(NO_RESULTS_TEXT)).toBeInTheDocument();
    });

    it('displays expected results when searching by label', async () => {
        const { resultList, setInputValue, queryByText, findAllByTestId } = setup();
        
        setInputValue('node');
        expect(resultList).toBeInTheDocument();
        expect(queryByText(NO_RESULTS_TEXT)).not.toBeInTheDocument();
        expect(await findAllByTestId(RESULT_ID)).toHaveLength(3);

        setInputValue('computer');
        expect(await findAllByTestId(RESULT_ID)).toHaveLength(1);
    });

    it('displays expected results when searching by objectid', async () => {
        const { resultList, setInputValue, queryByText, findAllByTestId } = setup();
        
        setInputValue('00');
        expect(resultList).toBeInTheDocument();
        expect(queryByText(NO_RESULTS_TEXT)).not.toBeInTheDocument();
        expect(await findAllByTestId(RESULT_ID)).toHaveLength(3);

        setInputValue('002');
        expect(await findAllByTestId(RESULT_ID)).toHaveLength(1);
    });

    it('displays the label for each result', async () => { 
        const { setInputValue, findByTestId } = setup();

        setInputValue('002');
        expect(await findByTestId(RESULT_ID)).toHaveTextContent('user_node');

        setInputValue('003');
        expect(await findByTestId(RESULT_ID)).toHaveTextContent('group_node');
    });

    it('fires onSelect prop when a result is clicked', async () => {
        const { setInputValue, findByTestId, testOnSelect } = setup();

        setInputValue('computer');
        const result = await findByTestId(RESULT_ID);
        fireEvent.click(result);
        expect(testOnSelect).toHaveBeenCalled();
    });
})