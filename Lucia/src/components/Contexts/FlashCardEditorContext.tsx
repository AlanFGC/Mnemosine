import {
  Dispatch, PropsWithChildren, createContext, useMemo, useReducer,
} from 'react';
import FlashCardReducer, { EditorState, FlashCardEditorActions, InitialState } from '../Reducers/FlashcardEditorReducer';

type ContextEditorState = {
  state: EditorState;
  dispatch: Dispatch<FlashCardEditorActions>;
};

const FlashCardEditorContext = createContext<ContextEditorState>(
  { state: InitialState } as ContextEditorState,
);

function FlashCardContextProvider({ children, propEditorState = InitialState }: PropsWithChildren<
// eslint-disable-next-line react/require-default-props
{ propEditorState?: EditorState }
>) {
  const [flashCardState, dispatch] = useReducer(FlashCardReducer, propEditorState);

  const memoizedContextValue = useMemo(
    () => ({ state: flashCardState, dispatch }),
    [flashCardState, dispatch],
  );

  return (
    <FlashCardEditorContext.Provider value={memoizedContextValue}>
      {children}
    </FlashCardEditorContext.Provider>
  );
}

export default FlashCardContextProvider;
